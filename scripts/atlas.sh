#!/bin/sh

# Extract DB_URL from the .env file
DEV_DB_URL=$(grep -E '^DEV_DB_URL=' ./.env | cut -d '=' -f 2-)

# Run atlas migrate with provided arguments
atlas migrate "$@" --var db_url="$DEV_DB_URL"
