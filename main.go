package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"gopkg.in/redis.v4"
)


func redisConnect() (s string) {

	// connect to redis server, set by redis environment variable
	client, err := redis.NewClient(&redis.Options{
    	Addr:     "redis:6379",
    	Password: "", // no password set
    	DB:       0,  // use default DB
	})

	if err == nil {
				s = "Connected"
	} else {
		s = "Not available"
	}

	return s
}


func indexHandler(w http.ResponseWriter, r *http.Request) {

	hostname, _ := os.Hostname()
	redisstatus := redisConnect()

	fmt.Fprintf(w, "<h1>Hello DBG IT Days</h1><br><b>Environment: </b>%s<br><b>Hostname: </b>%s<br><b>Redis Status: </b>%s", os.Getenv("NAME"), hostname, redisstatus)
	fmt.Println(hostname, "handled HTTP REQUEST at", time.Now(), "\n Redis Status:", redisstatus)
}


func main() {

	// start web server and listening on port 5080
	http.HandleFunc("/", indexHandler)
	fmt.Println("Listening on port 5080 for requests...")
	http.ListenAndServe(":5080", nil)
}

