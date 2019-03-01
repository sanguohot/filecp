package main

import (
	_ "github.com/CodyGuo/godaemon"
	"os"
)

func main() {
	done := make(chan os.Signal, 1)
	<-done
}
