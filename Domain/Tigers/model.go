package Tigers

type Tiger struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Dob      string  `json:"dob"`
	LastSeen int64   `json:"last_seen"` // unix timestamp
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
	Status   string  `json:"status"`
}

type StdResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EmailStruct struct {
	UserId    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	TigerName string `json:"tiger_name"`
	Location  string `json:"location"`
}
