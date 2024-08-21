package students

import "time"

type Student struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedOn time.Time `json:"updated_on"`
}
