package conversation_impl

import (
	"context"
	"fmt"
)

func (r *ConversationImpl) ListMemberIDs(ctx context.Context, conversationID string) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT user_id
		FROM conversation_members
		WHERE conversation_id = $1
	`, conversationID)
	if err != nil {
		return nil, fmt.Errorf("list member ids: %w", err)
	}
	defer rows.Close()

	memberIDs := make([]string, 0)
	for rows.Next() {
		var memberID string
		if err := rows.Scan(&memberID); err != nil {
			return nil, fmt.Errorf("scan member id: %w", err)
		}
		memberIDs = append(memberIDs, memberID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate member ids: %w", err)
	}

	return memberIDs, nil
}
