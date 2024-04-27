package responses

import "time"

type OrderResponse struct {
	Id        int       `json:"id"`
	OrderSum  int       `json:"ordersum"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `json:"user_id"`
}
