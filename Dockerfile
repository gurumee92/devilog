FROM golang:1.14-alpine AS builder
 
WORKDIR /app
COPY . .

RUN apk update && apk add git && apk add ca-certificates && apk add gcc
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/public/ /app/public/
COPY --from=builder /app/main /app/main

EXPOSE 8080

ARG APPLICATION_PATH
ARG DATABASE_DIALECT
ARG DATABASE_URL

CMD [ "/app/main" ]
