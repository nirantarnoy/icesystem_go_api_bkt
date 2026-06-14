package entity

import "time"

type OrderMaster struct {
	Id              int64
	OrderNo         string    `json:"order_no"`
	OrderDate       time.Time `json:"order_date"`
	OrderChannelId  int64     `json:"order_channel_id"`
	CustomerId      int64     `json:"customer_id"`
	SaleChannelId   int64     `json:"sale_channel_id"`
	CarRefId        int64     `json:"car_ref_id"`
	IssueId         int64     `json:"issue_id"`
	Status          int64     `json:"status"`
	CreatedBy       int64     `json:"created_by"`
	CompanyId       int64     `json:"company_id"`
	BranchId        int64     `json:"branch_id"`
	SaleFromMobile  int64     `json:"sale_from_mobile"`
	Emp_1           int64     `json:"emp_1"`
	Emp_2           int64     `json:"emp_2"`
	OrderDate2      time.Time `json:"order_date2"`
	OrderShift      int64     `json:"order_shift"`
	DiscountAmt     float64   `json:"discount_amt"`
	PaymentMethodId int64     `json:"payment_method_id"`
	PaymentStatus   int64     `json:"payment_status"`
	OrderTotalAmt   float64   `json:"order_total_amt"`
	Latlong         string    `json:"latlong"`
}
