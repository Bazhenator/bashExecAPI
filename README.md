[![Build](https://github.com/Bazhenator/bashExecAPI/actions/workflows/go.yaml/badge.svg)](https://github.com/Bazhenator/bashExecAPI/actions/workflows/go.yaml)
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
    - [Docker deployment](#docker-deployment)
    - [Run unit test and update coverage bage](#run-unit-test-and-update-coverage-bage)
    - [Build project](#build-project)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


# BashExecAPI

## About the project

This projects contains service that provides API to run bash scrips.

### API docs

Project has configured **swagger API documentation**, that can be accessed by endpoint `GET /swagger/`

  

### Status

The project is in ready for deployment.

## Getting started

### Download project

To download use:
```bash
git clone https://github.com/Bazhenator/bashExecAPI.git
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
make build-backend
```
or
```bash
make build-connector
```
for every node.
Executables will be in **/bin** folder
