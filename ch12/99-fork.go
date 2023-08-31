package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
    // Call fork to create a new process
    pid, err := syscall.ForkExec(os.Args[0], os.Args, nil)
    if err != nil {
        fmt.Println("Error forking process:", err)
        os.Exit(1)
    }

    if pid == 0 {
        // This is the child process
        fmt.Println("I am the child process with PID", os.Getpid())
        // Do some work here...
        os.Exit(0)
    } else {
        // This is the parent process
        fmt.Println("I am the parent process with PID", os.Getpid())
        fmt.Println("My child process has PID", pid)
        // Wait for the child process to exit
        var status syscall.WaitStatus
        _, err := syscall.Wait4(pid, &status, 0, nil)
        if err != nil {
            fmt.Println("Error waiting for child process:", err)
            os.Exit(1)
        }
        fmt.Println("Child process exited with status", status.ExitStatus())
    }
}