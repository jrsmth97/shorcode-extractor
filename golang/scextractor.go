package scextractor

import (
	"regexp"
	"strings"
)

type BracketType string

const (
	Curly       BracketType = "curly"
	Square      BracketType = "square"
	Parentheses BracketType = "parentheses"
)

var text string
var bracketType BracketType
var bracketOpen string
var bracketClose string
var regexPattern string
var variables map[string]interface{}

func main() {
	bracketType = "curly"
}

func SetText(txt string) {
	text = txt
}

func SetBracketType(brcktType BracketType) {
	bracketType = brcktType
}

func SetVariables(vars map[string]interface{}) {
	variables = vars
}

func Extract() string {
	regexBuild()
	r := regexp.MustCompile(regexPattern)
	matches := r.FindAllString(text, -1)

	re := regexp.MustCompile(`[` + bracketOpen + bracketClose + `]`)
	newText := text
	for _, m := range matches {
		key := re.ReplaceAllString(m, "")
		val := variables[key]

		if val != nil {
			newText = strings.Replace(newText, m, val.(string), 1)
		}
	}

	return newText
}

func regexBuild() {
	switch bracketType {
	case "curly":
		regexPattern = `{({*[^{}]*}*)}`
		bracketOpen = "{"
		bracketClose = "}"
	case "square":
		regexPattern = `\[([^\]\[\r\n]*)\]`
		bracketOpen = `\[`
		bracketClose = `\]`
	case "parentheses":
		regexPattern = `\(([^\)]+)\)`
		bracketOpen = "("
		bracketClose = ")"
	default:
		panic("Invalid bracket type")
	}
}
