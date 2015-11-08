#!/bin/sh
echo "{}" > cache.json
go build
while [ true ]
do
    ./maryo
    sleep 60
done
