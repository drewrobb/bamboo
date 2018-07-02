#!/usr/bin/env bash

label="upgrade-marathon-1.6-b15e455"
image="docker.strava.com/strava/bamboo:${label}"
echo $image
docker build -t $image .
docker push $image
