package models

type OneTeam struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Photo       *string `json:"photo"`
	IsVerified  bool    `json:"is_verified"`
}

type Member struct {
	UserId    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	UserPhoto string `json:"user_photo"`
}
