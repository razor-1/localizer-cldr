// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_kln_KE() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "kln_KE",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Mul", Feb: "Ngat", Mar: "Taa", Apr: "Iwo", May: "Mam", Jun: "Paa", Jul: "Nge", Aug: "Roo", Sep: "Bur", Oct: "Epe", Nov: "Kpt", Dec: "Kpa"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "M", Feb: "N", Mar: "T", Apr: "I", May: "M", Jun: "P", Jul: "N", Aug: "R", Sep: "B", Oct: "E", Nov: "K", Dec: "K"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Mulgul", Feb: "Ng’atyaato", Mar: "Kiptaamo", Apr: "Iwootkuut", May: "Mamuut", Jun: "Paagi", Jul: "Ng’eiyeet", Aug: "Rooptui", Sep: "Bureet", Oct: "Epeeso", Nov: "Kipsuunde ne taai", Dec: "Kipsuunde nebo aeng’"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Kts", Mon: "Kot", Tue: "Koo", Wed: "Kos", Thu: "Koa", Fri: "Kom", Sat: "Kol"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "T", Mon: "T", Tue: "O", Wed: "S", Thu: "A", Fri: "M", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Kotisap", Mon: "Kotaai", Tue: "Koaeng’", Wed: "Kosomok", Thu: "Koang’wan", Fri: "Komuut", Sat: "Kolo"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "krn", PM: "koosk"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "karoon", PM: "kooskoliny"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Rabisiekab Kibagegeitab arabuk", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Rabisiekab Angolan", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dolaitab Australian", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Rabisiekab Bahrain", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Rabisiekab Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Rabisiekab Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dolaitab Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Rabisiekab Congo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Rabisiekab Swiss", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Rabisiekab China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Rabisiekab Kepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Rabisiekab Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Rabisiekab Algerian", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Pauditab Misri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Rabisiekab Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Rabisiekab Ethiopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuroit", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "pounditab Uingereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Rabisiekab Ghana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Rabisiekab Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Rabisiekab Guinea", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rabisiekab India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Rabisiekab Japan", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Silingitab ya Kenya", Symbol: "Ksh"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Rabisiekab Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dolaitab Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Rabisiekab Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Rabisiekab Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Rabisiekab Moroccan", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Rabisiekab Malagasy", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Rabisiekab Mauritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Rabisiekab Mauritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rabisiekab Mauritius", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Rabisiekaby Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Rabisiekab Msumbiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dolaitab Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Rabisiekab Nigeria", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Rabisiekab Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Rabisiekab Saudia", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rabisiekab Shelisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Pouditab Sudan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Pouditab helena ne tilil", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leonit", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "siligitab Somalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Rabisiekab Sao Tome ak Principe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Rabisiekab Sao Tome ak Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangenit", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "RabisiekabTunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "silingitab Tanzania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Silingitab Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dolaitab ya Amareka", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Rabisiekab CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Rabisiekab CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Randitab Afrika nebo murot tai", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwachaitab Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwachaitab Zambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dolaitab ya Zimbabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "kutitab Aka",
			language.AM:  "kutitab Amariek",
			language.AR:  "kutitab Arabuk",
			language.BE:  "kutitab Belarusa",
			language.BG:  "kutitab Bulgaria",
			language.BN:  "kutitab Bengali",
			language.CS:  "kutitab Chek",
			language.DE:  "kutitab Chermani",
			language.EL:  "kutitab Greece",
			language.EN:  "kutitab Uingeresa",
			language.ES:  "kutitab Espianik",
			language.FA:  "kutitab Persia",
			language.FR:  "kutitab Kifaransa",
			language.HA:  "kutitab Hausa",
			language.HI:  "kutitab Maindiik",
			language.HU:  "kutitab Hangari",
			language.ID:  "kutitab Indonesia",
			language.IG:  "kutitab Igbo",
			language.IT:  "kutitab Talianek",
			language.JA:  "kutitap Japan",
			language.JV:  "kutitap Javanese",
			language.KLN: "Kalenjin",
			language.KM:  "kutitab Kher nebo Kwen",
			language.KO:  "kutitab Korea",
			language.MS:  "kutitab Malay",
			language.MY:  "kutitab Burma",
			language.NE:  "kutitab Nepali",
			language.NL:  "kutitab Boa",
			language.PA:  "kutitab Punjab",
			language.PL:  "kutitap Poland",
			language.PT:  "kutitab Portugal",
			language.RO:  "kutitab Romaniek",
			language.RU:  "kutitab Russia",
			language.RW:  "kutitab Kinyarwanda",
			language.SO:  "kutitab Somaliek",
			language.SV:  "kutitab Sweden",
			language.TA:  "kutitab Tamil",
			language.TH:  "kutitab Thailand",
			language.TR:  "kutitab Turkey",
			language.UK:  "kutitab Ukraine",
			language.UR:  "kutitab Urdu",
			language.VI:  "kutitab Vietnam",
			language.YO:  "kutitab Yoruba",
			language.ZH:  "kutitab China",
			language.ZU:  "kutitab Zulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Emetab Andorra",
			territory.AE: "Emetab kibagenge nebo arabuk",
			territory.AF: "Emetab Afghanistan",
			territory.AG: "Emetab Antigua ak Barbuda",
			territory.AI: "Emetab Anguilla",
			territory.AL: "Emetab Albania",
			territory.AM: "Emetab Armenia",
			territory.AO: "Emetab Angola",
			territory.AR: "Emetab Argentina",
			territory.AS: "Emetab American Samoa",
			territory.AT: "Emetab Austria",
			territory.AU: "Emetab Australia",
			territory.AW: "Emetab Aruba",
			territory.AZ: "Emetab Azerbaijan",
			territory.BA: "Emetab Bosnia ak Herzegovina",
			territory.BB: "Emetab Barbados",
			territory.BD: "Emetab Bangladesh",
			territory.BE: "Emetab Belgium",
			territory.BF: "Emetab Burkina Faso",
			territory.BG: "Emetab Bulgaria",
			territory.BH: "Emetab Bahrain",
			territory.BI: "Emetab Burundi",
			territory.BJ: "Emetab Benin",
			territory.BM: "Emetab Bermuda",
			territory.BN: "Emetab Brunei",
			territory.BO: "Emetab Bolivia",
			territory.BR: "Emetab Brazil",
			territory.BS: "Emetab Bahamas",
			territory.BT: "Emetab Bhutan",
			territory.BW: "Emetab Botswana",
			territory.BY: "Emetab Belarus",
			territory.BZ: "Emetab Belize",
			territory.CA: "Emetab Canada",
			territory.CD: "Emetab Congo - Kinshasa",
			territory.CF: "Emetab Afrika nebo Kwen",
			territory.CG: "Emetab Congo - Brazzaville",
			territory.CH: "Emetab Switzerland",
			territory.CI: "Emetab Côte d’Ivoire",
			territory.CK: "Ikwembeyotab Cook",
			territory.CL: "Emetab Chile",
			territory.CM: "Emetab Cameroon",
			territory.CN: "Emetab China",
			territory.CO: "Emetab Colombia",
			territory.CR: "Emetab Costa Rica",
			territory.CU: "Emetab Cuba",
			territory.CV: "Ikwembeyotab Cape Verde",
			territory.CY: "Emetab Cyprus",
			territory.CZ: "Emetab Czech Republic",
			territory.DE: "Emetab Geruman",
			territory.DJ: "Emetab Djibouti",
			territory.DK: "Emetab Denmark",
			territory.DM: "Emetab Dominica",
			territory.DO: "Emetab Dominican Republic",
			territory.DZ: "Emetab Algeria",
			territory.EC: "Emetab Ecuador",
			territory.EE: "Emetab Estonia",
			territory.EG: "Emetab Misiri",
			territory.ER: "Emetab Eritrea",
			territory.ES: "Emetab Spain",
			territory.ET: "Emetab Ethiopia",
			territory.FI: "Emetab Finland",
			territory.FJ: "Emetab Fiji",
			territory.FK: "Ikwembeyotab Falkland",
			territory.FM: "Emetab Micronesia",
			territory.FR: "Emetab France",
			territory.GA: "Emetab Gabon",
			territory.GB: "Emetab Kibagenge nebo Uingereza",
			territory.GD: "Emetab Grenada",
			territory.GE: "Emetab Georgia",
			territory.GF: "Emetab Guiana nebo Ufaransa",
			territory.GH: "Emetab Ghana",
			territory.GI: "Emetab Gibraltar",
			territory.GL: "Emetab Greenland",
			territory.GM: "Emetab Gambia",
			territory.GN: "Emetab Guinea",
			territory.GP: "Emetab Guadeloupe",
			territory.GQ: "Emetab Equatorial Guinea",
			territory.GR: "Emetab Greece",
			territory.GT: "Emetab Guatemala",
			territory.GU: "Emetab Guam",
			territory.GW: "Emetab Guinea-Bissau",
			territory.GY: "Emetab Guyana",
			territory.HN: "Emetab Honduras",
			territory.HR: "Emetab Croatia",
			territory.HT: "Emetab Haiti",
			territory.HU: "Emetab Hungary",
			territory.ID: "Emetab Indonesia",
			territory.IE: "Emetab Ireland",
			territory.IL: "Emetab Israel",
			territory.IN: "Emetab India",
			territory.IO: "Kebebertab araraitab indian Ocean nebo Uingeresa",
			territory.IQ: "Emetab Iraq",
			territory.IR: "Emetab Iran",
			territory.IS: "Emetab Iceland",
			territory.IT: "Emetab Italy",
			territory.JM: "Emetab Jamaica",
			territory.JO: "Emetab Jordan",
			territory.JP: "Emetab Japan",
			territory.KE: "Emetab Kenya",
			territory.KG: "Emetab Kyrgyzstan",
			territory.KH: "Emetab Cambodia",
			territory.KI: "Emetab Kiribati",
			territory.KM: "Emetab Comoros",
			territory.KN: "Emetab Saint Kitts ak Nevis",
			territory.KP: "Emetab Korea nebo murot katam",
			territory.KR: "Emetab korea nebo murot tai",
			territory.KW: "Emetab Kuwait",
			territory.KY: "Ikwembeyotab Cayman",
			territory.KZ: "Emetab Kazakhstan",
			territory.LA: "Emetab Laos",
			territory.LB: "Emetab Lebanon",
			territory.LC: "Emetab Lucia Ne",
			territory.LI: "Emetab Liechtenstein",
			territory.LK: "Emetab Sri Lanka",
			territory.LR: "Emetab Liberia",
			territory.LS: "Emetab Lesotho",
			territory.LT: "Emetab Lithuania",
			territory.LU: "Emetab Luxembourg",
			territory.LV: "Emetab Latvia",
			territory.LY: "Emetab Libya",
			territory.MA: "Emetab Morocco",
			territory.MC: "Emetab Monaco",
			territory.MD: "Emetab Moldova",
			territory.MG: "Emetab Madagascar",
			territory.MH: "Ikwembeiyotab Marshall",
			territory.ML: "Emetab Mali",
			territory.MM: "Emetab Myanmar",
			territory.MN: "Emetab Mongolia",
			territory.MP: "Ikwembeiyotab Mariana nebo murot katam",
			territory.MQ: "Emetab Martinique",
			territory.MR: "Emetab Mauritania",
			territory.MS: "Emetab Montserrat",
			territory.MT: "Emetab Malta",
			territory.MU: "Emetab Mauritius",
			territory.MV: "Emetab Maldives",
			territory.MW: "Emetab Malawi",
			territory.MX: "Emetab Mexico",
			territory.MY: "Emetab Malaysia",
			territory.MZ: "Emetab Mozambique",
			territory.NA: "Emetab Namibia",
			territory.NC: "Emetab New Caledonia",
			territory.NE: "Emetab niger",
			territory.NF: "Ikwembeiyotab Norfork",
			territory.NG: "Emetab Nigeria",
			territory.NI: "Emetab Nicaragua",
			territory.NL: "Emetab Holand",
			territory.NO: "Emetab Norway",
			territory.NP: "Emetab Nepal",
			territory.NR: "Emetab Nauru",
			territory.NU: "Emetab Niue",
			territory.NZ: "Emetab New Zealand",
			territory.OM: "Emetab Oman",
			territory.PA: "Emetab Panama",
			territory.PE: "Emetab Peru",
			territory.PF: "Emetab Polynesia nebo ufaransa",
			territory.PG: "Emetab Papua New Guinea",
			territory.PH: "Emetab Philippines",
			territory.PK: "Emetab Pakistan",
			territory.PL: "Emetab Poland",
			territory.PM: "Emetab Peter Ne titil ak Miquelon",
			territory.PN: "Emetab Pitcairn",
			territory.PR: "Emetab Puerto Rico",
			territory.PS: "Emetab Palestine",
			territory.PT: "Emetab Portugal",
			territory.PW: "Emetab Palau",
			territory.PY: "Emetab Paraguay",
			territory.QA: "Emetab Qatar",
			territory.RE: "Emetab Réunion",
			territory.RO: "Emetab Romania",
			territory.RU: "Emetab Russia",
			territory.RW: "Emetab Rwanda",
			territory.SA: "Emetab Saudi Arabia",
			territory.SB: "Ikwembeiyotab Solomon",
			territory.SC: "Emetab Seychelles",
			territory.SD: "Emetab Sudan",
			territory.SE: "Emetab Sweden",
			territory.SG: "Emetab Singapore",
			territory.SH: "Emetab Helena Ne tilil",
			territory.SI: "Emetab Slovenia",
			territory.SK: "Emetab Slovakia",
			territory.SL: "Emetab Sierra Leone",
			territory.SM: "Emetab San Marino",
			territory.SN: "Emetab Senegal",
			territory.SO: "Emetab Somalia",
			territory.SR: "Emetab Suriname",
			territory.ST: "Emetab São Tomé and Príncipe",
			territory.SV: "Emetab El Salvador",
			territory.SY: "Emetab Syria",
			territory.SZ: "Emetab Swaziland",
			territory.TC: "Ikwembeiyotab Turks ak Caicos",
			territory.TD: "Emetab Chad",
			territory.TG: "Emetab Togo",
			territory.TH: "Emetab Thailand",
			territory.TJ: "Emetab Tajikistan",
			territory.TK: "Emetab Tokelau",
			territory.TL: "Emetab Timor nebo Murot tai",
			territory.TM: "Emetab Turkmenistan",
			territory.TN: "Emetab Tunisia",
			territory.TO: "Emetab Tonga",
			territory.TR: "Emetab Turkey",
			territory.TT: "Emetab Trinidad ak Tobago",
			territory.TV: "Emetab Tuvalu",
			territory.TW: "Emetab Taiwan",
			territory.TZ: "Emetab Tanzania",
			territory.UA: "Emetab Ukrainie",
			territory.UG: "Emetab Uganda",
			territory.US: "Emetab amerika",
			territory.UY: "Emetab Uruguay",
			territory.UZ: "Emetab Uzibekistani",
			territory.VA: "Emetab Vatican",
			territory.VC: "Emetab Vincent netilil ak Grenadines",
			territory.VE: "Emetab Venezuela",
			territory.VG: "Ikwembeyotab British Virgin",
			territory.VI: "Ikwemweiyotab Amerika",
			territory.VN: "Emetab Vietnam",
			territory.VU: "Emetab Vanuatu",
			territory.WF: "Emetab Walis ak Futuna",
			territory.WS: "Emetab Samoa",
			territory.YE: "Emetab Yemen",
			territory.YT: "Emetab Mayotte",
			territory.ZA: "Emetab Afrika nebo Murot tai",
			territory.ZM: "Emetab Zambia",
			territory.ZW: "Emetab Zimbabwe",
		},
	}
}
