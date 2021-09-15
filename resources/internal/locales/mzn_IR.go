// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_mzn_IR() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "mzn_IR",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ژانویه", Feb: "فوریه", Mar: "مارس", Apr: "آوریل", May: "مه", Jun: "ژوئن", Jul: "ژوئیه", Aug: "اوت", Sep: "سپتامبر", Oct: "اکتبر", Nov: "نوامبر", Dec: "دسامبر"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ژانویه", Feb: "فوریه", Mar: "مارس", Apr: "آوریل", May: "مه", Jun: "ژوئن", Jul: "ژوئیه", Aug: "اوت", Sep: "سپتامبر", Oct: "اکتبر", Nov: "نوامبر", Dec: "دسامبر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "متحده عربی امارات ِدرهم", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "افغانستون ِافغانی", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "آلبانی ِلک", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "ارمنستون درهم", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "هلند ِآنتیل ِجزایر ِگویلدر", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "آنگولای ِکوانزا", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "آرژانتین ِپزو", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "آروبای ِفلورن", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "آذربایجون ِمنات", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "بوسنی و هرزگوین ِتبدیل\u200cبَیی مارک", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "باربادوس ِدولار", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "بنگلادش ِتاکا", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "بلغارستون ِلیوا", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "بحرین ِدینار", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "بوروندی ِفرانک", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "برمودای ِدولار", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "برونئی ِدولار", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "بولیوی ِبولیویانو", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "برزیل ِرئال", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "باهامای ِدولار", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "بوتان ِنگولتروم", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "بوتساوانای ِپولا", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "بلاروس ِروبل", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "بلاروس ِروبل (۲۰۰۰–۲۰۱۶)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "بلیز ِدولار", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "کانادای ِدولار", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "کنگوی ِفرانک", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "سوییس ِفرانک", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "شیلی ِپزو", Symbol: "CLP"},
				currency.CNY: cldr.Currency{DisplayName: "چین ِیوآن", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "کلمبیای ِپزو", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "کاستاریکای ِکولون", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "کوبای ِتبدیل\u200cبَیی پزو", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "کوبای ِپزو", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "عاج ِساحل ِایسکودو", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "چک ِکرون", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "جیبوتی ِفرانک", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "دانمارک ِکورن", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "دومینیکن ِپزو", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "الجزیره\u200cی ِدینار", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "مصر ِپوند", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "اریتره\u200cی ِناکفا", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "اتیوپی ِبیر", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "یورو", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "فالکلند ِجزایر ِپوند", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "بریتانیای ِپوند", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "گرجستون ِلاری", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "غنای ِسدی", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "جبل\u200cطارق ِپوند", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "گامبیای ِدالاسی", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "گینه\u200cی ِفرانک", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "گواتمالا کتزال", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "گویان ِدولار", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "هونگ کونگ ِدولار", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "هندوراس ِلمپیرا", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "کرواسی ِکونا", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "هائیتی ِگورد", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "مجارستون ِفروینت", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "اندونزی ِروپیه", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "اسراییل ِنو شِکِل", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "هند ِروپیه", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "عراق ِدینار", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ایران ریال", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "ایسلند کرونا", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "جاماییکای ِدولار", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "اردن ِدینار", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "جاپون ِین", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "کنیای ِشیلینگ", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "قرقیزستون ِسام", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "کامبوج ِریل", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "کامرون ِفرانک", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "شمالی کره\u200cی ِوون", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "جنوبی کُره\u200cی ِوون", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "کویت ِدینار", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "کایمن جزیره\u200cی ِدولار", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "قراقستون ِتنگ", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "لائوس ِکیپ", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "لبنان ِپوند", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "سریلانکا روپیه", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "لیبریای ِدولار", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "لیبی ِدینار", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "مراکش ِدرهم", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "مولداوی ِلئو", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ماداگاسکار ِآریاری", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "مقدونیه\u200cی ِدینار", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "میانمار ِکیات", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "مغلستون ِتوگریک", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "ماکائو ِپاتاجا", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "موریتانی ِاوگوئیا (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "موریتانی ِاوگوئیا", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "موریتیان ِروپیه", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "مالدیو ِروفیا", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "مالاوی ِکواچا", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "مکزیک ِپزو", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "مالزی ِرینگیت", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "موزامبیک متیکال", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "نامبیای ِدولار", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "نیجریه\u200cی ِنیارا", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "نیکاراگوئه\u200cی ِکوردوبا", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "نروژ ِکرون", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "نپال ِروپیه", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "عمان ِریال", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "پانامای ِبالبوا", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "پروی ِسول", Symbol: "PEN"},
				currency.PHP: cldr.Currency{DisplayName: "فیلیپین ِپزو", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "پاکستون روپیه", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "لهستون ِزلوتی", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "پاراگوئه\u200cی ِگوارانی", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "قطر ِریال", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "رومانی ِلئو", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "صربستون ِدینار", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "روسیه\u200cی ِروبل", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "روآندای ِفرانک", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "عربستون ِریال", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "سیشل ِروپیه", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "سودان ِپوند", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "سوئد ِکرون", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "سنگاپور ِدلار", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "سنت هلنای ِپوند", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "سیرالئون ِلئون", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "سومالی ِشیلینگ", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "سورینام ِدولار", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "جنوبی سودان ِپوند", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "سائوتومه و پرینسیپ ِدوبرا (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "سائوتومه و پرینسیپ ِدوبرا", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "سوریه\u200cی ِپوند", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "سوازیلند ِلیلانجنی", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "تایلند ِبات", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "تاجیکستون ِسامانی", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "ترکمنستون ِمنات", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "تونس ِدینار", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "ترکیه\u200cی ِلیره", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ترینیداد و توباگوی ِدولار", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "جدید ِتایوان ِدولار", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "تانزانیای ِشیلینگ", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "اکراین ِگریونا", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "اوگاندای ِشیلینگ", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "آمریکای ِدولار", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "اروگوئه\u200cی ِپزو", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "ازبکستون ِسام", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "ونزوئلایِ بولیوار (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "ونزوئلایِ بولیوار", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "ویتنام ِدنگ", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "میونی آفریقای ِسی\u200cاف\u200cای فرانک", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "شرقی کاراییب ِدولار", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "غربی آفریقای ِسی\u200cاف\u200cای فرانک", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "یمن ِریال", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "جنوبی آفریقای ِراند", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "زامبیای ِکواچا", Symbol: "ZMW"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AB:      "آبخازی",
			language.AF:      "آفریکانس",
			language.AGQ:     "آقم",
			language.AK:      "آکان",
			language.AM:      "امهری",
			language.AR:      "عربی",
			language.AR_001:  "مدرن استاندارد عربی",
			language.ARN:     "ماپوچه",
			language.AS:      "آسامی",
			language.ASA:     "آسو",
			language.AZ:      "آذری",
			language.AZ_ARAB: "جنوبی آذری ترکی",
			language.BA:      "باشقیری",
			language.BE:      "بلاروسی",
			language.BEM:     "بمبایی",
			language.BEZ:     "بنایی",
			language.BG:      "بلغاری",
			language.BGN:     "غربی بلوچی",
			language.BM:      "بامبارایی",
			language.BN:      "بنگالی",
			language.BO:      "تبتی",
			language.BR:      "برِتونی",
			language.BRX:     "بدویی",
			language.BS:      "بوسنیایی",
			language.CA:      "کاتالونی",
			language.CE:      "چچنی",
			language.CGG:     "چیگا",
			language.CHR:     "چروکیایی",
			language.CKB:     "میونی کوردی",
			language.CO:      "کورسیکان",
			language.CS:      "چکی",
			language.CV:      "چوواشی",
			language.CY:      "ولزی",
			language.DA:      "دانمارکی",
			language.DAV:     "تایتا",
			language.DE:      "آلمانی",
			language.DE_AT:   "اتریش ِآلمانی",
			language.DE_CH:   "سوییس ِآلمانی",
			language.DJE:     "زارمایی",
			language.DSB:     "پایین صربی",
			language.DUA:     "دوئالایی",
			language.DYO:     "جولا-فونی",
			language.DZ:      "دزونگخا",
			language.EBU:     "امبو",
			language.EE:      "اوه\u200cیی",
			language.EL:      "یونانی",
			language.EN:      "انگلیسی",
			language.EN_AU:   "استرالیای ِانگلیسی",
			language.EN_CA:   "کانادای ِانگلیسی",
			language.EN_GB:   "بریتانیای ِانگلیسی",
			language.EN_US:   "آمریکای ِانگلیسی",
			language.EO:      "اسپرانتو",
			language.ES:      "ایسپانیولی",
			language.ES_419:  "جنوبی آمریکای ِایسپانیولی",
			language.ES_ES:   "اروپای ِایسپانیولی",
			language.ES_MX:   "مکزیک ِایسپانیولی",
			language.ET:      "استونیایی",
			language.EU:      "باسکی",
			language.FA:      "فارسی",
			language.FI:      "فینیش",
			language.FIL:     "فیلیپینو",
			language.FJ:      "فیجیایی",
			language.FO:      "فارویی",
			language.FR:      "فرانسوی",
			language.FR_CA:   "کانادای ِفرانسوی",
			language.FR_CH:   "سوییس ِفرانسوی",
			language.FY:      "غربی فیریزی",
			language.GA:      "ایریش",
			language.GAG:     "گاگائوزی",
			language.GL:      "گالیک",
			language.GN:      "گورانی",
			language.GSW:     "سوییس آلمانی",
			language.GU:      "گجراتی",
			language.GUZ:     "گوسی",
			language.GV:      "مانکس",
			language.HA:      "هوسا",
			language.HAW:     "هاواییایی",
			language.HE:      "عبری",
			language.HI:      "هندی",
			language.HR:      "کرواتی",
			language.HSB:     "بالایی صربی",
			language.HT:      "هائتیایی",
			language.HU:      "مجاری",
			language.HY:      "ارمنی",
			language.ID:      "اندونزیایی",
			language.IG:      "ایگبو",
			language.II:      "سیچوئان یی",
			language.IS:      "ایسلندی",
			language.IT:      "ایتالیایی",
			language.IU:      "انوکتیتوت",
			language.JA:      "جاپونی",
			language.JGO:     "نگومبا",
			language.JMC:     "ماچامه",
			language.JV:      "جاوایی",
			language.KA:      "گرجی",
			language.KAB:     "قبایلی",
			language.KAM:     "کامبایی",
			language.KDE:     "ماکونده",
			language.KEA:     "کیپ وُردی",
			language.KHQ:     "کویرا چیینی",
			language.KI:      "کیکویو",
			language.KK:      "قزاقی",
			language.KL:      "کالائلیسوت",
			language.KLN:     "کالنجین",
			language.KM:      "خمری",
			language.KN:      "کانّادا",
			language.KO:      "کُره\u200cیی",
			language.KOI:     "کومی-پرمیاک",
			language.KOK:     "کونکانی",
			language.KS:      "کشمیری",
			language.KSB:     "شامبالا",
			language.KSF:     "بافیایی",
			language.KU:      "کوردی",
			language.KW:      "کورنیش",
			language.KY:      "قرقیزی",
			language.LA:      "لاتین",
			language.LAG:     "لانگی",
			language.LB:      "لوکزامبورگی",
			language.LG:      "گاندا",
			language.LKT:     "لاکوتا",
			language.LN:      "لینگالا",
			language.LO:      "لائویی",
			language.LRC:     "شمالی لُری",
			language.LT:      "لتونیایی",
			language.LU:      "لوبا-کاتانگا",
			language.LUO:     "لوئو",
			language.LUY:     "لوییا",
			language.LV:      "لاتویایی",
			language.MAS:     "ماسایی",
			language.MER:     "مِرویی",
			language.MFE:     "موریسین",
			language.MG:      "مالاگاسی",
			language.MGH:     "ماخوئا-میتو",
			language.MGO:     "مِتاء",
			language.MI:      "مائوری",
			language.MK:      "مقدونی",
			language.ML:      "مالایالام",
			language.MN:      "مغولی",
			language.MOH:     "موهاک",
			language.MR:      "ماراتی",
			language.MS:      "مالایی",
			language.MT:      "مالتی",
			language.MUA:     "موندانگ",
			language.MY:      "برمه\u200cیی",
			language.MZN:     "مازرونی",
			language.NAQ:     "ناما",
			language.NB:      "نروژی بوکمال",
			language.ND:      "شمالی ندبله",
			language.NDS:     "پایین آلمانی",
			language.NDS_NL:  "پایین ساکسونی",
			language.NE:      "نپالی",
			language.NL:      "هلندی",
			language.NL_BE:   "فلمیش",
			language.NMG:     "کوئاسیو",
			language.NN:      "نروژی نینورسک",
			language.NQO:     "نئکو",
			language.NUS:     "نوئر",
			language.NYN:     "نیانکوله",
			language.OM:      "اورومو",
			language.OR:      "اوریا",
			language.PA:      "پنجابی",
			language.PL:      "لهستونی",
			language.PS:      "پشتو",
			language.PT:      "پرتغالی",
			language.PT_BR:   "برزیل ِپرتغالی",
			language.PT_PT:   "اروپای ِپرتغالی",
			language.QU:      "قوئچوئا",
			language.QUC:     "کئیچه\u200cئی",
			language.RM:      "رومانش",
			language.RN:      "روندی",
			language.RO:      "رومانیایی",
			language.RO_MD:   "مولداوی",
			language.ROF:     "رومبو",
			language.RU:      "روسی",
			language.RW:      "کنیاروآندایی",
			language.RWK:     "روآیی",
			language.SA:      "سانسکریت",
			language.SAQ:     "سامبورو",
			language.SBP:     "سانگوو",
			language.SD:      "سندی",
			language.SDH:     "جنوبی کردی",
			language.SE:      "شمالی سامی",
			language.SEH:     "سِنایی",
			language.SES:     "کویرابورا سنی",
			language.SG:      "سانگو",
			language.SHI:     "تاچلهیت",
			language.SI:      "سینهالا",
			language.SK:      "اسلواکی",
			language.SL:      "اسلوونیایی",
			language.SMA:     "جنوبی سامی",
			language.SMJ:     "لوله سامی",
			language.SMN:     "ایناری سامی",
			language.SMS:     "سکولت سامی",
			language.SN:      "شونا",
			language.SO:      "سومالیایی",
			language.SQ:      "آلبانیایی",
			language.SR:      "صربی",
			language.SU:      "سوندانسی",
			language.SV:      "سوئدی",
			language.SW:      "سواحیلی",
			language.SW_CD:   "کنگو سواحیلی",
			language.TA:      "تامیلی",
			language.TE:      "تلوگویی",
			language.TEO:     "تسویی",
			language.TG:      "تاجیکی",
			language.TH:      "تایی",
			language.TI:      "تیگرینیایی",
			language.TK:      "ترکمونی",
			language.TO:      "تونگانی",
			language.TR:      "ترکی",
			language.TT:      "تاتاری",
			language.TWQ:     "تاساواقی",
			language.TZM:     "میونی اطلس تامزیقی",
			language.UG:      "ئوغوری",
			language.UK:      "اوکراینی",
			language.UND:     "نشناسی\u200cیه زوون",
			language.UR:      "اردو",
			language.UZ:      "ازبکی",
			language.VAI:     "وایی",
			language.VI:      "ویتنامی",
			language.VUN:     "وونجویی",
			language.WBP:     "والرپیری",
			language.WO:      "وولفی",
			language.XH:      "خوسا",
			language.XOG:     "سوگا",
			language.YO:      "یوروبا",
			language.ZGH:     "مراکش ِاستاندارد ِتامازیقتی",
			language.ZH:      "چینی",
			language.ZH_HANS: "ساده چینی",
			language.ZH_HANT: "سنتی چینی",
			language.ZU:      "زولو",
			language.ZXX:     "این زوون بشناسی\u200cیه نیّه",
		},
		Territories: cldr.Territories{
			territory.V_001: "جهون",
			territory.V_002: "آفریقا",
			territory.V_003: "شمالی آمریکا",
			territory.V_005: "جنوبی آمریکا",
			territory.V_009: "اوقیانوسیه",
			territory.V_011: "غربی آفریقا",
			territory.V_013: "میونی آمریکا",
			territory.V_014: "شرقی آفریقا",
			territory.V_015: "شمالی ۀفریقا",
			territory.V_017: "میونی آفریقا",
			territory.V_018: "جنوبی آفریقا",
			territory.V_019: "آمریکا",
			territory.V_021: "شمالی امریکا",
			territory.V_029: "کاراییب",
			territory.V_030: "شرقی آسیا",
			territory.V_034: "جنوبی آسیا",
			territory.V_035: "آسیای ِجنوب\u200cشرقی\u200cوَر",
			territory.V_039: "جنوبی اروپا",
			territory.V_053: "اوسترالزی",
			territory.V_054: "ملانزی",
			territory.V_057: "میکرونزی منقطه",
			territory.V_061: "پولی\u200cنزی",
			territory.V_142: "آسیا",
			territory.V_143: "میونی آسیا",
			territory.V_145: "غربی آسیا",
			territory.V_150: "اروپا",
			territory.V_151: "شرقی اروپا",
			territory.V_154: "شمالی اروپا",
			territory.V_155: "غربی اروپا",
			territory.V_419: "لاتین آمریکا",
			territory.AC:    "آسنسیون جزیره",
			territory.AD:    "آندورا",
			territory.AE:    "متحده عربی امارات",
			territory.AF:    "افغانستون",
			territory.AG:    "آنتیگوا و باربودا",
			territory.AI:    "آنگویلا",
			territory.AL:    "آلبانی",
			territory.AM:    "ارمنستون",
			territory.AO:    "آنگولا",
			territory.AQ:    "جنوبی یخ\u200cبزه قطب",
			territory.AR:    "آرژانتین",
			territory.AS:    "آمریکای ِساموآ",
			territory.AT:    "اتریش",
			territory.AU:    "استرالیا",
			territory.AW:    "آروبا",
			territory.AX:    "آلند جزیره",
			territory.AZ:    "آذربایجون",
			territory.BA:    "بوسنی و هرزگوین",
			territory.BB:    "باربادوس",
			territory.BD:    "بنگلادش",
			territory.BE:    "بلژیک",
			territory.BF:    "بورکینا فاسو",
			territory.BG:    "بلغارستون",
			territory.BH:    "بحرین",
			territory.BI:    "بوروندی",
			territory.BJ:    "بنین",
			territory.BL:    "سنت بارتلمی",
			territory.BM:    "برمودا",
			territory.BN:    "برونئی",
			territory.BO:    "بولیوی",
			territory.BQ:    "هلند ِکاراییبی جزایر",
			territory.BR:    "برزیل",
			territory.BS:    "باهاما",
			territory.BT:    "بوتان",
			territory.BV:    "بووت جزیره",
			territory.BW:    "بوتساوانا",
			territory.BY:    "بلاروس",
			territory.BZ:    "بلیز",
			territory.CA:    "کانادا",
			territory.CC:    "کوک (کیلینگ) جزایر",
			territory.CD:    "کنگو کینشاسا",
			territory.CF:    "مرکزی آفریقای جمهوری",
			territory.CG:    "کنگو برازاویل",
			territory.CH:    "سوییس",
			territory.CI:    "عاج ِساحل",
			territory.CK:    "کوک جزایر",
			territory.CL:    "شیلی",
			territory.CM:    "کامرون",
			territory.CN:    "چین",
			territory.CO:    "کلمبیا",
			territory.CP:    "کلیپرتون جزیره",
			territory.CR:    "کاستاریکا",
			territory.CU:    "کوبا",
			territory.CV:    "کیپ ورد",
			territory.CW:    "کوراسائو",
			territory.CX:    "کریسمس جزیره",
			territory.CY:    "قبرس",
			territory.CZ:    "چک جمهوری",
			territory.DE:    "آلمان",
			territory.DG:    "دیگو گارسیا",
			territory.DJ:    "جیبوتی",
			territory.DK:    "دانمارک",
			territory.DM:    "دومنیکا",
			territory.DO:    "دومنیکن جمهوری",
			territory.DZ:    "الجزیره",
			territory.EA:    "سوتا و ملیله",
			territory.EC:    "اکوادر",
			territory.EE:    "استونی",
			territory.EG:    "مصر",
			territory.EH:    "غربی صحرا",
			territory.ER:    "اریتره",
			territory.ES:    "ایسپانیا",
			territory.ET:    "اتیوپی",
			territory.EU:    "اروپا اتحادیه",
			territory.FI:    "فنلاند",
			territory.FJ:    "فیجی",
			territory.FK:    "فالکلند جزیره\u200cئون",
			territory.FM:    "میکرونزی",
			territory.FO:    "فارو جزایر",
			territory.FR:    "فرانسه",
			territory.GA:    "گابون",
			territory.GB:    "بریتانیا",
			territory.GD:    "گرانادا",
			territory.GE:    "گرجستون",
			territory.GF:    "فرانسه\u200cی ِگویان",
			territory.GG:    "گرنزی",
			territory.GH:    "غنا",
			territory.GI:    "جبل طارق",
			territory.GL:    "گرینلند",
			territory.GM:    "گامبیا",
			territory.GN:    "گینه",
			territory.GP:    "گوادلوپ",
			territory.GQ:    "استوایی گینه",
			territory.GR:    "یونان",
			territory.GS:    "جنوبی جورجیا و جنوبی ساندویچ جزایر",
			territory.GT:    "گواتمالا",
			territory.GU:    "گوئام",
			territory.GW:    "گینه بیسائو",
			territory.GY:    "گویان",
			territory.HK:    "هنگ کنگ",
			territory.HM:    "هارد و مک\u200cدونالد جزایر",
			territory.HN:    "هندوراس",
			territory.HR:    "کرواسی",
			territory.HT:    "هاییتی",
			territory.HU:    "مجارستون",
			territory.IC:    "قناری جزایر",
			territory.ID:    "اندونزی",
			territory.IE:    "ایرلند",
			territory.IL:    "ایسراییل",
			territory.IM:    "من ِجزیره",
			territory.IN:    "هند",
			territory.IO:    "بریتانیای هند ِاوقیانوس ِمناطق",
			territory.IQ:    "عراق",
			territory.IR:    "ایران",
			territory.IS:    "ایسلند",
			territory.IT:    "ایتالیا",
			territory.JE:    "جرسی",
			territory.JM:    "جاماییکا",
			territory.JO:    "اردن",
			territory.JP:    "جاپون",
			territory.KE:    "کنیا",
			territory.KG:    "قرقیزستون",
			territory.KH:    "کامبوج",
			territory.KI:    "کیریباتی",
			territory.KM:    "کومور",
			territory.KN:    "سنت کیتس و نویس",
			territory.KP:    "شمالی کُره",
			territory.KR:    "جنوبی کُره",
			territory.KW:    "کویت",
			territory.KY:    "کیمن جزیره\u200cئون",
			territory.KZ:    "قزاقستون",
			territory.LA:    "لائوس",
			territory.LB:    "لبنان",
			territory.LC:    "سنت لوسیا",
			territory.LI:    "لیختن اشتاین",
			territory.LK:    "سریلانکا",
			territory.LR:    "لیبریا",
			territory.LS:    "لسوتو",
			territory.LT:    "لتونی",
			territory.LU:    "لوکزامبورگ",
			territory.LV:    "لاتویا",
			territory.LY:    "لیبی",
			territory.MA:    "مراکش",
			territory.MC:    "موناکو",
			territory.MD:    "مولداوی",
			territory.ME:    "مونته\u200cنگرو",
			territory.MF:    "سنت مارتین",
			territory.MG:    "ماداگاسکار",
			territory.MH:    "مارشال جزایر",
			territory.ML:    "مالی",
			territory.MM:    "میانمار",
			territory.MN:    "مغولستون",
			territory.MO:    "ماکائو (چین دله)",
			territory.MP:    "شمالی ماریانا جزایر",
			territory.MQ:    "مارتینیک جزیره\u200cئون",
			territory.MR:    "موریتانی",
			territory.MS:    "مونتسرات",
			territory.MT:    "مالت",
			territory.MU:    "مورى تيوس",
			territory.MV:    "مالدیو",
			territory.MW:    "مالاوی",
			territory.MX:    "مکزیک",
			territory.MY:    "مالزی",
			territory.MZ:    "موزامبیک",
			territory.NA:    "نامبیا",
			territory.NC:    "نیو کالیدونیا",
			territory.NE:    "نیجر",
			territory.NF:    "نورفولک جزیره",
			territory.NG:    "نیجریه",
			territory.NI:    "نیکاراگوئه",
			territory.NL:    "هلند",
			territory.NO:    "نروژ",
			territory.NP:    "نپال",
			territory.NR:    "نائورو",
			territory.NU:    "نیئو",
			territory.NZ:    "نیوزلند",
			territory.OM:    "عمان",
			territory.PA:    "پاناما",
			territory.PE:    "پرو",
			territory.PF:    "فرانسه\u200cی پولی\u200cنزی",
			territory.PG:    "پاپوا نو گینه",
			territory.PH:    "فیلیپین",
			territory.PK:    "پاکستون",
			territory.PL:    "لهستون",
			territory.PM:    "سن پییر و میکلن",
			territory.PN:    "پیتکارین جزایر",
			territory.PR:    "پورتوریکو",
			territory.PS:    "فلسطین ِسرزمین",
			territory.PT:    "پرتغال",
			territory.PW:    "پالائو",
			territory.PY:    "پاراگوئه",
			territory.QA:    "قطر",
			territory.QO:    "اوقیانوسیه\u200cی ِپرت ِجائون",
			territory.RE:    "رئونیون",
			territory.RO:    "رومانی",
			territory.RS:    "صربستون",
			territory.RU:    "روسیه",
			territory.RW:    "روآندا",
			territory.SA:    "عربستون",
			territory.SB:    "سلیمون جزیره",
			territory.SC:    "سیشل",
			territory.SD:    "سودان",
			territory.SE:    "سوئد",
			territory.SG:    "سنگاپور",
			territory.SH:    "سنت هلنا",
			territory.SI:    "اسلوونی",
			territory.SJ:    "سوالبارد و يان ماين",
			territory.SK:    "اسلواکی",
			territory.SL:    "سیرالئون",
			territory.SM:    "سن مارینو",
			territory.SN:    "سنگال",
			territory.SO:    "سومالی",
			territory.SR:    "سورینام",
			territory.SS:    "جنوبی سودان",
			territory.ST:    "سائوتومه و پرینسیپ",
			territory.SV:    "السالوادور",
			territory.SX:    "سنت مارتن",
			territory.SY:    "سوریه",
			territory.SZ:    "سوازیلند",
			territory.TA:    "تریستان دا جونها",
			territory.TC:    "تورکس و کایکوس جزایر",
			territory.TD:    "چاد",
			territory.TF:    "فرانسه\u200cی جنوبی مناطق",
			territory.TG:    "توگو",
			territory.TH:    "تایلند",
			territory.TJ:    "تاجیکستون",
			territory.TK:    "توکلائو",
			territory.TL:    "تیمور شرقی",
			territory.TM:    "ترکمونستون",
			territory.TN:    "تونس",
			territory.TO:    "تونگا",
			territory.TR:    "ترکیه",
			territory.TT:    "ترینیداد و توباگو",
			territory.TV:    "تووالو",
			territory.TW:    "تایوان",
			territory.TZ:    "تانزانیا",
			territory.UA:    "اوکراین",
			territory.UG:    "اوگاندا",
			territory.UM:    "آمریکای پَرتِ\u200cپِلا جزیره\u200cئون",
			territory.US:    "متحده ایالات",
			territory.UY:    "اروگوئه",
			territory.UZ:    "ازبکستون",
			territory.VA:    "واتیکان",
			territory.VC:    "سنت وینسنت و گرنادین",
			territory.VE:    "ونزوئلا",
			territory.VG:    "بریتانیای ویرجین",
			territory.VI:    "آمریکای ویرجین",
			territory.VN:    "ویتنام",
			territory.VU:    "وانواتو",
			territory.WF:    "والیس و فوتونا",
			territory.WS:    "ساموآ",
			territory.XK:    "کوزوو",
			territory.YE:    "یمن",
			territory.YT:    "مایوت",
			territory.ZA:    "جنوبی افریقا",
			territory.ZM:    "زامبیا",
			territory.ZW:    "زیمبابوه",
			territory.ZZ:    "نامَیِّن منطقه",
		},
	}
}
