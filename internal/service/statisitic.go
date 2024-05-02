package service

import (
	pb "SystemStatDaemon/internal/api/grpc"
	cfg "SystemStatDaemon/internal/config"
	"context"
	"time"
)

const (
	defaultCollectInterval  = 1 * time.Second
	defaultResponseInterval = 15 * time.Second
)

type StatisticService struct {
	pb.UnimplementedSystemStatServer

	statistic *Statistic
	interval  time.Duration
	cfg       *cfg.Config
}

func (s *StatisticService) GetStat(cs *pb.CollectSettings, srv pb.SystemStat_GetStatServer) error {

	return nil
}

func NewService(ctx *context.Context, cfg *cfg.Config) *StatisticService {
	srv := &StatisticService{
		statistic: NewStorageStat(),
		interval:  defaultCollectInterval,
		cfg:       cfg,
	}
	//go srv.statistic.StartCollect(ctx, cfg)
	return srv
}
