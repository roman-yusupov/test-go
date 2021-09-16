#!/bin/bash

docker build --tag server .

docker images

export GRPC_PORT=5100
docker run --publish $GRPC_PORT:5100 server

