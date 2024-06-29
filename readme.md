## Init db
To initialize the db run the following command (change the appropriate db url properties) in the url shortener folder (requires [go migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate))
```shell
migrate -path=database/migrations -database "postgresql://postgres:postgres@localhost:5432/urlshortener?sslmode=disable" -verbose up
```
## Swagger url
[Swagger](http://localhost:8080/swagger/index.html)

## Sample Curls
cur to create short url:
```shell
curl -X POST -d "{\"longUrl\":\"http://www.google.com\"}" http://localhost:8080/urls
```
cur to get all urls:
```shell
curl http://localhost:8080/urls
```
cur to get the url visits:
```shell
curl http://localhost:8080/urls/XXX/visits
```
curl to redirect (replace XXX with the externalID value from the first curl)
```shell
curl -L  http://localhost:8081/XXX
```
## App config
Application configurations are on ***.env*** file