package Sightings

import (
	"TigerPopulation/Domain/Sightings"
	"TigerPopulation/Utils"
	"TigerPopulation/Utils/MessageQueues"
	"database/sql"
	"errors"

	"github.com/golang/glog"
	"github.com/panjf2000/ants"
)

func CreateSighting(req *Sightings.Sighting) error {
	// check if user name exists
	lat, long, err := Sightings.GetLastSightingLocation(req)
	if err != nil && err != sql.ErrNoRows {
		glog.Errorln("Error checking last tiger sighting from db : Err", err)
		return err
	}
	if err != sql.ErrNoRows { // check distance from last sighting
		distance := Utils.CalculateDistanceBetweenLocations(lat, long, req.Lat, req.Long)
		// if distance is less than 5 km sighting should not be submitted
		if distance < 5 {
			glog.V(2).Infoln("Sighting within 5 km from last sighting : skipping")
			return errors.New("sighting within 5 km of its prev sighting is not allowed")
		}
		// allow submitting sighting

	}
	//CalculateDistanceBetweenLocations(lat1, lon1, lat2, lon2 float64) float64 {
	err = Sightings.CreateTiger(req)
	if err != nil {
		return err
	}
	// send email to all the users who have reported sighting for this tiger
	defer ants.Submit(func() { MessageQueues.SendEmailToUsers(req.TigerId) })

	glog.V(2).Infoln("Sightings created , details = ", req)
	return nil
}

func GetTigerSightings(tigerId, page int) ([]Sightings.Sighting, error) {
	data, err := Sightings.GetTigerSightings(tigerId, page)
	if err != nil {
		glog.Errorln("Error fetching GetAllTigers : Err", err)
		return nil, err
	}
	return data, nil

}
