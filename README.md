[![Build](https://github.com/Bazhenator/bashExecAPI/actions/workflows/build.yaml/badge.svg)](https://github.com/Bazhenator/bashExecAPI/actions/workflows/build.yaml)
![Coverage](https://img.shields.io/badge/Coverage-84.5%25-green)
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
<!--   *generated with [DocToc](https://github.com/thlorenz/doctoc)* -->

- [BashExecAPI](#bashExecAPI)
  - [About the project](#about-the-project)
    - [API docs](#api-docs)
    - [Status](#status)
  - [Getting started](#getting-started)
    - [Download project](#download-project)
    - [Manual deployment](#manual-deployment)
    - [Docker deployment](#docker-deployment)
    - [Run unit test and update coverage bage](#run-unit-test-and-update-coverage-bage)
    - [Build project](#build-project)
    - [Update swagger documentation](#update-swagger-documentation)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


# BashExecAPI

## About the project

This project contains service that provides API to run bash scrips.

### API docs

Project has configured **swagger API documentation**, that can be accessed by endpoint `GET /swagger/`
<details>
  <summary>API screenshots</summary>
  <img src="https://github.com/Bazhenator/bashExecAPI/assets/113100166/080f13d3-ec6e-4bf9-87c1-4e940e711186">
</details>

  

### Status

The project is in ready for deployment.

## Getting started

### Download project

To download use:
```bash
git clone https://github.com/Bazhenator/bashExecAPI.git
```

### Manual deployment
> [!WARNING]
> You should use _Manual deployment_ if you want to run application without Docker/make etc.


> [!IMPORTANT]
> Follow this instruction to install **PostgreSQL** on (Linux).
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```
> Check PostgreSQL status:
```bash
service postgresql status
```
> [!IMPORTANT]
> To install pgAdmin4 (web/desktop) version on Linux follow [this](https://www.pgadmin.org/download/pgadmin-4-apt/) link:


> Create new server **bash** with database, user and password specified in configs/config.yaml


> Create new table 'commands', or put schema from init.sql in query of bash_db.


> [!WARNING]
> If your server's **host** or **port** differs from data in ./configs/config.yaml, please, replace it with actual. (postgres:5432 -> localhost:5432)
 

 
> [!IMPORTANT]
> To run application:
```bash
cd cmd
go run server.go
```
### Docker deployment

You can use docker deployment, to prepare docker images use:
> [!IMPORTANT]
> You should have installed **Docker** and **docker-compose** on your machine


> [!CAUTION]
> This will prune all unused images in the end of build to free up space after multistage Docker image build
```bash
make build-images
```

Then you can simply run
```bash
make start-dev
```
to run docker container using docker-compose, and

```bash
make stop-dev
```
to stop docker container

### Run unit tests and update coverage bage

This will **run unit tests** and **update link** for coverage bage in README
```bash
make unit-test
```
You can regenerate mocks with
```bash
make gogen
```

### Build project

You can build executable files with
```bash
make build-bash
```
Executables will be in **/bin** folder

### Update swagger documentation

> [!IMPORTANT]
> Project uses [swaggo](https://github.com/swaggo/swag), so you should install it

To update swagger documentation after adding new endpoints use:
```bash
make swag-bash
```
