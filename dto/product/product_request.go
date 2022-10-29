package productdto

type ProductRequest struct {
	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
	//Desc   string `json:"desc" gorm:"type:text" form:"desc"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	//Qty    int    `json:"qty" form:"qty" gorm:"type: int"`
	UserID int `json:"user_id" form:"user_id" gorm:"type: int"`
}

type UpdateProductRequest struct {
	Title string `json:"title" form:"title"`
	//Desc   string `json:"desc" form:"desc"`
	Price int    `json:"price" form:"price" `
	Image string `json:"image" form:"image" `
	//Qty    int    `json:"qty" form:"qty" `
	UserID int `json:"user_id" form:"user_id" `
}
