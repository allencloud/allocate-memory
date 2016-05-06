FROM ubuntu:14.04

RUN	apt-get update && \
	apt-get install -y wget g++ git && \
	wget -q https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz && \
	tar -C /usr/local -xzf go1.4.linux-amd64.tar.gz && \
	rm go1.4.linux-amd64.tar.gz && \
	apt-get clean

RUN mkdir /root/allocate-memory

ADD . /root/allocate-memory

WORKDIR /root/allocate-memory

RUN g++ memory.cpp -o memory.out

RUN export PATH=$PATH:/usr/local/go/bin && \
	export GOPATH=/root/allocate-memory && \
	go get -d && \
	go build allocate-memory.go

EXPOSE 8080

CMD ["./allocate-memory"]
