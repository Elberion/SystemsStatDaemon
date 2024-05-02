package service

import (
	pb "SystemStatDaemon/internal/api/grpc"
	cfg "SystemStatDaemon/internal/config"
	"context"
	"sync"
	"time"
)

type OSCommand interface {
	CollectCPU(stat *SystemStat, wg *sync.WaitGroup)
	CollectSpace(stat *SystemStat, wg *sync.WaitGroup)
}

type Statistic struct {
	mu      sync.RWMutex
	storage map[time.Time]SystemStat
}

func (s *Statistic) StartCollect(ctx *context.Context, config *cfg.Config) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(config.CollectorInterval):
			wg := &sync.WaitGroup{}
			stat := &SystemStat{}
			go s.CollectCPU(stat, wg)
			go s.CollectSpace(stat, wg)
			wg.Wait()
			s.mu.Lock()
			s.storage[time.Now()] = *stat
			s.mu.Unlock()
		}
	}
}

func (s *Statistic) CollectCPU(stat *SystemStat, wg *sync.WaitGroup) {
	defer wg.Done()
}

func (s *Statistic) CollectSpace(stat *SystemStat, wg *sync.WaitGroup) {
	defer wg.Done()
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
