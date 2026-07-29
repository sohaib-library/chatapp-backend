package conversation_impl

import (
	"chatapp-backend/models"
	"context"
	"fmt"
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
