package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myArgs := os.Args[1:]
	if len(myArgs) >= 1 {
		for i := 0; i < len(myArgs); i++ {
			content, err := os.ReadFile(myArgs[i])
			if err != nil {
				// err.Error("ERROR: open  no such file or directory \n exit status 1")
				errLine := "ERROR: open " + myArgs[i] + ": no such file or directory"
				for _, v := range errLine {
					z01.PrintRune(rune(v))
				}
				z01.PrintRune('\n')
				os.Exit(1)
			}
			for _, v := range content {
				z01.PrintRune(rune(v))
			}
		}
	} else {
		stdin := os.Stdin
		stdout, _ := ioutil.ReadAll(stdin)
		for _, v := range stdout {
			z01.PrintRune(rune(v))
		}
	}
}
