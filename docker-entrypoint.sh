#!/bin/sh
# Abort on any error (including if wait-for-it fails). set -e
# Wait for DB

 /go/src/app/wait-for-it.sh "postgres:5432" -t 20

 # Run the main container command.
exec "$@"