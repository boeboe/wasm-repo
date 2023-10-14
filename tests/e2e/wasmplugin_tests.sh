#!/bin/bash

BASE_URL="http://localhost:8080"
PLUGIN_ENDPOINT="$BASE_URL/plugins"

# Test: Create a new plugin
echo "Creating a new plugin..."
curl -X POST $PLUGIN_ENDPOINT \
     -H "Content-Type: application/json" \
     -d '{
         "Name": "TestPlugin",
         "Owner": "TestOwner",
         "Description": "This is a test plugin",
         "Type": "TestType"
     }'
echo

# Test: List all plugins
echo "Listing all plugins..."
curl -X GET $PLUGIN_ENDPOINT
echo

# Assuming the ID of the created plugin is 'test-id', you can retrieve, update, and delete using that ID.
# You might want to extract the ID dynamically from the creation response for a more robust test.

# Test: Retrieve a plugin by ID
echo "Retrieving a plugin by ID..."
curl -X GET "$PLUGIN_ENDPOINT/test-id"
echo

# Test: Update a plugin
echo "Updating a plugin..."
curl -X PUT "$PLUGIN_ENDPOINT/test-id" \
     -H "Content-Type: application/json" \
     -d '{
         "Name": "UpdatedPlugin",
         "Owner": "UpdatedOwner"
     }'
echo

# Test: Delete a plugin
echo "Deleting a plugin..."
curl -X DELETE "$PLUGIN_ENDPOINT/test-id"
echo
