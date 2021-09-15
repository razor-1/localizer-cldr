// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_os_GE() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "os_GE",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y 'аз'", Long: "d MMMM, y 'аз'", Medium: "dd MMM y 'аз'", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Янв.", Feb: "Февр.", Mar: "Март.", Apr: "Апр.", May: "Май", Jun: "Июнь", Jul: "Июль", Aug: "Авг.", Sep: "Сент.", Oct: "Окт.", Nov: "Нояб.", Dec: "Дек."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Январь", Feb: "Февраль", Mar: "Мартъи", Apr: "Апрель", May: "Май", Jun: "Июнь", Jul: "Июль", Aug: "Август", Sep: "Сентябрь", Oct: "Октябрь", Nov: "Ноябрь", Dec: "Декабрь"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Хцб", Mon: "Крс", Tue: "Дцг", Wed: "Ӕрт", Thu: "Цпр", Fri: "Мрб", Sat: "Сбт"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Х", Mon: "К", Tue: "Д", Wed: "Ӕ", Thu: "Ц", Fri: "М", Sat: "С"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Хуыцаубон", Mon: "Къуырисӕр", Tue: "Дыццӕг", Wed: "Ӕртыццӕг", Thu: "Цыппӕрӕм", Fri: "Майрӕмбон", Sat: "Сабат"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ӕмбисбоны размӕ", PM: "ӕмбисбоны фӕстӕ"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Бразилиаг реал", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "Евро", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Бритайнаг Фунт", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Лар", Symbol: "₾"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "", Symbol: "JP¥"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.MGA: cldr.Currency{DisplayName: "", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "Сом", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "АИШ-ы Доллар", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Нӕзонгӕ валютӕ", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AB:      "абхазаг",
			language.ADY:     "адыгейаг",
			language.AE:      "авестӕ",
			language.AF:      "африкаанс",
			language.ANG:     "рагон англисаг",
			language.AR:      "араббаг",
			language.AV:      "авайраг",
			language.AZ:      "тӕтӕйраг",
			language.BA:      "башкираг",
			language.BG:      "болгайраг",
			language.BS:      "босниаг",
			language.BUA:     "бурятаг",
			language.CA:      "каталайнаг",
			language.CE:      "цӕцӕйнаг",
			language.COP:     "коптаг",
			language.CS:      "чехаг",
			language.CV:      "чувашаг",
			language.DA:      "даниаг",
			language.DE:      "немыцаг",
			language.DE_AT:   "австралиаг немыцаг",
			language.DE_CH:   "швйецариаг немыцаг",
			language.EGY:     "рагон египтаг",
			language.EL:      "бердзейнаг",
			language.EN:      "англисаг",
			language.EN_AU:   "австралиаг англисаг",
			language.EN_CA:   "канадӕйаг англисаг",
			language.EN_GB:   "бритайнаг англисаг",
			language.EN_US:   "америкаг англисаг",
			language.EO:      "есперанто",
			language.ES:      "испайнаг",
			language.ES_419:  "латинаг америкаг англисаг",
			language.ES_ES:   "европӕйаг англисаг",
			language.ES_MX:   "мексикӕйаг испайнаг",
			language.ET:      "естойнаг",
			language.EU:      "баскаг",
			language.FA:      "персайнаг",
			language.FI:      "финнаг",
			language.FIL:     "филиппинаг",
			language.FJ:      "фиджи",
			language.FO:      "фарераг",
			language.FR:      "францаг",
			language.FR_CA:   "канадӕйаг францаг",
			language.FR_CH:   "швейцариаг францаг",
			language.FRO:     "рагон францаг",
			language.GA:      "ирландиаг",
			language.GRC:     "рагон бердзейнаг",
			language.HE:      "уираг",
			language.HR:      "хорватаг",
			language.HU:      "венгериаг",
			language.HY:      "сомихаг",
			language.INH:     "мӕхъӕлон",
			language.IT:      "италиаг",
			language.JA:      "япойнаг",
			language.KA:      "гуырдзиаг",
			language.KBD:     "кӕсгон",
			language.KRC:     "бӕлхъӕрон",
			language.KU:      "курдаг",
			language.KUM:     "хъуымыхъхъаг",
			language.LA:      "латинаг",
			language.LEZ:     "лекъаг",
			language.MK:      "мӕчъидон",
			language.OS:      "ирон",
			language.PT:      "португалиаг",
			language.PT_BR:   "бразилиаг португалиаг",
			language.PT_PT:   "европӕйаг полтугалиаг",
			language.ROM:     "цигайнаг",
			language.RU:      "уырыссаг",
			language.UND:     "нӕзонгӕ ӕвзаг",
			language.ZH:      "китайаг",
			language.ZH_HANS: "ӕнцонгонд китайаг",
			language.ZH_HANT: "традицион китайаг",
		},
		Territories: cldr.Territories{
			territory.V_001: "Дуне",
			territory.V_002: "Африкӕ",
			territory.V_009: "Океани",
			territory.V_019: "Америкӕ",
			territory.V_142: "Ази",
			territory.V_150: "Европӕ",
			territory.BR:    "Бразили",
			territory.CN:    "Китай",
			territory.DE:    "Герман",
			territory.FR:    "Франц",
			territory.GB:    "Стыр Британи",
			territory.GE:    "Гуырдзыстон",
			territory.IN:    "Инди",
			territory.IT:    "Итали",
			territory.JP:    "Япон",
			territory.RU:    "Уӕрӕсе",
			territory.US:    "АИШ",
			territory.ZZ:    "Нӕзонгӕ бӕстӕ",
		},
	}
}
