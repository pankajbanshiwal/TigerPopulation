package Users

import (
	"TigerPopulation/Domain/Users"
	"TigerPopulation/Utils/accessToken"
	"errors"

	"github.com/golang/glog"
)

func CreateUser(req *Users.CreateUser) error {
	// check if user name exists
	exists, err := Users.CheckIfUserNameExists(req.UserName)
	if err != nil {
		glog.Errorln("Error checking CheckIfUserNameExists : Err", err)
		return err
	}
	if exists {
		glog.V(2).Infoln("Email id already exists")
		return errors.New("duplicate Username")
	}
	// check if email id exists
	exists, err = Users.CheckIfEmailIdExists(req.Email)
	if err != nil {
		glog.Errorln("Error checking CheckIfEmailIdExists : Err", err)
		return err
	}
	if exists {
		glog.V(2).Infoln("Email id already exists")
		return errors.New("duplicate EmailId")
	}
	err = Users.CreateDbUser(req)
	if err != nil {
		return err
	}
	glog.V(2).Infoln("User created , details = ", req)
	return nil
}

func Login(req *Users.CreateUser) (*Users.CreateUser, error) {
	// first check if username or email id exists
	exists, err := Users.CheckIfUserNameExists(req.UserName)
	if err != nil {
		glog.Errorln("Error checking CheckIfUserNameExists : Err", err)
		return nil, err
	}
	// check if email id exists
	emailExists, err := Users.CheckIfEmailIdExists(req.Email)
	if err != nil {
		glog.Errorln("Error checking CheckIfEmailIdExists : Err", err)
		return nil, err
	}
	if !emailExists && !exists {
		glog.V(2).Infoln("Email id or username does not exists")
		return nil, errors.New("email id or username does not exists")
	}
	// log in is supported by username + password or emailid + password
	user, err := Users.CheckLogInDetails(req)
	if err != nil {
		glog.Errorln("Incorrect username or password : Err", err)
		return nil, errors.New("incorrect username or password ")
	}
	// now generate token and store it in db
	token, err := accessToken.GenerateToken(user.Id)
	if err != nil {
		glog.Errorln("Error creating access token GenerateToken : Err", err)
		return nil, err
	}
	// update token in db
	updated := Users.UpdateUserAuthToken(user.Id, token)
	if updated != nil {
		glog.Errorln("Error updating access token UpdateUserAuthToken : Err", updated)
		return nil, updated

	}
	user.Token = token
	return user, nil
}
