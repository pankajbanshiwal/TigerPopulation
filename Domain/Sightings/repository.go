package Sightings

import (
	"TigerPopulation/Utils"
	"TigerPopulation/Utils/dbConfig"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang/glog"
)

const (
	createSighting  = "INSERT INTO tiger_sightings  (user_id, tiger_id,loc,sight_time,image_url, created, updated) VALUES($1,$2,$3,$4,$5,$6,$7)"
	getLastSighting = `select loc[0],loc[1]  from tiger_sightings WHERE tiger_id=$1 AND status=$2  ORDER BY sight_time  desc  limit $3;`
	tigerSightings  = `select id,user_id,tiger_id,loc[0],loc[1],sight_time,image_url ,status  from tiger_sightings where status = $1 AND tiger_id = $2 ORDER BY sight_time  desc  limit $3 offset $4;`
)

func CreateTiger(req *Sighting) error {
	dateTime := Utils.GetDateTime()
	sightTimeTime := time.Unix(req.SightTime, 0)
	sightTime := sightTimeTime.Format("2006-01-02 15:04:05")

	_, err := dbConfig.DB.Exec(createSighting, req.UserId, req.TigerId, fmt.Sprintf("(%f,%f)", req.Lat, req.Long), sightTime, req.ImageUrl, dateTime, dateTime)
	if err != nil {
		glog.Errorln("\n", err.Error())
		return err
	}
	return nil
}

func GetLastSightingLocation(req *Sighting) (float64, float64, error) {
	var lat, long float64
	row := dbConfig.DB.QueryRow(getLastSighting, req.TigerId, "ACTIVE", 1)
	err := row.Scan(&lat, &long)
	if err == sql.ErrNoRows {
		return lat, long, err
	}
	return lat, long, err
}

func GetTigerSightings(tigerId, page int) ([]Sighting, error) {
	limit := 10
	offset := (page - 1) * limit
	rows, err := dbConfig.DB.Query(tigerSightings, "ACTIVE", tigerId, limit, offset)
	if err != nil {
		glog.Errorln("Error GetTigerSightings = Err ", err)

	}
	defer rows.Close()
	// Create a slice to store the results.
	objects := []Sighting{}
	// Iterate over the rows and scan the results into the slice.
	for rows.Next() {
		var object Sighting
		var sightTime time.Time
		err := rows.Scan(&object.Id, &object.UserId, &object.TigerId, &object.Lat, &object.Long, &sightTime, &object.ImageUrl, &object.Status)
		if err != nil {
			glog.Errorln("Error scaning tiger object , Err ", err)
			return objects, err
		}
		object.SightTime = sightTime.Unix()
		objects = append(objects, object)
	}
	return objects, nil
}
