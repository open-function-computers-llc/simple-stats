#!/bin/bash

echo -n "Killing any previously running build... "
killall simple-stats || true
echo "Done"

echo "Building"
go build -o dist/simple-stats

export OFCOSTATSTOKEN=SECRETTOKEN12345

./dist/simple-stats &
echo "App running..."
