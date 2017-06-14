#!/bin/sh

docker run --rm -v $PWD:/scripts williamyeh/wrk -t3000 -c9000 -d60s  -s /scripts/update-location.lua http://172.17.0.1:8000/location