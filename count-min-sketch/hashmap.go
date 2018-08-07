package main

import (
    "fmt"
    "runtime"
    "os"
    "bufio"
)

func main() {
    // load all ipaddresses to hashmap
    hashMap := make(map[string]int)

    inFile, _ := os.Open("./access.log")
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
    scanner.Split(bufio.ScanLines) 

    // if ipaddress exists, increments the counter
    for scanner.Scan() {
        hashMap[scanner.Text()] += 1
    }

    // print heap allocation after GC (retained heap memory)
    runtime.GC()
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("HeapAlloc ~= %v MB\n\n", m.HeapAlloc / 1024 / 1024) 

    // print frequency count of 3 random ipaddresses
    fmt.Println("frequency of 45.32.87.56:", hashMap["45.32.87.56"])
    fmt.Println("frequency of 43.167.194.248:", hashMap["43.167.194.248"])
    fmt.Println("frequency of 99.63.207.3:", hashMap["99.63.207.3"])
    fmt.Println("frequency of 1.1.1.1:", hashMap["1.1.1.1"])
}