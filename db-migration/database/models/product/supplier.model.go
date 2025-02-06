package product

import "time"

type Supplier struct {
	ID              string    `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	SupplierName    string    `gorm:"not null" json:"supplier_name,omitempty" valid:"type(string), required~Supplier name is required"`
	PhoneNumber     string    `json:"phone_number,omitempty" valid:"type(string)"`
	TelephoneNumber string    `json:"telephone_number,omitempty" valid:"type(string)"`
	CreatedAt       time.Time `json:"created_at"  valid:"-"`
	UpdatedAt       time.Time `json:"updated_at" valid:"-"`
}

type Suppliers []Supplier

func (Supplier) TableName() string {
	return "suppliers"
}
