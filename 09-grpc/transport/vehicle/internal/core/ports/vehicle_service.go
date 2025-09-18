package ports

import "context"

// VehicleServicePort defines the contract between the domain and any adapter.
// Notice: no protobuf types appear here — the domain only deals with plain Go types.
type VehicleServicePort interface {
	// FindAvailableVehicle looks up an available vehicle for a given route.
	// Returns vehicle ID, status (e.g. "available"), or an error if none found.
	FindAvailableVehicle(ctx context.Context, routeID string) (id string, status string, err error)

	// GetVehicleInfo retrieves type and status of a given vehicle.
	// Returns vehicle ID, type (bus, tram, taxi), status, or error.
	GetVehicleInfo(ctx context.Context, vehicleID string) (id string, vType string, status string, err error)

	// (Optional extension for exercises: support assignment streaming)
	// Assign handles a new assignment request (from StreamAssignments).
	// Returns whether the assignment was accepted.
	Assign(ctx context.Context, assignmentID, vehicleID, routeID string) (accepted bool, err error)
}
