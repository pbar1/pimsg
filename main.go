package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var tgt = []rune("CORPORATEMESSAGING")

func main() {
	fmt.Println("let's do our corporate messaging")

	pifile, err := os.Open("pifile.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(pifile)
	scanner.Split(bufio.ScanRunes)

	digitA := '0'
	digitB := '0'
	pos := 0
	// j := 0
	word := []rune("")
	for scanner.Scan() {
		b := scanner.Bytes()
		if len(b) != 1 {
			fmt.Fprintf(os.Stderr, "warning: skipping due to multiple bytes scanned: %s\n", b)
			continue
		}
		if b[0] == '.' {
			continue
		}

		digitA = digitB
		digitB = rune(b[0])
		pos++
		ascii, err := strconv.Atoi(fmt.Sprintf("%c%c", digitA, digitB))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		r := rune(ascii)
		if ascii >= 65 && ascii <= 90 {
			word = append(word, r)
		} else if len(word) > 0 {
			// flush word
			if len(word) > 2 {
				fmt.Printf("%s ", string(word))
			}
			word = []rune("")
		}

		// if rune(ascii) == tgt[j] {
		// 	// print only if found more than "CO"
		// 	if j > 1 {
		// 		fmt.Printf("digit: %d  -->  %s\n", pos, string(tgt[:j+1]))
		// 	}
		// 	j++
		// } else {
		// 	j = 0
		// }

		// if pos >= 100000 {
		// 	break
		// }
	}
}
