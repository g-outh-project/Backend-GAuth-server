
FROM golang:1.16-alpine AS builder

WORKDIR /build

COPY . . 
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .


WORKDIR /dist

RUN cp /build/main .
RUN cp /build/.env .

FROM scratch

COPY --from=builder /dist/main .
COPY --from=builder /dist/.env .

EXPOSE 4500
ENTRYPOINT ["sh", "./main"]