package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jasonlvhit/gocron"
	"log"
	"math/rand"
	"os"
	"time"
)

func WriteLog(message string) {
	text := fmt.Sprintf("Message: %s", message)
	f, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		panic(err)

	}
}

func myLogger() (*log.Logger, error) {

	f, err := os.OpenFile("logfile.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "Custom Logger", log.LstdFlags)
	logger.Println("Hanlding task here!")
	//logger.Println("more text to append")

	return logger, err

}

func InitCache() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

type Task struct {
	counter int
}

//func (t *Task) handler(msg int, inmuteable) {
func (t *Task) handler() {
	fmt.Println("Task is being performed.")

	t.counter += 1
	// set cache
	message := fmt.Sprintf("needs to be cached, %d, random number %d", t.counter, rand.Intn(100))
	client := InitCache()
	err := client.Set("market-price", message, 0).Err()
	if err != nil {
		panic(err)
	}
	// set db

	// handling log

	logger, _ := myLogger()
	logger.Println("logged by task handler")
}

func Runner() {

	s := gocron.NewScheduler()
	var task Task
	task.counter = 1
	s.Every(2).Seconds().Do(task.handler)
	<-s.Start()
}
func main() {

	go Runner()
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)

	}

}
