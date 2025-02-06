package product

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID string `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`

	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
	CategoryID string   `gorm:"not null" json:"category_id" valid:"uuidv4"`

	Inventory   Inventory `gorm:"foreignKey:InventoryID" json:"inventory"`
	InventoryID string    `gorm:"not null" json:"inventory_id" valid:"uuidv4"`

	Name        string `gorm:"not null" json:"name,omitempty" valid:"type(string), required~Product name is required"`
	Description string `json:"description,omitempty" valid:"-"`

	SalePrice       decimal.Decimal `gorm:"not null" json:"sale_price" valid:"type(float), required~Sale price is required"`
	QuantytyInStock uint            `gorm:"not null" json:"quantity_in_stock" valid:"type(uint), required~Stock quantity is required"`
	IsActive        bool            `json:"is_active"`
	IsDeleted       bool            `gorm:"default:false" json:"is_deleted"`
	CreatedAt       time.Time       `json:"created_at" valid:"-"`
	UpdatedAt       time.Time       `json:"updated_at" valid:"-"`
}

type Products []Product

func (Product) TableName() string {
	return "products"
}
