package models

// Comment комментарий к новости.
type Comment struct {
	ID          int    `json:"id" db:"id"`
	NewsID      int    `json:"news_id" db:"news_id"`
	ParentID    *int   `json:"parent_id,omitempty" db:"parent_id"`
	Text        string `json:"text" db:"text"`
	IsModerated bool   `json:"is_moderated" db:"is_moderated"`
}
