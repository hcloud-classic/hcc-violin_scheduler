package main

import (
	"hcc/violin-scheduler/action/graphql"
	schedulerEnd "hcc/violin-scheduler/end"
	schedulerInit "hcc/violin-scheduler/init"
	"hcc/violin-scheduler/lib/config"
	"hcc/violin-scheduler/lib/logger"
	"net/http"
	"strconv"
)

func init() {
	err := schedulerInit.MainInit()
	if err != nil {
		panic(err)
	}
}

func main() {
	defer func() {
		schedulerEnd.MainEnd()
	}()

	http.Handle("/graphql", graphql.GraphqlHandler)
	logger.Logger.Println("Opening server on port " + strconv.Itoa(int(config.HTTP.Port)) + "...")
	err := http.ListenAndServe(":"+strconv.Itoa(int(config.HTTP.Port)), nil)
	if err != nil {
		logger.Logger.Println(err)
		logger.Logger.Println("Failed to prepare http server!")
		return
	}
}
