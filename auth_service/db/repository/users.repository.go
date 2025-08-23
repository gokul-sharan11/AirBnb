package db

import (
	"database/sql"
	"fmt"
	"auth_service/models"
)

type UserRepository interface {
	GetByID(id string) (*models.User, error)
	CreateUser(username string, email string, encryptedPassword string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
} 

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository (_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db : _db,
	}
}

func (repository *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	query := "SELECT id, username, email, password, created_at, updated_at FROM users"

	rows, err := repository.db.Query(query)

	if err != nil {
		fmt.Println("Error fetching users", err)
		return nil, err
	}

	users := []*models.User{}

	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		) ; err != nil {
			fmt.Println("Error scanning user", err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}


func (repository *UserRepositoryImpl) DeleteByID(id int) error {
	query := "DELETE FROM users WHERE id = ?"

	result, err := repository.db.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting user by ID", err)
		return err
	}

	rowsAffected, rowsErr := result.RowsAffected()

	if rowsErr != nil {
		fmt.Println("Error getting rows affected", rowsErr)
		return rowsErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not deleted")
		return nil
	}

	fmt.Println("User deleted successfully, rows affected:", rowsAffected)
	return nil

}

func (repository *UserRepositoryImpl) GetByID(id string) (*models.User, error) {
	fmt.Println("Fetching user by ID")

	query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?"

	rowSet := repository.db.QueryRow(query, id)

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
	id, lastInsertIdErr := result.LastInsertId()

	if lastInsertIdErr != nil {
		fmt.Println("Error getting last insert id", lastInsertIdErr)
		return nil, lastInsertIdErr
	}

	if(err != nil){
		fmt.Println("Error geting rows affected", err)
		return nil, err
	}

	if(rowsAffected == 0){
		fmt.Println("No rows were affected, user not created", err)
		return nil, nil
	}

	fmt.Println("User created successfully, rows affected", rowsAffected)
	user := &models.User{
		ID : id,
		Username : username,
		Email : email,
	}
	return user, nil

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
