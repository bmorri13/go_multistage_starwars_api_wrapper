# Go Multistage Container Build for Star Wars API

This repository contains a Go application that serves data from the Star Wars API (SWAPI). It demonstrates the use of a multistage Docker container build to create a lightweight, secure final image.

## Prerequisites

- Docker installed on your machine. Visit [Docker's official installation guide](https://docs.docker.com/get-docker/) for instructions.

## Getting Started

Follow these steps to build and run the Star Wars API Go application inside a Docker container.

### Building the Docker Image

To build the Docker image, use the following command in the root directory of this project:

```sh
docker build -t go_multistage_starwars_api .
```

This command builds a Docker image named go_multistage_starwars_api using the Dockerfile located in the current directory.

### Running the Container
```sh
docker run -p 5002:5002 go_multistage_starwars_api
```
This command starts a container based on the go_multistage_starwars_api image. It maps port 5002 of the container to port 5002 on the host, allowing you to access the Go application by navigating to http://localhost:5002 in your web browser.


### API Endpoints
The application provides the following endpoints:

- '/ships': Fetches starship data from SWAPI. Supports a search query parameter for filtering.
- '/characters': Fetches character data from SWAPI. Also supports a search query parameter for filtering.
- '/': The home endpoint, listing all available APIs.

#### API Endpoints Search
- You can use the 'search' query parameter to filter results based on the name of the starship or character. 
```sh
http://localhost:5002/characters?luke
```


For detailed usage, refer to the API documentation section (if applicable).

### Additional Notes

- If you encounter any port conflicts, ensure the specified port is not already in use or adjust the docker run command to map to a different port.
