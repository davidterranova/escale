# Homework

## Questions
- each port seems to be identified by a code we also find in the `unlocs` property. What is this field ? Is it an identifier ? What woulld be the concept behinf a port having multiple unlocs ?
- 

## Current state
- 2 services :
  - store : acting as PortDomainService
  - client : acting as Client API
- API definitions are in the `api` folder
- Client is able to load (in stream mode) data from a JSON file and send them to store using gRPC communication
- Store saves received data in an in memory store
- REST API defined in openAPI format

## Improvements
- Implement the HTTP port and expose the search API defined in `api/http/port.yaml`
- Implement a repository able to store / fetch data to / from a real database (either SQL / NoSQL)
- Fix TODO in the project

## Requirements
- Go v1.16+
- docker / docker-compose
- make

## Run

### Lint / test / build / run

 ```sh
 make
 ```

 ### Lint

 ```sh
make lint
 ```

### Test

 ```sh
make test
 ```

### Build

```sh
make build
```
