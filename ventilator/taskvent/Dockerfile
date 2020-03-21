FROM python:3.8

ENV WORKDIR=/usr/src/app
ENV BIND_PORT=5557

WORKDIR $WORKDIR

RUN pip install zmq

COPY ventilator.py $WORKDIR

EXPOSE $BIND_PORT

CMD ["python", "ventilator.py"]
