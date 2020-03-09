import os
import zmq

context = zmq.Context()

receiver = context.socket(zmq.PULL)
source_port = os.getenv('SOURCE_PORT', '5558')
receiver.bind(f'tcp://0.0.0.0:{source_port}')

while True:
    result = receiver.recv()
    print(result.decode('utf-8'))
