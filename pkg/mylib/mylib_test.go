package mylib

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello, World! V1"
    if result := Hello(); result != expected {
        t.Errorf("Hello() = %v; want %v", result, expected)
    }
}
