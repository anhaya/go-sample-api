FROM golang

LABEL maintainer="Carlos Anhaya <carlos.anhaya@gmail.com>"

WORKDIR /app/src/go-sample-api

ENV GOPATH=/app

COPY . /app/src/go-sample-api/

RUN go build main.go

# Run builded app
ENTRYPOINT ["./main"]

EXPOSE 8081