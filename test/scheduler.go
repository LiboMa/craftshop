package main

import (
	"fmt"
	//	"reflect"
	"os"
	"time"
)

func JobRunner() func(t time.Duration, s string) {
	f := func(t time.Duration, s string) {
		ticker := time.NewTicker(t)
		for ; true; <-ticker.C {

			text := fmt.Sprintf("time start for each %s, dealing with message %s\n", t, s)

			f, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}

			defer f.Close()

			if _, err = f.WriteString(text); err != nil {
				panic(err)
			}
		}
	}

	return f
	// blocking
}

func ShowMsg(message string) {

	fmt.Println(message)

}
func main() {

	var dur time.Duration

	dur = 10 * time.Second

	runner := JobRunner()

	//fmt.Println(reflect.TypeOf(runner))
	runner(dur, "hello")

	//JobRunner(dur, ShowMsg("Job 1"))

	//var f ShowMsg

	//JobRunner(dur, f)

	//	ticker := time.NewTicker(time.Second * dur)
	//
	//	//os.File("text.log")
	//	for ; true; <-ticker.C {
	//		fmt.Printf("time start for each %s \n", dur*time.Second)
	//	}

}
