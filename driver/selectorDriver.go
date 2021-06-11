package driver

import (
	"fmt"
	"hcc/violin-scheduler/data"
	"hcc/violin-scheduler/lib/logger"
	"hcc/violin-scheduler/lib/scheduler"
	"hcc/violin-scheduler/model"
	"time"
)

// ScheduleNodes : Scheduling nodes
func ScheduleNodes(args map[string]interface{}) (interface{}, error) {
	listNodeData, err := AllNodes()
	nodes := listNodeData.(data.ListNodeData).Data.ListNode
	var userQuota model.Quota
	userQuota.ServerUUID = args["server_uuid"].(string)
	userQuota.CPU = args["cpu"].(int)
	userQuota.Memory = args["memory"].(int)
	userQuota.NumberOfNodes = args["nr_node"].(int)

	if err != nil {
		logger.Logger.Print(err)
		return nil, err
	}
	var testlist model.ScheduledNodes
	startTime := time.Now()
	selectedNodeList, err := scheduler.NodeListParser(nodes, userQuota)
	for _, selectedNodeUUID := range selectedNodeList {
		testlist.NodeList = append(testlist.NodeList, selectedNodeUUID)
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

	if err != nil {
		logger.Logger.Print(err)
		return nil, err
	}
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
