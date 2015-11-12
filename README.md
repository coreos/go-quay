# go-quay
Go-swagger generated client bindings for quay.io

# Generating Client Bindings

First install [go-swagger](https://github.com/go-swagger/go-swagger):

```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

Next run the script included in the repo to generate client bindings:

```
./generate_client
```

## Stability

This package is not considered stable, and breaking changes may occur. Go-swagger is still very early in it's support for generating client code from a schema, and may have bugs.
Additionally, the quay.io API is still in preview and does not make any guarantees to not make breaking changes.

Use at your own risk.
