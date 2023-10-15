#!/bin/bash

BASE_URL="http://localhost:8080"
PLUGIN_ENDPOINT="$BASE_URL/plugins"
FILES_ENDPOINT="$BASE_URL/files"

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
echo "Plugin Response: $PLUGIN_RESPONSE"
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
         "Description": "This is a test release"
     }')
echo "Release Response: $RELEASE_RESPONSE"
echo

# Extract the release ID from the response
RELEASE_ID=$(echo $RELEASE_RESPONSE | jq -r '.ID')

# Generate a random binary file
FILE_NAME="random_file_$(date +%s).bin"
dd if=/dev/urandom of=$FILE_NAME bs=1M count=1
echo "Generated file: $FILE_NAME"

# Test: Upload the binary file
echo "Uploading the binary file..."
UPLOAD_RESPONSE=$(curl -s -X POST $FILES_ENDPOINT \
     -F "file=@$FILE_NAME" \
     -F "pluginID=$PLUGIN_ID" \
     -F "releaseID=$RELEASE_ID")
echo "Upload Response: $UPLOAD_RESPONSE"
echo

# Extract the file ID from the response
FILE_ID=$(echo $UPLOAD_RESPONSE | jq -r '.ID')
if [[ "$FILE_ID" == "null" || -z "$FILE_ID" ]]; then
    echo "Failed to extract file ID from response."
    exit 1
fi

# Test: Download the uploaded file
echo "Downloading the uploaded file..."
DOWNLOAD_FILE_NAME="downloaded_$FILE_NAME"
curl -s -X GET "$FILES_ENDPOINT/$FILE_ID" -o $DOWNLOAD_FILE_NAME
echo "Downloaded file saved as: $DOWNLOAD_FILE_NAME"

# Compare the original and downloaded files
if cmp -s "$FILE_NAME" "$DOWNLOAD_FILE_NAME"; then
   echo "The files are identical."
else
   echo "The files are different."
fi

# Cleanup: Delete the generated files and the created plugin and release
rm $FILE_NAME $DOWNLOAD_FILE_NAME
curl -X DELETE "$PLUGIN_ENDPOINT/$PLUGIN_ID"

echo "End-to-end test completed."
