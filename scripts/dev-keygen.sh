#!/bin/bash

echo -n "123" > /tmp/eth_pk_password
echo "Private key password saved to /tmp/eth_pk_password"

# Generate new account with geth and capture the output
output=$(geth account new --lightkdf --datadir /tmp --password /tmp/eth_pk_password)

# Extract the file path of the secret key file using sed
file_path=$(echo "$output" | sed -n 's/Path of the secret key file: //p')

# Check if the file path was extracted successfully
if [ -n "$file_path" ]; then
    # Rename the file to /tmp/eth_pk_file
    mv "$file_path" /tmp/eth_pk_file
    export ETH_PRIVATE_KEY_FILE="/tmp/eth_pk_file"
    export ETH_PRIVATE_KEY_PASSWORD_FILE="/tmp/eth_pk_password"
    echo "Private key saved to /tmp/eth_pk_file"
else
    echo "Failed to extract the file path from the command output"
    exit 1
fi
