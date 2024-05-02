package service

import (
	pb "SystemStatDaemon/internal/api/grpc"
	"SystemStatDaemon/internal/collector/unix"
	"SystemStatDaemon/internal/collector/win"
	cfg "SystemStatDaemon/internal/config"
	"context"
	"sync"
	"time"
)

type OSCommand interface {
	CollectCPU(stat *SystemStat, wg *sync.WaitGroup, config cfg.Config)
	CollectSpace(stat *SystemStat, wg *sync.WaitGroup, config cfg.Config)
}

type Statistic struct {
	mu      sync.RWMutex
	storage map[time.Time]SystemStat
}

func (s *Statistic) StartCollect(ctx context.Context, config cfg.Config) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(config.CollectorInterval):
			wg := &sync.WaitGroup{}
			stat := &SystemStat{}
			go s.CollectCPU(stat, wg, config)
			go s.CollectSpace(stat, wg, config)
			wg.Wait()
			s.mu.Lock()
			s.storage[time.Now()] = *stat
			s.mu.Unlock()
		}
	}
}

func (s *Statistic) CollectCPU(stat *SystemStat, wg *sync.WaitGroup, config cfg.Config) {
	defer wg.Done()
	switch config.OSType {
	case 0:
		stat.CPU = unix.GetCPU()
	case 1:
		stat.CPU = win.GetCPU()
	}
}

func (s *Statistic) CollectSpace(stat *SystemStat, wg *sync.WaitGroup, config cfg.Config) {
	defer wg.Done()
	switch config.OSType {
	case 0:
		stat.Space = unix.GetSpace()
	case 1:
		stat.Space = win.GetSpace()
	}
}

type SystemStat struct {
	CPU   *pb.CPU
	Space *pb.Space
}

func NewStorageStat() *Statistic {
	return &Statistic{
		storage: make(map[time.Time]SystemStat),
	}
}
