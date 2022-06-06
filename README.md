## Description
Sample transaction API

## Tecnologies
Golang<br/>
Docker<br/>
Mysql<br/>
Gorilla Mux<br/>

## Setup

1. (Optional) Install Go, 1.16 version
2. (Mandatory) Install docker and docker-compose
3. (Optional) Install tools: </br>
      go get -u github.com/swaggo/swag/cmd/swag;
4. (Optional) go mod tidy

**Optional tools means that its only necessary if you are going to develop or run using IDE

## Run using docker

  ```
  #change env properties mysql to "host.docker.internal" (in case running docker in macOS) or "127.0.0.1"
  docker-compose up -d --build
  http://localhost:8081/swagger/index.html
  ```

### Throubleshoting
In case "swag command not found" ```export PATH=$PATH:$HOME/go/bin```<br/>