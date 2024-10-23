package dto

type AddTeam struct {
	Name   string `json:"name" validate:"required,min=3,max=32"`
	TypeId int    `json:"type_id" validate:"required,number"`
}

type Update struct {
	TeamId      int    `json:"team_id" validate:"required,number"`
	Name        string `json:"name" validate:"required,min=3,max=32"`
	Description string `json:"description"`
}

type UploadPhoto struct {
	TeamId int    `json:"team_id" validate:"required,number"`
	Photo  string `json:"photo" validate:"required,base64"`
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

type UpdateMember struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Photo    string `json:"photo"`
}
