package client

import (
	"context"
	"hcc/violin-scheduler/action/grpc/pb/rpcflute"
	"hcc/violin-scheduler/lib/config"
	"hcc/violin-scheduler/lib/logger"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

var fluteConn *grpc.ClientConn

func initFlute() error {
	var err error

	addr := config.Flute.ServerAddress + ":" + strconv.FormatInt(config.Flute.ServerPort, 10)
	fluteConn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	RC.flute = rpcflute.NewFluteClient(fluteConn)
	logger.Logger.Println("gRPC flute client ready")

	return nil
}

func closeFlute() {
	_ = fluteConn.Close()
}

// GetNode : Get infos of the node
func (rc *RPCClient) GetNode(uuid string) (*rpcflute.Node, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	node, err := rc.flute.GetNode(ctx, &rpcflute.ReqGetNode{UUID: uuid})
	if err != nil {
		return nil, err
	}

	return node.Node, nil
}

// GetNodeList : Get the list of nodes by server UUID.
func (rc *RPCClient) GetNodeList() ([]rpcflute.Node, error) {
	var nodeList []rpcflute.Node

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.Flute.RequestTimeoutMs)*time.Millisecond)
	defer cancel()
	pnodeList, err := rc.flute.GetNodeList(ctx, &rpcflute.ReqGetNodeList{Node: &rpcflute.Node{}})
	if err != nil {
		return nil, err
	}

	for _, pnode := range pnodeList.Node {
		if pnode.UUID != "" {
			nodeList = append(nodeList, rpcflute.Node{
				UUID:        pnode.UUID,
				ServerUUID:  pnode.ServerUUID,
				BmcMacAddr:  pnode.BmcMacAddr,
				BmcIP:       pnode.BmcIP,
				PXEMacAddr:  pnode.PXEMacAddr,
				Status:      pnode.Status,
				CPUCores:    pnode.CPUCores,
				Memory:      pnode.Memory,
				Description: pnode.Description,
				Active:      pnode.Active,
				CreatedAt:   pnode.CreatedAt,
			})
		}

	}

	return nodeList, nil
}
