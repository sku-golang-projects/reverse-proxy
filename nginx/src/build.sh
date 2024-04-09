#!/bin/bash

echo "Building home component..."
cd home
./build.sh
cd ..
echo "Exiting from home..."

echo "Building inventory component..."
cd inventory
./build.sh
cd ..
echo "Exiting from inventory..."

echo "Building order component..."
cd order
./build.sh
cd ..
echo "Exiting from order..."

echo "Building nodeserver component..."
cd nodeserver
./build.sh
cd ..
echo "Exiting from nodeserver..."
