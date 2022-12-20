#!/bin/bash
echo "starting Docker"

docker run -d --goapp -p 8081:8081 554248189203.dkr.ecr.us-east-1.amazonaws.com/goapp:demo
