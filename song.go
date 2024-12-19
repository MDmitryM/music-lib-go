package musiclib

type Song struct {
	Artist string `json:"artist" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Album  string `json:"album"`
	Year   string `json:"year"`
}
