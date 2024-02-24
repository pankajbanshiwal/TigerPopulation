package Tigers

import (
	"TigerPopulation/Utils"
	"TigerPopulation/Utils/dbConfig"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang/glog"
)

const (
	createUser           = "INSERT INTO tigers  (name, dob, last_seen,last_seen_location, created, updated) VALUES($1,$2,$3,$4,$5,$6)"
	tigerExists          = `SELECT id FROM tigers WHERE name=$1 AND dob=$2 AND status=$3;`
	getAllTigers         = `select id,name,dob,last_seen ,status , last_seen_location[0],last_seen_location[1]  from tigers ORDER BY last_seen  desc  limit $1 offset $2;`
	getTigerSightedUsers = "select u.id, u.user_name, u.email, (select name from tigers where id = $1) from tiger_sightings as ts inner join users u on (u.id = ts.user_id) where ts.tiger_id = $2 and u.status = 'ACTIVE' and ts.status = 'ACTIVE' group by u.id"
)

func CreateTiger(req *Tiger) error {
	dateTime := Utils.GetDateTime()
	date, parseErr := time.Parse("2006-01-02", req.Dob)
	if parseErr != nil {
		glog.Errorln("Error parsing date = Err ", parseErr)
		return parseErr
	}
	lastSeenTime := time.Unix(req.LastSeen, 0)
	lastSeen := lastSeenTime.Format("2006-01-02 15:04:05")

	_, err := dbConfig.DB.Exec(createUser, req.Name, date, lastSeen, fmt.Sprintf("(%f,%f)", req.Lat, req.Long), dateTime, dateTime)
	if err != nil {
		glog.Errorln("\n", err.Error())
		return err
	}
	return nil
}

func CheckIfTigerExists(name, dob string) (bool, error) {
	date, parseErr := time.Parse("2006-01-02", dob)
	if parseErr != nil {
		glog.Errorln("Error parsing date = Err ", parseErr)
		return false, parseErr
	}
	var id int
	row := dbConfig.DB.QueryRow(tigerExists, name, date, "ACTIVE")
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, err
}

func CheckIfTigerDobExists(dob string) (bool, error) {
	date, err := time.Parse("2006-01-02", dob)
	if err != nil {
		glog.Errorln("Error parsing date = Err ", err)
		return false, err
	}
	sqlStatement := `SELECT id FROM tigers WHERE dob=$1 AND status=$2;`
	var id int
	row := dbConfig.DB.QueryRow(sqlStatement, date, "ACTIVE")
	err = row.Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, err
}
func GetAllTigers(page int) ([]Tiger, error) {
	limit := 10
	offset := (page - 1) * limit
	rows, err := dbConfig.DB.Query(getAllTigers, limit, offset)
	if err != nil {
		glog.Errorln("Error GetAllTigers = Err ", err)

	}
	defer rows.Close()
	// Create a slice to store the results.
	objects := []Tiger{}
	// Iterate over the rows and scan the results into the slice.
	for rows.Next() {
		var object Tiger
		var lastSeen time.Time
		err := rows.Scan(&object.Id, &object.Name, &object.Dob, &lastSeen, &object.Status, &object.Lat, &object.Long)
		if err != nil {
			glog.Errorln("Error scaning tiger object , Err ", err)
			return objects, err
		}
		object.LastSeen = lastSeen.Unix()
		objects = append(objects, object)
	}
	return objects, nil
}

func GetUsersByTigerId(tigerId int) ([]EmailStruct, error) {
	rows, err := dbConfig.DB.Query(getTigerSightedUsers, tigerId, tigerId)
	if err != nil {
		glog.Errorln("Error GetAllTigers = Err ", err)

	}
	defer rows.Close()
	// Create a slice to store the results.
	objects := []EmailStruct{}
	// Iterate over the rows and scan the results into the slice.
	for rows.Next() {
		var object EmailStruct
		err := rows.Scan(&object.UserId, &object.UserName, &object.Email, &object.TigerName)
		if err != nil {
			glog.Errorln("Error scaning tiger object , Err ", err)
			return objects, err
		}
		objects = append(objects, object)
	}
	return objects, nil
}
