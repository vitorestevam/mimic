package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var fFile = flag.String("f", "", "file path")
	flag.Parse()

	if fFile == nil || *fFile == "" {
		panic("file path (-f) is required")
	}

	rules := flag.Args()

	fmt.Println(rules)

	file, err := os.ReadFile(*fFile)
	if err != nil {
		panic(err)
	}

	text := string(file)

	replacers := []Replacer{}

	for _, rule := range rules {
		replacer, err := GetReplacers(rule)
		if err != nil {
			panic(err)
		}

		replacers = append(replacers, replacer)
	}

	new_text := text
	for _, r := range replacers {
		new_text = ApplyReplacer(r, new_text)
	}

	f, err := os.Create("./output.txt")
	if err != nil {
		panic(err)
	}

	f.Write([]byte(new_text))
}
