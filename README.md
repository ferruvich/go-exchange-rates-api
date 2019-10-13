# Exchange Rates API

## Project Description

The application should:

- Have an endpoint that calls [this api](https://exchangeratesapi.io/) to get the latest exchange rates for the base currencies of GBP and USD;
- It should return the value of 1 GBP or 1 USD in euros;
- It should check that value against the historic rate for the last week and make a naive recommendation as to whether this is a good time to exchange money or not.

## What is missing

- Authentication;
- Logging, maybe with a logger such as [uber-go/zap](https://github.com/uber-go/zap) or [sirupsen/logrus](https://github.com/sirupsen/logrus);
- Metrics, written using [Prometheus client](https://github.com/prometheus/client_golang), collected using [Prometheus server](https://github.com/prometheus/prometheus) and visualized using a system such as [Grafana](https://github.com/grafana/grafana) where to show these metrics collected in the form of graphs, and where to setup alerts.

## Technologies

- [Go 1.13](https://golang.org/) with [go mod](https://github.com/golang/go/wiki/Modules) to manage dependencies;
- [gin-gonic/gin](https://github.com/gin-gonic/gin) as HTTP web framework;
- [golang/mock](https://github.com/golang/mock) to mock interfaces, for Unit testing;
- [swaggo/swag](https://github.com/swaggo/swag) to generate swagger files for documentation;
- [docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/) to deploy the project
- [make](https://www.gnu.org/software/make/manual/make.html), using Makefile to simplify commands.

## How to run the project

### What do you need

- Go 1.13;
- Docker and docker-compose;
- `make` to run Makefile commands.

### How to do it

- Install [docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/)
- Make sure that your `docker.service` is active;
- The application will be served in port `8000`. Please make sure that this port is available;
- Make sure your are in the project's root folder;
- Run `make run`.

## API Documentation

### Swagger

When project is running, you can go to `localhost:8000/swagger/index.html` or simply to `localhost:8000` to display the API documentation.

### How to perform API Calls

You can test API calls on Swagger documentation page, in `localhost:8000`. Otherwise, a [Postman](https://www.getpostman.com/) collection has been created in order to test them. It is available [here](https://github.com/ferruvich/go-exchange-rates-api/blob/master/api/postman_collection.json).
To use it, download it, open Postman, click on **import** and follow the instructions.
