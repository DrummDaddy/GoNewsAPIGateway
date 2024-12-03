package models

// NewsFullDetailed детальная информация о новости

type NewsFullDetailed struct {
	ID      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"contetnt"`
}

//NewsShortDetailed  краткая информация о новости.
type NewsShortDetailed struct {
	ID    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}

// Comment комментарий к новости.
type Comments struct {
	ID          int    `json:"id" db:"id"`
	NewsID      int    `json:"news_id" db:"news_id"`
	ParentID    *int   `json:"parent_id,omitempty" db:"parent_id"`
	Text        string `json:"text" db:"text"`
	IsModerated bool   `json:"is_moderated" db:"is_moderated"`
}

//News

type News struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
