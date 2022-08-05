package main

import (
	"gotdd/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{}
	sleeper.SetDuration(1 * time.Second)
	sleeper.SetSleep(time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreeterHandler)))
}
