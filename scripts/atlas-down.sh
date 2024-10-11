#!/bin/sh

# Extract DB_URL and DEV_DB_URL from the .env file
DB_URL=$(grep -E '^DB_URL=' ./.env | cut -d '=' -f 2-)
DEV_DB_URL=$(grep -E '^DEV_DB_URL=' ./.env | cut -d '=' -f 2-)

# Check if both variables were extracted correctly
if [ -z "$DB_URL" ] || [ -z "$DEV_DB_URL" ]; then
  echo "Error: DB_URL or DEV_DB_URL is not set in the .env file"
  exit 1
fi

# Run atlas migrate with provided arguments
atlas migrate "$@" --dir "file://internal/db/migrations" --url "$DB_URL" --dev-url "$DEV_DB_URL" 
