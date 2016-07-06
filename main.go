package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"strconv"
	"gopkg.in/redis.v4"
)


func redisConnect() (connected bool, countval int) {

	// connect to redis server, set by redis variable in docker-compose file!
	client := redis.NewClient(&redis.Options{
    	Addr:     "redis:6379",
    	Password: "", // no password set
		DB:       0,  // use default DB
	})

	// validate if connection worked
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	// set return value to indicate connectivity
	if err == nil {
		connected = true
		
		// get value from redis
		tempval, err := client.Get("counter").Result()

    	if err == redis.Nil {
        	fmt.Println("counter does not exists, set to 1")
        	countval = 1
        	client.Set("counter", countval, 0)
    	} else if err != nil {
        	panic(err)
    	} else {
	    	cval, err2 := strconv.Atoi(tempval)
        	fmt.Printf("current counter value: %v Error: %s\n",cval, err2)
        	countval = cval + 1
        	client.Set("counter", strconv.Itoa(countval), 0)
  		}
		
	} else {
		connected = false
		countval = 0
	}

	return connected, countval
}


func indexHandler(w http.ResponseWriter, r *http.Request) {

	// get current hostname
	hostname, _ := os.Hostname()
	redisstatus, counter := redisConnect()

	// generate the HTML output
	fmt.Fprintf(w, "<h1>Hello DBG IT Days</h1><br><i>Environment: </i><b>%s</b><br><i>Hostname: </i><b>%s</b><br><i>Redis Status: </i><b>%v</b><br><i>Redis Counter: </i><b>%v</b><br>", os.Getenv("NAME"), hostname, redisstatus, counter)
	fmt.Println(hostname, "handled HTTP REQUEST at", time.Now())
}


func main() {

	// start web server and listening on port 80
	http.HandleFunc("/", indexHandler)
	fmt.Println("Listening on port 80 for requests...")
	http.ListenAndServe(":80", nil)
}

