package main

import (
    "fmt"
    "runtime"
    "os"
    "bufio"
    "github.com/dustin/go-probably"
)

func main() {
    width := 1<<21
    depth := 2
    cms := probably.NewSketch(width, depth)

    // math is from the paper
    fmt.Printf("width: %d\n", width)
    fmt.Printf("depth: %d\n\n", depth)

    // load all the ipaddresses to count-min-sketch
    inFile, _ := os.Open("./access.log")
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        cms.Increment(scanner.Text())
    }

    // print heap allocation after GC (retained heap memory)
    runtime.GC()
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("HeapAlloc ~= %v MB\n\n", m.HeapAlloc / 1024 / 1024) 

    // print approximate frequency count of 3 random ipaddresses
    fmt.Println("approximate frequency of 45.32.87.56:", cms.Count("45.32.87.56"))
    fmt.Println("approximate frequency of 43.167.194.248:", cms.Count("43.167.194.248"))
    fmt.Println("approximate frequency of 99.63.207.3:", cms.Count("99.63.207.3"))
    fmt.Println("approximate frequency of 1.1.1.1:", cms.Count("1.1.1.1"))
}