FROM golang:latest as builder

WORKDIR /go/src/testapp

#RUN go get -d -v -u github.com/aws/aws-sdk-go

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o testapp .


FROM scratch   

WORKDIR /testapp/

COPY --from=builder /go/src/testapp/testapp .

ENTRYPOINT ["./testapp"]

EXPOSE 8086