package dbConfig

type EnvType int

const (
	LOCAL EnvType = 0
	DEV   EnvType = 1
	STAGE EnvType = 2
	LIVE  EnvType = 3
)

type DevConfig struct {
	Env            string
	Port           string
	UserContextKey string
	JwtSecretKey   string
	DbCreds        *DbConfig
}

type DbConfig struct {
	DbName     string
	DbUserName string
	DbPassword string
	DbPort     int
	DbHost     string
}
