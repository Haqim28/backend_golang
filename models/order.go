package models

type Cart struct {
	ID         int       `json:"id" gorm:""`
	Product_ID []int     `json:"-" form:"product_id" gorm:"-"`
	Product    []Product `json:"product" gorm:"many2many:product_cart;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Qty        int       `json:"qty" form"qty"`
}
