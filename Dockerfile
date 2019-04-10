FROM golang:1.12

WORKDIR /go/src/app
COPY . .

RUN go get -d --v ./...
RUN go install -v ./...

RUN go build -o server
# RUN /usr/local/go/bin/go build -o app .

ENV PORT=3001

CMD ["./server"]
