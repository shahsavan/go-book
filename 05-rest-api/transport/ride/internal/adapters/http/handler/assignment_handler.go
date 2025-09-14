package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yourname/transport/ride/internal/adapters/http/api"
	"github.com/yourname/transport/ride/internal/core/ports"
)

// AssignmentHandler is the HTTP adapter. It implements api.ServerInterface
// and delegates to the core AssignmentService (a hexagonal port).
type AssignmentHandler struct {
	service ports.AssignmentService
}

// NewAssignmentHandler constructs an AssignmentHandler with the given service.
func NewAssignmentHandler(service ports.AssignmentService) *AssignmentHandler {
	return &AssignmentHandler{service: service}
}

func (h *AssignmentHandler) ListAssignments(c *gin.Context, params api.ListAssignmentsParams) {

}
func (h *AssignmentHandler) CreateAssignment(c *gin.Context) {

}
func (h *AssignmentHandler) GetAssignment(c *gin.Context, id string) {}

// Ensure we implement the generated interface
var _ api.ServerInterface = (*AssignmentHandler)(nil)
