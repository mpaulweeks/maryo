#!/bin/sh
echo "{}" > cache.json
while [ true ]
do
    ./maryo
    sleep 60
done
