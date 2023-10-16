#!/bin/bash

BASE_URL="http://localhost:8080"
PLUGIN_ENDPOINT="$BASE_URL/plugins"

# Array of valid plugin types
VALID_PLUGIN_TYPES=("HttpFilter" "NetworkFilter" "WasmService")

# Generate a unique plugin name using a timestamp
PLUGIN_NAME="TestPlugin_$(date +%s)"
UPDATED_PLUGIN_NAME="UpdatedPlugin_$(date +%s)"

# Randomly select a valid plugin type
RANDOM_TYPE=${VALID_PLUGIN_TYPES[$RANDOM % ${#VALID_PLUGIN_TYPES[@]}]}
UPDATED_RANDOM_TYPE=${VALID_PLUGIN_TYPES[$RANDOM % ${#VALID_PLUGIN_TYPES[@]}]}

# Test: Create a new plugin
echo "Creating a new plugin..."
response=$(curl -s -X POST $PLUGIN_ENDPOINT \
     -H "Content-Type: application/json" \
     -d "{
         \"Name\": \"$PLUGIN_NAME\",
         \"Owner\": \"TestOwner\",
         \"Description\": \"This is a test plugin\",
         \"Type\": \"$RANDOM_TYPE\"
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
         \"Name\": \"$UPDATED_PLUGIN_NAME\",
         \"Owner\": \"UpdatedOwner\",
         \"Description\": \"This is a updated test plugin\",
         \"Type\": \"$UPDATED_RANDOM_TYPE\"
     }"
echo

# Test: Delete a plugin
echo "Deleting a plugin..."
curl -X DELETE "$PLUGIN_ENDPOINT/$PLUGIN_ID"
echo
