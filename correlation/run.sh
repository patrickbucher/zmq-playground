#!/usr/bin/bash

./client-a.sh >a.out 2>/dev/null &
./client-b.sh >b.out 2>/dev/null &
