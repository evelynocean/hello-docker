FROM golang:1.9.2

RUN git clone https://github.com/evelynocean/hello-docker.git /go/src/hello &&\
    cd /go/src/hello
RUN make docker build

WORKDIR /go/src/hello

EXPOSE 1031

ENTRYPOINT ["./hello"]