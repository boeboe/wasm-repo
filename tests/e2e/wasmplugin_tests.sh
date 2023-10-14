#!/bin/bash

BASE_URL="http://localhost:8080"
PLUGIN_ENDPOINT="$BASE_URL/plugins"

# Test: Create a new plugin
echo "Creating a new plugin..."
response=$(curl -s -X POST $PLUGIN_ENDPOINT \
     -H "Content-Type: application/json" \
     -d '{
         "Name": "TestPlugin",
         "Owner": "TestOwner",
         "Description": "This is a test plugin",
         "Type": "TestType"
     }')
echo $response
echo

# Extract the ID from the response
plugin_id=$(echo $response | jq -r '.ID')

# Test: List all plugins
echo "Listing all plugins..."
curl -X GET $PLUGIN_ENDPOINT
echo

# Test: Retrieve a plugin by ID
echo "Retrieving a plugin by ID..."
curl -X GET "$PLUGIN_ENDPOINT/$plugin_id"
echo

# Test: Update a plugin
echo "Updating a plugin..."
curl -X PUT "$PLUGIN_ENDPOINT/$plugin_id" \
     -H "Content-Type: application/json" \
     -d '{
         "Name": "UpdatedPlugin",
         "Owner": "UpdatedOwner"
     }'
echo

# Test: Delete a plugin
echo "Deleting a plugin..."
curl -X DELETE "$PLUGIN_ENDPOINT/$plugin_id"
echo
