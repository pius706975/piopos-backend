package product

import "time"

type Category struct {
	ID        string    `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	Name      string    `gorm:"not null" json:"name,omitempty" valid:"type(string), required~Category name is required"`
	CreatedAt time.Time `json:"created_at"  valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

type Categories []Category

func (Category) TableName() string {
	return "category"
}
