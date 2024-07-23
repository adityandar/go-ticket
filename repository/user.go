package repository

import (
	"database/sql"
	"go-ticket/structs"
	"log"
)

func InsertUser(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users (full_name, email, password, role) VALUES ($1, $2, $3, $4)"

	res, err := db.Exec(sql, user.FullName, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()

	log.Println("insert user:", rowsAffected)
	return nil
}

func GetAllUser(db *sql.DB) (users []structs.User, err error) {

	sql := "SELECT * FROM users"

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user = structs.User{}
		err := rows.Scan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	log.Println("get users:", len(users))
	return users, nil
}

func GetUserById(db *sql.DB, id int) (user structs.User, err error) {
	sql := "SELECT * FROM users WHERE id = $1"

	err = db.QueryRow(sql, id).Scan(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(db *sql.DB, email string) (user structs.User, err error) {
	sql := "SELECT id, full_name, email, password, role, created_at FROM users WHERE email = $1"

	err = db.QueryRow(sql, email).Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(db *sql.DB, user structs.User) error {
	sql := "UPDATE user SET full_name = $2, email = $3, password = $4, role = $5 WHERE ID = $1"

	res, err := db.Exec(sql, user.Id, user.FullName, user.Email, user.Password, user.Role)

	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	log.Println("update user:", rowsAffected)
	return nil

}

func DeleteUser(db *sql.DB, id int) error {
	sql := "DELETE FROM user WHERE id = $1"

	res, err := db.Exec(sql, id)
	if err != nil {
		return err
	}

	affectedRows, _ := res.RowsAffected()
	log.Println("affected rows:", affectedRows)

	return nil
}
