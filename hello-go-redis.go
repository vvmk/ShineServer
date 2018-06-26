package main

import (
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type Routine struct {
	Id            int    `json:"routine_id"`
	Title         string `json:"title"`
	TotalDuration int    `json:"total_duration"`
	CreatorId     int    `json:"creator_idk"`
	Drills        Drills `json:"drills"`
}
type Routines []Routine

type Drill struct {
	DrillTitle string `json:"drill_title"`
	Duration   int    `json:"duration"`
}
type Drills []Drill

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {

		nextId, err := redis.Int(conn.Do("INCR", "routine:id"))
		if err != nil {
			panic(err)
		}

		key := fmt.Sprintf("routine:%d", nextId)

		r := Routine{
			Id:            nextId,
			Title:         "Routine " + strconv.Itoa(nextId),
			TotalDuration: 60,
			CreatorId:     1,
			Drills: Drills{
				Drill{
					DrillTitle: "SHFFL",
					Duration:   15,
				},
				Drill{
					DrillTitle: "Shine",
					Duration:   15,
				},
				Drill{
					DrillTitle: "Repeat",
					Duration:   30,
				},
			},
		}

		//set
		conn.Do("SET", key, r)
	}

	// build key set
	keys := make([]interface{}, 10)

	for i := 1; i <= 10; i++ {
		keys = append(keys, fmt.Sprintf("routine:%d", i))
	}

	//get
	//reply, err := redis.Values(conn.Do("MGET", keys...))
	reply, err := redis.String(conn.Do("GET", "routine:8"))
	if err != nil {
		fmt.Println("key not found")
	}

	/*	var vals []string
		if _, err := redis.Scan(reply, vals); err != nil {
			fmt.Println(err)
		}
	*/
	fmt.Println(reply)
}
