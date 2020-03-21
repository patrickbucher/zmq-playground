import os
import sys
import zmq

if len(sys.argv) < 2:
    print('usage: worker.py [name]')
    sys.exit(1)

name = sys.argv[1]

context = zmq.Context()

receiver = context.socket(zmq.PULL)
source_uri = os.getenv('SOURCE_URI')
receiver.connect(source_uri)

sender = context.socket(zmq.PUSH)
sink_uri = os.getenv('SINK_URI')
sender.connect(sink_uri)

while True:
    task = receiver.recv().decode('utf-8')
    numbers = task.split(',')

    accumulator = 0
    for number in numbers:
        accumulator += int(number)

    sender.send_string(str(f'{name}: {accumulator}'))
