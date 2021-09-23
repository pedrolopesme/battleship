<h1 align="center">
  <br>
  <img src="https://github.com/pedrolopesme/battleship/blob/master/web/img/battleship_logo.png?raw=true" alt="Battleship" width="600">
  <br>
  Battleshiip
  <br>
  <br>
</h1>

<h4 align="center">A Battleship Game Server.</h4>
Battleship is a very simple multiplayer gaming server over HTTP 2 and Websocket

### TODO
In order to release `0.1`, Battleship Server needs to:

- [ ] Create an API to register a new Player
- [ ] Create an API to create a new Match
- [ ] Create a lobby to make players to wait until a match is ready.
- [ ] Separate players into matches of X configurable players.
- [ ] Create two Docker files: one with hot deploying for local development and other to run a stable version.
- [ ] Define Match structure containing Players and Events
- [ ] Create a throttling engine based on Time Frame to avoid API abuses
- [ ] Limit a Match by a configurable time and make an API endpoint to expose Match Results
- [ ] Create and API to register events in a given Match
- [ ] Make sure that the source code is well covered with tests (80%?)
- [ ] Use tools from [CNCF](https://www.cncf.io/) stack
    - [ ] Make a Kubernetes POD Configuration Yaml with all the stack needed to run and monitor the server 
    - [ ] Add a Gateway to protect the API by validating a JWT header

### Makefile

This project provides a Makefile with all common operations need to develop, test and build battleship-server.

* build: generates binaries
* test: runs all tests
* clean: removes binaries
* run: executes main func
* fmt: runs gofmt for all go files

### Running tests

Tests were write using [Testify](https://github.com/stretchr/testify). In order to run them, just type:

```shell
$ make test
```

### Contributing

 [CONTRIBUTING.md](CONTRIBUTING.md) 

### License

[Apache License 2.0](LICENSE)  