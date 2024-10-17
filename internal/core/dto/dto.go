package dto

type AddTeam struct {
	Name    string `json:"name" validate:"required,min=3,max=32"`
	TypeId  int    `json:"type_id" validate:"required,number"`
	OwnerId int    `json:"owner_id"`
}

type Update struct {
	MemberId    int
	TeamId      int    `json:"team_id" validate:"required,number"`
	Name        string `json:"name" validate:"required,min=3,max=32"`
	Description string `json:"description"`
}
