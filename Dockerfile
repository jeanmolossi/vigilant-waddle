FROM golang:1.18.4-alpine3.16

RUN apk add --no-cache git

WORKDIR /src/go

COPY . .

ENV CGO_ENABLED=0

RUN go install github.com/cespare/reflex@latest
RUN go install golang.org/x/tools/cmd/godoc@latest

RUN go mod tidy
RUN go build -o vigillant-waddle

EXPOSE 8080

CMD [ "./vigillant-waddle" ]
