#!/bin/bash

echo "Deleting old build"
rm -f daemon.run && \

set GOARCH=amd64 && \
set GOOS=linux && \

echo "Compiling code..." && \
go build -ldflags="-s -w" -o daemon.run && \

echo "Modifying permissions..." && \
chmod +x daemon.run && \

echo "Done"