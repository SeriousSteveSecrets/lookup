package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

// var (
// 	//file = flag.String("f", "", "Server list to ingest")
// )

func getlist(f string) []string {
	file, _ := os.Open(f)

	var addrs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		addrs = append(addrs, scanner.Text())
	}
	fmt.Println(addrs)
	return addrs
}

func nslook(addrs []string) {
	for _, addr := range addrs {
		out, _ := exec.Command("nslookup", addr).Output()
		// fmt.Println(cmd.Run())
		fmt.Println(string(out))
	}

}

func main() {
	var testfile string
	flag.StringVar(&testfile, "f", "", "File to consume")
	flag.Parse()
	nslook(getlist(testfile))
}
