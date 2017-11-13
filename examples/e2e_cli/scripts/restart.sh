#!/bin/bash

sudo docker stop $(sudo docker ps -aq)
sudo docker rm $(sudo docker ps -aq)

sudo docker rmi $(sudo docker images | grep dev)
echo
sudo docker-compose -f ../docker-compose-cli.yaml up -d