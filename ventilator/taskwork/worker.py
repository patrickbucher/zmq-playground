import os
import sys
import zmq

if len(sys.argv) < 2:
    print('usage: worker.py [name]')
    sys.exit(1)

name = sys.argv[1]

context = zmq.Context()

receiver = context.socket(zmq.PULL)
source_port = os.getenv('SOURCE_PORT', '5557')
receiver.connect(f'tcp://localhost:{source_port}')

sender = context.socket(zmq.PUSH)
sink_port = os.getenv('SINK_PORT', '5558')
sender.connect(f'tcp://localhost:{sink_port}')

while True:
    task = receiver.recv().decode('utf-8')
    numbers = task.split(',')

    accumulator = 0
    for number in numbers:
        accumulator += int(number)

    sender.send_string(str(f'{name}: {accumulator}'))
