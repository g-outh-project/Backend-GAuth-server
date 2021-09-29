
FROM golang:1.16-alpine AS builder

WORKDIR /build

COPY . . 
RUN go mod download
ENV CGO_ENABLED=0
RUN go build -o main


WORKDIR /dist

RUN cp /build/main .
RUN cp /build/.env .

FROM alpine

WORKDIR /app

COPY --from=builder /dist/main .
COPY --from=builder /dist/.env .

EXPOSE 8080
ENTRYPOINT ["./main"]