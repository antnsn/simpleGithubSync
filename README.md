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

The Docker image for this project is available on GitHub Container Registry (GHCR) with the `main` tag.

### Running with `docker run`

Run the Docker container with the following volumes:

- Mount the SSH key and known hosts file to the container.
- Mount the directories you want to monitor to `/repos`.

Example command:

```sh
docker run -d \
    -v /path/to/your/.ssh:/root/.ssh \
    -v /path/to/your/repo1:/repos/repo1 \
    ghcr.io/antnsn/simplegithubsync:main
```

### Running with `docker-compose`

Create a `docker-compose.yml` file in your project directory:

```yaml
version: "3.9"

services:
  simplegithubsync:
    container_name: githubSync
    volumes:
      - /path/to/your/.ssh:/root/.ssh
      - /path/to/your/repo1:/repos/repo1
      - /path/to/your/repo2:/repos/repo2
    image: ghcr.io/antnsn/simplegithubsync:main
```

Run the Docker Compose setup:

```sh
docker-compose up -d
```

### Environment Variables

This setup does not require environment variables for the SSH key because the key is mounted directly as a volume.

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
  docker logs githubSync
  ```

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

This `README.md` provides clear instructions for running the Docker container using the pre-built image from GHCR with the `main` tag and dynamically using all mounted volumes as Git repository paths, with generic paths for easier adaptation.