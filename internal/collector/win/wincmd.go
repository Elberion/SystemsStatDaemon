package win

import (
	pb "SystemStatDaemon/internal/api/grpc"
)

func GetCPU() *pb.CPU {
	cpu := pb.CPU{
		UserMode:   10,
		SystemMode: 3,
		Idle:       87,
	}
	return &cpu
}

func GetSpace() *pb.Space {
	space := pb.Space{
		UsageMB:    2514.5,
		UsageINode: 100,
	}
	return &space
}
