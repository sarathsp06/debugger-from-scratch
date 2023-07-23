package main

import (
	"bufio"
	"fmt"
	"errors"
	"os"
	"io"
	"strings"
)

func inputContinue(pid int) bool {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Printf("\n(C)ontinue, (S)tep or (Q)uit? > ")
	for {
		input,err:=scanner.ReadString('\n')
		if err != nil {
			if !errors.Is(err,io.EOF) {
				return false
			}
			return false
		}
		switch strings.ToUpper(input[:len(input)-1]) {
		case "C":
			return true
		case "S":
			return false
		case "Q":
			os.Exit(0)
		default:
			fmt.Printf("Unexpected input %s\n", input)
			fmt.Printf("\n(C)ontinue, (S)tep, set (B)reakpoint or (Q)uit? > ")
		}
	}
}

func setBreak(pid int, filename string, line int) (bool, []byte) {
	var err error
	pc, _, err = symTable.LineToPC(filename, line)
	if err != nil {
		fmt.Printf("Can't find breakpoint for %s, %d\n", filename, line)
		return false, []byte{}
	}

	// fmt.Printf("Stopping at %X\n", pc)
	return true, replaceCode(pid, pc, interruptCode)
}
