package flatten

import "time"

type History struct {
	Request   string
	Response  string
	Depth     int
	CreatedAt time.Time
}
