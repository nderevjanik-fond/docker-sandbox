# Docker Sandbox
A sandbox for experimenting with docker.

## Prerequisites

* [Docker Desktop](https://www.docker.com/products/docker-desktop/)

## Usage

### Development / Debug

1. Start the application in debug mode (note the use of a `debug` docker-compose file)
1. Identify a ***component*** (golang-service, node-service, react-web-app, etc) that you'd like to debug
1. Add breakpoints in VSCode
1. Launch "Attach to ***component***" (from VSCode's `Run and Debug` view)
1. Use the application so that it triggers the breakpoints

```sh
# Start and use the application in debug mode
> docker compose -f "docker-compose.debug.yml" up -d --build
> curl localhost:3000

# Stop the application
> docker compose -f "docker-compose.debug.yml" down
```

### Production

```sh
# Start and use the application
> docker compose -f "docker-compose.yml" up -d --build
> curl localhost:3000

# Stop the application
> docker compose -f "docker-compose.yml" down
```
