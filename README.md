# booking-fly
Booking Fly is a example app created with Torpedo framework

## How to run this application

```shell
ENVIRONMENT=dev go run main.go
```

### How to run application tests

```shell
ENVIRONMENT=test go test -race ./...
```

## How to generate Swagger info
The swagger info is already generated, but if you add changes to the project and
would like to re-generate the swagger info, please, follow the steps below:

First `swag` command is required.
```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

After that the command to run is:
```shell
swag init --parseDependency --parseInternal -o oas
```

Please for further information read the section: [Open API](https://darksubmarine.com/docs/torpedo/advanced_rest_api_oas.html)
