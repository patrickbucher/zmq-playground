#!/usr/bin/bash

for i in {1..100}
do
    curl 'http://localhost:8080/add?a=100&b=200'
    echo
done
