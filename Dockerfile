
FROM golang:1.16-alpine AS builder

WORKDIR /build

COPY . . 
RUN go mod download
RUN go build -o main .


WORKDIR /dist

RUN cp /build/main .
RUN cp /build/.env .

FROM scratch

COPY --from=builder /dist/main .
COPY --from=builder /dist/.env .

EXPOSE 8080
ENTRYPOINT ["/main"]