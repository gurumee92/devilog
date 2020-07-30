#!/bin/bash

if [ "$1" == "start" ]; then
    docker-compose up --build

elif [ "$1" == "stop" ]; then
    docker-compose down -v

else
    echo "please input start|stop"
fi