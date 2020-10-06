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
	NumOfNodes int // End of bmc ip address
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
	NodeOrder int // End of bmc ip address
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
	logger.Logger.Println("nodeListParser (%d): \n", len(nodes), nodes)
	for index := 0; index < len(nodes); index++ {
		// Later, Please Check Selected node limit equal or less than nodecount
		if nodes[index].Active == 0 && nodes[index].ServerUUID == "" {
			SetValue(nodeT, nodes[index].UUID, (nodes[index].CPUCores)*2, nodes[index].Memory, nodecount, IPsplitToInt(nodes[index].BmcIP))
			nodecount++
		}
	}
	// If Create server Action has not the number of nodes,but has cpu and memory quota
	// then filled maximun activation nodes number
	if userquota.NumberOfNodes == 0 {
		userquota.NumberOfNodes = nodecount + 1
	}
	for i, words := range nodeT {
		logger.Logger.Println(i, words.NodeUUID)
	}

	tmparr := make([]*nodeInfo, 0, len(nodeT))
	//*******Debug*******
	// logger.Logger.Println("nodes : ", nodes)
	for _, eachNode := range nodeT {
		// logger.Logger.Println("a : ", a, "eachNode : ", eachNode)
		tmparr = append(tmparr, eachNode)
	}
	// *******************
	for a, b := range tmparr {
		logger.Logger.Println("Comp :", a, *b)
	}
	//Sort bmp end of ip by Descending order
	sort.Sort(Weighting(tmparr))

	//*******Debug*******
	// for a, b := range tmparr {
	// 	logger.Logger.Println("Comp :", a, *b)
	// }
	// for index := 0; index < len(nodeT); index++ {
	// 	logger.Logger.Println("Node Num[", index, "]", nodeT[index])
	// }
	// *******************

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

// SetValue : set value
func SetValue(nodemap map[int]*nodeInfo, UUID string, cpus int, mems int, Index int, bmcip int) {
	nodemap[Index] = &nodeInfo{NodeUUID: UUID, CPU: cpus, Mem: mems, NodeOrder: bmcip}
}

// InputTest  : Test GraphQL argument
func InputTest(nodemap []*nodeInfo) ([]string, error) {
	var seletedNodeList []string
	logger.Logger.Println("Appending Selected nodes")
	for a, b := range nodemap {
		seletedNodeList = append(seletedNodeList, b.NodeUUID)
		logger.Logger.Println(a, *b)
	}
	return seletedNodeList, nil
}

// AllNodeClustering : All nodes clustering
func AllNodeClustering(numberOfNodes int, ServerUUID string) {

}

// To-Do : Select Cpu first
// CPUFritst
func CPUFritst() {

}

// MemFritst : asd
func MemFritst() {

}

func printMap(args map[int]*nodeInfo) {
	for key, value := range args {
		logger.Logger.Println("Key: [", key, "]  Value: [", *value, "]")
	}
}

func SelectorInit(nodemap []*nodeInfo, userquota model.Quota) ([]string, error) {

	// To-Do :BuildSliceInit input argument is the number of selectable nodes
	tmpmap := BuildSliceInit(userquota.NumberOfNodes)

	//*******Debug*******
	for a, b := range *tmpmap {
		logger.Logger.Println("tmpmap => ", a, b)
	}
	for a, b := range nodemap {
		logger.Logger.Println("a: ", a, "|| b : ", *b)
	}
	// *******************

	//To-Do : checkPathStatus.CPU  and checkPathStatus.Mem is filed nodemap cpu mem

	checkPathStatus.CPU = userquota.CPU
	checkPathStatus.Mem = userquota.Memory
	checkPathStatus.Depth = userquota.NumberOfNodes
	checkPathStatus.IsFind = false
	SearchPath(nodemap, tmpmap, 0, 0, 0)
	// logger.Logger.Println("result : ", checkPathStatus.NavigatePath)
	var nodeUUIDs []string

	if checkPathStatus.IsFind {
		for index := 0; index < len(checkPathStatus.NavigatePath); index++ {
			// logger.Logger.Println("node ", index, ": ", nodemap[checkPathStatus.NavigatePath[index]])
			nodeUUIDs = append(nodeUUIDs, nodemap[checkPathStatus.NavigatePath[index]].NodeUUID)
		}
	} else {
		// logger.Logger.Println("Not Satisfing Node")
		return nodeUUIDs, errors.New("Not Satisfing Node")
	}
	// logger.Logger.Println("Debug : ", nodeUUIDs)
	ResetGlobalVal()
	return nodeUUIDs, nil
}
func BuildSliceInit(size int) *[]int {
	dp := make([]int, size)
	for i := 0; i < len(dp); i++ {
		dp[i] = 0
	}
	// logger.Logger.Println("dp=>", &dp)
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

// SearchPath : visit Abailable nodes and Check out that The  node is Satisfy with quota
func SearchPath(nodemap []*nodeInfo, path *[]int, cpu int, mem int, depth int) {
	if !checkPathStatus.IsFind {
		for index := 0; index < len(nodemap); index++ {
			if (*path)[index] != 1 && IsvaildQuota(cpu+nodemap[index].CPU, mem+nodemap[index].Mem, depth+1) {
				(*path)[index] = 1
				//Debug
				// fmt.Printf("cpu[%d] mem[%d] depth[%d]\n", cpu, mem, depth)
				// fmt.Printf("index[%d]=>cpu[%d] map.cpu[%d] || mem[%d] map.mem[%d] || depth[%d]\n", index, cpu+nodemap[index].CPU, nodemap[index].CPU, mem+nodemap[index].Mem, nodemap[index].Mem, depth)
				// logger.Logger.Println("Processing : indes[", index, "]", (*path)[index])
				//Debug
				if IsoptimizedPath(cpu+nodemap[index].CPU, mem+nodemap[index].Mem, depth+1) {
					for triumphNumber := 0; triumphNumber < len(*path); triumphNumber++ {
						if (*path)[triumphNumber] == 1 {
							checkPathStatus.NavigatePath = append(checkPathStatus.NavigatePath, triumphNumber)
						}
					}
					// logger.Logger.Println(*path, " => ", cpu+nodemap[index].CPU, mem+nodemap[index].Mem, depth+1)
					break
				}
				SearchPath(nodemap, path, cpu+nodemap[index].CPU, mem+nodemap[index].Mem, depth+1)
				(*path)[index] = 0

				// logger.Logger.Println("Free : indes[", index, "]", (*path)[index])

			}
		}
	}

}

func ResetGlobalVal() {
	checkPathStatus.CPU = 0
	checkPathStatus.Mem = 0
	checkPathStatus.Depth = 0
	checkPathStatus.IsFind = false
	checkPathStatus.NavigatePath = checkPathStatus.NavigatePath[:0]
}
