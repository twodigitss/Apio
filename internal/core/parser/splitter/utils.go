package splitter

import (
	"encoding/json"
	"regexp"
	"strings"
)

// ponytail: inlined from deleted shared/variables.go
var httpMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}

func startsWithMethod(line string) bool {
	for _, m := range httpMethods {
		if strings.HasPrefix(strings.TrimSpace(line), m+" ") {
			return true
		}
	}
	return false
}

func removeComments(s string) string {
	lines := strings.Split(s, "\n")
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		if !strings.HasPrefix(strings.TrimSpace(line), "#") {
			result = append(result, line)
		}
	}
	return strings.Join(result, "\n")
}

var templateVarRe = regexp.MustCompile(`\{\{(\w+)\}\}`)

func resolveVariables(s string, vars map[string]string) string {
	return templateVarRe.ReplaceAllStringFunc(s,
		func(match string) string {
			key := templateVarRe.FindStringSubmatch(match)[1]
			if val, ok := vars[key]; ok {
				// return val
				b, _ := json.Marshal(val)
				return string(b[1 : len(b)-1])
			}
			return match // si no existe en el map, lo deja igual
		})
}
