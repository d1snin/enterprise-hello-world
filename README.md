# enterprise-hello-world
🚀 fast 🚀 robust 🚀 asynchronous🚀 lightweight🚀 production-ready🚀 cloud-native🚀 non-blocking 🚀 secure 🚀 hello world web server written in Go.

### Building
```sh
go build ./cmd -o enterprise-hello-world
```

### Running
```sh
./enterprise-hello-world
```
or
```sh
go run ./cmd
```

### Configuration
The following environment variables are supported:
| Environment variable | Description | Default value |
|----------------------|-------------|---------------|
| `HELLOWORLD_PORT`    | Server port.| `8080`        |
| `HELLOWORLD_MESSAGERECEIVER` | Whatever to use in the hello world message according to the following pattern: `Hello, {messagereceiver}!` | `World` |
| `HELLOWORLD_USERNAME` | Username to use for Basic auth. | `admin` |
| `HELLOWORLD_PASSWORD` | Password to use for Basic auth. | `admin` |

### Heartbeat
Heartbeat is located at `/ping`

### Why?
I made this to get more familiar with Go and for fun.
