package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/mrbitcoiner/gostretch/internal"
)

const Version string = "v0.1.0"

func GetArg(index int) string {
	if index >= len(os.Args) { return "" }
	return os.Args[index]
}

func main(){
	if len(os.Args) < 2 { printHelpMsg(); os.Exit(-1) }
	var (
		pw = ""
		size = uintptr(0)
		buf = []byte(nil)
	)
	switch os.Args[1] {
	case "low":
		size = internal.LOW
	case "mid":
		size = internal.MID
	case "high":
		size = internal.HIGH
	case "extreme":
		size = internal.EXTREME
	case "help":
		printHelpMsg()
		os.Exit(0)
	case "version":
		fmt.Println(Version)
	default:
		printHelpMsg()
		os.Exit(-1)
	}
	pw = getPasswd()
	buf = make([]byte, 0, size)
	buf = internal.Stretch(pw, size, buf)
	r := hex.EncodeToString(buf)
	fmt.Fprintln(os.Stdout, r)
}

func printHelpMsg(){
	fmt.Fprintf(
		os.Stderr,
		"Usage: < low(512MB) | mid(1024MB) | high(2048MB) | extreme(4096MB) > [ key ]\n",
	)
}

func getPasswd() string {
	var (
		pw string
	)
	pw = GetArg(2)
	if pw != "" { return pw }
	fmt.Printf("Your password: ")
	fmt.Scanf("%s", &pw)
	if pw == "" {
		fmt.Fprintln(os.Stderr, "empty password") 
		os.Exit(-1)
	}
	return pw
}
