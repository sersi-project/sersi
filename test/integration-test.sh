#!/bin/bash

# TODO: Implement integration tests
echo "Integration tests not implemented yet"

# Test scaffold
chmod +x ./bin/sersi
./bin/sersi frontend -n integration-test -f react -c tailwind -l js
