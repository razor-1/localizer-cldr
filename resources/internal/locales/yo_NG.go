// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_yo_NG() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "yo_NG",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMM y", Long: "d MMM y", Medium: "d MM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:m:s", Short: "H:m"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "WAT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ṣẹ́", Feb: "Èr", Mar: "Ẹr", Apr: "Ìg", May: "Ẹ̀b", Jun: "Òk", Jul: "Ag", Aug: "Òg", Sep: "Ow", Oct: "Ọ̀w", Nov: "Bé", Dec: "Ọ̀p"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "S", Feb: "È", Mar: "Ẹ", Apr: "Ì", May: "Ẹ̀", Jun: "Ò", Jul: "A", Aug: "Ò", Sep: "O", Oct: "Ọ̀", Nov: "B", Dec: "Ọ̀"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Ṣẹ́rẹ́", Feb: "Èrèlè", Mar: "Ẹrẹ̀nà", Apr: "Ìgbé", May: "Ẹ̀bibi", Jun: "Òkúdu", Jul: "Agẹmọ", Aug: "Ògún", Sep: "Owewe", Oct: "Ọ̀wàrà", Nov: "Bélú", Dec: "Ọ̀pẹ̀"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Àìk", Mon: "Aj", Tue: "Ìsẹ́g", Wed: "Ọjọ́r", Thu: "Ọjọ́b", Fri: "Ẹt", Sat: "Àbám"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "À", Mon: "A", Tue: "Ì", Wed: "Ọ", Thu: "Ọ", Fri: "Ẹ", Sat: "À"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Àìkú", Mon: "Ajé", Tue: "Ìsẹ́gun", Wed: "Ọjọ́rú", Thu: "Ọjọ́bọ", Fri: "Ẹtì", Sat: "Àbámẹ́ta"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Àìkú", Mon: "Ajé", Tue: "Ìsẹ́gun", Wed: "Ọjọ́rú", Thu: "Ọjọ́bọ", Fri: "Ẹtì", Sat: "Àbámẹ́ta"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Àárọ̀", PM: "Ọ̀sán"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "Àárọ̀", PM: "Ọ̀sán"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Àárọ̀", PM: "Ọ̀sán"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Diami ti Awon Orílẹ́ède Arabu", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afugánì Afuganísítàànì", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Lẹ́kẹ̀ Àlìbéníà", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Dírààmù Àmẹ́níà", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Gílídà Netherlands Antillean", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Wansa ti Orílẹ́ède Àngólà", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Pẹ́sò Agẹntínà", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dọla ti Orílẹ́ède Ástràlìá", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Fuloríìnì Àrúbà", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Mánààtì Àsàbáíjáì", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Àmi Yíyípadà Bosnia-Herzegovina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Dọ́là Bábádọ̀ọ̀sì", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Tákà Báńgíládẹ̀ẹ̀ṣì", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Owó Lẹ́fì Bọ̀lìgéríà", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Dina ti Orílẹ́ède Báránì", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Faransi ti Orílẹ́ède Bùùrúndì", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Dọ́là Bẹ̀múdà", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Dọ́là Bùrùnéì", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Bọlifiánò Bọ̀lífíà", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Owó ti Orílẹ̀-èdè Brazil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dọ́là Bàhámà", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Ìngọ́tírọ̀mù Bútàànì", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Pula ti Orílẹ́ède Bọ̀tìsúwánà", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Rọ́bù Bẹ̀lárùùsì", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "Dọ́là Bẹ̀lísè", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dọla ti Orílẹ́ède Kánádà", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faransi ti Orílẹ́ède Kóngò", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faransi ti Orílẹ́ède Siwisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Pẹ́sò Ṣílè", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Yúànì Sháínà", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Reminibi ti Orílẹ́ède ṣáínà", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Pẹ́sò Kòlóḿbíà", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Kólọ́ọ̀nì Kosita Ríkà", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Pẹ́sò Yíyípadà Kúbà", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Pẹ́sò Kúbà", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Kabofediano ti Orílẹ́ède Esuodo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Koruna Ṣẹ́ẹ̀kì", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faransi ti Orílẹ́ède Dibouti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Kírónì Dáníṣì", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Pẹ́sò Dòníníkà", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dina ti Orílẹ́ède Àlùgèríánì", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "pọọn ti Orílẹ́ède Egipiti", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakifa ti Orílẹ́ède Eriteriani", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Biri ti Orílẹ́ède Eutopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "owó Yúrò", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dọ́là Fíjì", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Pọ́n-ùn Erékùsù Falkland", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pọ́n-ùn ti Orilẹ̀-èdè Gẹ̀ẹ́sì", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Lárì Jọ́jíà", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "ṣidi ti Orílẹ́ède Gana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Sídì Ghanaian", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Pọ́n-ùn Gibraltar", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ti Orílẹ́ède Gamibia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Fírànkì Gíníànì", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Faransi ti Orílẹ́ède Gini", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Kúẹ́tísààlì Guatimílà", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Dọ́là Gùyánà", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Dọ́là Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lẹmipírà Ọ́ńdúrà", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kúnà Croatian", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Gọ́dì Àítì", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Fọ́ríǹtì Họ̀ngérí", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Rúpìyá Indonésíà", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "Ṣékélì Tuntun Ísírẹ̀ẹ̀lì", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupi ti Orílẹ́ède Indina", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dínárì Ìráákì", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Rial Iranian", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Kòrónà Icelandic", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Dọ́là Jàmáíkà", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Dínárì Jọ́dàànì", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yeni ti Orílẹ́ède Japani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "ṣiili ti Orílẹ́ède Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Sómú Kirijísítàànì", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Ráyò Kàm̀bọ́díà", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Faransi ti Orílẹ́ède ṣomoriani", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Wọ́ọ̀nù Àríwá Kòríà", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Wọ́ọ̀nù Gúúsù Kòríà", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dínárì Kuwaiti", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Dọ́là Erékùsù Cayman", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Tẹngé Kasakísítàànì", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Kíììpù Làótì", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Pọ́n-ùn Lebanese", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Rúpìì Siri Láńkà", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dọla ti Orílẹ́ède Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ti Orílẹ́ède Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dina ti Orílẹ́ède Libiya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirami ti Orílẹ́ède Moroko", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Owó Léhù Moldovan", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Faransi ti Orílẹ́ède Malagasi", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Dẹ́nà Masidóníà", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kíyàtì Myanmar", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Túgúrììkì Mòǹgólíà", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Pàtákà Màkáò", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya ti Orílẹ́ède Maritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya ti Orílẹ́ède Maritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupi ti Orílẹ́ède Maritiusi", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Rúfìyá Mọ̀lìdífà", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Kaṣa ti Orílẹ́ède Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Pẹ́sò Mẹ́síkò", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ríngìtì Màléṣíà", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metika ti Orílẹ́ède Mosamibiki", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Mẹ́tíkààlì Mòsáḿbíìkì", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dọla ti Orílẹ́ède Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Náìrà ti Orílẹ̀-èdè Nàìjíríà", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Kọ̀dóbà Naikarágúà", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Kírónì Nọ́ọ́wè", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Rúpìì Nẹ̵́pààlì", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Dọ́là New Zealand", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Ráyò Omani", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Bálíbóà Pànámà", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Sólì Pèrúù", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kínà Papua Guinea Tuntun", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Písò Fílípìnì", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Rúpìì Pakisitánì", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Sílọ̀tì Pọ́líṣì", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Gúáránì Párágúwè", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Ráyò Kàtárì", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Léhù Ròméníà", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dínárì Sàbíà", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Owó ruble ti ilẹ̀ Rọ́ṣíà", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Faransi ti Orílẹ́ède Ruwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riya ti Orílẹ́ède Saudi", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Dọ́là Erékùsù Sọ́lómọ́nì", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupi ti Orílẹ́ède Sayiselesi", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dina ti Orílẹ́ède Sudani", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Pọọun ti Orílẹ́ède Sudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Kòrónà Súwídìn", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Dọ́là Síngápọ̀", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Pọọun ti Orílẹ́ède ̣Elena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Lioni", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Sile ti Orílẹ́ède Somali", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Dọ́là Súrínámì", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Pọ́n-ùn Gúúsù Sùdáànì", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobira ti Orílẹ́ède Sao tome Ati Pirisipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobira ti Orílẹ́ède Sao tome Ati Pirisipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Pọ́n-ùn Sírìà", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Báàtì Tháì", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Sómónì Tajikístàànì", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Mánààtì Tọkimẹnístàànì", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Dina ti Orílẹ́ède Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Pàángà Tóńgà", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Lírà Tọ́kì", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Dọ́là Trinidad & Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Dọ́là Tàìwánì Tuntun", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Sile ti Orílẹ́ède Tansania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Ọrifiníyà Yukiréníà", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Siile ti Orílẹ́ède Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dọ́là", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "Pẹ́sò Úrúgúwè", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Sómú Usibẹkísítàànì", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Bọ̀lífà Fẹnẹsuẹ́là", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Dáhùn Vietnamese", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Fátù Vanuatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tálà Sàmóà", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Faransi ti Orílẹ́ède BEKA", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dọ́là Ilà Oòrùn Karíbíà", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Faransi ti Orílẹ́ède BIKEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Fírànkì CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "owóníná àìmọ̀", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Ráyò Yẹ́mẹ̀nì", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Randi ti Orílẹ́ède Ariwa Afirika", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kawaṣa ti Orílẹ́ède Saabia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kawaṣa ti Orílẹ́ède Saabia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dọla ti Orílẹ́ède Siibabuwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AF:      "Èdè Afrikani",
			language.AGQ:     "Ágẹ̀ẹ̀mù",
			language.AK:      "Èdè Akani",
			language.AM:      "Èdè Amariki",
			language.AR:      "Èdè Árábìkì",
			language.AS:      "Ti Assam",
			language.ASA:     "Asu",
			language.AST:     "Asturian",
			language.AZ:      "Èdè Azerbaijani",
			language.BAS:     "Basaa",
			language.BE:      "Èdè Belarusi",
			language.BEM:     "Béḿbà",
			language.BEZ:     "Bẹ́nà",
			language.BG:      "Èdè Bugaria",
			language.BM:      "Báḿbàrà",
			language.BN:      "Èdè Bengali",
			language.BO:      "Tibetán",
			language.BR:      "Èdè Bretoni",
			language.BRX:     "Bódò",
			language.BS:      "Èdè Bosnia",
			language.CA:      "Èdè Catala",
			language.CCP:     "Chakma",
			language.CE:      "Chechen",
			language.CEB:     "Cebuano",
			language.CGG:     "Chiga",
			language.CHR:     "Shẹ́rókiì",
			language.CKB:     "Ààrin Gbùngbùn Kurdish",
			language.CO:      "Corsican",
			language.CS:      "Èdè seeki",
			language.CU:      "Síláfííkì Ilé Ìjọ́sìn",
			language.CY:      "Èdè Welshi",
			language.DA:      "Èdè Ilẹ̀ Denmark",
			language.DAV:     "Táítà",
			language.DE:      "Èdè Jámánì",
			language.DE_AT:   "Èdè Jámánì (Ọ́síríà )",
			language.DE_CH:   "Èdè Ilẹ̀ Jámánì (Orílẹ́ède swítsàlandì)",
			language.DJE:     "Ṣárúmà",
			language.DSB:     "Ṣobíànù Ìpìlẹ̀",
			language.DUA:     "Duala",
			language.DYO:     "Jola-Fonyi",
			language.DZ:      "Dzongkha",
			language.EBU:     "Ẹmbù",
			language.EE:      "Ewè",
			language.EL:      "Èdè Giriki",
			language.EN:      "Èdè Gẹ̀ẹ́sì",
			language.EN_AU:   "Èdè Gẹ̀ẹ́sì (órílẹ̀-èdè Ọsirélíà)",
			language.EN_CA:   "Èdè Gẹ̀ẹ́sì (Orílẹ̀-èdè Kánádà)",
			language.EN_GB:   "Èdè Gẹ̀ẹ́sì (GB)",
			language.EN_US:   "Èdè Gẹ̀ẹ́sì (US)",
			language.EO:      "Èdè Esperanto",
			language.ES:      "Èdè Sípáníìṣì",
			language.ES_419:  "Èdè Sípáníìṣì (orílẹ̀-èdè Látìn-Amẹ́ríkà) ( Èdè Sípáníìshì (Látìn-Amẹ́ríkà)",
			language.ES_ES:   "Èdè Sípáníìṣì (orílẹ̀-èdè Yúróòpù)",
			language.ES_MX:   "Èdè Sípáníìṣì (orílẹ̀-èdè Mẹ́síkò)",
			language.ET:      "Èdè Estonia",
			language.EU:      "Èdè Baski",
			language.EWO:     "Èwóǹdò",
			language.FA:      "Èdè Pasia",
			language.FF:      "Èdè Fúlàní",
			language.FI:      "Èdè Finisi",
			language.FIL:     "Èdè Filipino",
			language.FO:      "Èdè Faroesi",
			language.FR:      "Èdè Faransé",
			language.FR_CA:   "Èdè Faransé (orílẹ̀-èdè Kánádà)",
			language.FR_CH:   "Èdè Faranṣé (Súwísàlaǹdì)",
			language.FUR:     "Firiúlíànì",
			language.FY:      "Èdè Frisia",
			language.GA:      "Èdè Ireland",
			language.GD:      "Èdè Gaelik ti Ilu Scotland",
			language.GL:      "Èdè Galicia",
			language.GN:      "Èdè Guarani",
			language.GSW:     "Súwísì ti Jámánì",
			language.GU:      "Èdè Gujarati",
			language.GUZ:     "Gusii",
			language.GV:      "Máǹkì",
			language.HA:      "Èdè Hausa",
			language.HAW:     "Hawaiian",
			language.HE:      "Èdè Heberu",
			language.HI:      "Èdè Híńdì",
			language.HMN:     "Hmong",
			language.HR:      "Èdè Kroatia",
			language.HSB:     "Sorbian Apá Òkè",
			language.HT:      "Haitian Creole",
			language.HU:      "Èdè Hungaria",
			language.HY:      "Èdè Ile Armenia",
			language.IA:      "Èdè pipo",
			language.ID:      "Èdè Indonéṣíà",
			language.IE:      "Iru Èdè",
			language.IG:      "Èdè Yíbò",
			language.II:      "Ṣíkuán Yì",
			language.IS:      "Èdè Icelandic",
			language.IT:      "Èdè Ítálì",
			language.JA:      "Èdè Jàpáànù",
			language.JGO:     "Ńgòmbà",
			language.JMC:     "Máṣámè",
			language.JV:      "Èdè Javanasi",
			language.KA:      "Èdè Georgia",
			language.KAB:     "Kabilè",
			language.KAM:     "Káńbà",
			language.KDE:     "Mákondé",
			language.KEA:     "Kabufadíánù",
			language.KHQ:     "Koira Ṣíínì",
			language.KI:      "Kíkúyù",
			language.KK:      "Kaṣakì",
			language.KKJ:     "Kàkó",
			language.KL:      "Kalaalísùtì",
			language.KLN:     "Kálẹnjín",
			language.KM:      "Èdè kameri",
			language.KN:      "Èdè Kannada",
			language.KO:      "Èdè Kòríà",
			language.KOK:     "Kónkánì",
			language.KS:      "Kaṣímirì",
			language.KSB:     "Ṣáńbálà",
			language.KSF:     "Báfíà",
			language.KSH:     "Colognian",
			language.KU:      "Kọdiṣì",
			language.KW:      "Kọ́nììṣì",
			language.KY:      "Kírígíìsì",
			language.LA:      "Èdè Latini",
			language.LAG:     "Láńgì",
			language.LB:      "Lùṣẹ́mbọ́ọ̀gì",
			language.LG:      "Ganda",
			language.LKT:     "Lákota",
			language.LN:      "Lìǹgálà",
			language.LO:      "Láò",
			language.LRC:     "Apáàríwá Lúrì",
			language.LT:      "Èdè Lithuania",
			language.LU:      "Lúbà-Katanga",
			language.LUY:     "Luyíà",
			language.LV:      "Èdè Latvianu",
			language.MAS:     "Másáì",
			language.MER:     "Mérù",
			language.MFE:     "Morisiyen",
			language.MG:      "Malagasì",
			language.MGH:     "Makhuwa-Meeto",
			language.MGO:     "Métà",
			language.MI:      "Màórì",
			language.MK:      "Èdè Macedonia",
			language.ML:      "Málàyálámù",
			language.MN:      "Mòngólíà",
			language.MR:      "Èdè marathi",
			language.MS:      "Èdè Malaya",
			language.MT:      "Èdè Malta",
			language.MUA:     "Múndàngì",
			language.MUL:     "Ọlọ́pọ̀ èdè",
			language.MY:      "Èdè Bumiisi",
			language.MZN:     "Masanderani",
			language.NAQ:     "Námà",
			language.NB:      "Nọ́ọ́wè Bokímàl",
			language.ND:      "Àríwá Ndebele",
			language.NDS:     "Jámánì ìpìlẹ̀",
			language.NE:      "Èdè Nepali",
			language.NL:      "Èdè Dọ́ọ̀ṣì",
			language.NMG:     "Kíwáṣíò",
			language.NN:      "Nọ́ọ́wè Nínọ̀sìkì",
			language.NNH:     "Ngiembùnù",
			language.NO:      "Èdè Norway",
			language.NUS:     "Núẹ̀",
			language.NY:      "Ńyájà",
			language.NYN:     "Ńyákọ́lè",
			language.OC:      "Èdè Occitani",
			language.OM:      "Òròmọ́",
			language.OR:      "Òdíà",
			language.OS:      "Ọṣẹ́tíìkì",
			language.PA:      "Èdè Punjabi",
			language.PL:      "Èdè Póláǹdì",
			language.PRG:     "Púrúṣíànù",
			language.PS:      "Páshítò",
			language.PT:      "Èdè Pọtogí",
			language.PT_BR:   "Èdè Pọtogí (Orilẹ̀-èdè Bràsíl)",
			language.PT_PT:   "Èdè Pọtogí (orílẹ̀-èdè Yúróòpù)",
			language.QU:      "Kúẹ́ńjùà",
			language.RM:      "Rómáǹṣì",
			language.RN:      "Rúńdì",
			language.RO:      "Èdè Romania",
			language.ROF:     "Róńbò",
			language.RU:      "Èdè Rọ́ṣíà",
			language.RW:      "Èdè Ruwanda",
			language.RWK:     "Riwa",
			language.SA:      "Èdè awon ara Indo",
			language.SAH:     "Sàkíhà",
			language.SAQ:     "Samburu",
			language.SBP:     "Sangu",
			language.SD:      "Èdè Sindhi",
			language.SE:      "Apáàríwá Sami",
			language.SEH:     "Ṣẹnà",
			language.SES:     "Koiraboro Seni",
			language.SG:      "Sango",
			language.SH:      "Èdè Serbo-Croatiani",
			language.SHI:     "Taṣelíìtì",
			language.SI:      "Èdè Sinhalese",
			language.SK:      "Èdè Slovaki",
			language.SL:      "Èdè Slovenia",
			language.SM:      "Sámóánù",
			language.SMN:     "Inari Sami",
			language.SN:      "Ṣọnà",
			language.SO:      "Èdè ara Somalia",
			language.SQ:      "Èdè Albania",
			language.SR:      "Èdè Serbia",
			language.ST:      "Èdè Sesoto",
			language.SU:      "Èdè Sudani",
			language.SV:      "Èdè Suwidiisi",
			language.SW:      "Èdè Swahili",
			language.TA:      "Èdè Tamili",
			language.TE:      "Èdè Telugu",
			language.TEO:     "Tẹ́sò",
			language.TG:      "Tàjíìkì",
			language.TH:      "Èdè Tai",
			language.TI:      "Èdè Tigrinya",
			language.TK:      "Èdè Turkmen",
			language.TLH:     "Èdè Klingoni",
			language.TO:      "Tóńgàn",
			language.TR:      "Èdè Tọọkisi",
			language.TT:      "Tatarí",
			language.TWQ:     "Tasawak",
			language.TZM:     "Ààrin Gbùngbùn Atlas Tamazight",
			language.UG:      "Yúgọ̀",
			language.UK:      "Èdè Ukania",
			language.UND:     "Èdè àìmọ̀",
			language.UR:      "Èdè Udu",
			language.UZ:      "Èdè Uzbek",
			language.VI:      "Èdè Jetinamu",
			language.VO:      "Fọ́lápùùkù",
			language.VUN:     "Funjo",
			language.WAE:     "Wọsà",
			language.WO:      "Wọ́lọ́ọ̀fù",
			language.XH:      "Èdè Xhosa",
			language.XOG:     "Ṣógà",
			language.YAV:     "Yangbẹn",
			language.YI:      "Èdè Yiddishi",
			language.YO:      "Èdè Yorùbá",
			language.YUE:     "Cantonese",
			language.ZGH:     "Àfẹnùkò Támásáìtì ti Mòrókò",
			language.ZH:      "Ṣáídà, Mandrínì",
			language.ZH_HANT: "Èdè Ìbílẹ̀ Ṣáínà",
			language.ZU:      "Èdè Ṣulu",
			language.ZXX:     "Kò sí àkóònú elédè",
		},
		Territories: cldr.Territories{
			territory.V_001: "Agbáyé",
			territory.V_002: "Áfíríkà",
			territory.V_003: "Àríwá Amẹ́ríkà",
			territory.V_005: "Gúúṣù Amẹ́ríkà",
			territory.V_009: "Oceania",
			territory.V_011: "Ìwọ̀ oorùn Afíríkà",
			territory.V_013: "Ààrin Gbùgbùn Àmẹ́ríkà",
			territory.V_014: "Ìlà Oorùn Áfíríkà",
			territory.V_015: "Northern Africa",
			territory.V_017: "Ààrín gbùngbùn Afíríkà",
			territory.V_018: "Apágúúsù Áfíríkà",
			territory.V_019: "Amẹ́ríkà",
			territory.V_021: "Apáàríwá Amẹ́ríkà",
			territory.V_029: "Káríbíànù",
			territory.V_030: "Ìlà Òòrùn Eṣíà",
			territory.V_034: "Gúúṣù Eṣíà",
			territory.V_035: "Gúúṣù ìlà òòrùn Éṣíà",
			territory.V_039: "Gúúṣù Yúróòpù",
			territory.V_053: "Ọṣirélaṣíà",
			territory.V_054: "Mẹlanéṣíà",
			territory.V_057: "Agbègbè Maikironéṣíà",
			territory.V_061: "Polineṣíà",
			territory.V_142: "Asia",
			territory.V_143: "Ààrin Gbùngbùn Éṣíà",
			territory.V_145: "Ìwọ̀ Òòrùn Eṣíà",
			territory.V_150: "Europe",
			territory.V_151: "Ìlà Òrùn Yúrópù",
			territory.V_154: "Northern Europe",
			territory.V_155: "Ìwọ̀ Òòrùn Yúrópù",
			territory.V_202: "Sàhárà Áfíríkà",
			territory.V_419: "Látín Amẹ́ríkà",
			territory.AC:    "Erékùsù Ascension",
			territory.AD:    "Orílẹ́ède Ààndórà",
			territory.AE:    "Orílẹ́ède Ẹmirate ti Awọn Arabu",
			territory.AF:    "Orílẹ́ède Àfùgànístánì",
			territory.AG:    "Orílẹ́ède Ààntígúà àti Báríbúdà",
			territory.AI:    "Orílẹ́ède Ààngúlílà",
			territory.AL:    "Orílẹ́ède Àlùbàníánì",
			territory.AM:    "Orílẹ́ède Améníà",
			territory.AO:    "Orílẹ́ède Ààngólà",
			territory.AQ:    "Antakítíkà",
			territory.AR:    "Orílẹ́ède Agentínà",
			territory.AS:    "Sámóánì ti Orílẹ́ède Àméríkà",
			territory.AT:    "Orílẹ́ède Asítíríà",
			territory.AU:    "Orílẹ́ède Ástràlìá",
			territory.AW:    "Orílẹ́ède Árúbà",
			territory.AX:    "Àwọn Erékùsù ti Åland",
			territory.AZ:    "Orílẹ́ède Asẹ́bájánì",
			territory.BA:    "Orílẹ́ède Bọ̀síníà àti Ẹtisẹgófínà",
			territory.BB:    "Orílẹ́ède Bábádósì",
			territory.BD:    "Orílẹ́ède Bángáládésì",
			territory.BE:    "Orílẹ́ède Bégíọ́mù",
			territory.BF:    "Orílẹ́ède Bùùkíná Fasò",
			territory.BG:    "Orílẹ́ède Bùùgáríà",
			territory.BH:    "Orílẹ́ède Báránì",
			territory.BI:    "Orílẹ́ède Bùùrúndì",
			territory.BJ:    "Orílẹ́ède Bẹ̀nẹ̀",
			territory.BL:    "St. Barthélemy",
			territory.BM:    "Orílẹ́ède Bémúdà",
			territory.BN:    "Orílẹ́ède Búrúnẹ́lì",
			territory.BO:    "Orílẹ́ède Bọ̀lífíyà",
			territory.BQ:    "Caribbean Netherlands",
			territory.BR:    "Orilẹ̀-èdè Bàràsílì",
			territory.BS:    "Orílẹ́ède Bàhámásì",
			territory.BT:    "Orílẹ́ède Bútánì",
			territory.BV:    "Erékùsù Bouvet",
			territory.BW:    "Orílẹ́ède Bọ̀tìsúwánà",
			territory.BY:    "Orílẹ́ède Bélárúsì",
			territory.BZ:    "Orílẹ́ède Bèlísẹ̀",
			territory.CA:    "Orílẹ́ède Kánádà",
			territory.CC:    "Erékùsù Cocos (Keeling)",
			territory.CD:    "Orilẹ́ède Kóngò",
			territory.CF:    "Orílẹ́ède Àrin gùngun Áfíríkà",
			territory.CG:    "Orílẹ́ède Kóngò",
			territory.CH:    "Orílẹ́ède switiṣilandi",
			territory.CI:    "Orílẹ́ède Kóútè forà",
			territory.CK:    "Orílẹ́ède Etíokun Kùúkù",
			territory.CL:    "Orílẹ́ède ṣílè",
			territory.CM:    "Orílẹ́ède Kamerúúnì",
			territory.CN:    "Orilẹ̀-èdè Ṣáínà",
			territory.CO:    "Orílẹ́ède Kòlómíbìa",
			territory.CP:    "Erékùsù Clipperston",
			territory.CR:    "Orílẹ́ède Kuusita Ríkà",
			territory.CU:    "Orílẹ́ède Kúbà",
			territory.CV:    "Orílẹ́ède Etíokun Kápé féndè",
			territory.CW:    "Curaçao",
			territory.CX:    "Erékùsù Christmas",
			territory.CY:    "Orílẹ́ède Kúrúsì",
			territory.CZ:    "Orílẹ́ède ṣẹ́ẹ́kì",
			territory.DE:    "Orílẹèdè Jámánì",
			territory.DG:    "Diego Gaṣia",
			territory.DJ:    "Orílẹ́ède Díbọ́ótì",
			territory.DK:    "Orílẹ́ède Dẹ́mákì",
			territory.DM:    "Orílẹ́ède Dòmíníkà",
			territory.DO:    "Orilẹ́ède Dòmíníkánì",
			territory.DZ:    "Orílẹ́ède Àlùgèríánì",
			territory.EA:    "Seuta àti Melilla",
			territory.EC:    "Orílẹ́ède Ekuádò",
			territory.EE:    "Orílẹ́ède Esitonia",
			territory.EG:    "Orílẹ́ède Égípítì",
			territory.EH:    "Ìwọ̀òòrùn Sàhárà",
			territory.ER:    "Orílẹ́ède Eritira",
			territory.ES:    "Orílẹ́ède Sipani",
			territory.ET:    "Orílẹ́ède Etopia",
			territory.EU:    "Ìṣọ̀kan Yúròpù",
			territory.EZ:    "Agbègbè Euro",
			territory.FI:    "Orílẹ́ède Filandi",
			territory.FJ:    "Orílẹ́ède Fiji",
			territory.FK:    "Orílẹ́ède Etikun Fakalandi",
			territory.FM:    "Orílẹ́ède Makoronesia",
			territory.FO:    "Àwọn Erékùsù ti Faroe",
			territory.FR:    "Orílẹ́ède Faranse",
			territory.GA:    "Orílẹ́ède Gabon",
			territory.GB:    "Orílẹ́èdè Gẹ̀ẹ́sì",
			territory.GD:    "Orílẹ́ède Genada",
			territory.GE:    "Orílẹ́ède Gọgia",
			territory.GF:    "Orílẹ́ède Firenṣi Guana",
			territory.GG:    "Guernsey",
			territory.GH:    "Orílẹ́ède Gana",
			territory.GI:    "Orílẹ́ède Gibaratara",
			territory.GL:    "Orílẹ́ède Gerelandi",
			territory.GM:    "Orílẹ́ède Gambia",
			territory.GN:    "Orílẹ́ède Gene",
			territory.GP:    "Orílẹ́ède Gadelope",
			territory.GQ:    "Orílẹ́ède Ekutoria Gini",
			territory.GR:    "Orílẹ́ède Geriisi",
			territory.GS:    "Gúúsù Georgia àti Gúúsù Àwọn Erékùsù Sandwich",
			territory.GT:    "Orílẹ́ède Guatemala",
			territory.GU:    "Orílẹ́ède Guamu",
			territory.GW:    "Orílẹ́ède Gene-Busau",
			territory.GY:    "Orílẹ́ède Guyana",
			territory.HK:    "Hong Kong SAR ti Ṣáìnà",
			territory.HM:    "Erékùsù Heard àti Erékùsù McDonald",
			territory.HN:    "Orílẹ́ède Hondurasi",
			territory.HR:    "Orílẹ́ède Kòróátíà",
			territory.HT:    "Orílẹ́ède Haati",
			territory.HU:    "Orílẹ́ède Hungari",
			territory.IC:    "Ẹrékùsù Kánárì",
			territory.ID:    "Orílẹ́ède Indonesia",
			territory.IE:    "Orílẹ́ède Ailandi",
			territory.IL:    "Orílẹ́ède Iserẹli",
			territory.IM:    "Isle of Man",
			territory.IN:    "Orílẹ́ède India",
			territory.IO:    "Orílẹ́ède Etíkun Índíánì ti Ìlú Bírítísì",
			territory.IQ:    "Orílẹ́ède Iraki",
			territory.IR:    "Orílẹ́ède Irani",
			territory.IS:    "Orílẹ́ède Aṣilandi",
			territory.IT:    "Orílẹ́ède Itáli",
			territory.JE:    "Jersey",
			territory.JM:    "Orílẹ́ède Jamaika",
			territory.JO:    "Orílẹ́ède Jọdani",
			territory.JP:    "Orílẹ́ède Japani",
			territory.KE:    "Orílẹ́ède Kenya",
			territory.KG:    "Orílẹ́ède Kuriṣisitani",
			territory.KH:    "Orílẹ́ède Kàmùbódíà",
			territory.KI:    "Orílẹ́ède Kiribati",
			territory.KM:    "Orílẹ́ède Kòmòrósì",
			territory.KN:    "Orílẹ́ède Kiiti ati Neefi",
			territory.KP:    "Orílẹ́ède Guusu Kọria",
			territory.KR:    "Orílẹ́ède Ariwa Kọria",
			territory.KW:    "Orílẹ́ède Kuweti",
			territory.KY:    "Orílẹ́ède Etíokun Kámánì",
			territory.KZ:    "Orílẹ́ède Kaṣaṣatani",
			territory.LA:    "Orílẹ́ède Laosi",
			territory.LB:    "Orílẹ́ède Lebanoni",
			territory.LC:    "Orílẹ́ède Luṣia",
			territory.LI:    "Orílẹ́ède Lẹṣitẹnisiteni",
			territory.LK:    "Orílẹ́ède Siri Lanka",
			territory.LR:    "Orílẹ́ède Laberia",
			territory.LS:    "Orílẹ́ède Lesoto",
			territory.LT:    "Orílẹ́ède Lituania",
			territory.LU:    "Orílẹ́ède Lusemogi",
			territory.LV:    "Orílẹ́ède Latifia",
			territory.LY:    "Orílẹ́ède Libiya",
			territory.MA:    "Orílẹ́ède Moroko",
			territory.MC:    "Orílẹ́ède Monako",
			territory.MD:    "Orílẹ́ède Modofia",
			territory.ME:    "Montenegro",
			territory.MF:    "St. Martin",
			territory.MG:    "Orílẹ́ède Madasika",
			territory.MH:    "Orílẹ́ède Etikun Máṣali",
			territory.MK:    "Àríwá Macedonia",
			territory.ML:    "Orílẹ́ède Mali",
			territory.MM:    "Orílẹ́ède Manamari",
			territory.MN:    "Orílẹ́ède Mogolia",
			territory.MO:    "Macao SAR ti Ṣáìnà",
			territory.MP:    "Orílẹ́ède Etikun Guusu Mariana",
			territory.MQ:    "Orílẹ́ède Matinikuwi",
			territory.MR:    "Orílẹ́ède Maritania",
			territory.MS:    "Orílẹ́ède Motserati",
			territory.MT:    "Orílẹ́ède Malata",
			territory.MU:    "Orílẹ́ède Maritiusi",
			territory.MV:    "Orílẹ́ède Maladifi",
			territory.MW:    "Orílẹ́ède Malawi",
			territory.MX:    "Orílẹ́ède Mesiko",
			territory.MY:    "Orílẹ́ède Malasia",
			territory.MZ:    "Orílẹ́ède Moṣamibiku",
			territory.NA:    "Orílẹ́ède Namibia",
			territory.NC:    "Orílẹ́ède Kaledonia Titun",
			territory.NE:    "Orílẹ́ède Nàìjá",
			territory.NF:    "Orílẹ́ède Etikun Nọ́úfókì",
			territory.NG:    "Orilẹ̀-èdè Nàìjíríà",
			territory.NI:    "Orílẹ́ède NIkaragua",
			territory.NL:    "Orílẹ́ède Nedalandi",
			territory.NO:    "Orílẹ́ède Nọọwii",
			territory.NP:    "Orílẹ́ède Nepa",
			territory.NR:    "Orílẹ́ède Nauru",
			territory.NU:    "Orílẹ́ède Niue",
			territory.NZ:    "Orílẹ́ède ṣilandi Titun",
			territory.OM:    "Orílẹ́ède Ọọma",
			territory.PA:    "Orílẹ́ède Panama",
			territory.PE:    "Orílẹ́ède Peru",
			territory.PF:    "Orílẹ́ède Firenṣi Polinesia",
			territory.PG:    "Orílẹ́ède Paapu ti Giini",
			territory.PH:    "Orílẹ́ède filipini",
			territory.PK:    "Orílẹ́ède Pakisitan",
			territory.PL:    "Orílẹ́ède Polandi",
			territory.PM:    "Orílẹ́ède Pẹẹri ati mikuloni",
			territory.PN:    "Orílẹ́ède Pikarini",
			territory.PR:    "Orílẹ́ède Pọto Riko",
			territory.PS:    "Agbègbè Palẹsítíànù",
			territory.PT:    "Orílẹ́ède Pọ́túgà",
			territory.PW:    "Orílẹ́ède Paalu",
			territory.PY:    "Orílẹ́ède Paraguye",
			territory.QA:    "Orílẹ́ède Kota",
			territory.QO:    "Agbègbè Oceania",
			territory.RE:    "Orílẹ́ède Riuniyan",
			territory.RO:    "Orílẹ́ède Romaniya",
			territory.RS:    "Serbia",
			territory.RU:    "Orílẹ́ède Rọṣia",
			territory.RW:    "Orílẹ́ède Ruwanda",
			territory.SA:    "Orílẹ́ède Saudi Arabia",
			territory.SB:    "Orílẹ́ède Etikun Solomoni",
			territory.SC:    "Orílẹ́ède seṣẹlẹsi",
			territory.SD:    "Orílẹ́ède Sudani",
			territory.SE:    "Orílẹ́ède Swidini",
			territory.SG:    "Orílẹ́ède Singapo",
			territory.SH:    "Orílẹ́ède Hẹlena",
			territory.SI:    "Orílẹ́ède Silofania",
			territory.SJ:    "Svalbard & Jan Mayen",
			territory.SK:    "Orílẹ́ède Silofakia",
			territory.SL:    "Orílẹ́ède Siria looni",
			territory.SM:    "Orílẹ́ède Sani Marino",
			territory.SN:    "Orílẹ́ède Sẹnẹga",
			territory.SO:    "Orílẹ́ède Somalia",
			territory.SR:    "Orílẹ́ède Surinami",
			territory.SS:    "Gúúsù Sudan",
			territory.ST:    "Orílẹ́ède Sao tomi ati piriiṣipi",
			territory.SV:    "Orílẹ́ède Ẹẹsáfádò",
			territory.SX:    "Sint Maarten",
			territory.SY:    "Orílẹ́ède Siria",
			territory.SZ:    "Orílẹ́ède Saṣiland",
			territory.TA:    "Tristan da Kunha",
			territory.TC:    "Orílẹ́ède Tọọki ati Etikun Kakọsi",
			territory.TD:    "Orílẹ́ède ṣààdì",
			territory.TF:    "Agbègbè Gúúsù Faranṣé",
			territory.TG:    "Orílẹ́ède Togo",
			territory.TH:    "Orílẹ́ède Tailandi",
			territory.TJ:    "Orílẹ́ède Takisitani",
			territory.TK:    "Orílẹ́ède Tokelau",
			territory.TL:    "Orílẹ́ède ÌlàOòrùn Tímọ̀",
			territory.TM:    "Orílẹ́ède Tọọkimenisita",
			territory.TN:    "Orílẹ́ède Tuniṣia",
			territory.TO:    "Orílẹ́ède Tonga",
			territory.TR:    "Orílẹ́ède Tọọki",
			territory.TT:    "Orílẹ́ède Tirinida ati Tobaga",
			territory.TV:    "Orílẹ́ède Tufalu",
			territory.TW:    "Orílẹ́ède Taiwani",
			territory.TZ:    "Orílẹ́ède Tàǹsáníà",
			territory.UA:    "Orílẹ́ède Ukarini",
			territory.UG:    "Orílẹ́ède Uganda",
			territory.UM:    "Àwọn Erékùsù Kékèké Agbègbè US",
			territory.UN:    "Ìṣọ̀kan àgbáyé",
			territory.US:    "Orílẹ̀-èdè Amẹrikà",
			territory.UY:    "Orílẹ́ède Nruguayi",
			territory.UZ:    "Orílẹ́ède Nṣibẹkisitani",
			territory.VA:    "Ìlú Vatican",
			territory.VC:    "Orílẹ́ède Fisẹnnti ati Genadina",
			territory.VE:    "Orílẹ́ède Fẹnẹṣuẹla",
			territory.VG:    "Orílẹ́ède Etíkun Fágínì ti ìlú Bírítísì",
			territory.VI:    "Orílẹ́ède Etikun Fagini ti Amẹrika",
			territory.VN:    "Orílẹ́ède Fẹtinami",
			territory.VU:    "Orílẹ́ède Faniatu",
			territory.WF:    "Orílẹ́ède Wali ati futuna",
			territory.WS:    "Orílẹ́ède Samọ",
			territory.XA:    "Pseudo-Accents",
			territory.XB:    "Pseudo-Bidi",
			territory.XK:    "Kòsófò",
			territory.YE:    "Orílẹ́ède yemeni",
			territory.YT:    "Orílẹ́ède Mayote",
			territory.ZA:    "Gúúṣù Áfíríkà",
			territory.ZM:    "Orílẹ́ède ṣamibia",
			territory.ZW:    "Orílẹ́ède ṣimibabe",
			territory.ZZ:    "Àgbègbè àìmọ̀",
		},
	}
}
