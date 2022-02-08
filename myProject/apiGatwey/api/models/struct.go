package models

type GetTask struct{
	Task Task `json:"task"`
}

type Task struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Summary string `json:"summary"`
	CreatedAt string `json:"createdAd"`
	User UserModel `json:"user"`

}

type UserModel struct{
	ID string `json:"id"`
	Name string `json:"name"`
}