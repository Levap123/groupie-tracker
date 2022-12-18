package entity

type Group struct {
	Id           int      `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Image        string   `json:"image,omitempty"`
	Members      []string `json:"members,omitempty"`
	CreationDate int      `json:"creationDate,omitempty"`
	FirstAlbum   string   `json:"firstAlbum,omitempty"`
}

type GroupModel struct {
	Id           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Image        string `json:"image,omitempty"`
	CreationDate int    `json:"creationDate,omitempty"`
}
