FROM golang:1.4

RUN mkdir /root/allocate-memory

ADD . /root/allocate-memory

WORKDIR /root/allocate-memory

RUN export GOPATH=/root/allocate-memory && \
	go get -d && \
	go build

CMD ["./allocate-memory"]
