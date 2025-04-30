package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kayrein/obviate/pkg/entities"
)

func main() {
	wholeMethod := flag.Bool("w", false, "capture events for the whole method, including all subtests")
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	testLines := make([]string, 0)
	testFailed := false

	for {
		cont := input.Scan()
		if !cont {
			break
		}
		var data entities.JsonLogItem
		err := json.Unmarshal(input.Bytes(), &data)
		if err != nil {
			log.Fatalf("Make sure you run tests with -json (%v)", err)
		}
		if data.Action == entities.Run && (!*wholeMethod || !strings.Contains(data.Test, "/")) {
			if testFailed {
				for _, l := range testLines {
					fmt.Print(l)
				}
			}
			testLines = []string{data.Output}
			testFailed = false
		} else {
			testLines = append(testLines, data.Output)
		}
		if data.Action == entities.Fail {
			testFailed = true
		}
	}

	if testFailed { // if the final test has failed, clear the buffer
		for _, l := range testLines {
			fmt.Print(l)
		}
	}
}
