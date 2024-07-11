
<p align="center">
  <img width=30%; src="./assets/logo.png">
</p>

# Simple GitHub Sync

A simple and efficient solution to synchronize multiple Git repositories using Docker. This container periodically fetches, pulls, commits, and pushes changes to ensure your repositories are always up-to-date.

## 📋 Prerequisites

Before you get started, ensure you have the following:

1. **Docker**: Install Docker on your system. Refer to the [Docker installation guide](https://docs.docker.com/get-docker/) for instructions.

2. **SSH Keys**: Generate SSH keys if you don't have them already. You can generate SSH keys using the following command:

    ```sh
    ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
    ```

    - Follow the prompts to save the key in the default location (`/home/your_user/.ssh/id_rsa`).
    - Add the SSH key to your GitHub account. Follow the instructions [here](https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh).

3. **Git Configuration**: Ensure you have a `.gitconfig` file in your home directory with the following details:

    ```ini
    [user]
        email = your_email@example.com
        name = Your Name
    ```

    This configuration allows Git to recognize your commits with the correct user information.

4. **Pre-existing Git Repositories**: This tool is designed to sync pre-existing Git repositories. Ensure your repositories are cloned and available on your local system.

## 🚀 Usage

To use this Docker container, follow the steps below:

### Docker Compose Configuration

Here's an example `docker-compose.yml` configuration:

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
    image: ghcr.io/antnsn/simplegithubsync:latest
networks: {}
```

### Docker Run Command

Alternatively, you can run the container using `docker run`:

```sh
docker run -d \
  --name githubSync \
  -e TZ=Europe/Oslo \
  -v /path/to/your/.ssh:/root/.ssh \
  -v /path/to/your/home:/root \
  -v /path/to/repo/repo1:/repos/repo1 \
  -v /path/to/repo/repo2:/repos/repo2 \
  ghcr.io/antnsn/simplegithubsync:latest
```

### Configuration Details

1. **SSH Keys**: Mount your SSH keys directory to `/root/.ssh` inside the container to ensure Git can authenticate with GitHub.

    ```yaml
    - /path/to/your/.ssh:/root/.ssh
    ```

2. **Git Configuration**: Mount your home directory containing the `.gitconfig` file to `/root` inside the container for Git user configuration.

    ```yaml
    - /path/to/your/home:/root
    ```

3. **Repositories**: Mount each of your Git repositories to `/repos` inside the container, ensuring each repository is in its own directory.

    ```yaml
    - /path/to/repo/repo1:/repos/repo1
    - /path/to/repo/repo2:/repos/repo2
    ```

### Running the Container

To run the container using Docker Compose, navigate to the directory containing your `docker-compose.yml` file and execute:

```sh
docker-compose up -d
```

To run the container using `docker run`, use the command provided above.

### 🌍 Environment Variables

- `TZ`: Set the timezone for the container (e.g., `Europe/Oslo`).

### 📖 Logging

Logs for the synchronization process will be available in the Docker container's logs. You can view them using:

```sh
docker logs -f githubSync
```

## 🤝 Contributing

Feel free to open issues or submit pull requests if you have suggestions for improvements or bug fixes.

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Happy syncing! 🚀
