package main

import (
    "fmt"
    "sync"
    "time"
    "math/rand"
)

func sender(wg *sync.WaitGroup, S <-chan string, N chan int) {
    for s := range S {
        wg.Add(1)
        go func(v string) {     // Don't write the way below or you'll get lots of goroutines having the same value of s.
            defer wg.Done()     // go func() {
            N <- len(v)         //   N <- len(s)
        }(s)                    // }
    }
}

func receiver(N <-chan int, R chan<- int){
    var sum int
    for n := range N {
        fmt.Printf("%d <-N\n",n)
        sum += n
    }
    R <- sum
}

func main() {
    const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    s := make(chan string)
    r := make(chan int)
    n := make(chan int)
    var wg sync.WaitGroup
    rand.Seed(time.Now().UnixNano())
    go sender(&wg,s,n)                  // Important: use pointer of wg.
    go receiver(n,r)
    for i:=0;i<10;i++ {
        size := 1 + rand.Intn(20)
        b := make([]byte, size)
        for i := range b {
            b[i] = charset[rand.Intn(len(charset))]
        }
        s <- string(b)
        fmt.Printf("s<- %s\n", b)
    }
    wg.Wait()
    close(s)
    close(n)
    sum := <- r
    fmt.Println(sum)
}
