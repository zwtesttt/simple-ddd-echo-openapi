package converter

import (
	"cosslan/internal/domain/line"
	"cosslan/internal/infra/persistence/po"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func ToPO(l *line.Line) *po.Line {
	// Convert []byte to primitive.ObjectID
	userID, err := primitive.ObjectIDFromHex(l.UserID)
	if err != nil {
		fmt.Println("Error converting NextNodeID to ObjectID:", err)
	}

	nextNodeID, err := primitive.ObjectIDFromHex(l.NextNodeID)
	if err != nil {
		fmt.Println("Error converting NextNodeID to ObjectID:", err)
	}

	// Convert int64 to time.Time
	createdAt := time.Unix(0, l.CreatedAt)
	updatedAt := time.Unix(0, l.UpdatedAt)

	// Convert int64 to *time.Time
	var deletedAt *time.Time
	if l.DeletedAt != 0 {
		dt := time.Unix(0, l.DeletedAt)
		deletedAt = &dt
	}

	return &po.Line{
		Name:       l.Name,
		Type:       l.Type,
		Subnets:    l.Subnets,
		UserID:     userID,
		NextNodeID: nextNodeID,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeletedAt:  deletedAt,
	}
}

// FromPO converts a PO to a Line entity.
func FromPO(p *po.Line) *line.Line {
	// Convert primitive.ObjectID to []byte
	userID := p.UserID.Bytes()
	nextNodeID := p.NextNodeID.Bytes()

	// Convert time.Time to int64
	createdAt := p.CreatedAt.UnixNano()
	updatedAt := p.UpdatedAt.UnixNano()

	// Convert *time.Time to int64
	var deletedAt int64
	if p.DeletedAt != nil {
		deletedAt = p.DeletedAt.UnixNano()
	}

	return &line.Line{
		Name:       p.Name,
		Type:       p.Type,
		Subnets:    p.Subnets,
		UserID:     userID,
		NextNodeID: nextNodeID,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeletedAt:  deletedAt,
	}
}
