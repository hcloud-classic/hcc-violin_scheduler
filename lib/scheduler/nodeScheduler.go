package scheduler

import (
	"hcc/violin-scheduler/lib/logger"
	"hcc/violin-scheduler/model"
	"sort"
)

// If user want specified CPU,MEM or the number of nodes,
// using the quotaSpec Structure at selecting algorithms
type quotaSpec struct {
	CPU        int
	Mem        int
	NumOfNodes int
}

type PathStatus struct {
	CPU          int
	Mem          int
	Depth        int
	IsFind       bool
	NavigatePath []int
}
type nodeInfo struct {
	NodeUUID  string
	CPU       int
	Mem       int
	NodeOrder int
}

type keyValue struct {
	nodemap map[string]nodeInfo
}

var checkPathStatus PathStatus

type Weighting []*nodeInfo

func (a Weighting) Len() int           { return len(a) }
func (a Weighting) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Weighting) Less(i, j int) bool { return a[i].NodeOrder < a[j].NodeOrder }

func NodeListParser(nodes []model.Node, userquota model.Quota) ([]string, error) {
	var nodecount int = 0
	var nodeT = map[int]*nodeInfo{}
	for index := 0; index < len(nodes); index++ {
		if nodes[index].Active == 0 && nodes[index].ServerUUID == "" {
			SetValue(nodeT, nodes[index].UUID, (nodes[index].CPUCores)*2, nodes[index].Memory, nodecount, IPsplitToInt(nodes[index].BmcIP))
			nodecount++
		}
	}
	if userquota.NumberOfNodes == 0 {
		userquota.NumberOfNodes = nodecount + 1
	}

	tmparr := make([]*nodeInfo, 0, len(nodeT))
	for _, eachNode := range nodeT {
		tmparr = append(tmparr, eachNode)
	}

	//Sort bmp end of ip by Descending order
	sort.Sort(Weighting(tmparr))

	seletedNodeList, err := SelectorInit(tmparr, userquota)
	if err != nil {
		logger.Logger.Println(err)
	}
	logger.Logger.Println("Server Scheduled : ", seletedNodeList)
	return seletedNodeList, nil
}
