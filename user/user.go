package user

import (
	"database/sql"
	"errors"

	CO "github.com/lsgser/go-social/config"
)

type User struct {
	ID              int64  `json:"-"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at,omitempty"`
	Password        string `json:"password,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
}
type LogInUser struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func NewUser() *User {
	return new(User)
}
func NewLogInUser() *LogInUser {
	return new(LogInUser)
}

/*
Saves a new user to the database
*/
func (u *User) SaveUser() error {
	db, err := CO.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()
	/*
	   Check if the user already exists
	*/
	var (
		username string
		email    string
	)
	stmt, err := db.Prepare("SELECT username,email FROM users WHERE username = ? OR email = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.Username, u.Email).Scan(&username, &email)
	if err != nil {
		//The user does not exist
		if err == sql.ErrNoRows {
			//Hash the users password
			hashedPassword, err := CO.HashPassword(u.Password)
			if err != nil {
				return err
			}
			//Add the new user
			insert_stmt, err := db.Prepare("INSERT users (name,surname,username,email,password) VALUES (?,?,?,?,?)")
			if err != nil {
				return err
			}
			defer insert_stmt.Close()
			_, err = insert_stmt.Exec(u.Name, u.Surname, u.Username, u.Email, hashedPassword)

			return err
		} else {
			return err
		}
	} else {
		err = errors.New("User already exists")
		return err
	}
	return err
}

/*
Login a user
If everything goes well then return the users username along with a nil error
The username will be used to generate a new JWT Token
A user should be able to login via their unique username or as an alternative
they can also user their unique email address to login
*/
func (u *LogInUser) UserLogin() (string, error) {
	db, err := CO.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return "", err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT username,password FROM users WHERE username = ? OR email = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	var (
		username string
		password string
	)
	err = stmt.QueryRow(u.User, u.User).Scan(&username, &password)

	if err == sql.ErrNoRows {
		err = errors.New("The user does not exist")
		return "", err
	} else if err != nil {
		return "", err
	}
	err = CO.CheckPassword(password, u.Password)

	if err != nil {
		err = errors.New("Wrong password")
		return "", err
	}
	return username, err
}
