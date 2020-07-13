package scheduler

import (
	"errors"
	"hcc/violin-scheduler/lib/logger"
	"hcc/violin-scheduler/model"
	"sort"
	"strconv"
	"strings"
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

func IPsplitToInt(ip string) int {
	tmptp := strings.Split(ip, ".")
	if len(tmptp) == 4 {
		intIP, err := strconv.Atoi(tmptp[3])
		if err != nil {
			return 0
		}
		return intIP
	}
	return 0
}

func SetValue(nodemap map[int]*nodeInfo, UUID string, cpus int, mems int, Index int, bmcip int) {
	nodemap[Index] = &nodeInfo{NodeUUID: UUID, CPU: cpus, Mem: mems, NodeOrder: bmcip}
}

func InputTest(nodemap []*nodeInfo) ([]string, error) {
	var seletedNodeList []string
	logger.Logger.Println("Appending Selected nodes")
	for a, b := range nodemap {
		seletedNodeList = append(seletedNodeList, b.NodeUUID)
		logger.Logger.Println(a, *b)
	}
	return seletedNodeList, nil
}

func SelectorInit(nodemap []*nodeInfo, userquota model.Quota) ([]string, error) {

	tmpmap := BuildSliceInit(userquota.NumberOfNodes)
	checkPathStatus.CPU = userquota.CPU
	checkPathStatus.Mem = userquota.Memory
	checkPathStatus.Depth = userquota.NumberOfNodes
	checkPathStatus.IsFind = false
	SearchPath(nodemap, tmpmap, 0, 0, 0)
	var nodeUUIDs []string

	if checkPathStatus.IsFind {
		for index := 0; index < len(checkPathStatus.NavigatePath); index++ {
			nodeUUIDs = append(nodeUUIDs, nodemap[checkPathStatus.NavigatePath[index]].NodeUUID)
		}
	} else {
		return nodeUUIDs, errors.New("Not Satisfing Node")
	}
	ResetGlobalVal()
	return nodeUUIDs, nil
}
func BuildSliceInit(size int) *[]int {
	dp := make([]int, size)
	for i := 0; i < len(dp); i++ {
		dp[i] = 0
	}
	return &dp
}

func IsoptimizedPath(cpu int, mem int, depth int) bool {
	if (cpu == checkPathStatus.CPU && mem == checkPathStatus.Mem && checkPathStatus.IsFind == false) || (checkPathStatus.CPU == 0 && checkPathStatus.Mem == 0 && checkPathStatus.Depth == depth) {
		checkPathStatus.IsFind = true
		return true
	}
	return false
}
func IsvaildQuota(cpu int, mem int, depth int) bool {
	if (cpu <= checkPathStatus.CPU && mem <= checkPathStatus.Mem && depth <= checkPathStatus.Depth) || (checkPathStatus.CPU == 0 && checkPathStatus.Mem == 0 && depth <= checkPathStatus.Depth) {
		return true
	}
	return false
}

func SearchPath(nodemap []*nodeInfo, path *[]int, cpu int, mem int, depth int) {
	if !checkPathStatus.IsFind {
		for index := 0; index < len(*path); index++ {
			if (*path)[index] != 1 && IsvaildQuota(cpu+nodemap[index].CPU, mem+nodemap[index].Mem, depth+1) {
				(*path)[index] = 1
				if IsoptimizedPath(cpu+nodemap[index].CPU, mem+nodemap[index].Mem, depth+1) {
					for triumphNumber := 0; triumphNumber < len(*path); triumphNumber++ {
						if (*path)[triumphNumber] == 1 {
							checkPathStatus.NavigatePath = append(checkPathStatus.NavigatePath, triumphNumber)
						}
					}
					break
				}
				SearchPath(nodemap, path, cpu+nodemap[index].CPU, mem+nodemap[index].Mem, depth+1)
				(*path)[index] = 0
			}
		}
	}

}
