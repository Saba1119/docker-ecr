#!/bin/bash
echo "updating file permissions"
chown -R ubuntu:ubuntu /home/ubuntu/goapp/
chmod +x /home/ubuntu/goapp/main
