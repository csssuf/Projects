package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	args := os.Args;
	if len(args) < 2 {
		fmt.Printf("USAGE: %s <delay> [delay2] ... [delayn]\n", args[0]);
		os.Exit(1);
	}
	for i := 1; i < len(args); i++ {
		pd, pderr := time.ParseDuration(args[i]);
		if pderr != nil { os.Exit(-1); }
		time.Sleep(pd);
	}
	fmt.Printf("\a");
}
