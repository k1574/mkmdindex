package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

var(
	delim = '\n'
)

func
ppShowPath(s string) string {
	s = strings.ReplaceAll(s, "_", "\\_")	
	s = strings.ReplaceAll(s, "*", "\\*")	
	return s
}

func
main() {
	r := bufio.NewReader(os.Stdin)
	for {
		p, e := r.ReadString(byte(delim))	
		if e!=nil {
			break
		}
		p = p[:len(p)-1]
		fmt.Printf("[%s](%s)  \n", ppShowPath(p), p)
	}
}

