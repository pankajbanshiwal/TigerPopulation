//go:build live
// +build live

package dbConfig

import (
	"flag"

	"github.com/golang/glog"
)

func init() {
	flag.Set("v", "2")
	flag.Set("stderrthreshold", "0")
	flag.Parse()
	glog.V(2).Infoln("Now its Live")
	SetEnvMode(STAGE)
	devConfig := ViperConfigStage()
	InitDB(devConfig.DbCreds.DbName, devConfig.DbCreds.DbUserName, devConfig.DbCreds.DbPassword, devConfig.DbCreds.DbHost, devConfig.DbCreds.DbPort)
}
