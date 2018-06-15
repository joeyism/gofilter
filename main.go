package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var beginsWith = flag.String("begins-with", "", "filter the inputs that begins with this")
var notBeginsWith = flag.String("not-begins-with", "", "filter the inputs that doesn't begin with this")
var endsWith = flag.String("ends-with", "", "filter the inputs that ends with this")
var notEndsWith = flag.String("not-ends-with", "", "filter the inputs that doesn't end with this")

func main() {
	flag.Parse()

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage:")
		fmt.Println("  cat yourfile.txt | gofilter --begins-with=<pattern>")
	} else {
		reader := bufio.NewReader(os.Stdin)
		line := 1
		for {
			input, err := reader.ReadString('\n')
			if err != nil && err == io.EOF {
				break
			}
			input = strings.TrimSpace(input)
			gofilter(input)
			line++
		}
	}
}

func gofilter(input string) {
	//fmt.Println(doesBeginWith(input))
	//fmt.Println(doesEndWith(input))
	//fmt.Println(doesNotBeginWith(input))
	//fmt.Println(doesNotEndWith(input))
	//fmt.Println("-------")
	if doesBeginWith(input) && doesEndWith(input) && doesNotBeginWith(input) && doesNotEndWith(input) {
		fmt.Println(input)
	}
}

func doesBeginWith(input string) bool {
	return strings.HasPrefix(input, *beginsWith)
}

func doesEndWith(input string) bool {
	return strings.HasSuffix(input, *endsWith)
}

func doesNotBeginWith(input string) bool {
	if *notBeginsWith == "" {
		return true
	} else {
		return !strings.HasPrefix(input, *notBeginsWith)
	}
}

func doesNotEndWith(input string) bool {
	if *notEndsWith == "" {
		return true
	} else {
		return !strings.HasSuffix(input, *notEndsWith)
	}
}
