#!/bin/sh

# Extract DB_URL from the .env file
DB_URL=$(grep -E '^DB_URL=' ./.env | cut -d '=' -f 2-)

atlas schema clean -u "$DB_URL"
