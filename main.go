// Program uni translates tex-escaped text to unicode
//
// The program prints a line of translated text for each input argument
// to stdout.
// If no arguments are given, it reads lines from stdin.
//
// Line endings and missing final newline from stdin are preserved.
//
// Uni translates only known patterns. Unrecognized patterns such as \n
// are printed as is.
//
//
// Examples:
//	$ uni "x = \alpha+\beta"
//	x = α+β
//
//	# show all patterns in a human readable form
//	uni -h
//
// 	# list all patterns in the order of the replacements
//	uni -l
//
// Uni can be used in ktye/editor by selecting text and middle-clicking on |uni in the tag bar.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-h" {
		fmt.Println(gen)
	} else if len(os.Args) == 2 && os.Args[1] == "-l" {
		for _, e := range table {
			fmt.Printf("%s %s\n", e[0], e[1])
		}
	} else if len(os.Args) > 1 {
		for _, line := range os.Args[1:] {
			translate(line + "\n")
		}
	} else {
		// We want to keep "\r\n" line endings.
		// If no final newline exists, we don't add one.
		r := bufio.NewReader(os.Stdin)
		for {
			if s, err := r.ReadString('\n'); err != nil {
				// The final line may have no newline.
				if err == io.EOF {
					translate(s)
					break
				} else {
					log.Fatal(err)
				}
			} else {
				translate(s)
			}
		}
	}
}

// Translate replaces all known patterns in s.
func translate(s string) {
	for _, e := range table {
		s = strings.Replace(s, e[1], e[0], -1)
	}
	if _, err := io.WriteString(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
}
