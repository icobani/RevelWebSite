package model

import (
	"fmt"
	"log"
	"time"
	"github.com/icobani/RevelWebSite/app"
)

type Expense struct {
	Id           int64     `json:"id"`
	Type         string    `json:"type" sql:"type:varchar(20);"`
	Date         time.Time `json:"date"`
	CurrencyCode string    `json:"currency"`
	Kdv          float32   `json:"kdv"`
	BillNumber   string    `json:"bill_number" sql:"type:varchar(100);"`
	MileageKm    float32   `json:"mileage_km" `
	ReceiptId    int64     `json:"-"`

	CompanyId    int64 `json:"-"`
	DepartmentId int64 `json:"department"`
	ReportId     int64 `json:"-"`
	UserId       int64 `json:"-"`

	CategoryId int64             `json:"-"`
	Category   *SExpenseCategory `json:"category" sql:"-" ss:"-"`

	VatPercentage     float32    `json:"-"`
	VatAmount         float32    `json:"-"`
	Amount            float32    `json:"amount"`
	AmountLcy         float32    `json:"-"`
	PaymentTypeId     int64      `json:"payment_type"`
	Reimbursable      bool       `json:"reimbursable"`
	Billable          bool       `json:"billable"`
	MerchantName      string     `sql:"type:varchar(250);" json:"merchant_name"`
	MerchantTaxOffice string     `sql:"type:varchar(100);" json:"-"`
	MerchantTaxId     string     `sql:"type:varchar(100);" json:"-"`
	InvoiceNumber     string     `sql:"type:varchar(100);" json:"-"`
	MileageCategoryId int64      `json:"-"`
	ProjectId         int64      `json:"project"`
	Country           string     `sql:"type:varchar(100);" json:"country"`
	MileageFrom       string     `sql:"type:varchar(250);" json:"mileage_from"`
	MileageTo         string     `sql:"type:varchar(250);" json:"mileage_to"`
	MileageDistance   float32    `json:"-"`
	Note              string     `sql:"type:text" json:"note"`
	TagId             int64      `json:"tag"`
	TtHours           float32    `json:"tt_hours"`
	TtHoursRate       float32    `json:"tt_hours_rate"`
	DocumentUrl       string     `sql:"type:varchar(500);" json:"-"`
	DocumentUrl2      string     `sql:"type:varchar(500);" json:"-"`
	DocumentUrl3      string     `sql:"type:varchar(500);" json:"-"`
	ApprovalDate      *time.Time `json:"-"`
	ApprovedUser      int64      `json:"-"`

	Status      string `json:"status" sql:"type:varchar(20);"`
	GhostStatus string `json:"-" sql:"type:varchar(20);"`

	Receipt *Receipt `json:"receipt,omitempty" sql:"-" ss:"-"` //todo: make this visible on json response

	User *ExpenseUser `json:"user,omitempty" ss:"-"`

	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"-"`
	// DeletedAt time.Time `json:"-"`
}

func (this Expense) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Expense Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Expense Table Created")
}



type Receipt struct {
	Id        int64 `json:"id"`
	UserId    int64 `json:"-"`
	ExpenseId int64 `json:"-"` //todo: expense uygulanırken aitlik kontrolü yap

	OriginalName string `json:"original_name"`
	Name         string `json:"name"`
	Size         int64  `json:"size"`
	Url          string `json:"url" sql:"-" ss:"-"`
}

func (this Receipt) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Receipt Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Receipt Table Created")
}

type Import struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"-"`
	File      string    `json:"-"`
	Errors    string    `json:"errors" sql:"type:longtext"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

func (this Import) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Import Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Import Table Created")
}


type ExpenseUser struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	LastName        string `json:"last_name"`
	ProfileImage    string `json:"-"`
	ProfileImageUrl string `json:"profile_image_url"`
}

func (this ExpenseUser) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("ExpenseUser Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("ExpenseUser Table Created")
}

