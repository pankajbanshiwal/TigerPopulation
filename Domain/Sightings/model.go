package Sightings

type Sighting struct {
	Id        int     `json:"id"`
	UserId    int     `json:"user_id"`
	TigerId   int     `json:"tiger_id"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"long"`
	SightTime int64   `json:"sight_time"` // unix timestamp
	ImageUrl  string  `json:"image_url"`
	Status    string  `json:"status"`
}
