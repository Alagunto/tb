package telegram

// Currency represents a currency supported by Telegram Payments.
type Currency struct {
	Code         string
	Title        string
	Symbol       string
	Native       string
	ThousandsSep string
	DecimalSep   string
	SymbolLeft   bool
	SpaceBetween bool
	Exp          int
	MinAmount    int64
	MaxAmount    int64
}

// SupportedCurrencies contains all currencies supported by Telegram Payments.
var SupportedCurrencies = map[string]Currency{
	"AED": {Code: "AED", Title: "United Arab Emirates Dirham", Symbol: "AED", Native: "د.إ.\u200f", ThousandsSep: ",", DecimalSep: ".", SymbolLeft: true, SpaceBetween: true, Exp: 2, MinAmount: 367, MaxAmount: 3673200},
	"AFN": {Code: "AFN", Title: "Afghan Afghani", Symbol: "AFN", Native: "\u060b", ThousandsSep: ",", DecimalSep: ".", SymbolLeft: true, SpaceBetween: false, Exp: 2, MinAmount: 7554, MaxAmount: 75540495},
	"ALL": {Code: "ALL", Title: "Albanian Lek", Symbol: "ALL", Native: "Lek", ThousandsSep: ".", DecimalSep: ",", SymbolLeft: false, SpaceBetween: false, Exp: 2, MinAmount: 10908, MaxAmount: 109085036},
	// Add other currencies as needed - moved from payments_data.go
}

