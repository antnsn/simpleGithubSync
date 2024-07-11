package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "strings"
    "time"
)

var repoDirs []string

func init() {
    paths := os.Getenv("FOLDER_PATHS")
    repoDirs = strings.Split(paths, ",")
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
    for {
        for _, repoDir := range repoDirs {
            syncRepo(repoDir)
        }
        time.Sleep(60 * time.Second)
    }
}
