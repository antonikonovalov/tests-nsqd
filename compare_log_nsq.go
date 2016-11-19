package main

import (
	"flag"
	"io/ioutil"
	"strings"
)

var (
	a = flag.String("a", "", "first log")
	b = flag.String("b", "", "second log")
)

func main() {
	flag.Parse()

	fileA, err := ioutil.ReadFile(*a)
	if err != nil {
		panic(err)
	}
	fileB, err := ioutil.ReadFile(*b)
	if err != nil {
		panic(err)
	}

	matcher := make(map[string]int, 0)
	totalA := 0
	for _, str := range strings.Split(string(fileA), "\n") {
		if strings.Trim(str, ` `) == "" {
			continue
		}
		matcher[str] = 1
		totalA++
	}

	totalB := 0
	for _, str := range strings.Split(string(fileB), "\n") {
		if strings.Trim(str, ` `) == "" {
			continue
		}
		val, ok := matcher[str]
		if ok {
			matcher[str] = val + 1
		} else {
			matcher[str] = 1
		}
		totalB++
	}
	println("message\t|count\t")
	println("_____________________")
	matches := 0
	for key, count := range matcher {
		if count > 1 {
			println(key, "\t|", count)
			matches++
		}
	}
	println("_____________________")
	println("total uniq in A:", *a, totalA)
	println("total uniq in B:", *b, totalB)

	println("total doubles - ", matches)
	println("total uniq - ", len(matcher))
}
