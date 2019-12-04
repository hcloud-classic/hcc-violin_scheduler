package driver

import (
	"fmt"
	"hcc/violin-scheduler/data"
	"hcc/violin-scheduler/lib/logger"
	"hcc/violin-scheduler/lib/scheduler"
	"hcc/violin-scheduler/model"
	"time"
)

//**node scheduling argument */
// cpu, mem, end of bmc ip address

// ScheduleNodes : Scheduling nodes
func ScheduleNodes(args map[string]interface{}) (interface{}, error) {
	listNodeData, err := AllNodes()
	nodes := listNodeData.(data.ListNodeData).Data.ListNode
	var userQuota model.Quota
	userQuota.ServerUUID = args["server_uuid"].(string)
	userQuota.CPU = args["cpu"].(int)
	userQuota.Memory = args["memory"].(int)
	userQuota.NumberOfNodes = args["nr_node"].(int)

	//Debug
	// logger.Logger.Println("ScheduleNodes : ", listNodeData)
	// fmt.Println("server_uuid: ", args["server_uuid"], "\ncpu : ", args["cpu"], "\nmemory : ", args["memory"], "\nnumber_of_nodes : ", args["nr_node"])
	if err != nil {
		logger.Logger.Print(err)
		return nil, err
	}
	// logger.Logger.Println("nodes : ", len(nodes))

	// logger.Logger.Println(nodes)
	// var nodelist model.ScheduledNodes

	//testing**********************
	// var nodelist data.ScheduledNodeData
	var testlist model.ScheduledNodes
	startTime := time.Now()
	selectedNodeList, err := scheduler.NodeListParser(nodes, userQuota)
	for _, selectedNodeUUID := range selectedNodeList {
		testlist.NodeList = append(testlist.NodeList, selectedNodeUUID)
		// fmt.Println("nodelist.NodeList: ", nodelist.NodeList)
	}
	var returnData data.ScheduledNodeData
	returnData.Data.ScheduledNode = testlist
	elapsedTime := time.Since((startTime))
	logger.Logger.Println("[Create Server Scheduling Action]\nServer UUID : ", args["server_uuid"], "  Scheduling Elapse Time : ", elapsedTime)
	return returnData.Data.ScheduledNode, err
}

func TestSchedule(args map[string]interface{}) (interface{}, error) {
	listNodeData, err := AllNodes()
	nodes := listNodeData.(data.ListNodeData).Data.ListNode
	var userQuota model.Quota
	userQuota.ServerUUID = args["server_uuid"].(string)
	userQuota.CPU = args["cpu"].(int)
	userQuota.Memory = args["memory"].(int)
	userQuota.NumberOfNodes = args["nr_node"].(int)

	//Debug
	// logger.Logger.Println("ScheduleNodes : ", listNodeData)
	// fmt.Println("server_uuid: ", args["server_uuid"], "\ncpu : ", args["cpu"], "\nmemory : ", args["memory"], "\nnumber_of_nodes : ", args["nr_node"])
	if err != nil {
		logger.Logger.Print(err)
		return nil, err
	}
	// logger.Logger.Println("nodes : ", len(nodes))

	// logger.Logger.Println(nodes)

	//testing**********************
	var testlist model.ScheduledNodes
	startTime := time.Now()
	selectedNodeList, err := scheduler.NodeListParser(nodes, userQuota)
	for _, selectedNodeUUID := range selectedNodeList {
		testlist.NodeList = append(testlist.NodeList, selectedNodeUUID)
		fmt.Println("nodelist.NodeList: ", testlist.NodeList)
	}
	elapsedTime := time.Since((startTime))
	logger.Logger.Println("[Create Server Scheduling Action]\nServer UUID : ", args["server_uuid"], "  Scheduling Elapse Time : ", elapsedTime)
	return testlist.NodeList, err
}
