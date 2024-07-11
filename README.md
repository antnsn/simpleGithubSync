# Git Sync Docker Container

This project provides a Docker container that monitors multiple directories for changes, and syncs them with a remote Git repository. The container periodically checks for changes in the local directories and the remote repository, ensuring both are in sync.

## Features

- Monitors multiple directories for changes.
- Adds, commits, and pushes changes to a remote Git repository.
- Pulls changes from the remote repository if there are updates.
- Configurable via environment variables.

## Getting Started

### Prerequisites

- Docker installed on your system.
- A Git repository to sync with.
- SSH key for authentication with the Git repository.

### Building the Docker Image

1. Clone this repository or copy the Dockerfile and entrypoint.sh to your working directory.
2. Build the Docker image:

```sh
docker build -t git-sync .
```

### Running the Docker Container

Run the Docker container with the following environment variables:

- `FOLDER_PATHS`: Comma-separated list of directory paths to monitor.
- `SSH_KEY`: Your private SSH key for authentication with the Git repository.

Example command:

```sh
docker run -d \
    -e FOLDER_PATHS="/path/to/dir1,/path/to/dir2,/path/to/dir3" \
    -e SSH_KEY="$(cat /path/to/your/private/key)" \
    git-sync
```

## Environment Variables

- `FOLDER_PATHS`: Comma-separated list of absolute paths to directories that need to be monitored and synced.
- `SSH_KEY`: The private SSH key for accessing your Git repository. This key should be provided as a single string.

## How It Works

1. The `entrypoint.sh` script sets up the SSH key for authentication.
2. The script monitors the specified directories for changes.
3. If changes are detected locally, it adds, commits, and pushes the changes to the remote repository.
4. If changes are detected in the remote repository, it pulls the updates to the local directories.
5. The script runs in an infinite loop, checking for changes every 60 seconds.

## Customization

You can adjust the interval for checking changes by modifying the `sleep` duration in the `entrypoint.sh` script.

```sh
sleep 60  # Adjust the interval as needed
```

## Troubleshooting

- Ensure the SSH key has the correct permissions:

```sh
chmod 600 /path/to/your/private/key
```

- Verify that the directories specified in `FOLDER_PATHS` exist and are accessible.
- Check the logs of the Docker container for any error messages:

```sh
docker logs <container_id>
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

