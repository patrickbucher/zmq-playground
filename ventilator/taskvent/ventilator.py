import os
import random
import time
import zmq

context = zmq.Context()

sender = context.socket(zmq.PUSH)
bind_port = os.getenv('BIND_PORT', '5557')
sender.bind(f'tcp://0.0.0.0:{bind_port}')

sink = context.socket(zmq.PUSH)
sink_port = os.getenv('SINK_PORT', '5558')
sink.connect(f'tcp://localhost:{sink_port}')

random.seed()
while True:
    items = []
    for item in range(10):
        items.append(str(random.randint(1, 11)))
    sender.send_string(','.join(items))

time.sleep(3)
