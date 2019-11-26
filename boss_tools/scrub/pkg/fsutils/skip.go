package fsutils

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

func PathElemRe(e string) string {
	return pathElemReImpl(e, filepath.Separator)
}

var StdExcludeDirRegexps = []string{
	PathElemRe("\\.git"), // make sure to escape dots
	PathElemRe("vendor"),
	PathElemRe("third_party"),
	PathElemRe("testdata"),
	PathElemRe("examples"),
	PathElemRe("Godeps"),
	PathElemRe("builtin"),
}

func OptionallyEscape(in string) string {
	if strings.Contains(in, "\\") {
		// assume proper regexp - no need to escape
		return in
	}
	out := strings.Replace(in, ".", "\\.", -1)
	out = strings.Replace(out, "*", ".*", -1)
	return out
}

func pathElemReImpl(e string, sep rune) string {
	escapedSep := regexp.QuoteMeta(string(sep))
	return fmt.Sprintf(`(^|%s)%s($|%s)`, escapedSep, e, escapedSep)
}
