#!/bin/bash

docker build -f Dockerfile -t nodeserver:0.1 .
#docker save -o nodeserver.tar nodeserver:0.1
