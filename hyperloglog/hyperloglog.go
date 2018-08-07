package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/axiomhq/hyperloglog"
)

func main() {
	// default uses 2^14 registers (byte[16384])
	hll := hyperloglog.New()

	// load all ipaddresses to hyperloglog
	inFile, _ := os.Open("./access.log")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	// read row by row and add ipaddress to hyperloglog
	for scanner.Scan() {
		// 1) string -> 64 bit hashcode
		// 2) extract LSB 14 bits which is index into byte[] register
		// 3) in remaining 50 bits, count leading zeros from left to right
		// 4) store the leading zero count in register byte[index].
		// 5) if byte[index] already has a value only retain the max value
		hll.Insert([]byte(scanner.Text()))
	}

	// print heap allocation after GC (retained heap memory)
	runtime.GC()
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("HeapAlloc ~= %v KB\n\n", mem.HeapAlloc/1024)

	// print approximate count
	// iterate byte[] register, find overall max leading zeroes found
	// estimate cardinality ~= 2^(max+1)
	fmt.Println("approximate count:", hll.Estimate())
}
