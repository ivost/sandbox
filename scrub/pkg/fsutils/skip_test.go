package fsutils

import (
	"log"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://regex-golang.appspot.com/assets/html/index.html

func TestPathElemRe(t *testing.T) {
	matches := [][]string{
		{"vendor"},
		{"root", "vendor"},
		{"root", "vendor", "subvendor"},
		{"vendor", "subvendor"},
	}
	noMatches := [][]string{
		{"novendor"},
		{"vendorno"},
		{"root", "vendorno"},
		{"root", "novendor"},
		{"root", "vendorno", "subvendor"},
		{"root", "novendor", "subvendor"},
		{"vendorno", "subvendor"},
		{"novendor", "subvendor"},
	}
	check(t, "vendor", matches, noMatches)
}

func TestDirWithDot(t *testing.T) {
	matches := [][]string{
		{".git"},
		{"root", ".git"},
		{"root", ".git", "config"},
	}
	noMatches := [][]string{
		{"git"},
		{".it"},
		{".gito"},
		{"root", ".gito", "subdir"},
		{"root", ".it", "subdir"},
		{"root", "git", "subdir"},
	}
	// make sure to escape dots
	//dir := `\.git`
	dir := "\\.git"
	check(t, dir, matches, noMatches)
}

func TestStdExcludeDir(t *testing.T) {
	in := []string{".git", "vendor", "root/vendor", "/root/vendor", "/root/vendor/"}
	matches := 0
	for _, p := range in {
		for _, r := range StdExcludeDirRegexps {
			re := regexp.MustCompile(r)
			hit := re.Match([]byte(p))
			//log.Printf("%v matches %v: %v", r, p, hit)
			if hit {
				matches++
				break
			}
		}
	}
	assert.Equal(t, len(in), matches)
}

func TestOptionallyEscape(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"t0", "", ""},
		{"t1", "\\foo", "\\foo"},
		{"t2", ".git", "\\.git"},
		{"t3", "*.go", ".*\\.go"},
		{"t4", "*_gen.go", ".*_gen\\.go"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptionallyEscape(tt.arg); got != tt.want {
				t.Errorf("OptionallyEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileMatch(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		re   string
		want bool
	}{
		{"t1", "foo.go", "*.go", true},
		{"t2", "foo.got", "*.go", false},
		{"t3", "foo-go", "*.go", false},
		{"t4", "foo.go.txt", "*.go", false},
		{"t5", "", "*.go", false},
		{"t6", "/bar/foo.go", "*.go", true},
		{"t6", "bar/foo.go", "*.go", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := OptionallyEscape(tt.re)
			r = PathElemRe(r)
			re := regexp.MustCompile(r)
			hit := re.Match([]byte(tt.arg))
			assert.Equal(t, tt.want, hit)
		})
	}
}

func check(t *testing.T, name string, matches, noMatches [][]string) {
	for _, sep := range []rune{'/' /* , '\\' */} {
		reStr := pathElemReImpl(name, sep)
		re := regexp.MustCompile(reStr)
		for _, m := range matches {
			p := "/" + strings.Join(m, string(sep))
			log.Printf("%v matches %v: %v", reStr, p, re.Match([]byte(p)))
			assert.Regexp(t, re, p)
			assert.Regexp(t, re, "/"+p)
		}
		for _, m := range noMatches {
			p := strings.Join(m, string(sep))
			log.Printf("%v matches %v: %v", reStr, p, re.Match([]byte(p)))
			assert.NotRegexp(t, re, p)
			assert.NotRegexp(t, re, "/"+p)
		}
	}
}
