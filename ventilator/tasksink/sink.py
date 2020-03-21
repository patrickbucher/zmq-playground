import os
import zmq

context = zmq.Context()

receiver = context.socket(zmq.PULL)
source_port = os.getenv('BIND_URI', 'tcp://0.0.0.0:5558')
receiver.bind(source_port)

while True:
    result = receiver.recv()
    print(result.decode('utf-8'))
