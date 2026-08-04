package conversation_impl

import (
	"context"
	"fmt"

	"github.com/sohaib-library/chatapp-backend/models"
)

type dmRow struct {
	ConvID     string `gorm:"column:conv_id"`
	ConvType   string `gorm:"column:conv_type"`
	ConvAt     string `gorm:"column:conv_at"`
	OtherID    string `gorm:"column:other_id"`
	OtherName  string `gorm:"column:other_name"`
	OtherEmail string `gorm:"column:other_email"`
}

func (r *ConversationImpl) ListDMConversations(ctx context.Context, userID string) ([]models.Conversation, error) {
	var rows []dmRow

	result := r.db.WithContext(ctx).Raw(`
		SELECT
			c.id        AS conv_id,
			c.type      AS conv_type,
			c.created_at AS conv_at,
			u.id        AS other_id,
			u.name      AS other_name,
			u.email     AS other_email
		FROM conversations c
		JOIN conversation_members me    ON me.conversation_id    = c.id AND me.user_id    = ?
		JOIN conversation_members other ON other.conversation_id = c.id AND other.user_id <> ?
		JOIN users u ON u.id = other.user_id
		WHERE c.type = 'dm'
		ORDER BY c.created_at DESC
	`, userID, userID).Scan(&rows)

	if result.Error != nil {
		return nil, fmt.Errorf("list dm conversations: %w", result.Error)
	}

	conversations := make([]models.Conversation, 0, len(rows))
	for _, row := range rows {
		conv := models.Conversation{
			ID:   row.ConvID,
			Type: row.ConvType,
			OtherUser: &models.UserInfo{
				ID:    row.OtherID,
				Name:  row.OtherName,
				Email: row.OtherEmail,
			},
		}
		conversations = append(conversations, conv)
	}

	return conversations, nil
}
