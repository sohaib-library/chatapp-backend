package auth_impl

import (
	"chatapp-backend/database"
	"chatapp-backend/models"
)
 
func (a *AuthImpl) Login(login models.Login) (*models.Users, error) {
	// var user int64

	row := database.DB.QueryRow(`
	
		SELECT id , 
		name , 
		email , 
		password
		FROM users 
		WHERE email = $1 
	`, login.Email)

	var Resp models.Users

	err := row.Scan(
		&Resp.Id,
		&Resp.Name,
		&Resp.Email,
		&Resp.Password,
	)

	if err != nil {
		return nil, err
	}

	return &Resp, err
}
