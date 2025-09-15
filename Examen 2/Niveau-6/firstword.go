package main

import "fmt"

func FirstWord(s string) string {
	started := false
	word := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && !started {
			started = true
		}
		if started {
			if s[i] == ' ' {
				break
			}
			word += string(s[i])
		}
	}

	return word + "\n"
}

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}
