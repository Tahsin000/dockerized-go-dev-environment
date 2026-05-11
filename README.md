# Dockerized Go Development Environment

A beginner-friendly Go/Golang development environment powered by Docker, Docker Compose, and Air live reload.

This project allows you to write and run Go code without installing Go directly on your local machine. The Go environment runs inside a Docker container, while your source code stays on your local machine.

## Features

- No need to install Go locally
- Go environment runs inside Docker
- Live reload support using Air
- Local source code is synced with the Docker container
- Docker Compose based workflow
- Port `8080` exposed for web server practice
- Suitable for Go learning, practice, and small projects

## Requirements

Before running this project, make sure you have the following installed:

- Docker Desktop
- Git
- Terminal or command line access

You do not need to install Go on your local machine.

## Project Structure

```text
.
├── Dockerfile
├── compose.yaml
├── .air.toml
├── go.mod
├── main.go
└── README.md
```

## Getting Started

### 1. Clone the repository

```bash
git clone <your-repository-url>
```

Go to the project directory:

```bash
cd <repository-folder-name>
```

Example:

```bash
cd dockerized-go-dev-environment
```

### 2. Make sure Docker is running

Check Docker version:

```bash
docker --version
docker compose version
```

Check whether the Docker daemon is running:

```bash
docker info
```

If Docker is not running on macOS, open Docker Desktop:

```bash
open -a Docker
```

### 3. Build and run the project

For the first run, build and start the container:

```bash
docker compose up --build
```

For later runs, you can use:

```bash
docker compose up
```

### 4. Open the application

Once the server is running, open your browser and visit:

```text
http://localhost:8080
```

## Live Reload

This project uses Air for live reload.

When you update and save any Go file, Air automatically rebuilds and restarts the application inside the Docker container.

You do not need to manually run:

```bash
go run .
```

Just keep Docker Compose running:

```bash
docker compose up
```

Then edit your Go files locally and refresh the browser.

## Common Commands

### Run the project

```bash
docker compose up
```

### Run the project with rebuild

```bash
docker compose up --build
```

### Stop the project

Press:

```text
Ctrl + C
```

Or stop and remove the running container:

```bash
docker compose down
```

### Run in detached mode

```bash
docker compose up -d
```

### View logs

```bash
docker compose logs -f
```

### Enter the container shell

```bash
docker compose run --rm go bash
```

### Format Go code

```bash
docker compose run --rm go go fmt ./...
```

### Run tests

```bash
docker compose run --rm go go test ./...
```

### Clean and tidy Go modules

```bash
docker compose run --rm go go mod tidy
```

## Go Module Setup

If the project does not already have a `go.mod` file, initialize a Go module using:

```bash
docker compose run --rm go go mod init example.com/golang-practice
```

Then tidy the module dependencies:

```bash
docker compose run --rm go go mod tidy
```

## How It Works

The local project directory is mounted into the Docker container at `/app`.

```yaml
volumes:
  - .:/app
```

This means any code change made on your local machine is immediately available inside the Docker container.

Air watches the project files inside the container and automatically rebuilds and restarts the Go application when changes are detected.

## Port Configuration

The project uses port `8080` by default.

The port mapping is defined in `compose.yaml`:

```yaml
ports:
  - "8080:8080"
```

The application will be available at:

```text
http://localhost:8080
```

## Troubleshooting

### Docker daemon is not running

If you see an error like:

```text
Cannot connect to the Docker daemon
```

Start Docker Desktop.

On macOS, you can run:

```bash
open -a Docker
```

Then verify Docker is running:

```bash
docker info
```

### Port 8080 is already in use

If port `8080` is already being used by another application, update the port mapping in `compose.yaml`:

```yaml
ports:
  - "8081:8080"
```

Then access the application at:

```text
http://localhost:8081
```

### Rebuild the container

If you change the `Dockerfile`, install new tools, or update dependency-related configuration, rebuild the container:

```bash
docker compose up --build
```

## Notes

- Go does not need to be installed locally.
- You can use any code editor such as VS Code, GoLand, Sublime Text, or Vim.
- Keep `docker compose up` running during development to enable live reload.
- Local code changes are synced into the container through Docker volume mounting.

## License

This project is created for learning and practice purposes.