#!/bin/bash

BASE_URL="http://localhost:8080"
PLUGIN_ENDPOINT="$BASE_URL/plugins"

# Generate a unique plugin name using a timestamp
PLUGIN_NAME="TestPlugin_$(date +%s)"
UPDATE_PLUGIN_NAME="UpdatedPlugin_$(date +%s)"

# Test: Create a new plugin
echo "Creating a new plugin..."
response=$(curl -s -X POST $PLUGIN_ENDPOINT \
     -H "Content-Type: application/json" \
     -d "{
         \"Name\": \"$PLUGIN_NAME\",
         \"Owner\": \"TestOwner\",
         \"Description\": \"This is a test plugin\",
         \"Type\": \"TestType\"
     }")
echo $response
echo

# Extract the ID from the response
PLUGIN_ID=$(echo $response | jq -r '.ID')
if [[ "$PLUGIN_ID" == "null" || -z "$PLUGIN_ID" ]]; then
    echo "Failed to extract plugin ID from response."
    exit 1
fi

# Test: List all plugins
echo "Listing all plugins..."
curl -X GET $PLUGIN_ENDPOINT
echo

# Test: Retrieve a plugin by ID
echo "Retrieving a plugin by ID..."
curl -X GET "$PLUGIN_ENDPOINT/$PLUGIN_ID"
echo

# Test: Update a plugin
echo "Updating a plugin..."
curl -X PUT "$PLUGIN_ENDPOINT/$PLUGIN_ID" \
     -H "Content-Type: application/json" \
     -d "{
         \"Name\": \"$UPDATE_PLUGIN_NAME\",
         \"Owner\": \"UpdatedOwner\",
         \"Description\": \"This is a updated test plugin\",
         \"Type\": \"TestType\"
     }"
echo

# Test: Delete a plugin
echo "Deleting a plugin..."
curl -X DELETE "$PLUGIN_ENDPOINT/$PLUGIN_ID"
echo
