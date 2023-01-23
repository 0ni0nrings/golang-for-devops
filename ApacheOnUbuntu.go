package main

import (
    "fmt"
    "os/exec"
)

func main() {
    // Install Apache
    installCmd := exec.Command("apt-get", "install", "-y", "apache2")
    _, err := installCmd.Output()
    if err != nil {
        fmt.Println("Error installing Apache:", err)
        return
    }
    fmt.Println("Apache installed successfully")

    // Start Apache
    startCmd := exec.Command("systemctl", "start", "apache2")
    _, err = startCmd.Output()
    if err != nil {
        fmt.Println("Error starting Apache:", err)
        return
    }
    fmt.Println("Apache started successfully")

    // Enable Apache to start on boot
    enableCmd := exec.Command("systemctl", "enable", "apache2")
    _, err = enableCmd.Output()
    if err != nil {
        fmt.Println("Error enabling Apache:", err)
        return
    }
    fmt.Println("Apache enabled successfully")
}
