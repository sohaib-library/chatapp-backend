package error

import "errors"


var (
	ErrInvalidRequest       = errors.New("invalid signup request")
	ErrInvalidEmail         = errors.New("enter valid email")
	ErrUserExists           = errors.New("user already exists")
	ErrInvalidCredentials   = errors.New("invalid email or password")
	ErrUserNotFound         = errors.New("user not found")
	ErrConversationNotFound = errors.New("conversation not found")
	ErrNotConversationMember = errors.New("you are not a member of this conversation")
	ErrInvalidMessage       = errors.New("message content is required")
	ErrCannotDMYourself     = errors.New("cannot start a dm with yourself")
)