package converter

import (
	"regexp"
	"strings"
)

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

// rxNumberSequence is used to add boundaries between word and numbers.
var rxNumberSequence = regexp.MustCompile(`(\w)(\d+)(\w?)`)

// toCamelCase convert str to camel case, for example :
// hello_world => HelloWorld
// run_1_car => Run1Car
// run.car => RunCar
func toCamelCase(str string) string {
	// Find common initialism and make it caps
	str = strings.ReplaceAll(str, "_", " ")
	str = strings.ReplaceAll(str, "-", " ")
	str = strings.ReplaceAll(str, ".", " ")
	str = rxNumberSequence.ReplaceAllString(str, "$1 $2 $3")
	str = strings.TrimSpace(str)

	words := strings.Fields(str)
	for i, word := range words {
		word = strings.ToUpper(word)
		if _, exist := commonInitialisms[word]; exist {
			words[i] = word
		}
	}

	str = strings.Join(words, " ")

	// Create final camel case
	result := ""
	capsNext := true

	for _, ch := range str {
		// If it's already uppercase, put it as it is
		if ch >= 'A' && ch <= 'Z' {
			result += string(ch)
			continue
		}

		// 	If it's number, put it as it is and capitalize next char
		if ch >= '0' && ch <= '9' {
			result += string(ch)
			capsNext = true
			continue
		}

		// 	If it's space, don't write it and capitalize next char
		if ch == ' ' {
			capsNext = true
			continue
		}

		// At this point, char must be lowercase
		// Just put it as it is and capitalize if needed
		if ch >= 'a' && ch <= 'z' {
			if capsNext {
				result += strings.ToUpper(string(ch))
				capsNext = false
			} else {
				result += string(ch)
			}
		}
	}

	return result
}
