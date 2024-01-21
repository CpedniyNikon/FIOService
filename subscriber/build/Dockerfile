FROM golang:1.21.5

WORKDIR /app

COPY ../.. .

RUN go mod download

RUN go env -w GO111MODULE=on
RUN go build -o subscriber ./cmd

EXPOSE 8080


CMD [ "./subscriber" ]