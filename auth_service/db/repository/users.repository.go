package db

import (
	"database/sql"
	"fmt"
	"auth_service/models"
)

type UserRepository interface {
	GetByID() (*models.User, error)
	CreateUser(username string, email string, encryptedPassword string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
} 

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository (_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db : _db,
	}
}

func (repository *UserRepositoryImpl) GetByID() (*models.User, error) {
	fmt.Println("Fetching user by ID")

	query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?"

	rowSet := repository.db.QueryRow(query, 1)

	user := &models.User{}

	err := rowSet.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found")
			return nil, nil
		}else{
			fmt.Println("Error fetching user by ID", err)
			return nil, err
		}
	}

	fmt.Println(user)
	return user, nil
}


func (repository *UserRepositoryImpl) CreateUser(username string, email string, encryptedPassword string) (*models.User, error) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	fmt.Println("Encrpyted Pass : ", encryptedPassword)

	result, err :=repository.db.Exec(query, username, email, encryptedPassword)

	if err != nil {
		fmt.Println("Error inserting user", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()

	if(err != nil){
		fmt.Println("Error geting rows affected", err)
		return nil, err
	}

	if(rowsAffected == 0){
		fmt.Println("No rows were affected, user not created", err)
		return nil, nil
	}

	fmt.Println("User created successfully, rows affected", rowsAffected)
	return nil, nil

}

func (repository *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = ?"

	result := repository.db.QueryRow(query, email)

	user := models.User{}

	err := result.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found")
			return nil, nil
		}else{
			fmt.Println("Error fetching user by ID", err)
			return nil, err
		}
	}

	return &user,nil
}
