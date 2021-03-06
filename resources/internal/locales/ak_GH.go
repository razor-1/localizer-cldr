// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_ak_GH() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ak_GH",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "yy/MM/dd"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "S-Ɔ", Feb: "K-Ɔ", Mar: "E-Ɔ", Apr: "E-O", May: "E-K", Jun: "O-A", Jul: "A-K", Aug: "D-Ɔ", Sep: "F-Ɛ", Oct: "Ɔ-A", Nov: "Ɔ-O", Dec: "M-Ɔ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Sanda-Ɔpɛpɔn", Feb: "Kwakwar-Ɔgyefuo", Mar: "Ebɔw-Ɔbenem", Apr: "Ebɔbira-Oforisuo", May: "Esusow Aketseaba-Kɔtɔnimba", Jun: "Obirade-Ayɛwohomumu", Jul: "Ayɛwoho-Kitawonsa", Aug: "Difuu-Ɔsandaa", Sep: "Fankwa-Ɛbɔ", Oct: "Ɔbɛsɛ-Ahinime", Nov: "Ɔberɛfɛw-Obubuo", Dec: "Mumu-Ɔpɛnimba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Kwe", Mon: "Dwo", Tue: "Ben", Wed: "Wuk", Thu: "Yaw", Fri: "Fia", Sat: "Mem"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "K", Mon: "D", Tue: "B", Wed: "W", Thu: "Y", Fri: "F", Sat: "M"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Kwesida", Mon: "Dwowda", Tue: "Benada", Wed: "Wukuda", Thu: "Yawda", Fri: "Fida", Sat: "Memeneda"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AN", PM: "EW"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AN", PM: "EW"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "¤#,##0.00", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Ɛmirete Arab Nkabɔmu Deram", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Angola Kwanza", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Ɔstrelia Dɔla", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Baren Dina", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Burundi Frank", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Botswana Pula", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Kanada Dɔla", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongo Frank", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Ɛskudo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Gyebuti Frank", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Ɔlgyeria Dina", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Egypt Pɔn", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Ɛretereya Nakfa", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Itiopia Bir", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Iro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Breten Pɔn", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Ghana Sidi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Ghana Sidi", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Gambia Dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Gini Frank", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "India Rupi", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Gyapan Yɛn", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenya Hyelen", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Komoro Frank", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Laeberia Dɔla", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lesoto Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Libya Dina", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Moroko Diram", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Madagasi Frank", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Mɔretenia Ouguiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Mɔretenia Ouguiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mɔrehyeɔs Rupi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Malawi Kwacha", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Mozambik Metical", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Namibia Dɔla", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naegyeria Naira", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Rewanda Frank", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudi Riyal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Seyhyɛls Rupi", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Sudan Dina", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Sudan Pɔn", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "St Helena Pɔn", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Somailia Hyelen", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Sao Tome ne Principe Dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Sao Tome ne Principe Dobra", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Tunisia Dina", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzania Hyelen", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Uganda Hyelen", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Amɛrika Dɔla", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Sefa", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Afrika Anaafo Rand", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambia Kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Zambia Kwacha", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwe Dɔla", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK: "Akan",
			language.AM: "Amarik",
			language.AR: "Arabik",
			language.BE: "Belarus kasa",
			language.BG: "Bɔlgeria kasa",
			language.BN: "Bengali kasa",
			language.CS: "Kyɛk kasa",
			language.DE: "Gyaaman",
			language.EL: "Greek kasa",
			language.EN: "Borɔfo",
			language.ES: "Spain kasa",
			language.FA: "Pɛɛhyia kasa",
			language.FR: "Frɛnkye",
			language.HA: "Hausa",
			language.HI: "Hindi",
			language.HU: "Hangri kasa",
			language.ID: "Indonihyia kasa",
			language.IG: "Igbo",
			language.IT: "Italy kasa",
			language.JA: "Gyapan kasa",
			language.JV: "Gyabanis kasa",
			language.KM: "Kambodia kasa",
			language.KO: "Korea kasa",
			language.MS: "Malay kasa",
			language.MY: "Bɛɛmis kasa",
			language.NE: "Nɛpal kasa",
			language.NL: "Dɛɛkye",
			language.PA: "Pungyabi kasa",
			language.PL: "Pɔland kasa",
			language.PT: "Pɔɔtugal kasa",
			language.RO: "Romenia kasa",
			language.RU: "Rahyia kasa",
			language.RW: "Rewanda kasa",
			language.SO: "Somalia kasa",
			language.SV: "Sweden kasa",
			language.TA: "Tamil kasa",
			language.TH: "Taeland kasa",
			language.TR: "Tɛɛki kasa",
			language.UK: "Ukren kasa",
			language.UR: "Urdu kasa",
			language.VI: "Viɛtnam kasa",
			language.YO: "Yoruba",
			language.ZH: "Kyaena kasa",
			language.ZU: "Zulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andora",
			territory.AE: "United Arab Emirates",
			territory.AF: "Afganistan",
			territory.AG: "Antigua ne Baabuda",
			territory.AI: "Anguila",
			territory.AL: "Albenia",
			territory.AM: "Aamenia",
			territory.AO: "Angola",
			territory.AR: "Agyɛntina",
			territory.AS: "Amɛrika Samoa",
			territory.AT: "Ɔstria",
			territory.AU: "Ɔstrelia",
			territory.AW: "Aruba",
			territory.AZ: "Azebaegyan",
			territory.BA: "Bosnia ne Hɛzegovina",
			territory.BB: "Baabados",
			territory.BD: "Bangladɛhye",
			territory.BE: "Bɛlgyium",
			territory.BF: "Bɔkina Faso",
			territory.BG: "Bɔlgeria",
			territory.BH: "Baren",
			territory.BI: "Burundi",
			territory.BJ: "Bɛnin",
			territory.BM: "Bɛmuda",
			territory.BN: "Brunae",
			territory.BO: "Bolivia",
			territory.BR: "Brazil",
			territory.BS: "Bahama",
			territory.BT: "Butan",
			territory.BW: "Bɔtswana",
			territory.BY: "Bɛlarus",
			territory.BZ: "Beliz",
			territory.CA: "Kanada",
			territory.CD: "Kongo (Zair)",
			territory.CF: "Afrika Finimfin Man",
			territory.CG: "Kongo",
			territory.CH: "Swetzaland",
			territory.CI: "La Côte d’Ivoire",
			territory.CK: "Kook Nsupɔw",
			territory.CL: "Kyili",
			territory.CM: "Kamɛrun",
			territory.CN: "Kyaena",
			territory.CO: "Kolombia",
			territory.CR: "Kɔsta Rika",
			territory.CU: "Kuba",
			territory.CV: "Kepvɛdfo Islands",
			territory.CY: "Saeprɔs",
			territory.CZ: "Kyɛk Kurokɛse",
			territory.DE: "Gyaaman",
			territory.DJ: "Gyibuti",
			territory.DK: "Dɛnmak",
			territory.DM: "Dɔmeneka",
			territory.DO: "Dɔmeneka Kurokɛse",
			territory.DZ: "Ɔlgyeria",
			territory.EC: "Ikuwadɔ",
			territory.EE: "Ɛstonia",
			territory.EG: "Nisrim",
			territory.ER: "Ɛritrea",
			territory.ES: "Spain",
			territory.ET: "Ithiopia",
			territory.FI: "Finland",
			territory.FJ: "Figyi",
			territory.FK: "Fɔlkman Aeland",
			territory.FM: "Maekronehyia",
			territory.FR: "Frɛnkyeman",
			territory.GA: "Gabɔn",
			territory.GB: "Ahendiman Nkabom",
			territory.GD: "Grenada",
			territory.GE: "Gyɔgyea",
			territory.GF: "Frɛnkye Gayana",
			territory.GH: "Gaana",
			territory.GI: "Gyebralta",
			territory.GL: "Greenman",
			territory.GM: "Gambia",
			territory.GN: "Gini",
			territory.GP: "Guwadelup",
			territory.GQ: "Gini Ikuweta",
			territory.GR: "Greekman",
			territory.GT: "Guwatemala",
			territory.GU: "Guam",
			territory.GW: "Gini Bisaw",
			territory.GY: "Gayana",
			territory.HN: "Hɔnduras",
			territory.HR: "Krowehyia",
			territory.HT: "Heiti",
			territory.HU: "Hangari",
			territory.ID: "Indɔnehyia",
			territory.IE: "Aereland",
			territory.IL: "Israel",
			territory.IN: "India",
			territory.IO: "Britenfo Hɔn Man Wɔ India Po No Mu",
			territory.IQ: "Irak",
			territory.IR: "Iran",
			territory.IS: "Aesland",
			territory.IT: "Itali",
			territory.JM: "Gyameka",
			territory.JO: "Gyɔdan",
			territory.JP: "Gyapan",
			territory.KE: "Kɛnya",
			territory.KG: "Kɛɛgestan",
			territory.KH: "Kambodia",
			territory.KI: "Kiribati",
			territory.KM: "Kɔmɔrɔs",
			territory.KN: "Saint Kitts ne Nɛves",
			territory.KP: "Etifi Koria",
			territory.KR: "Anaafo Koria",
			territory.KW: "Kuwete",
			territory.KY: "Kemanfo Islands",
			territory.KZ: "Kazakstan",
			territory.LA: "Laos",
			territory.LB: "Lɛbanɔn",
			territory.LC: "Saint Lucia",
			territory.LI: "Lektenstaen",
			territory.LK: "Sri Lanka",
			territory.LR: "Laeberia",
			territory.LS: "Lɛsutu",
			territory.LT: "Lituwenia",
			territory.LU: "Laksembɛg",
			territory.LV: "Latvia",
			territory.LY: "Libya",
			territory.MA: "Moroko",
			territory.MC: "Mɔnako",
			territory.MD: "Mɔldova",
			territory.MG: "Madagaska",
			territory.MH: "Marshall Islands",
			territory.ML: "Mali",
			territory.MM: "Miyanma",
			territory.MN: "Mɔngolia",
			territory.MP: "Northern Mariana Islands",
			territory.MQ: "Matinik",
			territory.MR: "Mɔretenia",
			territory.MS: "Mantserat",
			territory.MT: "Mɔlta",
			territory.MU: "Mɔrehyeɔs",
			territory.MV: "Maldives",
			territory.MW: "Malawi",
			territory.MX: "Mɛksiko",
			territory.MY: "Malehyia",
			territory.MZ: "Mozambik",
			territory.NA: "Namibia",
			territory.NC: "Kaledonia Foforo",
			territory.NE: "Nigyɛ",
			territory.NF: "Nɔfolk Aeland",
			territory.NG: "Naegyeria",
			territory.NI: "Nekaraguwa",
			territory.NL: "Nɛdɛland",
			territory.NO: "Nɔɔwe",
			territory.NP: "Nɛpɔl",
			territory.NR: "Naworu",
			territory.NU: "Niyu",
			territory.NZ: "Ziland Foforo",
			territory.OM: "Oman",
			territory.PA: "Panama",
			territory.PE: "Peru",
			territory.PF: "Frɛnkye Pɔlenehyia",
			territory.PG: "Papua Guinea Foforo",
			territory.PH: "Philippines",
			territory.PK: "Pakistan",
			territory.PL: "Poland",
			territory.PM: "Saint Pierre ne Miquelon",
			territory.PN: "Pitcairn",
			territory.PR: "Puɛto Riko",
			territory.PS: "Palestaen West Bank ne Gaza",
			territory.PT: "Pɔtugal",
			territory.PW: "Palau",
			territory.PY: "Paraguay",
			territory.QA: "Kata",
			territory.RE: "Reyuniɔn",
			territory.RO: "Romenia",
			territory.RU: "Rɔhyea",
			territory.RW: "Rwanda",
			territory.SA: "Saudi Arabia",
			territory.SB: "Solomon Islands",
			territory.SC: "Seyhyɛl",
			territory.SD: "Sudan",
			territory.SE: "Sweden",
			territory.SG: "Singapɔ",
			territory.SH: "Saint Helena",
			territory.SI: "Slovinia",
			territory.SK: "Slovakia",
			territory.SL: "Sierra Leone",
			territory.SM: "San Marino",
			territory.SN: "Senegal",
			territory.SO: "Somalia",
			territory.SR: "Suriname",
			territory.ST: "São Tomé and Príncipe",
			territory.SV: "Ɛl Salvadɔ",
			territory.SY: "Siria",
			territory.SZ: "Swaziland",
			territory.TC: "Turks ne Caicos Islands",
			territory.TD: "Kyad",
			territory.TG: "Togo",
			territory.TH: "Taeland",
			territory.TJ: "Tajikistan",
			territory.TK: "Tokelau",
			territory.TL: "Timɔ Boka",
			territory.TM: "Tɛkmɛnistan",
			territory.TN: "Tunihyia",
			territory.TO: "Tonga",
			territory.TR: "Tɛɛki",
			territory.TT: "Trinidad ne Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwan",
			territory.TZ: "Tanzania",
			territory.UA: "Ukren",
			territory.UG: "Uganda",
			territory.US: "Amɛrika",
			territory.UY: "Yurugwae",
			territory.UZ: "Uzbɛkistan",
			territory.VA: "Vatican Man",
			territory.VC: "Saint Vincent ne Grenadines",
			territory.VE: "Venezuela",
			territory.VG: "Britainfo Virgin Islands",
			territory.VI: "Amɛrika Virgin Islands",
			territory.VN: "Viɛtnam",
			territory.VU: "Vanuatu",
			territory.WF: "Wallis ne Futuna",
			territory.WS: "Samoa",
			territory.YE: "Yɛmen",
			territory.YT: "Mayɔte",
			territory.ZA: "Afrika Anaafo",
			territory.ZM: "Zambia",
			territory.ZW: "Zembabwe",
		},
	}
}
