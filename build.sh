#!/bin/sh
#build golang, build to linux 386 because docker mac os use this env
#
env GOOS=linux GOARCH=386 go build
docker build --rm -f Dockerfile -t gosample:latest .