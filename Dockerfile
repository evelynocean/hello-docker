FROM golang:1.9.2

RUN git clone https://github.com/evelynocean/hello-docker.git /go/src/hello
#RUN mkdir -p /go/src/hello
#COPY ./ /go/src/hello

RUN cd /go/src/hello && go build

WORKDIR /go/src/hello

EXPOSE 1031

ENTRYPOINT ["./hello"]

# docker build -t hello -f Dockerfile .
# docker run --name hello -d -p 127.0.0.1:1031:1031 hello