# Standard Go Service Layout


## Overview

This layout is one of the possible ways to structure Go services on disk.

I came up with this layout while working on [memesis.app](https://memesis.app)
project, which consists of multiple Go services. The layout has proven to
work really well, as unification removes a lot of small, routine tasks and
allows to focus on the actual logic.

The layout is largely inspired by
[golang-standards/project-layout](https://github.com/golang-standards/project-layout),
the way Uber structures Go services, as well as ideas from Clean Architecture and MVC.


## Go Directories

### `/internal`

Private service and library code.

#### `/internal/handler`

Initial processing of incoming requests.

##### `/internal/handler/argsvalidator`

Argument validation.

#### `/internal/controller`

Service business logic.

##### `/internal/controller/hydrator`

Domain data enrichment.

#### `/internal/repository`

Database access logic.

#### `/internal/model`

Domain data structures.

#### `/internal/config`

Configuration deserialization.

#### `/internal/cron`

Scheduled task execution.

#### `/internal/workflow`

Long-running asynchronous workflows.

### `/cmd`

Service main function and utilities.

### `/api`

GRPC service contracts and client libraries in multiple languages for
communicating with the service. GRPC can be replaced with OpenAPI/Swagger
specifications, JSON Schema files, etc.

### `/config`

Service configuration files.

### `/deployments`

IaaS, PaaS, system and container orchestration deployment configurations and
templates (Dockerfile, docker-compose, kubernetes/helm, terraform).


## Demo

The repository contains a minimal buildable example of the inventory service,
which manages an inventory fleet. It can be built, tested, and deployed using
make. For this purpose, the following targets are implemented:

* fmt - run go fmt for service sources

* generate-proto - generate all protobuf code

* generate-mocks - generate mocks for unit tests

* build - compile the service binary

* test - run unit tests

* run-client-demo - run the demo client

* clean - clean build artifacts and temporary files

* images - build the Docker image for the service

* examine - start a shell inside the built Docker image

* up - run the service container

* down - stop and remove the running service container

* deploy - rebuild image and redeploy the container

* restart - restart the running service container

* ps - show status of the service container

* ssh - ssh inside the running container

* logs - show service logs
