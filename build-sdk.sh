#!/bin/bash

docker run --rm -v "${PWD}:/local" -v /Users/abewang/Downloads/console-apidocs.json:/console-apidocs.json openapitools/openapi-generator-cli generate \
    -i /console-apidocs.json \
    -g go \
    -o "/local"