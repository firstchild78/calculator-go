FROM golang:latest as builder
WORKDIR /goCalculatorService
COPY . .
COPY ./vendor/github.com/gorilla/mux $GOPATH/src/github.com/gorilla/mux
COPY ./vendor/github.com/stretchr/testify $GOPATH/src/github.com/stretchr/testify
RUN go test -v -cover
RUN CGO_ENABLED=0 go build

FROM alpine:latest
WORKDIR /app
COPY --from=builder goCalculatorService .
CMD ["./goCalculatorService"]