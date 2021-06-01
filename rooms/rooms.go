package rooms

import (
	"time"
)

type Rooms struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Status    int64     `json:"status"`
	CreatedBy int64     `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
