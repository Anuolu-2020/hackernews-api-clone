#!/bin/sh

# Extract DB_URL from the .env file
DB_URL=$(grep -E '^DB_URL=' ./.env | cut -d '=' -f 2-)

# Run atlas migrate with provided arguments
atlas migrate "$@" --dir "file://internal/db/migrations" --url "$DB_URL" --revisions-schema public
