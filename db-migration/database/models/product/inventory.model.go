package product

import (
	"time"

	"github.com/shopspring/decimal"
)

type Inventory struct {
	ID string `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`

	Supplier   Supplier `gorm:"foreignKey:SupplierID" json:"supplier"`
	SupplierID string   `gorm:"not null" json:"supplier_id" valid:"uuidv4"`

	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
	CategoryID string   `gorm:"not null" json:"category_id" valid:"uuidv4"`

	Name              string          `gorm:"not null" json:"name,omitempty" valid:"type(string), required~Name is required"`
	QuantytyInStock   uint            `gorm:"not null" json:"quantity_in_stock" valid:"type(uint), required~Stock quantity is required"`
	QuantitySold      uint            `json:"quantity_sold" valid:"type(uint)"`
	RestockThreshold  uint            `gorm:"not null" json:"restock_threshold" valid:"type(uint), required~Restock threshold is required"`
	PurchasePrice     decimal.Decimal `gorm:"not null" json:"purchase_price" valid:"type(float), required~Purchase price is required"`
	SalePrice         decimal.Decimal `gorm:"not null" json:"sale_price" valid:"type(float), required~Sale price is required"`
	Image             string          `json:"image,omitempty" valid:"type(string)"`	
	IsActive          bool            `json:"is_active"`
	IsDeleted         bool            `gorm:"default:false" json:"is_deleted"`
	WarehouseLocation string          `gorm:"not null" json:"warehouse_location" valid:"required~Warehouse location is required"`
	CreatedAt         time.Time       `json:"created_at"  valid:"-"`
	UpdatedAt         time.Time       `json:"updated_at" valid:"-"`
}

type Inventories []Inventory

func (Inventory) TableName() string {
	return "inventories"
}
