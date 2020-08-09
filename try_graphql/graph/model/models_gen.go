// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewVideo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      int    `json:"userId"`
	URL         string `json:"url"`
}

type Screenshot struct {
	ID      int    `json:"id"`
	VideoID int    `json:"videoId"`
	URL     string `json:"url"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
