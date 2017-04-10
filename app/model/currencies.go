package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
	"github.com/icobani/RevelWebSite/app/modelViews"
)

type Currency struct {
	Code     string `json:"code" option:"id" gorm:"primary_key" sql:"type:varchar(10);" CaptionML:"enu=Code;trk=Code"`
	FullName string `json:"-" option:"value" CaptionML:"enu=Full Name;trk=Tam Adı"`
	Symbol     string `json:"name" sql:"type:varchar(50);" CaptionML:"enu=Symbol;trk=Sembol"`
}

func (this Currency) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Currency Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Currency Table Created")
	app.MakeCaptionML(this)
}

// Combolara çıkacak olan değerler burada hazırlanıyor.
func (this Currency) GetComboValues(user User, master *modelViews.ModelReferance) []modelViews.ComboItem {
	var CurrenciesList []Currency
	var ComboItems []modelViews.ComboItem
	app.DB.Select("code, full_name").Order("code").Find(&CurrenciesList)

	for _, item := range CurrenciesList {
		ComboItems = append(ComboItems, modelViews.ComboItem{Code:item.Code, Value:item.FullName})
	}
	return ComboItems
}


func (this Currency) InitTable() {

	sql := `
		DELETE From Currencies;
		/* INSERT QUERY NO: 1 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'AED', 'د.إ;', 'UAE dirham'
		);

		/* INSERT QUERY NO: 2 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'AFN', 'Afs', 'Afghan afghani'
		);

		/* INSERT QUERY NO: 3 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'ALL', 'L', 'Albanian lek'
		);

		/* INSERT QUERY NO: 4 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'AMD', 'AMD', 'Armenian dram'
		);

		/* INSERT QUERY NO: 5 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'ANG', 'NAƒ', 'Netherlands Antillean gulden'
		);

		/* INSERT QUERY NO: 6 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'AOA', 'Kz', 'Angolan kwanza'
		);

		/* INSERT QUERY NO: 7 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'ARS', '$', 'Argentine peso'
		);

		/* INSERT QUERY NO: 8 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'AUD', '$', 'Australian dollar'
		);

		/* INSERT QUERY NO: 9 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'AWG', 'ƒ', 'Aruban florin'
		);

		/* INSERT QUERY NO: 10 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'AZN', 'AZN', 'Azerbaijani manat'
		);

		/* INSERT QUERY NO: 11 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BAM', 'KM', 'Bosnia and Herzegovina konvertibilna marka'
		);

		/* INSERT QUERY NO: 12 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BBD', 'Bds$', 'Barbadian dollar'
		);

		/* INSERT QUERY NO: 13 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BDT', '৳', 'Bangladeshi taka'
		);

		/* INSERT QUERY NO: 14 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BGN', 'BGN', 'Bulgarian lev'
		);

		/* INSERT QUERY NO: 15 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BHD', '.د.ب', 'Bahraini dinar'
		);

		/* INSERT QUERY NO: 16 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BIF', 'FBu', 'Burundi franc'
		);

		/* INSERT QUERY NO: 17 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BMD', 'BD$', 'Bermudian dollar'
		);

		/* INSERT QUERY NO: 18 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BND', 'B$', 'Brunei dollar'
		);

		/* INSERT QUERY NO: 19 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BOB', 'Bs.', 'Bolivian boliviano'
		);

		/* INSERT QUERY NO: 20 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BRL', 'R$', 'Brazilian real'
		);

		/* INSERT QUERY NO: 21 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BSD', 'B$', 'Bahamian dollar'
		);

		/* INSERT QUERY NO: 22 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BTN', 'Nu.', 'Bhutanese ngultrum'
		);

		/* INSERT QUERY NO: 23 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BWP', 'P', 'Botswana pula'
		);

		/* INSERT QUERY NO: 24 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BYR', 'Br', 'Belarusian ruble'
		);

		/* INSERT QUERY NO: 25 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'BZD', 'BZ$', 'Belize dollar'
		);

		/* INSERT QUERY NO: 26 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'CAD', '$', 'Canadian dollar'
		);

		/* INSERT QUERY NO: 27 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'CDF', 'F', 'Congolese franc'
		);

		/* INSERT QUERY NO: 28 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'CHF', 'Fr.', 'Swiss franc'
		);

		/* INSERT QUERY NO: 29 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'CLP', '$', 'Chilean peso'
		);

		/* INSERT QUERY NO: 30 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'CNY', '¥', 'Chinese/Yuan renminbi'
		);

		/* INSERT QUERY NO: 31 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'COP', 'Col$', 'Colombian peso'
		);

		/* INSERT QUERY NO: 32 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'CRC', '₡', 'Costa Rican colon'
		);

		/* INSERT QUERY NO: 33 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'CUC', '$', 'Cuban peso'
		);

		/* INSERT QUERY NO: 34 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'CVE', 'Esc', 'Cape Verdean escudo'
		);

		/* INSERT QUERY NO: 35 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'CZK', 'Kč', 'Czech koruna'
		);

		/* INSERT QUERY NO: 36 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'DJF', 'Fdj', 'Djiboutian franc'
		);

		/* INSERT QUERY NO: 37 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'DKK', 'Kr', 'Danish krone'
		);

		/* INSERT QUERY NO: 38 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'DOP', 'RD$', 'Dominican peso'
		);

		/* INSERT QUERY NO: 39 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'DZD', 'د.ج', 'Algerian dinar'
		);

		/* INSERT QUERY NO: 40 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'EEK', 'KR', 'Estonian kroon'
		);

		/* INSERT QUERY NO: 41 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'EGP', '£', 'Egyptian pound'
		);

		/* INSERT QUERY NO: 42 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'ERN', 'Nfa', 'Eritrean nakfa'
		);

		/* INSERT QUERY NO: 43 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'ETB', 'Br', 'Ethiopian birr'
		);

		/* INSERT QUERY NO: 44 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'EUR', '€', 'European Euro'
		);

		/* INSERT QUERY NO: 45 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'FJD', 'FJ$', 'Fijian dollar'
		);

		/* INSERT QUERY NO: 46 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'FKP', '£', 'Falkland Islands pound'
		);

		/* INSERT QUERY NO: 47 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'GBP', '£', 'British pound'
		);

		/* INSERT QUERY NO: 48 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'GEL', 'GEL', 'Georgian lari'
		);

		/* INSERT QUERY NO: 49 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'GHS', 'GH₵', 'Ghanaian cedi'
		);

		/* INSERT QUERY NO: 50 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'GIP', '£', 'Gibraltar pound'
		);

		/* INSERT QUERY NO: 51 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'GMD', 'D', 'Gambian dalasi'
		);

		/* INSERT QUERY NO: 52 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'GNF', 'FG', 'Guinean franc'
		);

		/* INSERT QUERY NO: 53 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'GQE', 'CFA', 'Central African CFA franc'
		);

		/* INSERT QUERY NO: 54 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'GTQ', 'Q', 'Guatemalan quetzal'
		);

		/* INSERT QUERY NO: 55 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'GYD', 'GY$', 'Guyanese dollar'
		);

		/* INSERT QUERY NO: 56 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'HKD', 'HK$', 'Hong Kong dollar'
		);

		/* INSERT QUERY NO: 57 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'HNL', 'L', 'Honduran lempira'
		);

		/* INSERT QUERY NO: 58 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'HRK', 'kn', 'Croatian kuna'
		);

		/* INSERT QUERY NO: 59 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'HTG', 'G', 'Haitian gourde'
		);

		/* INSERT QUERY NO: 60 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'HUF', 'Ft', 'Hungarian forint'
		);

		/* INSERT QUERY NO: 61 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'IDR', 'Rp', 'Indonesian rupiah'
		);

		/* INSERT QUERY NO: 62 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'ILS', '₪', 'Israeli new sheqel'
		);

		/* INSERT QUERY NO: 63 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'INR', '₹', 'Indian rupee'
		);

		/* INSERT QUERY NO: 64 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'IQD', 'د.ع', 'Iraqi dinar'
		);

		/* INSERT QUERY NO: 65 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'IRR', 'IRR', 'Iranian rial'
		);

		/* INSERT QUERY NO: 66 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'ISK', 'kr', 'Icelandic króna'
		);

		/* INSERT QUERY NO: 67 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'JMD', 'J$', 'Jamaican dollar'
		);

		/* INSERT QUERY NO: 68 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'JOD', 'JOD', 'Jordanian dinar'
		);

		/* INSERT QUERY NO: 69 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'JPY', '¥', 'Japanese yen'
		);

		/* INSERT QUERY NO: 70 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'KES', 'KSh', 'Kenyan shilling'
		);

		/* INSERT QUERY NO: 71 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'KGS', 'сом', 'Kyrgyzstani som'
		);

		/* INSERT QUERY NO: 72 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'KHR', '៛', 'Cambodian riel'
		);

		/* INSERT QUERY NO: 73 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'KMF', 'KMF', 'Comorian franc'
		);

		/* INSERT QUERY NO: 74 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'KPW', 'W', 'North Korean won'
		);

		/* INSERT QUERY NO: 75 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'KRW', 'W', 'South Korean won'
		);

		/* INSERT QUERY NO: 76 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'KWD', 'KWD', 'Kuwaiti dinar'
		);

		/* INSERT QUERY NO: 77 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'KYD', 'KY$', 'Cayman Islands dollar'
		);

		/* INSERT QUERY NO: 78 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'KZT', 'T', 'Kazakhstani tenge'
		);

		/* INSERT QUERY NO: 79 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'LAK', 'KN', 'Lao kip'
		);

		/* INSERT QUERY NO: 80 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'LBP', '£', 'Lebanese lira'
		);

		/* INSERT QUERY NO: 81 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'LKR', 'Rs', 'Sri Lankan rupee'
		);

		/* INSERT QUERY NO: 82 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'LRD', 'L$', 'Liberian dollar'
		);

		/* INSERT QUERY NO: 83 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'LSL', 'M', 'Lesotho loti'
		);

		/* INSERT QUERY NO: 84 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'LTL', 'Lt', 'Lithuanian litas'
		);

		/* INSERT QUERY NO: 85 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'LVL', 'Ls', 'Latvian lats'
		);

		/* INSERT QUERY NO: 86 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'LYD', 'LD', 'Libyan dinar'
		);

		/* INSERT QUERY NO: 87 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MAD', 'MAD', 'Moroccan dirham'
		);

		/* INSERT QUERY NO: 88 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MDL', 'MDL', 'Moldovan leu'
		);

		/* INSERT QUERY NO: 89 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MGA', 'FMG', 'Malagasy ariary'
		);

		/* INSERT QUERY NO: 90 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MKD', 'MKD', 'Macedonian denar'
		);

		/* INSERT QUERY NO: 91 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MMK', 'K', 'Myanma kyat'
		);

		/* INSERT QUERY NO: 92 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MNT', '₮', 'Mongolian tugrik'
		);

		/* INSERT QUERY NO: 93 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MOP', 'P', 'Macanese pataca'
		);

		/* INSERT QUERY NO: 94 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MRO', 'UM', 'Mauritanian ouguiya'
		);

		/* INSERT QUERY NO: 95 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MUR', 'Rs', 'Mauritian rupee'
		);

		/* INSERT QUERY NO: 96 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MVR', 'Rf', 'Maldivian rufiyaa'
		);

		/* INSERT QUERY NO: 97 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MWK', 'MK', 'Malawian kwacha'
		);

		/* INSERT QUERY NO: 98 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MXN', '$', 'Mexican peso'
		);

		/* INSERT QUERY NO: 99 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MYR', 'RM', 'Malaysian ringgit'
		);

		/* INSERT QUERY NO: 100 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'MZM', 'MTn', 'Mozambican metical'
		);

		/* INSERT QUERY NO: 101 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'NAD', 'N$', 'Namibian dollar'
		);

		/* INSERT QUERY NO: 102 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'NGN', '₦', 'Nigerian naira'
		);

		/* INSERT QUERY NO: 103 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'NIO', 'C$', 'Nicaraguan córdoba'
		);

		/* INSERT QUERY NO: 104 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'NOK', 'kr', 'Norwegian krone'
		);

		/* INSERT QUERY NO: 105 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'NPR', 'NRs', 'Nepalese rupee'
		);

		/* INSERT QUERY NO: 106 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'NZD', 'NZ$', 'New Zealand dollar'
		);

		/* INSERT QUERY NO: 107 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'OMR', 'OMR', 'Omani rial'
		);

		/* INSERT QUERY NO: 108 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'PAB', 'B./', 'Panamanian balboa'
		);

		/* INSERT QUERY NO: 109 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'PEN', 'S/.', 'Peruvian nuevo sol'
		);

		/* INSERT QUERY NO: 110 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'PGK', 'K', 'Papua New Guinean kina'
		);

		/* INSERT QUERY NO: 111 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'PHP', '₱', 'Philippine peso'
		);

		/* INSERT QUERY NO: 112 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'PKR', 'Rs.', 'Pakistani rupee'
		);

		/* INSERT QUERY NO: 113 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'PLN', 'zł', 'Polish zloty'
		);

		/* INSERT QUERY NO: 114 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'PYG', '₲', 'Paraguayan guarani'
		);

		/* INSERT QUERY NO: 115 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'QAR', 'QR', 'Qatari riyal'
		);

		/* INSERT QUERY NO: 116 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'RON', 'L', 'Romanian leu'
		);

		/* INSERT QUERY NO: 117 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'RSD', 'din.', 'Serbian dinar'
		);

		/* INSERT QUERY NO: 118 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'RUB', 'R', 'Russian ruble'
		);

		/* INSERT QUERY NO: 119 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SAR', 'SR', 'Saudi riyal'
		);

		/* INSERT QUERY NO: 120 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SBD', 'SI$', 'Solomon Islands dollar'
		);

		/* INSERT QUERY NO: 121 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SCR', 'SR', 'Seychellois rupee'
		);

		/* INSERT QUERY NO: 122 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SDG', 'SDG', 'Sudanese pound'
		);

		/* INSERT QUERY NO: 123 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SEK', 'kr', 'Swedish krona'
		);

		/* INSERT QUERY NO: 124 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SGD', 'S$', 'Singapore dollar'
		);

		/* INSERT QUERY NO: 125 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SHP', '£', 'Saint Helena pound'
		);

		/* INSERT QUERY NO: 126 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SLL', 'Le', 'Sierra Leonean leone'
		);

		/* INSERT QUERY NO: 127 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SOS', 'Sh.', 'Somali shilling'
		);

		/* INSERT QUERY NO: 128 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SRD', '$', 'Surinamese dollar'
		);

		/* INSERT QUERY NO: 129 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SYP', 'LS', 'Syrian pound'
		);

		/* INSERT QUERY NO: 130 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'SZL', 'E', 'Swazi lilangeni'
		);

		/* INSERT QUERY NO: 131 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'THB', '฿', 'Thai baht'
		);

		/* INSERT QUERY NO: 132 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'TJS', 'TJS', 'Tajikistani somoni'
		);

		/* INSERT QUERY NO: 133 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'TMT', 'm', 'Turkmen manat'
		);

		/* INSERT QUERY NO: 134 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'TND', 'DT', 'Tunisian dinar'
		);

		/* INSERT QUERY NO: 135 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'TRY', '₺', 'Turkish new lira'
		);

		/* INSERT QUERY NO: 136 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'TTD', 'TT$', 'Trinidad and Tobago dollar'
		);

		/* INSERT QUERY NO: 137 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'TWD', 'NT$', 'New Taiwan dollar'
		);

		/* INSERT QUERY NO: 138 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'TZS', 'TZS', 'Tanzanian shilling'
		);

		/* INSERT QUERY NO: 139 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'UAH', 'UAH', 'Ukrainian hryvnia'
		);

		/* INSERT QUERY NO: 140 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'UGX', 'USh', 'Ugandan shilling'
		);

		/* INSERT QUERY NO: 141 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'USD', 'US$', 'United States dollar'
		);

		/* INSERT QUERY NO: 142 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'UYU', '$U', 'Uruguayan peso'
		);

		/* INSERT QUERY NO: 143 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'UZS', 'UZS', 'Uzbekistani som'
		);

		/* INSERT QUERY NO: 144 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'VEB', 'Bs', 'Venezuelan bolivar'
		);

		/* INSERT QUERY NO: 145 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'VND', '₫', 'Vietnamese dong'
		);

		/* INSERT QUERY NO: 146 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'VUV', 'VT', 'Vanuatu vatu'
		);

		/* INSERT QUERY NO: 147 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'WST', 'WS$', 'Samoan tala'
		);

		/* INSERT QUERY NO: 148 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'XAF', 'CFA', 'Central African CFA franc'
		);

		/* INSERT QUERY NO: 149 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'XCD', 'EC$', 'East Caribbean dollar'
		);

		/* INSERT QUERY NO: 150 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'XDR', 'SDR', 'Special Drawing Rights'
		);

		/* INSERT QUERY NO: 151 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'XOF', 'CFA', 'West African CFA franc'
		);

		/* INSERT QUERY NO: 152 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'XPF', 'F', 'CFP franc'
		);

		/* INSERT QUERY NO: 153 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'YER', 'YER', 'Yemeni rial'
		);

		/* INSERT QUERY NO: 154 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'ZAR', 'R', 'South African rand'
		);

		/* INSERT QUERY NO: 155 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'ZMK', 'ZK', 'Zambian kwacha'
		);

		/* INSERT QUERY NO: 156 */
		INSERT INTO Currencies(Code, Symbol, Full_Name)
		VALUES
		(
		'ZWR', 'Z$', 'Zimbabwean dollar'
		);


	`

	app.DB.Exec(sql)


}