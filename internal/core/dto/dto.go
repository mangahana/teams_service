package dto

type CreateTeam struct {
	Name    string `json:"name"`
	TypeId  int    `json:"type_id"`
	OwnerId int    `json:"owner_id"`
}

type Update struct {
}
