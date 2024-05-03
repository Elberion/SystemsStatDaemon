package win

import (
	pb "SystemStatDaemon/internal/api/grpc"
	"os/exec"
	"strconv"
	"strings"
)

func GetCPU() *pb.CPU {
	cpu := pb.CPU{
		UserMode:   0,
		SystemMode: 0,
		Idle:       0,
	}
	cmd := exec.Command("powershell", "wmic cpu get loadpercentage /value")
	out, err := cmd.CombinedOutput()
	if err != nil {
		// TODO log
		return &cpu
	}
	str := strings.Trim(string(out[:]), "\n\r")
	str, _ = strings.CutPrefix(str, "LoadPercentage=")
	load, err := strconv.Atoi(str)
	if err != nil {
		// TODO log
		return &cpu
	}
	cpu.UserMode = float32(load)
	cpu.Idle = float32(100 - load)

	return &cpu
}

func GetSpace() *pb.Space {
	space := pb.Space{
		UsageMB:    2514.5,
		UsageINode: 100,
	}
	return &space
}
