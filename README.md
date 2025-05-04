# HTTP from TCP

This repo contains the project files for HTTP from TCP course on bootdev.
[https://www.boot.dev/courses/learn-http-protocol-golang](https://www.boot.dev/courses/learn-http-protocol-golang)

## The need to create docker setup

this course requires netcat to test tcp and udp listeners. Unfortunately, the gnu-netcat is not behaving as expected. So, it's better to use openbsd-netcat.

### Notes:

1. the docker image copies your bootdev cli config into the container (So no need to login bootdev cli inside the container). Change the source config path if neccessary

2. All go files are now in src directory

## Running the project

- Run the project using the following command:
```
docker-compose up -w
```
This will start the project in a docker container with watch mode enabled.
- Exec into the container using the following command:
```
docker-compose exec -it go sh
```
Now you can run `go run ..` commands and netcat commands.
