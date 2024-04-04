#!/bin/bash

docker build -f Dockerfile -t inventory:0.1 .
#docker save -o inventory.tar inventory:0.1
