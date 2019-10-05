# itoa

A microservice to handle natural language processing. The present implementation is limited to string to integer conversions.

## development

To build and run the service, use `make run`. A test suite is included, run it with `make test`. 
Before submitting a PR, please add tests and run `go fmt`.

## use

```bash
$ curl -s http://localhost:8080/api/v1/str/itoa -d '{"integer": 123}'
```

```json
{
  "string": "123"
}
```