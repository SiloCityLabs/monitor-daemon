#!/bin/bash

echo "Deleting old build"
rm -f plex-daemon.run && \

set GOARCH=amd64 && \
set GOOS=linux && \

echo "Compiling code..." && \
go build -ldflags="-s -w" -o plex-daemon.run && \

echo "Modifying permissions..." && \
chmod +x plex-daemon.run && \

echo "Done"