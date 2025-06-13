#!/bin/sh

# Integration tests
set -e

command=${1:-frontend}
if [ "$command" != "frontend" ] && [ "$command" != "backend" ]; then
    echo "Invalid command. Must be 'frontend' or 'backend'"
    exit 1
fi

if [ "$command" = "backend" ]; then
    lang=${2:-js}
    framework=${3:-react}
    database=${4:-none}
else
    framework=${2:-react}
    css=${3:-tailwind}
    lang=${4:-js}
fi

echo "Running integration test for $command"
if [ "$command" = "backend" ]; then
    echo "Command: $command"
    echo "Framework: $framework"
    echo "CSS: $css"
    echo "Language: $lang"
    echo "Database: $database"
else
    echo "Command: $command"
    echo "Framework: $framework"
    echo "CSS: $css"
    echo "Language: $lang"
fi

# Clean up any existing test directory
rm -rf integration-test

if [ "$command" = "backend" ]; then
    echo "Building backend"
    sersi create $command -n integration-test --framework $framework --lang $lang --database $database
    tree integration-test
else
    echo "Building frontend"
    sersi create $command -n integration-test --framework $framework --css $css --lang $lang
    tree integration-test
    echo "Checking frontend build"
    cd integration-test
    npm install
    npm run build
fi

# Clean up after tests
rm -rf integration-test
exit 0