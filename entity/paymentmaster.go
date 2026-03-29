package entity

import "time"

type PaymentMaster struct {
	Id         uint64    `json:"id"`
	TransDate  time.Time `json:"trans_date"`
	CustomerId uint64    `json:"customer_id"`
	JournalNo  string    `json:"journal_no"`
	Status     uint64    `json:"status"`
	CompanyId  uint64    `json:"company_id"`
	BranchId   uint64    `json:"branch_id"`
	CratedBy   uint64    `json:"crated_by"`
	CreatedAt  uint64    `json:"created_at"`
	SlipDoc    string    `json:"slip_doc"`
}

type SlipDoc struct {
	Image []string `json:"image"`
}
