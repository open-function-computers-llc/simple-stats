#!/bin/bash

echo -n "Killing any previously running build... "
killall simple-stats || true
echo "Done"

echo "Building"
go build -o dist/simple-stats
./dist/simple-stats &
echo "App running..."
