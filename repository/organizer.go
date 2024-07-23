package repository

import (
	"database/sql"
	"go-ticket/structs"
	"log"
)

func InsertOrganizer(db *sql.DB, organizer structs.Organizer) (err error) {
	var userID int
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	userSql := "INSERT INTO users (full_name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id"
	audienceSql := "INSERT INTO organizers (user_id, company_name) VALUES ($1, $2)"

	// res, err := db.Exec(sql, user.FullName, user.Email, user.Password, user.Role)
	err = tx.QueryRow(userSql, organizer.User.FullName, organizer.User.Email, organizer.User.Password, organizer.User.Role).
		Scan(&userID)
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Exec(audienceSql, userID, organizer.CompanyName)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()

	log.Println("insert user:", rowsAffected)
	return nil
}

func GetAllOrganizer(db *sql.DB) (organizers []structs.Organizer, err error) {

	sql := `
		SELECT users.id, users.full_name, users.email, users.password, users.role, organizers.user_id, organizers.company_name
		FROM users
		JOIN organizers ON users.id = organizers.user_id 
	`

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var organizer = structs.Organizer{}
		err := rows.Scan(&organizer.User.Id, &organizer.User.FullName, &organizer.User.Email, &organizer.User.Password,
			&organizer.User.Role, &organizer.UserId, &organizer.CompanyName)
		if err != nil {
			return nil, err
		}
		organizers = append(organizers, organizer)
	}
	log.Println("get users:", len(organizers))
	return organizers, nil
}

func GetOrganizerById(db *sql.DB, id int) (organizer structs.Organizer, err error) {
	sql := `
		SELECT users.id, users.full_name, users.email, users.password, users.role, organizers.user_id, organizers.company_name
		FROM users
		JOIN organizers ON users.id = organizers.user_id 
		WHERE id = $1
	`

	err = db.QueryRow(sql, id).Scan(&organizer.User.Id, &organizer.User.FullName, &organizer.User.Email, &organizer.User.Password,
		&organizer.User.Role, &organizer.UserId, &organizer.CompanyName)
	if err != nil {
		return organizer, err
	}
	return organizer, nil
}

func GetOrganizerByEmail(db *sql.DB, email string) (organizer structs.Organizer, err error) {
	sql := `
		SELECT users.id, users.full_name, users.email, users.password, users.role, organizers.user_id, organizers.company_name
		FROM users
		JOIN organizers ON users.id = organizers.user_id 
		WHERE email = $1
	`

	err = db.QueryRow(sql, email).Scan(&organizer.User.Id, &organizer.User.FullName, &organizer.User.Email, &organizer.User.Password,
		&organizer.User.Role, &organizer.UserId, &organizer.CompanyName)
	if err != nil {
		return organizer, err
	}
	return organizer, nil
}

func UpdateOrganizer(db *sql.DB, organizer structs.Organizer) error {
	var user = organizer.User
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// update user
	userSql := "UPDATE users SET full_name = $2, email = $3, password = $4, role = $5 WHERE ID = $1"
	res, err := tx.Exec(userSql, user.Id, user.FullName, user.Email, user.Password, user.Role)
	if err != nil {
		tx.Rollback()
		return err
	}
	organizerSql := "UPDATE organizers SET company_name = $2 WHERE user_id = $1"
	res, err = tx.Exec(organizerSql, user.Id, organizer.CompanyName)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	log.Println("update organizer:", rowsAffected)
	return nil

}

func DeleteOrganizer(db *sql.DB, id int) error {
	sql := "DELETE FROM users WHERE id = $1"

	res, err := db.Exec(sql, id)
	if err != nil {
		return err
	}

	affectedRows, _ := res.RowsAffected()
	log.Println("affected rows:", affectedRows)

	return nil
}
