package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type TestAction string

var (
	Run      TestAction = "run"
	Output   TestAction = "output"
	Pass     TestAction = "pass"
	Fail     TestAction = "fail"
	Start    TestAction = "start"
	Skip     TestAction = "skip"
	Pause    TestAction = "pause"
	Bench    TestAction = "bench"
	Continue TestAction = "cont"
)

type JsonLogItem struct {
	Time    time.Time
	Action  TestAction
	Test    string
	Output  string
	Elapsed *float64
}

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
		var data JsonLogItem
		err := json.Unmarshal(input.Bytes(), &data)
		if err != nil {
			log.Fatalf("Make sure you run tests with -json (%v)", err)
		}
		if data.Action == Start || (data.Action == Run && (!*wholeMethod || !strings.Contains(data.Test, "/"))) {
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
		if data.Action == Fail {
			testFailed = true
		}
	}

	if testFailed { // if the final test has failed, clear the buffer
		for _, l := range testLines {
			fmt.Print(l)
		}
	}
}
