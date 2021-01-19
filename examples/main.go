package main

import (
	"fmt"
	"time"

	"github.com/johnjmartin/multilog"
)

func main() {
	// relative paths used -- if running this example should be executed from the examples/ dir
	mtl, err := multilog.NewMultiLog(map[string]string{
		"server1":         "server1.log",
		"database server": "db_server.log",
	})
	if err != nil {
		fmt.Printf("Error initializing MultiLog %s, exiting", err)
		return
	}

	t, _ := time.Parse(time.RFC3339, "2020-02-28T05:20:55Z")
	fmt.Printf("query 1:\nQuery(2020-02-28T05:20:55Z, 10, []string{\"server1\", \"database server\"}, multilog.WARN)\n")
	res, _ := mtl.Query(t, 10, []string{"server1", "database server"}, multilog.WARN)

	fmt.Printf("result:\n%s\n", res)

	t, _ = time.Parse(time.RFC3339, "2020-02-28T05:20:17Z")
	fmt.Printf("query 2:\nQuery(2020-02-28T05:20:17Z, 5, []string{\"server1\", \"database server\"}, multilog.INFO)\n")
	res, _ = mtl.Query(t, 5, []string{"server1", "database server"}, multilog.INFO)

	fmt.Printf("result:\n%s\n", res)

	t, _ = time.Parse(time.RFC3339, "2020-02-28T05:20:57Z")
	fmt.Printf("query 2:\nQuery(2020-02-28T05:20:57Z, 3, []string{\"server1\", \"database server\"}, multilog.WARN)\n")
	res, _ = mtl.Query(t, 3, []string{"server1", "database server"}, multilog.WARN)

	fmt.Printf("result:\n%s\n", res)
}
