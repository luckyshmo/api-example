FROM golang:1.16.2-buster

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go version
ENV GOPATH=/

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go build -o api-example ./cmd/main.go

CMD ["./api-example"]