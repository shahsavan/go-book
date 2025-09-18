package grpcadapter

import (
	"context"

	"github.com/yourname/transport/vehicle/internal/adapters/grpc/vehiclepb"
	"github.com/yourname/transport/vehicle/internal/core/ports"
)

// VehicleServer is our gRPC adapter. It implements pb.VehicleServiceServer
// and delegates requests to the domain service (ports.VehicleServicePort).
type VehicleServer struct {
	vehiclepb.UnimplementedVehicleServiceServer                          // forward-compatibility
	svc                                         ports.VehicleServicePort // domain port
}

// FindAvailableVehicle adapts a gRPC request to a domain call.
func (s *VehicleServer) FindAvailableVehicle(
	ctx context.Context,
	req *vehiclepb.FindRequest,
) (*vehiclepb.FindResponse, error) {
	id, status, err := s.svc.FindAvailableVehicle(ctx, req.RouteId)
	if err != nil {
		return nil, err
	}
	return &vehiclepb.FindResponse{VehicleId: id, Status: status}, nil
}

// GetVehicleInfo adapts gRPC InfoRequest to the domain.
func (s *VehicleServer) GetVehicleInfo(
	ctx context.Context,
	req *vehiclepb.InfoRequest,
) (*vehiclepb.InfoResponse, error) {
	id, vType, status, err := s.svc.GetVehicleInfo(ctx, req.VehicleId)
	if err != nil {
		return nil, err
	}
	return &vehiclepb.InfoResponse{
		VehicleId: id,
		Type:      vType,
		Status:    status,
	}, nil
}

// StreamAssignments demonstrates a streaming RPC.
// In practice, it reads from a domain channel or queue.
func (s *VehicleServer) StreamAssignments(
	stream vehiclepb.VehicleService_StreamAssignmentsServer,
) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		// Imagine passing to domain: s.svc.Assign(...)
		ack := &vehiclepb.AssignmentAck{
			AssignmentId: req.AssignmentId,
			Accepted:     true,
		}
		if err := stream.Send(ack); err != nil {
			return err
		}
	}
}
