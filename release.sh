#!/bin/bash

set -e

# gf pack resource internal/packed/resource.go

# gf build -a amd64 -s linux -n xyhelper-gpt

# rm -rf ./internal/packed/resource.go
# docker build -f Dockerfile.release -t xyhelper/xyhelper-gpt:latest .
# docker push xyhelper/xyhelper-gpt:latest

gf docker main.go -p -t xyhelper/xyhelper-gpt:latest