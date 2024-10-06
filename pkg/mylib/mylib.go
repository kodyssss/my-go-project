package mylib

import (
    "os"
)

func Hello() string {
    message := os.Getenv("GREETING_MESSAGE")
    if message == "" {
        message = "Hello, World! V1"
    }
    return message
}