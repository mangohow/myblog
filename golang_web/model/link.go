package model


type Link struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Desc       string `json:"desc" db:"desc"`
	Url        string `json:"url" db:"url"`
	CategoryId int    `json:"categoryId" db:"category_id"`
	Icon       string `json:"icon" db:"icon"`
}

type LinkCategory struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
