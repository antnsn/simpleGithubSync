package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "time"
)

// Function to get mounted volumes
func getMountedVolumes() ([]string, error) {
    var volumes []string
    files, err := ioutil.ReadDir("/repos")
    if err != nil {
        return nil, err
    }
    for _, file := range files {
        if file.IsDir() {
            volumes = append(volumes, filepath.Join("/repos", file.Name()))
        }
    }
    return volumes, nil
}

func setupSSH() error {
    sshDir := "/root/.ssh"
    if _, err := os.Stat(filepath.Join(sshDir, "id_rsa")); os.IsNotExist(err) {
        return fmt.Errorf("SSH key not found in %s", sshDir)
    }

    // Add GitHub to known hosts to avoid manual fingerprint verification
    knownHostsCmd := exec.Command("ssh-keyscan", "github.com")
    knownHosts, err := knownHostsCmd.Output()
    if err != nil {
        return fmt.Errorf("failed to scan GitHub SSH keys: %v", err)
    }

    err = ioutil.WriteFile(filepath.Join(sshDir, "known_hosts"), knownHosts, 0600)
    if err != nil {
        return fmt.Errorf("failed to write known_hosts file: %v", err)
    }

    return nil
}

func syncRepo(repoDir string) {
    cmd := exec.Command("git", "fetch", "origin")
    cmd.Dir = repoDir
    if err := cmd.Run(); err != nil {
        log.Printf("Failed to fetch in %s: %v", repoDir, err)
        return
    }

    local := exec.Command("git", "rev-parse", "@")
    local.Dir = repoDir
    localSHA, err := local.Output()
    if err != nil {
        log.Printf("Failed to get local SHA in %s: %v", repoDir, err)
        return
    }

    remote := exec.Command("git", "rev-parse", "@{u}")
    remote.Dir = repoDir
    remoteSHA, err := remote.Output()
    if err != nil {
        log.Printf("Failed to get remote SHA in %s: %v", repoDir, err)
        return
    }

    if string(localSHA) != string(remoteSHA) {
        log.Printf("Remote changes detected in %s. Pulling changes...", repoDir)
        cmd := exec.Command("git", "pull", "origin", "main")
        cmd.Dir = repoDir
        if err := cmd.Run(); err != nil {
            log.Printf("Failed to pull in %s: %v", repoDir, err)
            return
        }
    }

    cmd = exec.Command("git", "add", ".")
    cmd.Dir = repoDir
    if err := cmd.Run(); err != nil {
        log.Printf("Failed to add changes in %s: %v", repoDir, err)
        return
    }

    cmd = exec.Command("git", "commit", "-m", fmt.Sprintf("Automated commit %s", time.Now().Format("2006-01-02 15:04:05")))
    cmd.Dir = repoDir
    if err := cmd.Run(); err != nil {
        log.Printf("No changes to commit in %s", repoDir)
        return
    }

    cmd = exec.Command("git", "push", "origin", "main")
    cmd.Dir = repoDir
    if err := cmd.Run(); err != nil {
        log.Printf("Failed to push in %s: %v", repoDir, err)
        return
    }

    log.Printf("Successfully synced %s", repoDir)
}

func main() {
    err := setupSSH()
    if err != nil {
        log.Fatalf("SSH setup failed: %v", err)
    }

    for {
        repoDirs, err := getMountedVolumes()
        if err != nil {
            log.Fatalf("Failed to get mounted volumes: %v", err)
        }

        for _, repoDir := range repoDirs {
            syncRepo(repoDir)
        }
        time.Sleep(60 * time.Second)
    }
}
