package main

import (
	"fmt"
	"os"
	"time"
)

type ErrorWithTime struct {
	text string
	time string
}

func New(text string) error {
	t := time.Now()
	return &ErrorWithTime{
		text: text,
		time: t.Format(time.RFC3339),
	}
}
func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s\ntime error:\n%s", e.text, e.time)
}

func main() {

	var err error
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	_, err = os.Open("max.go")
	if err != nil {
		err = New("my error")
		panic(err)
	}

}
