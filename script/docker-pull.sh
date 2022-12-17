#!/bin/bash
echo "Pulling image"
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 554248189203.dkr.ecr.us-east-1.amazonaws.com
docker pull 554248189203.dkr.ecr.us-east-1.amazonaws.com/goapp:new
