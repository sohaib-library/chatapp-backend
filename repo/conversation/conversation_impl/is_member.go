package conversation_impl

import (
	"context"
	"fmt"
)

func (r *ConversationImpl) IsMember(ctx context.Context, conversationID, userID string) (bool, error) {
	var exists bool
	err := r.db.QueryRowContext(ctx, `
		SELECT EXISTS(
			SELECT 1 FROM conversation_members
			WHERE conversation_id = $1 AND user_id = $2
		)
	`, conversationID, userID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check conversation member: %w", err)
	}
	return exists, nil
}
