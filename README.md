# ko with go

Learn how to use ko with golang.

## packages

When `GOPATH` points additionally to project root, then all dirs under `src` are packages. Package `app` contains `main()` method.

```bash
export GOPATH=$GOPATH:$(pwd)
```

## go basics

```bash
# run the app package
go run app
# build binary of the app package
go build -o bin/app app
bin/app
```

## containerizing

```bash
# build local docker image without hashed import path and tag v2 using current date
export SOURCE_DATE_EPOCH=$(date +%s) 
ko publish app -L -P -t v2
# push to dockerhub with hashed import path
docker login
export KO_DOCKER_REPO=docker.io/alitari
ko publish app -t v2
```

## k8s

```bash
#create deployment from local import path
ko resolve -f deployment_local.yaml -t v2
```