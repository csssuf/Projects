package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args;
	if len(args) != 4 {
		fmt.Printf("USAGE: %s <rhost> <lowerbound> <upperbound>\n", args[0]);
		os.Exit(1);
	}
	ipaddr := args[1];
	lport, ce1 := strconv.Atoi(args[2]);
	uport, ce2 := strconv.Atoi(args[3]);
	if ce1 != nil || ce2 != nil {
		fmt.Printf("Invalid port received.");
	}
	to, toerr := time.ParseDuration("3s");
	if toerr != nil {}
	for i := lport; i <= uport; i++ {
		fmt.Printf("Scanning port %d\n", i);
		conn, err := net.DialTimeout("tcp", ipaddr + ":" + strconv.Itoa(i), to);
		if err == nil {
			fmt.Printf("%d open on TCP.\n", i);
			continue;
		}
		if conn == nil {}
	}
}
