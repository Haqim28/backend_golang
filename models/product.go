package models

type Product struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
	// Desc   string               `json:"desc" gorm:"type:text" form:"desc"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	// Qty    int                  `json:"qty" form:"qty"`
	UserID int                  `json:"user_id" form:"user_id"`
	User   UsersProfileResponse `json:"user"`
	//OrderID int                  `json:"order_id" form:"order_id"`
	//Order   []OrderResponse      `json:"order"`
}

type ProductResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	// Desc   string               `json:"desc"`
	Price int    `json:"price"`
	Image string `json:"image"`
	// Qty    int                  `json:"qty"`
	UserID int                  `json:"user_id"`
	User   UsersProfileResponse `json:"user"`
}

type ProductUserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	Qty    int    `json:"qty"`
	UserID int    `json:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}
