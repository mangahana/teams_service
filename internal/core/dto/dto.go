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

type UploadPhoto struct {
	MemberId int
	TeamId   int    `json:"team_id" validate:"required,number"`
	Photo    string `json:"photo" validate:"required,base64"`
}

type UpdateMemberPermissions struct {
	OwnerId     int
	TeamId      int      `json:"team_id" validate:"required,number"`
	MemberId    int      `json:"member_id" validate:"required,number"`
	Permissions []string `json:"permissions" validate:"required"`
}

type CreateInvite struct {
	OwnerId int
	TeamId  int `json:"team_id" validate:"required,number"`
	UserId  int `json:"user_id" validate:"required,number"`
}
