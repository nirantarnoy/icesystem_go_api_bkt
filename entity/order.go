package entity

type OrderLineStruct struct {
	OrderId      int     `json:"order_id"`
	ProductId    int     `json:"product_id"`
	Qty          float64 `json:"qty"`
	Price        float64 `json:"price"`
	PriceGroupId int     `json:"price_group_id"`
}

type OrderClose struct {
	RouteId       uint64 `json:"route_id" form:"route_id"`
	UserId        uint64 `json:"user_id" form:"user_id"`
	CompanyId     uint64 `json:"company_id" form:"company_id"`
	BranchId      uint64 `json:"branch_id"`
	IsReturnStock uint64 `json:"return_stock"`
}

type OrderCreate struct {
	CustomerId       uint64            `json:"customer_id" form:"customer_id" binding:"required"`
	UserId           uint64            `json:"user_id" form:"user_id" binding:"required"`
	EmpId            uint64            `json:"emp_id" form:"emp_id"`
	EmpId2           uint64            `json:"emp2_id" form:"emp2_id"`
	RouteId          uint64            `json:"route_id" form:"route_id"`
	CarId            uint64            `json:"car_id" form:"car_id"`
	PaymentTypeId    uint64            `json:"payment_type_id" form:"payment_type_id"`
	CompanyId        uint64            `json:"company_id" form:"company_id"`
	BranchId         uint64            `json:"branch_id"`
	DataList         []OrderLineStruct `json:"data"`
	Discount         float64           `json:"discount"`
	RouteCode        string            `json:"route_code"`
	RunNo            string            `json:"runno"`
	IssueId          uint64            `json:"issue_id"`
	OrderNo          string            `json:"order_no"`
	OrderTotalAmount float64           `json:"order_total_amount"`
	LoginShift       string            `json:"login_shift"`
	Image            string            `json:"image"`
	SaleTypeError    string            `json:"sale_type_error"`
	CustomerName     string            `json:"customer_name"`
}
