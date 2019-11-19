package data

import "hcc/violin-scheduler/model"

// ListNodeData : Data structure of list_node
type ListNodeData struct {
	Data struct {
		ListNode []model.Node `json:"all_node"`
	} `json:"data"`
}

//SelectedNodeData : Selected nodes
type SelectedNodeData struct {
	Data struct {
		ListNode []model.Nodes `json:"all_node"`
	} `json:"data"`
}
