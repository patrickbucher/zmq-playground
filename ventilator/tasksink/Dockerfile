FROM python:3.8

ENV WORKDIR=/usr/src/app
ENV SOURCE_PORT=5558

WORKDIR $WORKDIR

RUN pip install zmq

COPY sink.py $WORKDIR

EXPOSE $SOURCE_PORT

CMD ["python", "sink.py"]
