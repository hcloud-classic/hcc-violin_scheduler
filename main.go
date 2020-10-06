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

	// http.Handle("/graphql", graphql.GraphqlHandler)
	// logger.Logger.Println("Opening server on port " + strconv.Itoa(int(config.HTTP.Port)) + "...")
	// err := http.ListenAndServe(":"+strconv.Itoa(int(config.HTTP.Port)), nil)
	// if err != nil {
	// 	logger.Logger.Println(err)
	// 	logger.Logger.Println("Failed to prepare http server!")
	// 	return
	// }
}
