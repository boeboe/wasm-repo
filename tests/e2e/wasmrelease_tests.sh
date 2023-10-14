#!/bin/bash

BASE_URL="http://localhost:8080"
PLUGIN_ENDPOINT="$BASE_URL/plugins"

# Test: Create a new plugin
echo "Creating a new plugin..."
PLUGIN_RESPONSE=$(curl -s -X POST $PLUGIN_ENDPOINT \
     -H "Content-Type: application/json" \
     -d '{
         "Name": "TestPlugin",
         "Owner": "TestOwner",
         "Description": "This is a test plugin",
         "Type": "TestType"
     }')
echo $PLUGIN_RESPONSE
echo

# Extract the plugin ID from the response
PLUGIN_ID=$(echo $PLUGIN_RESPONSE | jq -r '.ID')
RELEASE_ENDPOINT="$PLUGIN_ENDPOINT/$PLUGIN_ID/releases"

# Test: Create a new release for the plugin
echo "Creating a new release for the plugin..."
RELEASE_RESPONSE=$(curl -s -X POST $RELEASE_ENDPOINT \
     -H "Content-Type: application/json" \
     -d '{
         "Version": "1.0.0",
         "Sha256": "test-sha256",
         "Description": "This is a test release",
         "Size": 100
     }')
echo $RELEASE_RESPONSE
echo

# Extract the release ID from the response
RELEASE_ID=$(echo $RELEASE_RESPONSE | jq -r '.ID')

# Test: List all releases for the plugin
echo "Listing all releases for the plugin..."
curl -X GET $RELEASE_ENDPOINT
echo

# Test: Retrieve a release by ID
echo "Retrieving a release by ID..."
curl -X GET "$RELEASE_ENDPOINT/$RELEASE_ID"
echo

# Test: Update a release
echo "Updating a release..."
curl -X PUT "$RELEASE_ENDPOINT/$RELEASE_ID" \
     -H "Content-Type: application/json" \
     -d '{
         "Version": "1.0.1",
         "Description": "This is an updated test release"
     }'
echo

# Cleanup: Delete the release
echo "Deleting the release..."
curl -X DELETE "$RELEASE_ENDPOINT/$RELEASE_ID"
echo

# Cleanup: Delete the plugin
echo "Deleting the plugin..."
curl -X DELETE "$PLUGIN_ENDPOINT/$PLUGIN_ID"
echo

echo "End-to-end test completed."
