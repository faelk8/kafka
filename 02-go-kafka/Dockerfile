FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y
    # librdkakfa - lib em C para utilizar o Kafka em outras linguaguens, sem ele não funciona para as outras linguagens

CMD ["tail", "-f", "/dev/null"]