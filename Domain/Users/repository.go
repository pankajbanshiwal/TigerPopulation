package Users

import (
	"TigerPopulation/Utils"
	"TigerPopulation/Utils/dbConfig"
	"database/sql"

	"github.com/golang/glog"
)

const (
	createUser = "INSERT INTO users  (user_name, password, email, created, updated) VALUES($1,$2,$3,$4,$5)"
)

func CreateDbUser(req *CreateUser) error {
	dateTime := Utils.GetDateTime()
	_, err := dbConfig.DB.Exec(createUser, &req.UserName, &req.Password, &req.Email, dateTime, dateTime)
	if err != nil {
		glog.Errorln("\n", err.Error())
		return err
	}
	return nil
}

func CheckIfUserNameExists(userName string) (bool, error) {
	sqlStatement := `SELECT id FROM users WHERE user_name=$1 AND status=$2;`
	var id int
	row := dbConfig.DB.QueryRow(sqlStatement, userName, "ACTIVE")
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, err

}

func CheckIfEmailIdExists(email string) (bool, error) {
	sqlStatement := `SELECT id FROM users WHERE email=$1 AND status=$2;`
	var id int
	row := dbConfig.DB.QueryRow(sqlStatement, email, "ACTIVE")
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, err

}

func CheckLogInDetails(user *CreateUser) (*CreateUser, error) {
	sqlStatement := `SELECT id,user_name,password,email,status FROM users WHERE (user_name=$1 AND password=$2 AND status=$3) OR (email=$4 AND password=$5 AND status=$6);`
	var id int
	var userName, password, email, status string
	row := dbConfig.DB.QueryRow(sqlStatement, user.UserName, user.Password, "ACTIVE", user.Email, user.Password, "ACTIVE")
	err := row.Scan(&id, &userName, &password, &email, &status)
	if err == sql.ErrNoRows {
		return nil, err
	}
	dbUser := &CreateUser{}
	dbUser.Email = email
	dbUser.Password = password
	dbUser.UserName = userName
	dbUser.Id = id
	dbUser.Status = status
	return dbUser, err
}

func UpdateUserAuthToken(userId int, token string) error {
	dateTime := Utils.GetDateTime()
	sqlStatement := `UPDATE users SET auth_token = $2, updated=$3 WHERE id = $1;`
	_, err := dbConfig.DB.Exec(sqlStatement, userId, token, dateTime)
	if err != nil {
		return err
	}
	return nil
}
