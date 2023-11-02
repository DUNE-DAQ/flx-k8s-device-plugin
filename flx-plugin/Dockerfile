FROM golang:1.21 AS builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go ./
ENV CGO_ENABLED=0 
RUN go build -v -o /usr/local/bin/app ./main.go

FROM alpine:3.18.0 as baseImage
COPY --from=builder /usr/local/bin/app /usr/local/bin/app

CMD ["app"]
