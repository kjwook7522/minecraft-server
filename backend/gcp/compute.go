package gcp

import (
	"context"
	"fmt"
	"time"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/option"
)

type ComputeClient struct {
	instancesClient *compute.InstancesClient
	projectID       string
	zone            string
	instanceName    string
}

func NewComputeClient(ctx context.Context, projectID, zone, instanceName string) (*ComputeClient, error) {
	instancesClient, err := compute.NewInstancesRESTClient(ctx, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to create compute client: %w", err)
	}

	return &ComputeClient{
		instancesClient: instancesClient,
		projectID:       projectID,
		zone:            zone,
		instanceName:    instanceName,
	}, nil
}

func (c *ComputeClient) StartInstance(ctx context.Context) error {
	req := &computepb.StartInstanceRequest{
		Project:  c.projectID,
		Zone:     c.zone,
		Instance: c.instanceName,
	}

	op, err := c.instancesClient.Start(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to start instance: %w", err)
	}

	return c.waitForOperation(ctx, op)
}

func (c *ComputeClient) StopInstance(ctx context.Context) error {
	req := &computepb.StopInstanceRequest{
		Project:  c.projectID,
		Zone:     c.zone,
		Instance: c.instanceName,
	}

	op, err := c.instancesClient.Stop(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to stop instance: %w", err)
	}

	return c.waitForOperation(ctx, op)
}

func (c *ComputeClient) GetInstanceStatus(ctx context.Context) (string, error) {
	req := &computepb.GetInstanceRequest{
		Project:  c.projectID,
		Zone:     c.zone,
		Instance: c.instanceName,
	}

	instance, err := c.instancesClient.Get(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to get instance: %w", err)
	}

	return instance.GetStatus(), nil
}

func (c *ComputeClient) waitForOperation(ctx context.Context, op *compute.Operation) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			status := op.Proto().GetStatus()
			if status == computepb.Operation_Status(computepb.Operation_DONE) {
				if err := op.Proto().GetError(); err != nil {
					return fmt.Errorf("operation failed: %s", err)
				}
				return nil
			}
		}
	}
}

func (c *ComputeClient) Close() error {
	return c.instancesClient.Close()
}
