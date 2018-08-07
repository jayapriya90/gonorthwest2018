package main

import (
    "fmt"
    "runtime"
    "os"
    "bufio"
)

func main() {
    // load all ipaddresses to hashmap and print size of hashmap
    hashMap := make(map[string]bool)

    inFile, _ := os.Open("./access.log")
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
    scanner.Split(bufio.ScanLines) 

    // no hashset in go, so using dummy boolean value
    for scanner.Scan() {
        hashMap[scanner.Text()] = true
    }

    // print heap allocation after GC (retained heap memory)
    runtime.GC()
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("HeapAlloc ~= %v MB\n\n", m.HeapAlloc / 1024 / 1024) 

    fmt.Println("count:", len(hashMap))
}