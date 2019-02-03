# citta-server
Citta Server is a very simple multiplayer gaming server over HTTP 2 and [Server Sent Events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)

### Makefile

This project provides a Makefile with all common operations need to develop, test and build call-it.

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