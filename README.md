# Docker Sandbox
A sandbox for experimenting with docker.

## Prerequisites

* [Docker Desktop](https://www.docker.com/products/docker-desktop/)

## Usage

### Development / Debug

VSCode debugging is built-in:
1. Start the Node Web App in debug mode (note the use of a `debug` docker-compose file)
1. Add breakpoint(s) in VSCode
1. Launch "Docker: Attach to Node Web App" (from VSCode's `Run and Debug` view)
1. Use the Node Web App so that it triggers the breakpoint(s)

```sh
# Start and use the Node Web App in debug mode
> docker compose -f "docker-compose.debug.yml" up -d --build
> curl localhost:3000

# Stop the Node Web App
> docker compose -f "docker-compose.debug.yml" down
```

### Production

```sh
# Start and use the Node Web App
> docker compose -f "docker-compose.yml" up -d --build
> curl localhost:3000

# Stop the Node Web App
> docker compose -f "docker-compose.yml" down
```
