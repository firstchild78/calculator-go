FROM golang:latest as builder
WORKDIR /goCalculatorWeb
COPY main.go .
RUN CGO_ENABLED=0 go build

FROM alpine:latest
WORKDIR /app
COPY --from=builder goCalculatorWeb .
COPY index.html .
CMD ["./goCalculatorWeb"]



