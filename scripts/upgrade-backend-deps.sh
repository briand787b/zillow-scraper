#!/bin/bash

# reposition to docker-compose dir
cd ..

docker-compose run \
    --rm \
    --no-deps \
    -v ${PWD}/backend:/