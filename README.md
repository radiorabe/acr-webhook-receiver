# ACRCloud Webhook Receiver

Spec and a server for receiving ACRCloud events and storing them to PostgreSQL.

* [spec](./swagger.yaml)

The received data is validated and stored in a JSONB field. There is a simple
API to query the database.

## Requirements

* PostgreSQL server
* ACRCloud account with Broadcast Monitoring subscription

## Usage

```
Usage:
  acr-webhooks-server [OPTIONS]

Hooks to receive ACRCloud events on and a simple API to query the received data.
Stores all data in a PostgreSQL database for later querying.


Application Options:
      --scheme=            the listeners to enable, this can be repeated and defaults to the schemes in the swagger spec
      --cleanup-timeout=   grace period for which to wait before killing idle connections (default: 10s)
      --graceful-timeout=  grace period for which to wait before shutting down the server (default: 15s)
      --max-header-size=   controls the maximum number of bytes the server will read parsing the request header's keys and values, including the request line. It does not limit the size of the request body. (default: 1MiB)
      --socket-path=       the unix socket to listen on (default: /var/run/acr-webhooks.sock)
      --host=              the IP to listen on (default: localhost) [$HOST]
      --port=              the port to listen on for insecure connections, defaults to a random value [$PORT]
      --listen-limit=      limit the number of outstanding requests
      --keep-alive=        sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download) (default: 3m)
      --read-timeout=      maximum duration before timing out read of the request (default: 30s)
      --write-timeout=     maximum duration before timing out write of the response (default: 60s)
      --tls-host=          the IP to listen on for tls, when not specified it's the same as --host [$TLS_HOST]
      --tls-port=          the port to listen on for secure connections, defaults to a random value [$TLS_PORT]
      --tls-certificate=   the certificate to use for secure connections [$TLS_CERTIFICATE]
      --tls-key=           the private key to use for secure connections [$TLS_PRIVATE_KEY]
      --tls-ca=            the certificate authority file to be used with mutual tls auth [$TLS_CA_CERTIFICATE]
      --tls-listen-limit=  limit the number of outstanding requests
      --tls-keep-alive=    sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download)
      --tls-read-timeout=  maximum duration before timing out read of the request
      --tls-write-timeout= maximum duration before timing out write of the response

Webhook Settings:
      --api-key=           API key transmitted by ACRCloud with each request [$ACR_API_KEY]

Database Settings:
      --postgres-dsn=      PostgreSQL DSN for gorm.io/driver/postgresql [$POSTGRES_DSN]

Help Options:
  -h, --help               Show this help message
```

Once the server is running and accessible to ACRCloud you can configure the results webhook in ACRCloud:

1. Navigate to your Broadcast Monitoring Stream
2. Select it by clicking it and choose "Set Result Callback" from "Actions"
3. Use `https://exmaple.org/v1/_webhooks/results?api_key=YOURAPIKEY` as "Result Callback URL"", substitute YOURAPIKEY with what you passed in the `--api-key` parameter when you started the server
4. Check if you are receiving data on the `https://example.org/v1/results` endpoint.

## Development

```bash
# clone the repo
git clone https://github.com/radiorabe/acr-webhook-receiver.git
cd acr-webhook-receiver

# run the command line locally
go run cmd/acr-webhooks-server/main.go --help

# build a binary
go build cmd/acr-webhooks-server/main.go -o acr-webhooks-server
```

### pre-commit hook

#### pre-commit configuration

```bash
pre-commit install
pre-commit install --hook-type commit-msg
pre-commit install-hooks
```

### Release Process

Create a git tag and push it to this repo or use the git web ui.

This is build on GitHub Actions and uses a `GH_PAT_TOKEN` secret to work. The access key must
have repo, read:packages, write:packages and delete:packages in it's scope.

## License
This software is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, version 3 of the License.

## Copyright
Copyright (c) 2020 [Radio Bern RaBe](http://www.rabe.ch)
