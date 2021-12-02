package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func StrToInt(s string, fallback int) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return fallback
	}
	return v
}

func StrSliceToIntSlice(strs []string) []int {
	ints := make([]int, len(strs))
	for in, v := range strs {
		i := StrToInt(v, 0)
		ints[in] = i
	}
	return ints
}

func StrToSlice(s string, delim string) []string {
	return strings.Split(s, delim)
}

func StrToTuple(s string, delim string) (string, string) {
	x := strings.Split(s, delim)
	return x[0], x[1]
}

func SumSlice(s []int) int {
	result := 0
	for _, v := range s {
		result += v
	}
	return result
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Readfile(day int) string {
	filename := fmt.Sprintf("day%02d/input.txt", day)
	file, err := os.Open(filename)
	PanicOnErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	contents, err := ioutil.ReadAll(reader)
	PanicOnErr(err)

	return strings.TrimSuffix(string(contents), "\n")
}

func Assert(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("Expected %d, received %d", want, got)
	} else {
		t.Logf("Got %d, want %d. Good job ✨.", got, want)
	}
}
