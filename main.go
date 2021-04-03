package main

import(
	"fmt"
	"bufio"
	"os"
)

var(
	delim = '\n'
)

func
main() {
	r := bufio.NewReader(os.Stdin)
	for {
		p, e := r.ReadString(byte(delim))	
		if e!=nil {
			break
		}
		p = p[:len(p)-1]
		fmt.Printf("[%s](%s)  \n", p, p)
	}
}

