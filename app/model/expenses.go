package model

import (
	"fmt"
	"time"
	"github.com/icobani/RevelWebSite/app"
)

//TODO: Burdaki alanları netleştirelim.
type Expense struct {
	Id                int64     `json:"id" CaptionML:"enu=ID;trk=ID"`
	Type              string    `json:"type" sql:"type:varchar(20);" CaptionML:"enu=Type;trk=Tip"`
	Date              time.Time `json:"date" CaptionML:"enu=Date;trk=Tarih"`
	CurrencyCode      string    `json:"currency" CaptionML:"enu=Currency Code;trk=Döviz Kodu"`
	InvoiceNumber     string    `json:"bill_number" sql:"type:varchar(100);" CaptionML:"enu=Invoice Number;trk=Fiş/Fatura No."`
	MileageDistance   float32   `json:"mileage_km" CaptionML:"enu=Mileage Distange;trk=Mesafe"`
	CompanyId         int64 `json:"-" CaptionML:"enu=Company;trk=Şirket"`
	DepartmentId      int64 `json:"department" CaptionML:"enu=Department;trk=Departman"`
	ReportId          int64 `json:"-" CaptionML:"enu=Report;trk=Rapor"`
	UserId            int64 `json:"-" CaptionML:"enu=User;trk=Personel"`
	CategoryId        int64             `json:"-" CaptionML:"enu=Category;trk=Kategori"`
	Category          *SExpenseCategory `json:"category" sql:"-" ss:"-" CaptionML:"enu=Category;trk=Kategori"`
	Vat               float32    `json:"-" CaptionML:"enu=VAT %;trk=KDV %"`
	VatAmount         float32    `json:"-" CaptionML:"enu=VAT Amount;trk=Kdv Tutarı"`
	Amount            float32    `json:"amount" CaptionML:"enu=Amount;trk=Tutar"`
	AmountLCY         float32    `json:"-" CaptionML:"enu=Amount LCY;trk=Tutar (LPB)"`
	PaymentTypeId     int64      `json:"payment_type" CaptionML:"enu=Paymet Type;trk=Ödeme Türü"`
	Reimbursable      bool       `json:"reimbursable" CaptionML:"enu=Reimbursable;trk=Geri Ödenebilir"`
	Billable          bool       `json:"billable" CaptionML:"enu=Billable;trk=Faturalanabilir"`
	MerchantName      string     `sql:"type:varchar(250);" json:"merchant_name" CaptionML:"enu=Merchant Name;trk=Satıcı No."`
	MerchantTaxOffice string     `sql:"type:varchar(100);" json:"-" CaptionML:"enu=Merhant Tax Office;trk=Satıcı Vergi Dairesi"`
	MerchantTaxId     string     `sql:"type:varchar(100);" json:"-" CaptionML:"enu=Merchant Tax ID;trk=Satıcı Vergi No."`
	MileageCategoryId int64      `json:"-" CaptionML:"enu=;trk="`
	ProjectId         int64      `json:"project" CaptionML:"enu=;trk="`
	Country           string     `sql:"type:varchar(100);" json:"country" CaptionML:"enu=;trk="`
	MileageFrom       string     `sql:"type:varchar(250);" json:"mileage_from" CaptionML:"enu=;trk="`
	MileageTo         string     `sql:"type:varchar(250);" json:"mileage_to" CaptionML:"enu=;trk="`
	Note              string     `sql:"type:text" json:"note" CaptionML:"enu=;trk="`
	TagId             int64      `json:"tag" CaptionML:"enu=;trk="`
	TtHours           float32    `json:"tt_hours" CaptionML:"enu=;trk="`
	TtHoursRate       float32    `json:"tt_hours_rate" CaptionML:"enu=;trk="`
	DocumentUrl       string     `sql:"type:varchar(500);" json:"-" CaptionML:"enu=;trk="`
	DocumentUrl2      string     `sql:"type:varchar(500);" json:"-" CaptionML:"enu=;trk="`
	DocumentUrl3      string     `sql:"type:varchar(500);" json:"-" CaptionML:"enu=;trk="`
	ApprovalDate      *time.Time `json:"-" CaptionML:"enu=;trk="`
	ApprovedUser      int64      `json:"-" CaptionML:"enu=;trk="`
	Status            string `json:"status" sql:"type:varchar(20);" CaptionML:"enu=;trk="`
	GhostStatus       string `json:"-" sql:"type:varchar(20);" CaptionML:"enu=;trk="`
	Receipt           *Receipt `json:"receipt,omitempty" sql:"-" ss:"-" CaptionML:"enu=;trk="` //todo: make this visible on json response
	User              *ExpenseUser `json:"user,omitempty" ss:"-" CaptionML:"enu=;trk="`
	CreatedAt         time.Time `json:"created_at" CaptionML:"enu=;trk="`
	UpdatedAt         time.Time `json:"-" CaptionML:"enu=;trk="`
																																														 // DeletedAt time.Time `json:"-"`
}

func (this Expense) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Expense Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Expense Table Created")
	app.MakeCaptionML(this)
}

type Receipt struct {
	Id           int64 `json:"id"`
	UserId       int64 `json:"-"`
	ExpenseId    int64 `json:"-"` //todo: expense uygulanırken aitlik kontrolü yap

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
	app.MakeCaptionML(this)
}

type Import struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"-"`
	File      string    `json:"-"`
	Errors    string    `json:"errors" sql:"type:text"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

func (this Import) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Import Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Import Table Created")
	app.MakeCaptionML(this)
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
	app.MakeCaptionML(this)
}

