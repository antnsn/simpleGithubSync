# Simple GitHub Sync

This project provides a simple solution to sync multiple Git repositories using Docker. It periodically fetches, pulls, commits, and pushes changes to ensure all repositories are up-to-date.

## Usage

To use this Docker container, follow the steps below:

### Prerequisites

- Docker installed on your system.
- SSH keys set up for GitHub access.
- `.gitconfig` file configured with your Git user name and email.

### Docker Compose Configuration

Here is an example of a `docker-compose.yml` configuration:

```yaml
version: "3.9"
services:
  simplegithubsync:
    container_name: githubSync
    environment:
      - TZ=Europe/Oslo
    volumes:
      - /path/to/your/.ssh:/root/.ssh
      - /path/to/your/home:/root
      - /path/to/repo/repo1:/repos/repo1
      - /path/to/repo/repo2:/repos/repo2
    image: ghcr.io/antnsn/simplegithubsync:v1.0.3
networks: {}
```

### Docker Run Command

If you prefer to use `docker run` instead of Docker Compose, you can use the following command:

```sh
docker run -d \
  --name githubSync \
  -e TZ=Europe/Oslo \
  -v /path/to/your/.ssh:/root/.ssh \
  -v /path/to/your/home:/root \
  -v /path/to/repo/repo1:/repos/repo1 \
  -v /path/to/repo/repo2:/repos/repo2 \
  ghcr.io/antnsn/simplegithubsync:v1.0.3
```

### Configuration

1. **SSH Keys**: Mount your SSH keys directory to `/root/.ssh` inside the container. Ensure your SSH keys have the correct permissions.

   ```yaml
   - /path/to/your/.ssh:/root/.ssh
   ```

2. **Git Configuration**: Mount your home directory containing the `.gitconfig` file to `/root` inside the container.

   ```yaml
   - /path/to/your/home:/root
   ```

3. **Repositories**: Mount each of your Git repositories to `/repos` inside the container. Each repository should be in its own directory.

   ```yaml
   - /path/to/repo/repo1:/repos/repo1
   - /path/to/repo/repo2:/repos/repo2
   ```

### Running the Container

To run the container using Docker Compose, navigate to the directory containing your `docker-compose.yml` file and execute:

```sh
docker-compose up -d
```

This will start the `simplegithubsync` container in detached mode. The container will periodically sync your repositories according to the script logic.

To run the container using `docker run`, execute the command provided in the Docker Run Command section.

### Environment Variables

- `TZ`: Set the timezone for the container (e.g., `Europe/Oslo`).

### Logging

Logs for the synchronization process will be available in the Docker container's logs. You can view them using:

```sh
docker logs -f githubSync
```

### Contributing

Feel free to open issues or submit pull requests if you have suggestions for improvements or bug fixes.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
