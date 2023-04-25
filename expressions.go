package main

import (
	"fmt"
	"regexp"
)

var (
	RegexDate string = `\d{4}-[01]\d-[0-3]\dT[0-2]\d:[0-5]\d:[0-5]\d(?:\.\d+)?Z?`
	RegexUUID string = `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`
)

type Replacer struct {
	Expression string
	Generator  func() string
}

func ApplyReplacer(replacer Replacer, text string) string {
	re := regexp.MustCompile(replacer.Expression)

	textReplaced := re.ReplaceAllFunc([]byte(text), func(b []byte) []byte {
		value := replacer.Generator()
		return []byte(value)
	})

	return string(textReplaced)
}

func GetReplacers(replacerType string) (Replacer, error) {
	if replacerType == "date" {
		return Replacer{
			Expression: RegexDate,
			Generator:  GetRandomDate,
		}, nil
	}
	if replacerType == "uuid" {
		return Replacer{
			Expression: RegexUUID,
			Generator:  GetRandomUUID,
		}, nil
	}
	return Replacer{}, fmt.Errorf("invalid replacer type passed %s", replacerType)
}
