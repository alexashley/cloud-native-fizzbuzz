# mod 3

This is a cutting edge microservice written in Go that acts as an oracle for congruence relations to the prime number 3.

## development

To build and run the service, use `make run`. A test suite is included, run it with `make test`. 
Before submitting a PR, please add tests and run `go fmt`.

## use

Currently, a single API endpoint is exposed for the oracle.

```bash
$ curl -s http://localhost:8080/api/v1/math/mod/3 -d '{"value": 5}'
```

```json
{
  "divisibleByThree": false
}
```