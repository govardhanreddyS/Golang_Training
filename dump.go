package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    pingCh := make(chan string)
    pongCh := make(chan string)

    wg.Add(1) // One WaitGroup for the first goroutine
    go func() {
        defer wg.Done()

        for i := 0; i < 26; i++ {
            letter := string(rune('A' + i)) // Convert integer to corresponding letter
            fmt.Println("Ping", letter, i+1) // Print the letter and its corresponding iteration count
            pongCh <- letter
            <-pingCh
        }
    }()

    go func() {
        wg.Add(1) // One WaitGroup for the second goroutine
        defer wg.Done()
        for i := 0; i < 26; i++ {
            letter := <-pongCh
            fmt.Println("Pong", letter, i+1) // Print the letter and its corresponding iteration count
            pingCh <- string(rune(letter[0] + 1)) // Send next letter in sequence
        }
    }()

    wg.Wait() // Wait for the first goroutine to finish
    close(pingCh)
    close(pongCh)
}
