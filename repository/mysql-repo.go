package repository

import (
	"database/sql"
	"log"

	"github.com/JuliusIvander/case-02/entity"
	// For identifying mysql
	_ "github.com/go-sql-driver/mysql"
)

type (
	userRepoImpl struct {
		db *sql.DB
	}
)

// NewMySQL function
func NewMySQL(url string) UserRepositry {
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &userRepoImpl{db}
}

// GetUserByID Method
func (m *userRepoImpl) GetUserByID(id string) (*entity.User, error) {
	user := entity.User{}
	rows, err := m.db.Query("SELECT nama, email, nomorhp, pekerjaan FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Nama, &user.Email, &user.NomorHp, &user.Pekerjaan)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return &user, nil
}

// AddUser Method
func (m *userRepoImpl) AddUser(user *entity.User) error {
	stmt, err := m.db.Prepare("INSERT INTO users(id, nama, email, nomorhp, pekerjaan) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println(stmt)
	_, err = stmt.Exec(user.ID, user.Nama, user.Email, user.NomorHp, user.Pekerjaan)
	return err
}
