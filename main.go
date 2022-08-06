package main

import (
	"gotdd/httpselect"
	"gotdd/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{}
	sleeper.SetDuration(1 * time.Second)
	sleeper.SetSleep(time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
	httpselect.Racer("https://facebook.com", "https://github.com")
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreeterHandler)))
}
