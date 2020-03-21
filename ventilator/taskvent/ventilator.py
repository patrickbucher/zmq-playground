import os
import random
import time
import zmq

context = zmq.Context()

sender = context.socket(zmq.PUSH)
bind_uri = os.getenv('BIND_URI', 'tcp://0.0.0.0:5557')
sender.bind(bind_uri)

random.seed()
while True:
    items = []
    for item in range(10):
        items.append(str(random.randint(1, 11)))
    sender.send_string(','.join(items))

time.sleep(3)
