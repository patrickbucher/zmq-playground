#!/usr/bin/bash

source env/bin/activate

python taskvent/ventilator.py &
vent_pid="$!"

python taskwork/worker.py 1 &
worker_1_pid="$!"

python taskwork/worker.py 2 &
worker_2_pid="$!"

python taskwork/worker.py 3 &
worker_3_pid="$!"

python tasksink/sink.py &
sink_pid="$!"

sleep 1
kill "$vent_pid"

sleep 1
kill "$worker_1_pid"
kill "$worker_2_pid"
kill "$worker_3_pid"

sleep 1
kill "$sink_pid"
