package main

import (
    "fmt"
    "runtime"
    "os"
    "bufio"
    "github.com/tylertreat/BoomFilters"
)

func main() {
    // 10000000 expected entries
    // 5% false positive rate (error rate)
    bf := boom.NewBloomFilter(10000000, 0.05)
    
    // print bitset size and number of hash functions used
    fmt.Printf("m: %d bits (~%d MB)\n", bf.Capacity(), bf.Capacity() / 8 / 1024 / 1024)
    fmt.Printf("k: %d hash functions\n\n", bf.K())

    // load all US ipaddresss into bloomfilter
    inFile, _ := os.Open("./access.log")
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
    scanner.Split(bufio.ScanLines) 

    // read line by line and add to bloomfilter
    for scanner.Scan() {
        bf.Add([]byte(scanner.Text()))
    }

    // print heap allocation after GC (retained heap memory)
    runtime.GC()
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("HeapAlloc ~= %v MB\n\n", m.HeapAlloc / 1024 / 1024)

    // check bloomfilter for existence of ipaddress
    if bf.Test([]byte("221.194.84.27")) {
        fmt.Println("may contain ipaddress: 221.194.84.27")
    } else {
        fmt.Println("does not contain ipaddress: 221.194.84.27")
    }
    
    if bf.Test([]byte("61.200.37.35")) {
        fmt.Println("may contain ipaddress: 61.200.37.35")
    } else {
        fmt.Println("does not contain ipaddress: 61.200.37.35")
    }

    if bf.Test([]byte("127.0.0.1")) {
        fmt.Println("may contain ipaddress: 127.0.0.1")
    } else {
        fmt.Println("does not contain ipaddress: 127.0.0.1")
    }

    if bf.Test([]byte("0.0.0.0")) {
        fmt.Println("may contain ipaddress: 0.0.0.0")
    } else {
        fmt.Println("does not contain ipaddress: 0.0.0.0")
    }
}