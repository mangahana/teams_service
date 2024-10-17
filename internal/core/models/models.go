package models

type Session struct {
	UserID      int      `json:"user_id"`
	IsBanned    bool     `json:"is_banned"`
	Permissions []string `json:"permissions"`
}

type OneTeam struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Photo       *string `json:"photo"`
	IsVerified  bool    `json:"is_verified"`
	OwnerID     int     `json:"owner_id"`
}

type Member struct {
	Permissions []string `json:"-"`
	UserId      int      `json:"user_id"`
	UserName    string   `json:"user_name"`
	UserPhoto   string   `json:"user_photo"`
}

type TeamType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
