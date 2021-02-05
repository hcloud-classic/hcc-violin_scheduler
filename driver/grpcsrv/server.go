package grpcsrv

import (
	"fmt"
	hccerr "github.com/hcloud-classic/hcc_errors"
	"github.com/hcloud-classic/pb"
	"hcc/violin-scheduler/action/grpc/client"
	"hcc/violin-scheduler/lib/logger"
	"hcc/violin-scheduler/lib/scheduler"
	"hcc/violin-scheduler/model"
	"time"
)

func reformatPBReqToPBServer(contents *pb.ReqScheduleHandler) *pb.ScheduleServer {
	pbServer := contents.GetServer()
	return &pb.ScheduleServer{
		ScheduleServer: pbServer.GetScheduleServer(),
		NumOfNodes:     pbServer.GetNumOfNodes(),
	}
}

func reformatPBScheduleServerToModeQuota(pbServer *pb.ScheduleServer, userQuota *model.Quota) {
	parsePb := pbServer.GetScheduleServer()
	userQuota.CPU = int(parsePb.GetCPU())
	userQuota.Memory = int(parsePb.GetMemory())
	userQuota.ServerUUID = parsePb.GetUUID()
	userQuota.UserUUID = parsePb.GetUserUUID()
	userQuota.NumberOfNodes = int(pbServer.GetNumOfNodes())
}

func reformatStringListToPBNodes(nodesList []string) *pb.ScheduledNodes {
	pbNodePtr := make([]*pb.Node, len(nodesList))

	for i, args := range nodesList {
		pbNodePtr[i] = new(pb.Node)
		pbNodePtr[i].UUID = args

	}
	// fmt.Println("Second Debug : \n", pbNodePtr)
	return &pb.ScheduledNodes{
		ShceduledNode: pbNodePtr,
	}
}

func reformatNodeListToPBNodes(nodesList *model.Nodes) *pb.ScheduledNodes {
	pbNodePtr := make([]*pb.Node, len(nodesList.Nodes))

	for i, args := range nodesList.Nodes {
		pbNodePtr[i] = new(pb.Node)
		pbNodePtr[i].ServerUUID = args.ServerUUID
		pbNodePtr[i].UUID = args.UUID
		pbNodePtr[i].CPUCores = int32(args.CPUCores)
		pbNodePtr[i].Memory = int32(args.Memory)
	}
	// fmt.Println("Second Debug : \n", pbNodePtr)
	return &pb.ScheduledNodes{
		ShceduledNode: pbNodePtr,
	}
}

func reformatPBNodesToModelNodes(pbNodes []pb.Node) []model.Node {
	var nodes []model.Node

	for _, args := range pbNodes {
		if args.CPUCores > 0 && args.Memory > 0 && args.BmcIP != "" {
			nodes = append(nodes, model.Node{
				ServerUUID: args.ServerUUID,
				UUID:       args.UUID,
				CPUCores:   int(args.CPUCores),
				Memory:     int(args.Memory),
				BmcIP:      args.BmcIP,
			})
		}
	}
	return nodes
}

// func reformatNodeListToPBNode(node *model.Node) *pb.Node {
// 	return &pb.Node{
// 		ServerUUID: node.ServerUUID,
// 		UUID:       node.UUID,
// 		CPUCores:   int32(node.CPUCores),
// 		Memory:     int32(node.Memory),
// 	}

// }

// ScheduleHandler : Manipulate Schedule Nodes
func ScheduleHandler(contents *pb.ReqScheduleHandler) (*pb.ScheduledNodes, *hccerr.HccErrorStack) {
	// var err error
	// var uuid string
	errStack := hccerr.NewHccErrorStack()
	pbServer := reformatPBReqToPBServer(contents)
	var startTime time.Time
	var elapsedTime time.Duration
	var userQuota model.Quota
	var selectedNodeList []string
	reformatPBScheduleServerToModeQuota(pbServer, &userQuota)
	logger.Logger.Println("Resolving: Scheduler")
	pbNodes, err := client.RC.GetNodeList()

	modelNodes := reformatPBNodesToModelNodes(pbNodes)

	// sort.Slice(modelNodes, func(i, j int) bool {

	// 	a, _ := strconv.Atoi(modelNodes[i].BmcIP)
	// 	b, _ := strconv.Atoi(modelNodes[j].BmcIP)
	// 	fmt.Println("comp : ", a, "<>", b)
	// 	return a < b
	// })

	for _, args := range pbNodes {
		fmt.Println("pbNodes : ", args.BmcIP, "\n", args.BmcMacAddr)
	}
	logger.Logger.Println("Flute : ", pbNodes, "\n", modelNodes)
	if err != nil {
		logger.Logger.Print(err)
		goto ERROR
	}
	// fmt.Println("Debug: \n", modelNodes)

	//Debug
	// qwe := new(model.Nodes)
	// qwe.Nodes = make([]model.Node, 5)
	// var j int
	// j = 0
	// for j < 5 {
	// 	newNode := new(model.Node)
	// 	newNode.ServerUUID = strconv.Itoa(999)
	// 	qwe.Nodes[j] = *newNode
	// 	j++
	// 	if j > 10 {
	// 		goto ERROR
	// 	}
	// }
	//Debug
	// fmt.Println("Debug: \n", qwe)

	startTime = time.Now()
	selectedNodeList, err = scheduler.NodeListParser(modelNodes, userQuota)
	// for _, selectedNodeUUID := range selectedNodeList {
	// 	testlist.NodeList = append(testlist.NodeList, selectedNodeUUID)
	// 	// fmt.Println("nodelist.NodeList: ", nodelist.NodeList)
	// }

	elapsedTime = time.Since((startTime))
	logger.Logger.Println("[Create Server Scheduling Action]\nServer UUID : ", userQuota.ServerUUID, "  Scheduling Elapse Time : ", elapsedTime)

	return reformatStringListToPBNodes(selectedNodeList), errStack.ConvertReportForm()

ERROR:
	_ = errStack.Push(hccerr.NewHccError(hccerr.ShcedulerHandlerFaild, "ScheduleHandler(): Failed to schedule"))

	return nil, errStack.ConvertReportForm()
}
