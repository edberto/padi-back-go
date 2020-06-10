FROM golang:1.13

WORKDIR /go/src/padi-back-go
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["padi-back-go"]