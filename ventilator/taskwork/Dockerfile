FROM python:3.8

ENV WORKDIR=/usr/src/app
ENV SOURCE_PORT=5557
ENV SINK_PORT=5558

WORKDIR $WORKDIR

RUN pip install zmq

COPY worker.py $WORKDIR

EXPOSE $SOURCE_PORT $SINK_PORT

CMD ["python", "worker.py", "worker"]
