FROM golang

ENV SRC_DIR=/go/src/testapp/

ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o testapp

ENTRYPOINT ["./testapp"]

EXPOSE 8086