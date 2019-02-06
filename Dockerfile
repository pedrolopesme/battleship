FROM golang:1.11.5

WORKDIR /go/src/github.com/pedrolopesme/citta-server

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./" -command="./citta-server"