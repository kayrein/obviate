package entities

import "time"

type TestAction string

var (
	Run    TestAction = "run"
	Output TestAction = "output"
	Pass   TestAction = "pass"
	Fail   TestAction = "fail"
)

type JsonLogItem struct {
	Time    time.Time
	Action  TestAction
	Test    string
	Output  string
	Elapsed *float64
}
