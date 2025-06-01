package service

import (
	"context"
	"fmt"
	"sync"

	"minecraft-server/config"
	"minecraft-server/gcp"
)

type InstanceService struct {
	computeClient *gcp.ComputeClient
	mu            sync.Mutex
}

var (
	instance *InstanceService
	once     sync.Once
)

func NewInstanceService(ctx context.Context) (*InstanceService, error) {
	var err error
	once.Do(func() {
		cfg := config.Get()
		client, e := gcp.NewComputeClient(ctx,
			cfg.GCP.ProjectID,
			cfg.GCP.Zone,
			cfg.GCP.InstanceName,
		)
		if e != nil {
			err = fmt.Errorf("failed to create compute client: %w", e)
			return
		}
		instance = &InstanceService{
			computeClient: client,
		}
	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (s *InstanceService) StartServer(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	status, err := s.computeClient.GetInstanceStatus(ctx)
	if err != nil {
		return err
	}

	if status == "RUNNING" {
		return fmt.Errorf("server is already running")
	}

	return s.computeClient.StartInstance(ctx)
}

func (s *InstanceService) StopServer(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	status, err := s.computeClient.GetInstanceStatus(ctx)
	if err != nil {
		return err
	}

	if status == "TERMINATED" {
		return fmt.Errorf("server is already stopped")
	}

	return s.computeClient.StopInstance(ctx)
}

func (s *InstanceService) GetServerStatus(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.computeClient.GetInstanceStatus(ctx)
}
