package server

import (
	"context"
	"github.com/hcloud-classic/pb"
	"hcc/violin-scheduler/action/grpc/errconv"
	"hcc/violin-scheduler/driver/grpcsrv"
	"hcc/violin-scheduler/lib/logger"
)

type schedulerServer struct {
	pb.UnimplementedSchedulerServer
}

func (s *schedulerServer) ScheduleHandler(_ context.Context, in *pb.ReqScheduleHandler) (*pb.ResScheduleHandler, error) {
	logger.Logger.Println("Request received: Scheduling Nodes()")
	// fmt.Println("Grpc : \n", &pb.ResVolumeHandler{Volume: &pb.Volume{}, HccErrorStack: errconv.HccStackToGrpc(nil)})
	nodeList, errStack := grpcsrv.ScheduleHandler(in)
	if nodeList == nil {
		return &pb.ResScheduleHandler{Nodes: &pb.ScheduledNodes{}, HccErrorStack: errconv.HccStackToGrpc(errStack)}, nil
	}

	return &pb.ResScheduleHandler{Nodes: nodeList}, nil
}
