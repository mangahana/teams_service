package dto

type AddTeam struct {
	Name    string `json:"name" validate:"required,min=3,max=32"`
	TypeId  int    `json:"type_id" validate:"required,number"`
	OwnerId int    `json:"owner_id"`
}

type Update struct {
}
