package Tigers

import (
	"TigerPopulation/Domain/Tigers"
	"errors"

	"github.com/golang/glog"
)

func CreateTiger(req *Tigers.Tiger) error {
	// check if user name exists
	exists, err := Tigers.CheckIfTigerExists(req.Name, req.Dob)
	if err != nil {
		glog.Errorln("Error checking CheckIfTigerNameExists : Err", err)
		return err
	}
	if exists {
		glog.V(2).Infoln("Tiger with same name & dob already exists")
		return errors.New("duplicate tiger")
	}
	err = Tigers.CreateTiger(req)
	if err != nil {
		return err
	}
	glog.V(2).Infoln("Tiger created , details = ", req)
	return nil
}

func GetAllTigers(page int) ([]Tigers.Tiger, error) {
	data, err := Tigers.GetAllTigers(page)
	if err != nil {
		glog.Errorln("Error fetching GetAllTigers : Err", err)
		return nil, err
	}
	return data, nil

}
