package user

import (
	"database/sql"
	"fmt"
	"formative-14/configs/database"
)

func GetAllUser() ([]User, error) {
	var users []User

	query := "SELECT * FROM users"

	rows, err := database.DB.Query(query)

	if err != nil {
		return []User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User

		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName)

		if err != nil {
			return []User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserById(id int) (User, error) {
	var user User

	query := "SELECT * FROM users WHERE id = $1"

	err := database.DB.QueryRow(query, id).
		Scan(&user.Id, &user.FirstName, &user.LastName)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("failed to get user data, user with id \"%d\" not found", id)
		}

		return User{}, err
	}

	return user, nil
}

func CreateUser(user User) (User, error) {
	query := "INSERT INTO users (first_name, last_name) VALUES ($1, $2) RETURNING *"

	err := database.DB.QueryRow(query, user.FirstName, user.LastName).
		Scan(&user.Id, &user.FirstName, &user.LastName)

	if err != nil {
		return User{}, err
	}
	
	return user, err
}

func UpdateUserById(id int, user User) (User, error) {
	query := "UPDATE users SET first_name = $2, last_name = $3 WHERE id = $1 RETURNING *"

	err := database.DB.QueryRow(query, id, user.FirstName, user.LastName).
		Scan(&user.Id, &user.FirstName, &user.LastName)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("failed updating user, user with id \"%d\" not found", id)
		}

		return User{}, err
	}

	return user, nil
}

func DeleteUserById(id int) (User, error) {
	var deletedUser User

	query := "DELETE FROM users WHERE id = $1 RETURNING *"

	err := database.DB.QueryRow(query, id).
		Scan(&deletedUser.Id, &deletedUser.FirstName, &deletedUser.LastName)

	if err != nil {
		if err == sql.ErrNoRows {
			return deletedUser, fmt.Errorf("failed deleting user, user with id \"%d\" not found", id)
		}

		return User{}, err
	}

	return deletedUser, nil
}