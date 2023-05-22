FROM golang:1.17 AS builder

WORKDIR /usr/src/app

ARG HTTP_PROXY=local
ARG HTTPS_PROXY=local
ENV HTTP_PROXY=${HTTP_PROXY}
ENV HTTPS_PROXY=${HTTPS_PROXY}


# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go ./
ENV CGO_ENABLED=0 
RUN go build -v -o /usr/local/bin/app ./main.go

FROM alpine:3.18.0 as baseImage
COPY --from=builder /usr/local/bin/app /usr/local/bin/app

CMD ["app"]
