package conversation_impl

import (
	"context"
	"fmt"

	"github.com/sohaib-library/chatapp-backend/models"
)

func (r *ConversationImpl) ListMemberIDs(ctx context.Context, conversationID string) ([]string, error) {
	var members []models.ConversationMember

	result := r.db.WithContext(ctx).
		Where("conversation_id = ?", conversationID).
		Find(&members)

	if result.Error != nil {
		return nil, fmt.Errorf("list member ids: %w", result.Error)
	}

	ids := make([]string, 0, len(members))
	for _, m := range members {
		ids = append(ids, m.UserID)
	}

	return ids, nil
}
