package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jasonlvhit/gocron"
)

func main() {

	shutdown_time := os.Args[1]
	if len(shutdown_time) == 0 {
		shutdown_time = "18:30"
	}
	fmt.Printf("Bilgisayar %s tarihinde kapanacak!\n", shutdown_time)

	s := gocron.NewScheduler()
	s.Every(1).Day().At(shutdown_time).Do(shutdownPc)
	<-s.Start()

}

func shutdownPc() {
	if err := exec.Command("cmd", "/C", "shutdown", "/s").Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}
