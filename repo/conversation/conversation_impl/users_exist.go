package conversation_impl

import (
	"context"
	"fmt"

	"github.com/lib/pq"
)

func (r *ConversationImpl) UsersExist(ctx context.Context, userIDs []string) (bool, error) {
	if len(userIDs) == 0 {
		return false, nil
	}

	var count int
	err := r.db.QueryRowContext(ctx, `
		SELECT COUNT(*)
		FROM users
		WHERE id = ANY($1::uuid[])
	`, pq.Array(userIDs)).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("check users exist: %w", err)
	}

	return count == len(userIDs), nil
}
