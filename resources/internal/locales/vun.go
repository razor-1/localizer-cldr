// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_vun() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "vun",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Machi", Apr: "Aprilyi", May: "Mei", Jun: "Junyi", Jul: "Julyai", Aug: "Agusti", Sep: "Septemba", Oct: "Oktoba", Nov: "Novemba", Dec: "Desemba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jpi", Mon: "Jtt", Tue: "Jnn", Wed: "Jtn", Thu: "Alh", Fri: "Iju", Sat: "Jmo"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "J", Mon: "J", Tue: "J", Wed: "J", Thu: "A", Fri: "I", Sat: "J"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Jumapilyi", Mon: "Jumatatuu", Tue: "Jumanne", Wed: "Jumatanu", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Jumamosi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "utuko", PM: "kyiukonyi"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "utuko", PM: "kyiukonyi"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "¤#,##0.00", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham ya Falme za Kiarabu", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza ya Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dola ya Australia", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dinari ya Bahareni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Faranga ya Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pula ya Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dola ya Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faranga ya Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faranga ya Uswisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Renminbi ya China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Eskudo ya Kepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faranga ya Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinari ya Aljeria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Pauni ya Misri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa ya Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Bir ya Uhabeshi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pauni ya Uingereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi ya Ghana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ya Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Faranga ya Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupia ya India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Sarafu ya Kijapani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilingi ya Kenya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Faranga ya Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dola ya Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ya Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinari ya Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirham ya Moroko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Faranga ya Bukini", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ugwiya ya Moritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ugwiya ya Moritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupia ya Morisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha ya Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikali ya Msumbiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dola ya Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira ya Nijeria", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Faranga ya Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal ya Saudia", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia ya Shelisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dinari ya Sudani", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Pauni ya Sudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Pauni ya Santahelena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leoni", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilingi ya Somalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinari ya Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilingi ya Tanzania", Symbol: "TSh"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Shilingi ya Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dola ya Marekani", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Faranga CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Faranga CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Randi ya Afrika Kusini", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha ya Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha ya Zambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dola ya Zimbabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Kiakanyi",
			language.AM:  "Kiamharyi",
			language.AR:  "Kyiarabu",
			language.BE:  "Kyibelarusi",
			language.BG:  "Kyibulgaryia",
			language.BN:  "Kyibangla",
			language.CS:  "Kyichecki",
			language.DE:  "Kyijerumani",
			language.EL:  "Kyigiriki",
			language.EN:  "Kyingereza",
			language.ES:  "Kyihispania",
			language.FA:  "Kyiajemi",
			language.FR:  "Kyifaransa",
			language.HA:  "Kyihausa",
			language.HI:  "Kyihindi",
			language.HU:  "Kyihungari",
			language.ID:  "Kyiindonesia",
			language.IG:  "Kyiigbo",
			language.IT:  "Kyiitaliano",
			language.JA:  "Kyijapani",
			language.JV:  "Kyijava",
			language.KM:  "Kyikambodia",
			language.KO:  "Kyikorea",
			language.MS:  "Kyimalesia",
			language.MY:  "Kyiburma",
			language.NE:  "Kyinepali",
			language.NL:  "Kyiholanzi",
			language.PA:  "Kyipunjabi",
			language.PL:  "Kyipolandi",
			language.PT:  "Kyireno",
			language.RO:  "Kyiromania",
			language.RU:  "Kyirusi",
			language.RW:  "Kyinyarwanda",
			language.SO:  "Kyisomalyi",
			language.SV:  "Kyiswidi",
			language.TA:  "Kyitamil",
			language.TH:  "Kyitailandi",
			language.TR:  "Kyiturukyi",
			language.UK:  "Kyiukrania",
			language.UR:  "Kyiurdu",
			language.VI:  "Kyivietinamu",
			language.VUN: "Kyivunjo",
			language.YO:  "Kyiyoruba",
			language.ZH:  "Kyichina",
			language.ZU:  "Kyizulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andora",
			territory.AE: "Falme za Kiarabu",
			territory.AF: "Afuganistani",
			territory.AG: "Antigua na Barbuda",
			territory.AI: "Anguilla",
			territory.AL: "Albania",
			territory.AM: "Armenia",
			territory.AO: "Angola",
			territory.AR: "Ajentina",
			territory.AS: "Samoa ya Marekani",
			territory.AT: "Austria",
			territory.AU: "Australia",
			territory.AW: "Aruba",
			territory.AZ: "Azabajani",
			territory.BA: "Bosnia na Hezegovina",
			territory.BB: "Babadosi",
			territory.BD: "Bangladeshi",
			territory.BE: "Ubelgiji",
			territory.BF: "Bukinafaso",
			territory.BG: "Bulgaria",
			territory.BH: "Bahareni",
			territory.BI: "Burundi",
			territory.BJ: "Benini",
			territory.BM: "Bermuda",
			territory.BN: "Brunei",
			territory.BO: "Bolivia",
			territory.BR: "Brazili",
			territory.BS: "Bahama",
			territory.BT: "Butani",
			territory.BW: "Botswana",
			territory.BY: "Belarusi",
			territory.BZ: "Belize",
			territory.CA: "Kanada",
			territory.CD: "Jamhuri ya Kidemokrasia ya Kongo",
			territory.CF: "Jamhuri ya Afrika ya Kati",
			territory.CG: "Kongo",
			territory.CH: "Uswisi",
			territory.CI: "Kodivaa",
			territory.CK: "Visiwa vya Cook",
			territory.CL: "Chile",
			territory.CM: "Kameruni",
			territory.CN: "China",
			territory.CO: "Kolombia",
			territory.CR: "Kostarika",
			territory.CU: "Kuba",
			territory.CV: "Kepuvede",
			territory.CY: "Kuprosi",
			territory.CZ: "Jamhuri ya Cheki",
			territory.DE: "Ujerumani",
			territory.DJ: "Jibuti",
			territory.DK: "Denmaki",
			territory.DM: "Dominika",
			territory.DO: "Jamhuri ya Dominika",
			territory.DZ: "Aljeria",
			territory.EC: "Ekwado",
			territory.EE: "Estonia",
			territory.EG: "Misri",
			territory.ER: "Eritrea",
			territory.ES: "Hispania",
			territory.ET: "Uhabeshi",
			territory.FI: "Ufini",
			territory.FJ: "Fiji",
			territory.FK: "Visiwa vya Falkland",
			territory.FM: "Mikronesia",
			territory.FR: "Ufaransa",
			territory.GA: "Gaboni",
			territory.GB: "Uingereza",
			territory.GD: "Grenada",
			territory.GE: "Jojia",
			territory.GF: "Gwiyana ya Ufaransa",
			territory.GH: "Ghana",
			territory.GI: "Jibralta",
			territory.GL: "Grinlandi",
			territory.GM: "Gambia",
			territory.GN: "Gine",
			territory.GP: "Gwadelupe",
			territory.GQ: "Ginekweta",
			territory.GR: "Ugiriki",
			territory.GT: "Gwatemala",
			territory.GU: "Gwam",
			territory.GW: "Ginebisau",
			territory.GY: "Guyana",
			territory.HN: "Hondurasi",
			territory.HR: "Korasia",
			territory.HT: "Haiti",
			territory.HU: "Hungaria",
			territory.ID: "Indonesia",
			territory.IE: "Ayalandi",
			territory.IL: "Israeli",
			territory.IN: "India",
			territory.IO: "Eneo la Uingereza katika Bahari Hindi",
			territory.IQ: "Iraki",
			territory.IR: "Uajemi",
			territory.IS: "Aislandi",
			territory.IT: "Italia",
			territory.JM: "Jamaika",
			territory.JO: "Yordani",
			territory.JP: "Japani",
			territory.KE: "Kenya",
			territory.KG: "Kirigizistani",
			territory.KH: "Kambodia",
			territory.KI: "Kiribati",
			territory.KM: "Komoro",
			territory.KN: "Santakitzi na Nevis",
			territory.KP: "Korea Kaskazini",
			territory.KR: "Korea Kusini",
			territory.KW: "Kuwaiti",
			territory.KY: "Visiwa vya Kayman",
			territory.KZ: "Kazakistani",
			territory.LA: "Laosi",
			territory.LB: "Lebanoni",
			territory.LC: "Santalusia",
			territory.LI: "Lishenteni",
			territory.LK: "Sirilanka",
			territory.LR: "Liberia",
			territory.LS: "Lesoto",
			territory.LT: "Litwania",
			territory.LU: "Lasembagi",
			territory.LV: "Lativia",
			territory.LY: "Libya",
			territory.MA: "Moroko",
			territory.MC: "Monako",
			territory.MD: "Moldova",
			territory.MG: "Bukini",
			territory.MH: "Visiwa vya Marshal",
			territory.ML: "Mali",
			territory.MM: "Myama",
			territory.MN: "Mongolia",
			territory.MP: "Visiwa vya Mariana vya Kaskazini",
			territory.MQ: "Martiniki",
			territory.MR: "Moritania",
			territory.MS: "Montserrati",
			territory.MT: "Malta",
			territory.MU: "Morisi",
			territory.MV: "Modivu",
			territory.MW: "Malawi",
			territory.MX: "Meksiko",
			territory.MY: "Malesia",
			territory.MZ: "Msumbiji",
			territory.NA: "Namibia",
			territory.NC: "Nyukaledonia",
			territory.NE: "Nijeri",
			territory.NF: "Kisiwa cha Norfok",
			territory.NG: "Nijeria",
			territory.NI: "Nikaragwa",
			territory.NL: "Uholanzi",
			territory.NO: "Norwe",
			territory.NP: "Nepali",
			territory.NR: "Nauru",
			territory.NU: "Niue",
			territory.NZ: "Nyuzilandi",
			territory.OM: "Omani",
			territory.PA: "Panama",
			territory.PE: "Peru",
			territory.PF: "Polinesia ya Ufaransa",
			territory.PG: "Papua",
			territory.PH: "Filipino",
			territory.PK: "Pakistani",
			territory.PL: "Polandi",
			territory.PM: "Santapieri na Mikeloni",
			territory.PN: "Pitkairni",
			territory.PR: "Pwetoriko",
			territory.PS: "Ukingo wa Magharibi na Ukanda wa Gaza wa Palestina",
			territory.PT: "Ureno",
			territory.PW: "Palau",
			territory.PY: "Paragwai",
			territory.QA: "Katari",
			territory.RE: "Riyunioni",
			territory.RO: "Romania",
			territory.RU: "Urusi",
			territory.RW: "Rwanda",
			territory.SA: "Saudi",
			territory.SB: "Visiwa vya Solomon",
			territory.SC: "Shelisheli",
			territory.SD: "Sudani",
			territory.SE: "Uswidi",
			territory.SG: "Singapoo",
			territory.SH: "Santahelena",
			territory.SI: "Slovenia",
			territory.SK: "Slovakia",
			territory.SL: "Siera Leoni",
			territory.SM: "Samarino",
			territory.SN: "Senegali",
			territory.SO: "Somalia",
			territory.SR: "Surinamu",
			territory.ST: "Sao Tome na Principe",
			territory.SV: "Elsavado",
			territory.SY: "Siria",
			territory.SZ: "Uswazi",
			territory.TC: "Visiwa vya Turki na Kaiko",
			territory.TD: "Chadi",
			territory.TG: "Togo",
			territory.TH: "Tailandi",
			territory.TJ: "Tajikistani",
			territory.TK: "Tokelau",
			territory.TL: "Timori ya Mashariki",
			territory.TM: "Turukimenistani",
			territory.TN: "Tunisia",
			territory.TO: "Tonga",
			territory.TR: "Uturuki",
			territory.TT: "Trinidad na Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwani",
			territory.TZ: "Tanzania",
			territory.UA: "Ukraini",
			territory.UG: "Uganda",
			territory.US: "Marekani",
			territory.UY: "Urugwai",
			territory.UZ: "Uzibekistani",
			territory.VA: "Vatikani",
			territory.VC: "Santavisenti na Grenadini",
			territory.VE: "Venezuela",
			territory.VG: "Visiwa vya Virgin vya Uingereza",
			territory.VI: "Visiwa vya Virgin vya Marekani",
			territory.VN: "Vietinamu",
			territory.VU: "Vanuatu",
			territory.WF: "Walis na Futuna",
			territory.WS: "Samoa",
			territory.YE: "Yemeni",
			territory.YT: "Mayotte",
			territory.ZA: "Afrika Kusini",
			territory.ZM: "Zambia",
			territory.ZW: "Zimbabwe",
		},
	}
}
