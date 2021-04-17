# Blackbox HelloWorld Responder

[![Go Report Card](https://goreportcard.com/badge/github.com/KohlsTechnology/blackbox-helloworld-responder)](https://goreportcard.com/report/github.com/KohlsTechnology/blackbox-helloworld-responder)
[![codecov](https://codecov.io/gh/KohlsTechnology/blackbox-helloworld-responder/branch/master/graph/badge.svg)](https://codecov.io/gh/KohlsTechnology/blackbox-helloworld-responder) [![Join the chat at https://gitter.im/KohlsTechnology/helloworld-responder](https://badges.gitter.im/KohlsTechnology/helloworld-responder.svg)](https://gitter.im/KohlsTechnology/helloworld-responder?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

The blackbox-helloworld-responder's sole purpose is to respond with "Hello World!" on either a HTTP or TCP connection. It is intended to work as a simple responder for blackbox or synthetic transaction monitoring, using tools like the [Prometheus Blackbox Exporter](https://github.com/prometheus/blackbox_exporter). It purposely is written very simple, so that it can run securely and with minimal resources.

## License

See [LICENSE](LICENSE) for details.

## Code of Conduct

See [CODE_OF_CONDUCT.md](.github/CODE_OF_CONDUCT.md)
for details.

# Execution

The executable supports overwriting of the ports via environment variables and will log all connections to stdout.

## Settings

| Protocol | Default Port | Environment Variable |
| --- | --- | --- |
| HTTP | 8080 | HELLO_WORLD_HTTP_PORT |
| TCP | 8081 | HELLO_WORLD_TCP_PORT |

## Sample process output
```
./blackbox-helloworld-responder

2020/07/28 13:57:39 Starting HTTP Server on 8080
2020/07/28 13:57:39 Starting TCP Server on 8081
2020/07/28 13:57:46 HTTP 127.0.0.1:60513
2020/07/28 13:57:50 TCP 127.0.0.1:60514
```

## Testing
```
# HTTP
curl -v 127.0.0.1:8080

# TCP
nc 127.0.0.1 8081
```

# Development

## Binary

If you just need a local version to test, then the simplest way is to execute:

```
make build
```

## Image

In order to build the docker image, simply execute:

```
make image
```

## Testing

### Unit Test
```
make test
```

## Releasing

This project is using [goreleaser](https://goreleaser.com). GitHub release creation is automated using GitHub
Actions. New releases are automatically created when new tags are pushed to the repo.
```
$ TAG=v0.0.2 make tag
```

How to manually create a release without relying on GitHub Actions.
```
$ TAG=v0.0.2 make tag
$ GITHUB_TOKEN=xxx make clean release
```

### Test Release

```
make test-release
```
