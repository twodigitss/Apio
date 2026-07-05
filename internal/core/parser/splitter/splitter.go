package splitter

import (
	"regexp"
	"strings"
)

var varLineRe = regexp.MustCompile(`^@(\w+)\s*=\s*(.+)$`)
var blockCommentRe = regexp.MustCompile(`^[!@#$%&].*`)

// Recovers all possible variables inside a .http / .rest file.
// Variables should be declared as @var = value per line.
func GetVariables(file []byte) map[string]string {
	var vars map[string]string = make(map[string]string)

	for line := range strings.SplitSeq(string(file), "\n") {
		if strings.HasPrefix(line, "@") {
			match := varLineRe.FindStringSubmatch(line)
			if match == nil {
				continue
			}

			vars[match[1]] = match[2]
		}
	}
	return vars
}

// Obtains raw request blocks delimited by ###
// ignores comments and variable definitions
func RequestSplitter(file []byte) []string {
	var cleanBlocks []string

	for block := range strings.SplitSeq(string(file), "###") {
		block = strings.NewReplacer("\r\n", "\n", "\r", "\n").Replace(block)
		block = strings.TrimSpace(block)
		if block == "" {
			continue
		}

		// remove unwanted comments before anything
		if blockCommentRe.MatchString(block) {
			continue
		}

		// at this point, the blocks may contain a comment inside the ### separator.
		// it may be treated as "<space> GET Weather" or just "<space>"
		// the first line will always be a comment or at least an empty space
		// it need to be ignored
		firstLine, rest, found := strings.Cut(block, "\n")

		// formatting the rest of the block by trimming, excluding comments, etc.
		if found && !startsWithMethod(firstLine) {
			block = rest
		}

		block = removeComments(block)
		block = resolveVariables(block, GetVariables(file))
		block = strings.TrimSpace(block)

		if block != "" && startsWithMethod(block) {
			cleanBlocks = append(cleanBlocks, block)
		}
	}

	return cleanBlocks
}
