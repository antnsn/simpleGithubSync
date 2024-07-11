
# Simple GitHub Sync

This project provides a Docker container that monitors multiple directories for changes and syncs them with a remote Git repository. The container periodically checks for changes in the local directories and the remote repository, ensuring both are in sync.

## Features

- Monitors multiple directories for changes.
- Adds, commits, and pushes changes to a remote Git repository.
- Pulls changes from the remote repository if there are updates.
- Configurable via environment variables.
- Automatically uses all mounted volumes as Git repository paths.

## Prerequisites

- Docker installed on your system.
- A Git repository to sync with.
- SSH key for authentication with the Git repository.

## Getting Started

### Using the Pre-built Image

The Docker image for this project is available on GitHub Container Registry (GHCR) with the `latest` tag.

### Running with `docker run`

Run the Docker container with the following environment variables:

- `SSH_KEY`: Your private SSH key for authentication with the Git repository.

Example command:

```sh
docker run -d \
    -e SSH_KEY="$(cat /path/to/private_key)" \
    -v /path/to/dir1:/repos/dir1 \
    -v /path/to/dir2:/repos/dir2 \
    ghcr.io/USERNAME/simplegithubsync:latest
```

### Running with `docker-compose`

Create a `docker-compose.yml` file in your project directory:

```yaml
version: '3.8'

services:
  simple-github-sync:
    image: ghcr.io/USERNAME/simplegithubsync:latest
    environment:
      - SSH_KEY=${SSH_KEY}
    volumes:
      - path/to/dir1:/repos/dir1
      - /path/to/dir2:/repos/dir2
```

Create a `.env` file in the same directory to securely pass your SSH key:

```dotenv
SSH_KEY=$(cat /path/to/private_key)
```

Run the Docker Compose setup:

```sh
docker-compose up -d
```

### Environment Variables

- `SSH_KEY`: The private SSH key for accessing your Git repository. This key should be provided as a single string.

### How It Works

1. The entrypoint script sets up the SSH key for authentication.
2. The script lists all directories mounted to `/repos` and uses them as Git repository paths.
3. If changes are detected locally, it adds, commits, and pushes the changes to the remote repository.
4. If changes are detected in the remote repository, it pulls the updates to the local directories.
5. The script runs in an infinite loop, checking for changes every 60 seconds.

### Troubleshooting

- Ensure the SSH key has the correct permissions:

  ```sh
  chmod 600 /path/to/your/private/key
  ```

- Verify that the directories mounted to `/repos` exist and are accessible.
- Check the logs of the Docker container for any error messages:

  ```sh
  docker logs <container_id>
  ```

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

