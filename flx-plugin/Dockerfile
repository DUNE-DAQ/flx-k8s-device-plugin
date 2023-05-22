FROM golang:1.17

WORKDIR /usr/src/app

ARG HTTP_PROXY=local
ARG HTTPS_PROXY=local
ENV HTTP_PROXY=${HTTP_PROXY}
ENV HTTPS_PROXY=${HTTPS_PROXY}


# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go ./
RUN go build -v -o /usr/local/bin/app ./main.go

CMD ["app"]
