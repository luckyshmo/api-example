FROM golang:1.16.2-alpine as builder

# RUN go version
ENV GOPATH=/

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go build -o api-example ./cmd/main.go

#Build destination container
FROM alpine:latest

# install psql
RUN apk --update add postgresql-client

# copy bin and pg-wait script
COPY --from=builder /go/api-example /go/wait-for-postgres.sh ./

# copy PG migrations
COPY --from=builder /go/pkg/repository/pg/migrations/*.sql ./migrations/

EXPOSE 8080
