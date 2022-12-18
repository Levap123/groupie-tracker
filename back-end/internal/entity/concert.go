package entity

type Concert struct {
	Id        int                 `json:"id"`
	Locations map[string][]string `json:"datesLocations,omitempty"`
}
