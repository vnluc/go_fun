package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fn := "test.py"
	fh, _ := os.OpenFile(fn, os.O_RDONLY, 0777)
	defer fh.Close()
	sc := bufio.NewScanner(fh)
	for sc.Scan() {
		data := sc.Text()
		fmt.Println(data)
	}

	s := "to test split func"
	sa := strings.Split(s, " ")

	for i, v := range sa {
		fmt.Printf("%d = %s\n", i, v)
	}
}
