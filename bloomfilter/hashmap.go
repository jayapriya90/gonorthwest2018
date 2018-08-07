package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	// load all ipaddresss to hashmap
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
	fmt.Printf("HeapAlloc ~= %v MB\n\n", m.HeapAlloc/1024/1024)

	// check bloomfilter for existence of ipaddresses
	if hashMap["221.194.84.27"] {
		fmt.Println("contains ipaddress: 221.194.84.27")
	} else {
		fmt.Println("does not contain ipaddress: 221.194.84.27")
	}

	if hashMap["61.200.37.35"] {
		fmt.Println("may contain ipaddress: 61.200.37.35")
	} else {
		fmt.Println("does not contain ipaddress: 61.200.37.35")
	}

	if hashMap["127.0.0.1"] {
		fmt.Println("may contain ipaddress: 127.0.0.1")
	} else {
		fmt.Println("does not contain ipaddress: 127.0.0.1")
	}

	if hashMap["0.0.0.0"] {
		fmt.Println("may contain ipaddress: 0.0.0.0")
	} else {
		fmt.Println("does not contain ipaddress: 0.0.0.0")
	}
}
