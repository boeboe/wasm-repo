#!/bin/bash

BASE_URL="http://localhost:8080"
FILES_ENDPOINT="$BASE_URL/files"

# Generate a random binary file
FILE_NAME="random_file_$(date +%s).bin"
dd if=/dev/urandom of=$FILE_NAME bs=1M count=1

echo "Generated file: $FILE_NAME"

# Test: Upload the binary file
echo "Uploading the binary file..."
UPLOAD_RESPONSE=$(curl -s -X POST $FILES_ENDPOINT -F "file=@$FILE_NAME")
echo "Upload Response: $UPLOAD_RESPONSE"
echo

# Extract the file ID from the response
FILE_ID=$(echo $UPLOAD_RESPONSE | jq -r '.ID' 2>/dev/null)
if [[ "$FILE_ID" == "null" || -z "$FILE_ID" ]]; then
    echo "Failed to extract file ID from response."
    exit 1
fi

# Test: Download the uploaded file
echo "Downloading the uploaded file..."
DOWNLOAD_FILE_NAME="downloaded_$FILE_NAME"
curl -s -X GET "$FILES_ENDPOINT/$FILE_ID" -o $DOWNLOAD_FILE_NAME
echo "Downloaded file saved as: $DOWNLOAD_FILE_NAME"

# Optionally, you can add logic to compare the original and downloaded files
if cmp -s "$FILE_NAME" "$DOWNLOAD_FILE_NAME"; then
   echo "The files are identical."
else
   echo "The files are different."
fi

# Cleanup: Delete the generated files
rm $FILE_NAME $DOWNLOAD_FILE_NAME

echo "End-to-end test completed."
