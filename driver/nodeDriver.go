package driver

import (
	"hcc/violin-scheduler/data"
	"hcc/violin-scheduler/http"
)

// AllNodes : Search for all nodes in Pam
func AllNodes() (interface{}, error) {
	query := "query {\n" +
		"	all_node {\n" +
		"		server_uuid\n" +
		"		uuid\n" +
		"		bmc_mac_addr\n" +
		"		bmc_ip\n" +
		"		pxe_mac_addr\n" +
		"		status\n" +
		"		cpu_cores\n" +
		"		memory\n" +
		"		description\n" +
		"		created_at\n" +
		"		active\n" +
		"	}\n" +
		"}"
	var listNodeData data.ListNodeData

	result, err := http.DoHTTPRequest("flute", true, listNodeData, query)
	if err != nil {
		return listNodeData, err
	}

	return result, nil
}
