package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := dial()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	s, err := redis.Strings(c.Do("HGETALL", "album:1"))
	fmt.Printf("%#v\n", s)
	// s is []string{field1,value1,field2,value2...}
}

func dial() (redis.Conn, error) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return nil, nil
	}
	return c, nil
}
