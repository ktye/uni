// Program uni translates tex-escaped text to unicode
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var debug bool

func main() {
	args := os.Args

	if len(args) == 2 && args[1] == "-h" {
		fmt.Println(gen)
		return
	} else if len(args) == 2 && args[1] == "-l" {
		for _, e := range table {
			fmt.Printf("%s %s\n", e[0], e[1])
		}
		return
	}

	if len(args) > 1 && args[1] == "-d" {
		debug = true
		args = args[1:]
	}

	if len(args) > 1 {
		for i, line := range args[1:] {
			translate(line)
			if i < len(args[1:])-1 {
				fmt.Println()
			}
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
	if debug {
		for _, r := range s {
			fmt.Printf("%q %+q 0x%x 0%o\n", r, r, r, r)
		}
	} else {
		if _, err := io.WriteString(os.Stdout, s); err != nil {
			log.Fatal(err)
		}
	}
}
