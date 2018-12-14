package gof

import (
	"strings"
)

func Explode(input string, separator string) []string {
	dataResult := strings.Split(input, separator)
	dataResultNew := make([]string, 0, len(dataResult))
	for _, singleResult := range dataResult {
		singleResult = strings.Trim(singleResult, " ")
		if len(singleResult) == 0 {
			continue
		}
		dataResultNew = append(dataResultNew, singleResult)
	}
	return dataResultNew
}

func Implode(data []string, separator string) string {
	return strings.Join(data, separator)
}
