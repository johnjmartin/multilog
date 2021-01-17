package main

import (
	"fmt"

	"github.com/johnjmartin/multilog"
)

func main() {
	// relative paths used -- if running this example should be executed from the examples/ dir
	mtl, err := multilog.NewMultiLog(map[string]string{
		"server1":   "server1.log",
		"db server": "db_server.log",
		"server2":   "server2.log",
	})
	if err != nil {
		fmt.Printf("Error initializing MultiLog %s, exiting", err)
		return
	}

	fmt.Println(mtl.Query())
}
