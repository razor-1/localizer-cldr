// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_twq() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "twq",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Žan", Feb: "Fee", Mar: "Mar", Apr: "Awi", May: "Me", Jun: "Žuw", Jul: "Žuy", Aug: "Ut", Sep: "Sek", Oct: "Okt", Nov: "Noo", Dec: "Dee"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Ž", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Ž", Jul: "Ž", Aug: "U", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Žanwiye", Feb: "Feewiriye", Mar: "Marsi", Apr: "Awiril", May: "Me", Jun: "Žuweŋ", Jul: "Žuyye", Aug: "Ut", Sep: "Sektanbur", Oct: "Oktoobur", Nov: "Noowanbur", Dec: "Deesanbur"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Alh", Mon: "Ati", Tue: "Ata", Wed: "Ala", Thu: "Alm", Fri: "Alz", Sat: "Asi"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "H", Mon: "T", Tue: "T", Wed: "L", Thu: "L", Fri: "L", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Alhadi", Mon: "Atinni", Tue: "Atalaata", Wed: "Alarba", Thu: "Alhamiisa", Fri: "Alzuma", Sat: "Asibti"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Subbaahi", PM: "Zaarikay b"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Subbaahi", PM: "Zaarikay b"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00¤", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Laaraw Immaara Margantey Dirham", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Angoola Kwanza", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Ostraali Dollar", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Bahareen Dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Burundi Fraŋ", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Botswaana Pund", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Kanaada Dollar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongo Fraŋ", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Swisu Fraŋ", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Sinwa Yuan Renminbi", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Kapuver Escudo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Jibuuti Fraŋ", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Alžeeri Dinar", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Misra Pund", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Eritree Nafka", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ecioopi Birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Eero", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Britin Pund", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Gaana Šiidi", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Gambi Dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Ginee Fraŋ", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indu Rupii", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Jaapoŋ Yen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Keeniya Šiiliŋ", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Komoor Fraŋ", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Liberia Dollar", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Leezoto Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Liibi Dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Maarok Dirham", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Malgaaši Fraŋ", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Mooritaani Ugiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Mooritaani Ugiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mooris Rupii", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Malaawi Kwaca", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Mozambik Metikal", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Naamibi Dollar", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naajiriya Neera", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Rwanda Fraŋ", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudiya Riyal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Seešel Rupii", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Suudaŋ Dinar", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Suudaŋ Pund", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Seŋ Helena Fraŋ", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leeon", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Somaali Šiiliŋ", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Sao Tome nda Prinsipe Dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Sao Tome nda Prinsipe Dobra", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Tunizi Dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzaani Šiiliŋ", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Uganda Šiiliŋ", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Ameriki Dollar", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "CFA Fraŋ (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "CFA Fraŋ (BCEAO)", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Hawasa Afriki Rand", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambi Kwaca (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Zambi Kwaca", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwe Dollar", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Akan senni",
			language.AM:  "Amhaarik senni",
			language.AR:  "Laaraw senni",
			language.BE:  "Belaruus senni",
			language.BG:  "Bulagaari senni",
			language.BN:  "Bengali senni",
			language.CS:  "Cek senni",
			language.DE:  "Almaŋ senni",
			language.EL:  "Grek senni",
			language.EN:  "Inglisi senni",
			language.ES:  "Espaaɲe senni",
			language.FA:  "Farsi senni",
			language.FR:  "Fransee senni",
			language.HA:  "Hawsance senni",
			language.HI:  "Induu senni",
			language.HU:  "Hungaari senni",
			language.ID:  "Indoneesi senni",
			language.IG:  "Iboo senni",
			language.IT:  "Itaali senni",
			language.JA:  "Japonee senni",
			language.JV:  "Javanee senni",
			language.KM:  "Kmeer senni, Game here",
			language.KO:  "Koree senni",
			language.MS:  "Maleezi senni",
			language.MY:  "Burme senni",
			language.NE:  "Neepal senni",
			language.NL:  "Holandee senni",
			language.PA:  "Punjaabi sennii",
			language.PL:  "Polonee senni",
			language.PT:  "Portugee senni",
			language.RO:  "Rumaani senni",
			language.RU:  "Ruusi senni",
			language.RW:  "Rwanda senni",
			language.SO:  "Somaali senni",
			language.SV:  "Suweede senni",
			language.TA:  "Tamil senni",
			language.TH:  "Taailandu senni",
			language.TR:  "Turku senni",
			language.TWQ: "Tasawaq senni",
			language.UK:  "Ukreen senni",
			language.UR:  "Urdu senni",
			language.VI:  "Vietnaam senni",
			language.YO:  "Yorbance senni",
			language.ZH:  "Sinuwa senni, Mandareŋ",
			language.ZU:  "Zulu senni",
		},
		Territories: cldr.Territories{
			territory.AD: "Andoora",
			territory.AE: "Laaraw Imaarawey Margantey",
			territory.AF: "Afgaanistan",
			territory.AG: "Antigua nda Barbuuda",
			territory.AI: "Angiiya",
			territory.AL: "Albaani",
			territory.AM: "Armeeni",
			territory.AO: "Angoola",
			territory.AR: "Argentine",
			territory.AS: "Ameriki Samoa",
			territory.AT: "Otriši",
			territory.AU: "Ostraali",
			territory.AW: "Aruuba",
			territory.AZ: "Azerbaayijaŋ",
			territory.BA: "Bosni nda Herzegovine",
			territory.BB: "Barbaados",
			territory.BD: "Bangladeši",
			territory.BE: "Belgiiki",
			territory.BF: "Burkina faso",
			territory.BG: "Bulgaari",
			territory.BH: "Bahareen",
			territory.BI: "Burundi",
			territory.BJ: "Beniŋ",
			territory.BM: "Bermuda",
			territory.BN: "Bruunee",
			territory.BO: "Boolivi",
			territory.BR: "Breezil",
			territory.BS: "Bahamas",
			territory.BT: "Buutaŋ",
			territory.BW: "Botswaana",
			territory.BY: "Biloriši",
			territory.BZ: "Beliizi",
			territory.CA: "Kanaada",
			territory.CD: "Kongoo demookaratiki laboo",
			territory.CF: "Centraafriki koyra",
			territory.CG: "Kongoo",
			territory.CH: "Swisu",
			territory.CI: "Kudwar",
			territory.CK: "Kuuk gungey",
			territory.CL: "Šiili",
			territory.CM: "Kameruun",
			territory.CN: "Šiin",
			territory.CO: "Kolombi",
			territory.CR: "Kosta rika",
			territory.CU: "Kuuba",
			territory.CV: "Kapuver gungey",
			territory.CY: "Šiipur",
			territory.CZ: "Cek labo",
			territory.DE: "Almaaɲe",
			territory.DJ: "Jibuuti",
			territory.DK: "Danemark",
			territory.DM: "Doominiki",
			territory.DO: "Doominiki laboo",
			territory.DZ: "Alžeeri",
			territory.EC: "Ekwateer",
			territory.EE: "Estooni",
			territory.EG: "Misra",
			territory.ER: "Eritree",
			territory.ES: "Espaaɲe",
			territory.ET: "Ecioopi",
			territory.FI: "Finlandu",
			territory.FJ: "Fiji",
			territory.FK: "Kalkan gungey",
			territory.FM: "Mikronezi",
			territory.FR: "Faransi",
			territory.GA: "Gaabon",
			territory.GB: "Albaasalaama Marganta",
			territory.GD: "Grenaada",
			territory.GE: "Gorgi",
			territory.GF: "Faransi Guyaan",
			territory.GH: "Gaana",
			territory.GI: "Gibraltar",
			territory.GL: "Grinland",
			territory.GM: "Gambi",
			territory.GN: "Gine",
			territory.GP: "Gwadeluup",
			territory.GQ: "Ginee Ekwatorial",
			territory.GR: "Greece",
			territory.GT: "Gwatemaala",
			territory.GU: "Guam",
			territory.GW: "Gine-Bisso",
			territory.GY: "Guyaane",
			territory.HN: "Honduras",
			territory.HR: "Krwaasi",
			territory.HT: "Haiti",
			territory.HU: "Hungaari",
			territory.ID: "Indoneezi",
			territory.IE: "Irlandu",
			territory.IL: "Israyel",
			territory.IN: "Indu laboo",
			territory.IO: "Britiši Indu teekoo laama",
			territory.IQ: "Iraak",
			territory.IR: "Iraan",
			territory.IS: "Ayseland",
			territory.IT: "Itaali",
			territory.JM: "Jamaayik",
			territory.JO: "Urdun",
			territory.JP: "Jaapoŋ",
			territory.KE: "Keeniya",
			territory.KG: "Kyrgyzstan",
			territory.KH: "kamboogi",
			territory.KI: "Kiribaati",
			territory.KM: "Komoor",
			territory.KN: "Seŋ Kitts nda Nevis",
			territory.KP: "Kooree, Gurma",
			territory.KR: "Kooree, Hawsa",
			territory.KW: "Kuweet",
			territory.KY: "Kayman gungey",
			territory.KZ: "Kaazakstan",
			territory.LA: "Laawos",
			territory.LB: "Lubnaan",
			territory.LC: "Seŋ Lussia",
			territory.LI: "Liechtenstein",
			territory.LK: "Srilanka",
			territory.LR: "Liberia",
			territory.LS: "Leesoto",
			territory.LT: "Lituaani",
			territory.LU: "Luxembourg",
			territory.LV: "Letooni",
			territory.LY: "Liibi",
			territory.MA: "Maarok",
			territory.MC: "Monako",
			territory.MD: "Moldovi",
			territory.MG: "Madagascar",
			territory.MH: "Maršal gungey",
			territory.ML: "Maali",
			territory.MM: "Maynamar",
			territory.MN: "Mongooli",
			territory.MP: "Mariana Gurma Gungey",
			territory.MQ: "Martiniiki",
			territory.MR: "Mooritaani",
			territory.MS: "Montserrat",
			territory.MT: "Malta",
			territory.MU: "Mooris gungey",
			territory.MV: "Maldiivu",
			territory.MW: "Malaawi",
			territory.MX: "Mexiki",
			territory.MY: "Maleezi",
			territory.MZ: "Mozambik",
			territory.NA: "Naamibi",
			territory.NC: "Kaaledooni Taagaa",
			territory.NE: "Nižer",
			territory.NF: "Norfolk Gungoo",
			territory.NG: "Naajiriia",
			territory.NI: "Nikaragwa",
			territory.NL: "Hollandu",
			territory.NO: "Norveej",
			territory.NP: "Neepal",
			territory.NR: "Nauru",
			territory.NU: "Niue",
			territory.NZ: "Zeelandu Taaga",
			territory.OM: "Omaan",
			territory.PA: "Panama",
			territory.PE: "Peeru",
			territory.PF: "Faransi Polineezi",
			territory.PG: "Papua Ginee Taaga",
			territory.PH: "Filipine",
			territory.PK: "Paakistan",
			territory.PL: "Poloɲe",
			territory.PM: "Seŋ Piyer nda Mikelon",
			territory.PN: "Pitikarin",
			territory.PR: "Porto Riko",
			territory.PS: "Palestine Dangay nda Gaaza",
			territory.PT: "Portugaal",
			territory.PW: "Palu",
			territory.PY: "Paraguwey",
			territory.QA: "Kataar",
			territory.RE: "Reenioŋ",
			territory.RO: "Rumaani",
			territory.RU: "Iriši laboo",
			territory.RW: "Rwanda",
			territory.SA: "Saudiya",
			territory.SB: "Solomon Gungey",
			territory.SC: "Seešel",
			territory.SD: "Suudaŋ",
			territory.SE: "Sweede",
			territory.SG: "Singapur",
			territory.SH: "Seŋ Helena",
			territory.SI: "Sloveeni",
			territory.SK: "Slovaaki",
			territory.SL: "Seera Leon",
			territory.SM: "San Marino",
			territory.SN: "Senegal",
			territory.SO: "Somaali",
			territory.SR: "Surinaam",
			territory.ST: "Sao Tome nda Prinsipe",
			territory.SV: "Salvador laboo",
			territory.SY: "Suuria",
			territory.SZ: "Swaziland",
			territory.TC: "Turk nda Kayikos Gungey",
			territory.TD: "Caadu",
			territory.TG: "Togo",
			territory.TH: "Taayiland",
			territory.TJ: "Taažikistan",
			territory.TK: "Tokelau",
			territory.TL: "Timoor hawsa",
			territory.TM: "Turkmenistaŋ",
			territory.TN: "Tunizi",
			territory.TO: "Tonga",
			territory.TR: "Turki",
			territory.TT: "Trinidad nda Tobaago",
			territory.TV: "Tuvalu",
			territory.TW: "Taayiwan",
			territory.TZ: "Tanzaani",
			territory.UA: "Ukreen",
			territory.UG: "Uganda",
			territory.US: "Ameriki Laabu Margantey",
			territory.UY: "Uruguwey",
			territory.UZ: "Uzbeekistan",
			territory.VA: "Vaatikan Laama",
			territory.VC: "Seŋvinsaŋ nda Grenadine",
			territory.VE: "Veneezuyeela",
			territory.VG: "Britiši Virgin gungey",
			territory.VI: "Ameerik Virgin Gungey",
			territory.VN: "Vietnaam",
			territory.VU: "Vanautu",
			territory.WF: "Wallis nda Futuna",
			territory.WS: "Samoa",
			territory.YE: "Yaman",
			territory.YT: "Mayooti",
			territory.ZA: "Hawsa Afriki Laboo",
			territory.ZM: "Zambi",
			territory.ZW: "Zimbabwe",
		},
	}
}
