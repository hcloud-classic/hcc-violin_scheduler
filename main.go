package main

import (
	"hcc/violin-scheduler/action/grpc/server"
	schedulerEnd "hcc/violin-scheduler/end"
	schedulerInit "hcc/violin-scheduler/init"

	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := schedulerInit.MainInit()
	if err != nil {
		panic(err)
	}
}

func main() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		schedulerEnd.MainEnd()
		fmt.Println("Exiting violin module...")
		os.Exit(0)
	}()
	server.Init()

}
