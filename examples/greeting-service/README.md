# Greeting Service

Greeting Service is a complete end-to-end client/server example, demonstrating how to use the Edelweiss compiler.

## Build and run

Build and run the server:
```
cd server
go build
./server
```

Build and run the client:
```
cd client
go build
./client
```

## Source directory

### API definition and generation

The service API is defined in `api/gen-greeting-api.go`. This is a Go program which generates the API client-server code for Go. For your convenience, the generated code is already provided in package `api/proto`. To generate the code yourself:

```
cd api
rm -rf proto
go build
./api
```

### Service implementation

A specific implementation of the Greeting Service is provided in package `service`.
### Command-line binaries

Subdirectories `client` and `server` hold the command-line client and server programs for the Greeting Service.

### Unit tests

An example of how one would write a client-server unit test for the Greeting Service is found in `clientserver_test.go`.
