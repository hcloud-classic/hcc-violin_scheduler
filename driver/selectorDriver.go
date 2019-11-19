package driver

import (
	"fmt"
	"hcc/violin-scheduler/data"
	"hcc/violin-scheduler/lib/logger"

	"github.com/graphql-go/graphql"
)

// ScheduleNodes : Scheduling nodes
func ScheduleNodes(params graphql.ResolveParams) (interface{}, error) {
	listNodeData, err := AllNodes()
	nodes := listNodeData.(data.ListNodeData).Data.ListNode
	// fmt.Println("ScheduleNodes : ", listNodeData)

	if err != nil {
		logger.Logger.Print(err)
		return nil, err
	}
	fmt.Println(len(nodes))
	var nrNodes int = len(nodes)
	var nodeUUIDs []string
	var nodeSelected = 0
	for _, node := range nodes {
		if nodeSelected > nrNodes {
			break
		}
		nodeUUIDs = append(nodeUUIDs, node.UUID)

		nodeSelected++
	}

	fmt.Println("nodeSelected: ", nodeSelected, " nodeUUIDs : ", nodeUUIDs)
	fmt.Println(nodes)
	return nodeUUIDs, nil
}
