package repositories

import (
	"github.com/rymccue/fifa-tracker/models"
	"github.com/rymccue/fifa-tracker/utils"
)

func GetUserByEmail(email string) (*models.User, error) {
	return models.UserByEmail(db, email)
}

func CreateUser(email, password, firstName, lastName string) (*models.User, error) {
	salt := utils.GenerateSalt()
	hashedPassword := utils.HashPassword(password, salt)
	u := &models.User{
		Email:     email,
		Password:  hashedPassword,
		Salt:      salt,
		FirstName: firstName,
		LastName:  lastName,
	}
	err := u.Insert(db)
	return u, err
}

func GetUserByID(id int) (*models.User, error) {
	return models.UserByID(db, id)
}

func GetAllUsers(page, resultsPerPage int) ([]*models.User, error) {
	return models.GetAllUsers(db, page, resultsPerPage)
}

func SaveUser(u *models.User) error {
	return u.Save(db)
}

func DeleteUser(u *models.User) error {
	return u.Delete(db)
}
