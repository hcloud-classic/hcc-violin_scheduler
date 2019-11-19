package scheduler

import (
	"fmt"
	"hcc/violin-scheduler/model"
	"sort"
)

type nodeInfo struct {
	NodeUUID  string
	CPU       int
	Mem       int
	NodeOrder string // End of bmc ip address
}

type keyValue struct {
	nodemap map[string]nodeInfo
}

func NodeListParser(nodes []model.Node) {
	var nodeT = map[int]*nodeInfo{}
	// fmt.Println("nodeListParser : \n", nodes)
	for index := 0; index < len(nodes); index++ {
		fmt.Println("#### : ", nodes[index])
		SetValue(nodeT, nodes[index].UUID, (nodes[index].CPUCores)*2, nodes[index].Memory, index)
	}

	// for i, words := range nodeT {
	// 	fmt.Println(i, words.NodeUUID)
	// }
	nodeT[1].CPU = 12
	nodeT[2].CPU = 2
	nodeT[3].CPU = 93
	fmt.Println("nodes : ", nodes)
	testarr := make([]*nodeInfo, 0, len(nodeT))
	for _, uuid := range nodeT {
		testarr = append(testarr, uuid)
	}
	sort.Slice(testarr, func(i, j int) bool {
		fmt.Println("re : ", i, j)
		return nodeT[i].CPU > nodeT[j].CPU
	})
	// fmt.Println(testarr)
	fmt.Println("After")
	// printMap(nodeT)
	for a, b := range testarr {
		fmt.Println(a, *b)
	}

}

func SchedulerInit() (interface{}, error) {
	var selectedNodes []string

	return selectedNodes, nil
}
func ActionParser(numberOfNodes int, NodeUUID string) {

	// testarr := make([]*nodeInfo, 0, len(nodeT))
	// for _, uuid := range nodeT {
	// 	testarr = append(testarr, uuid)
	// }

	// fmt.Println("Before")
	// for a, b := range testarr {
	// 	fmt.Println(a, *b)
	// }

	// sort.Slice(testarr, func(i, j int) bool {
	// 	fmt.Println("re : ", i, j)
	// 	return nodeT[i].CPU > nodeT[j].CPU
	// })
	// fmt.Println(testarr)
	// fmt.Println("After")

	// for a, b := range testarr {
	// 	fmt.Println(a, *b)
	// }

}

// SetValue : set value
func SetValue(nodemap map[int]*nodeInfo, UUID string, cpus int, mems int, Index int) {
	nodemap[Index] = &nodeInfo{NodeUUID: UUID, CPU: cpus, Mem: mems}
}

// ServerFirst  : asdasdasd
func ServerFirst(nodemap []int) {
	for i := range nodemap {
		nodemap[i] = i
	}
}

// AllNodeClustering : All nodes clustering
func AllNodeClustering(numberOfNodes int, ServerUUID string) {

}

// CPUFritst : asd
func CPUFritst() {

}

// MemFritst : asd
func MemFritst() {

}

func printMap(args map[int]*nodeInfo) {
	for key, value := range args {
		fmt.Println("Key: [", key, "]  Value: [", *value, "]")
	}
}
