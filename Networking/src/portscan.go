package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func checkPort(p int, address string, t time.Duration, chanind int, chans []chan int) {
	conn, err := net.DialTimeout("tcp", address + ":" + strconv.Itoa(p), t);
	if err == nil {
		fmt.Printf("%d open on TCP.\n", p);
		chans[chanind] <- 0;
	}
	if conn == nil {}
	//fmt.Printf("%d Nope %d\n", p, chanind);
	chans[chanind] <- 0;
}
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
	to, toerr := time.ParseDuration("300ms");
	var quits_received = make([]chan int, 500, 1000);
	for quit := range quits_received {
		quits_received[quit] = make(chan int)
	}
	if toerr != nil {}
	for i := lport; i <= uport; i++ {
		go checkPort(i, ipaddr, to, uport - i, quits_received);
		//fmt.Printf("Scanning port %d\n", i);
	}
	for j := lport; j <= uport; j++ {
		<-quits_received[uport - j];
	}
}
