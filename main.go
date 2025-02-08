package main

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "syscall"
    "time"
    "github.com/gen2brain/beeep"
    "github.com/olebedev/when"
    "github.com/olebedev/when/rules/common"
    "github.com/olebedev/when/rules/en"
)

const (
    markName  = "GOLANG_CLI_REMINDER"
    markValue = "1"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: golang-cli-reminder <time> <message>")
        fmt.Println("Example: golang-cli-reminder \"in 5 minutes\" \"Take a break\"")
        os.Exit(1)
    }

    now := time.Now()
    w := when.New(nil)
    w.Add(en.All...)
    w.Add(common.All...)

    t, err := w.Parse(os.Args[1], now)
    if err != nil {
        fmt.Printf("Error parsing time: %v\n", err)
        os.Exit(2)
    }

    if t == nil {
        fmt.Println("Unable to parse time. Please use a valid time format.")
        fmt.Println("Examples: \"in 5 minutes\", \"tomorrow at 3pm\", \"14:30\"")
        os.Exit(2)
    }

    if now.After(t.Time) {
        fmt.Println("Please set a time in the future")
        os.Exit(3)
    }

    diff := t.Time.Sub(now)

    if os.Getenv(markName) == markValue {
        time.Sleep(diff)
        
        exePath, err := os.Executable()
        if err != nil {
            fmt.Printf("Error getting executable path: %v\n", err)
            os.Exit(4)
        }
        
        iconPath := filepath.Join(filepath.Dir(exePath), "assets", "information.png")
        
        err = beeep.Alert("Reminder", os.Args[2], iconPath)
        if err != nil {
            fmt.Printf("Error showing notification: %v\n", err)
            os.Exit(4)
        }
    } else {
        cmd := exec.Command(os.Args[0], os.Args[1:]...)
        cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
        
        if runtime.GOOS == "windows" {
            cmd.SysProcAttr = &syscall.SysProcAttr{
                CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
            }
        }
        
        if err := cmd.Start(); err != nil {
            fmt.Printf("Error starting reminder process: %v\n", err)
            os.Exit(5)
        }

        time.Sleep(100 * time.Millisecond)
        
        fmt.Printf("Reminder will be displayed after %v\n", diff.Round(time.Second))
        os.Exit(0)
    }
}