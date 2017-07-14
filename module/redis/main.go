package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"time"

	_ "net/http/pprof"

	"github.com/garyburd/redigo/redis"
	"github.com/pkg/profile"
)

type data struct {
	Desc    string
	Expired bool
}

type user struct {
	Name string
	Date time.Time
	Data data
}

var cpuprofile = flag.String("cpu", "", "Writes some cpu profile to file.")

func init() {
	fmt.Println("Initiating...")
}

func UseRedis() {
	fmt.Println("In redis....")
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal("Connect to redis failed.")
	}
	defer c.Close()

	ud := &user{
		Name: "Ray",
		Date: time.Now(),
		Data: data{
			Desc:    "Good boy",
			Expired: false,
		},
	}
	jd, _ := json.Marshal(*ud)
	c.Do("AUTH", "123456")
	obj, err := c.Do("SET", "struct_data", jd)
	if err != nil {
		log.Println("Redis sets value error.")
	}
	fmt.Printf("The returned interface: %#v\n", obj)
	data, errx := redis.String(c.Do("GET", "struct_data"))
	if errx != nil {
		log.Println("Redis gets value error.")
	}
	fmt.Printf("The gotten value: %#v\n", data)
	fmt.Println("The gotten value in string: ", string(data))
	fmt.Println("Print a meaningless content here...")
	fmt.Println()
}

func main() {
	p := profile.Start(profile.CPUProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	defer p.Stop()

	// flag.Parse()
	// if *cpuprofile != "" {
	// 	f, err := os.Create(*cpuprofile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pprof.StartCPUProfile(f)
	// 	defer pprof.StopCPUProfile()
	// }
	http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
		UseRedis()
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		UseRedis()
		w.Write([]byte("Hello"))
	})
	http.ListenAndServe(":8080", nil)
}
