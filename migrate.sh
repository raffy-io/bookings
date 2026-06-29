#!/bin/sh
if [ -z "$1" ]; then
    echo "❌ Error: Please provide a goose command (e.g., ./migrate.sh up)"
    exit 1
fi

# Get the directory where this script actually lives
SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="postgres://postgres:secret@localhost:5431/bookings"
# Use the absolute path based on the script location
export GOOSE_MIGRATION_DIR="$SCRIPT_DIR/internal/schema/migrations"

echo "🚀 Running goose $1..."
goose "$@"