#!/usr/bin/bash

for i in {1..100}
do
    curl 'http://localhost:8080/add?a=1&b=2'
    echo
done
