package helpers

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	rnd "math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// CleanStringSlice takes an input and will attempt to "clean" the input by removing "empty" or "whitespace-only"
// values.  You will always get a []string response, even if you pass in a nil
func CleanStringSlice(in []string) (out []string) {
	var inVal, outVal string

	// if in nil or empty...
	if nil == in || 0 == len(in) {
		return
	}

	for _, inVal = range in {
		outVal = strings.TrimSpace(inVal)
		if "" != outVal {
			out = append(out, outVal)
		}
	}

	return
}

// UniqueStringSlice takes a string array and returns a slice with duplicates removed
// This also strips "empty" values out.  You will always get a []string response, even if you pass in a nil
func UniqueStringSlice(in []string) (out []string) {
	var inVal, outVal string

	// if in nil or empty...
	if nil == in || 0 == len(in) {
		return
	}

UniqueLoop:
	for _, inVal = range in {
		// see if empty, but don't change values for non-empty
		inVal = strings.TrimSpace(inVal)
		if "" == inVal {
			continue UniqueLoop
		}

		for _, outVal = range out {
			if outVal == inVal {
				continue UniqueLoop
			}
		}

		out = append(out, inVal)
	}

	return
}

// CombineStringSlices takes 1..* string arrays and combines them into a single slice.
// This will also unique and strip "empty" values out, so y'know...there you go.
func CombineStringSlices(ins ...[]string) (out []string, delta int) {
	var i int
	var in []string
	var inVal, outVal string

	// if there was no input...
	if 0 == len(ins) {
		return
	}

	for i, in = range ins {
		// if "empty", just move on
		if nil == in || 0 == len(in) {
			continue
		}

	ValueLoop:
		for _, inVal = range UniqueStringSlice(in) {
			for _, outVal = range out {
				if outVal == inVal {
					continue ValueLoop
				}
			}

			out = append(out, inVal)
			if 0 < i {
				delta++
			}
		}
	}

	return
}

// RemoveStringsFromSlice will attempt to remove values present in "remove" from "root"
func RemoveStringsFromSlice(root, remove []string) (out []string, delta int) {
	var rootVal, removeVal string

	// if root is nil, just return.
	if nil == root {
		return
	}

	// if remove list empty, copy root to out
	if nil == remove || 0 == len(remove) {
		out = make([]string, len(root))
		copy(out, root)
		return
	}

RootLoop:
	for _, rootVal = range root {
		for _, removeVal = range remove {
			if rootVal == removeVal {
				delta++
				continue RootLoop
			}
		}
		out = append(out, rootVal)
	}

	return
}

// AllUniqueWord will return a bool value based on
// the string containing only unique characters or not.
func AllUniqueWord(str string) bool {
	result := make(map[rune]int)
	for _, v := range str {
		result[v]++
	}
	for _, v := range result {
		if v > 1 {
			return false
		}
	}
	return true
}

// ReverseStr returns the string reversed rune-wise left to right
// Source: https://github.com/golang/example/tree/master/stringutil
func ReverseStr(s string) (result string) {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// StrPermutation checks if two strings have the same characters in the same amount
func StrPermutation(a string, b string) bool {
	ma := make(map[rune]int)
	mb := make(map[rune]int)

	for _, v := range a {
		ma[v]++
	}
	for _, v := range b {
		mb[v]++
	}
	return reflect.DeepEqual(ma, mb)
}

// ReplaceSpacesWSymb expects two strings, first arg being the string itself
// Second argument is the symbol to replace spaces
// e.g. s := "emir ozer" symb := "%20"  / Result will be "emir%20ozer"
func ReplaceSpacesWSymb(s string, symb string) (result string) {
	return strings.Join(strings.Fields(s), symb)
}

//PseudoUUID generates a uuid like string
func PseudoUUID() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}

// RandomString returns a random string of size l
// Source: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func RandomString(l int) string {
	const (
		letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		size    = int64(len(letters))
	)
	rnd.Seed(time.Now().UnixNano())
	bytes := make([]byte, l)
	for i := range bytes {
		bytes[i] = letters[rnd.Int63()%size]
	}
	return string(bytes)
}

// SHA1hash calculates a SHA1 given a string
func SHA1hash(cad string) string {
	data := sha1.Sum([]byte(cad))
	return hex.EncodeToString(data[:])
}

// UpperCaseFirst changes a string's first character to uppercase
func UpperCaseFirst(s string) string {
	phrase := []rune(s)
	phrase[0] = unicode.ToUpper(phrase[0])
	return string(phrase)
}

// UpperCaseWords changes all words' first character to uppercase (Delimited by single space)
func UpperCaseWords(s string) string {
	var stack []string
	r := regexp.MustCompile("[^\\s]+")
	words := r.FindAllString(s, -1)

	for index := range words {
		first := strings.ToUpper(string(words[index][0]))
		rest := words[index][1:]
		stack = append(stack, first+rest)
	}

	return strings.Join(stack, " ")
}

// RemoveDuplicatesFromIntSlice receives a slice of integers and iterates through them
// eliminating duplicate items in the slice.
func RemoveDuplicatesFromIntSlice(a []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

// Will return the first N chars of provided string
func LastNCharacters(s string, n int) string {
	return string(s[len(s)-n:])
}

// Will return the last N chars of provided string
func FirstNCharacters(s string, n int) string {
	return string(s[:n])
}
