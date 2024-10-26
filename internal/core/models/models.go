package models

type User struct {
	ID          int      `json:"id"`
	Username    string   `json:"username"`
	Photo       string   `json:"photo"`
	IsBanned    bool     `json:"is_banned"`
	Permissions []string `json:"permissions"`
}

type Team struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

type OneTeam struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	IsVerified  bool   `json:"is_verified"`
	IsModerated bool   `json:"is_moderated"`
	OwnerID     int    `json:"owner_id"`
}

type Member struct {
	Permissions []string `json:"-"`
	UserID      int      `json:"user_id"`
	Username    string   `json:"user_name"`
	UserPhoto   string   `json:"user_photo"`
}

type TeamType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
