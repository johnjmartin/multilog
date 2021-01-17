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
		"server2":         "server2.log",
	})
	if err != nil {
		fmt.Printf("Error initializing MultiLog %s, exiting", err)
		return
	}

	t, _ := time.Parse(time.RFC3339, "2020-01-02T00:00:00Z")
	res, _ := mtl.Query(t, 10, []string{"server1", "database server"}, multilog.WARN)
	fmt.Println(res)
}
