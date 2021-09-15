// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_eo_001() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "eo_001",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d-'a' 'de' MMMM y", Long: "y-MMMM-dd", Medium: "y-MMM-dd", Short: "yy-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "H-'a' 'horo' 'kaj' m:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "UTC{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "aŭg", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januaro", Feb: "februaro", Mar: "marto", Apr: "aprilo", May: "majo", Jun: "junio", Jul: "julio", Aug: "aŭgusto", Sep: "septembro", Oct: "oktobro", Nov: "novembro", Dec: "decembro"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "di", Mon: "lu", Tue: "ma", Wed: "me", Thu: "ĵa", Fri: "ve", Sat: "sa"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "Ĵ", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "dimanĉo", Mon: "lundo", Tue: "mardo", Wed: "merkredo", Thu: "ĵaŭdo", Fri: "vendredo", Sat: "sabato"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "atm", PM: "ptm"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "atm", PM: "ptm"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a00K", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Aŭstralia dolaro", Symbol: "AU$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Brazila realo", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Kanada dolaro", Symbol: "CA$"},
				currency.CHF: cldr.Currency{DisplayName: "Svisa franko", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Ĉina juano", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "Dana krono", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "Eŭro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Brita pundo", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Honkonga dolaro", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Indonezia rupio", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Barata rupio", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Japana eno", Symbol: "JP¥"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Sud-korea ŭono", Symbol: "₩"},
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
				currency.MXN: cldr.Currency{DisplayName: "Meksika peso", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Norvega krono", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Pola zloto", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "Rusa rublo", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Sauda rialo", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "Sveda krono", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "Taja bahto", Symbol: "฿"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Turka liro", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Nova tajvana dolaro", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "Usona dolaro", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "arĝento", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "oro", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "eŭropa monunuo", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XFO: cldr.Currency{DisplayName: "franca ora franko", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "paladio", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "plateno", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Nekonata valuto", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Sud-afrika rando", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "afara",
			language.AB:      "abĥaza",
			language.AF:      "afrikansa",
			language.AK:      "akana",
			language.AM:      "amhara",
			language.AR:      "araba",
			language.AR_001:  "moderna norma araba",
			language.ARN:     "mapuĉa",
			language.AS:      "asama",
			language.AY:      "ajmara",
			language.AZ:      "azerbajĝana",
			language.BA:      "baŝkira",
			language.BE:      "belorusa",
			language.BG:      "bulgara",
			language.BI:      "bislamo",
			language.BM:      "bambara",
			language.BN:      "bengala",
			language.BO:      "tibeta",
			language.BR:      "bretona",
			language.BRX:     "bodoa",
			language.BS:      "bosnia",
			language.CA:      "kataluna",
			language.CHR:     "ĉeroka",
			language.CKB:     "sorana",
			language.CO:      "korsika",
			language.CS:      "ĉeĥa",
			language.CY:      "kimra",
			language.DA:      "dana",
			language.DE:      "germana",
			language.DE_AT:   "aŭstra germana",
			language.DE_CH:   "svisa germana",
			language.DSB:     "malsuprasoraba",
			language.DV:      "mahla",
			language.DZ:      "dzonko",
			language.EFI:     "ibibioefika",
			language.EL:      "greka",
			language.EN:      "angla",
			language.EN_AU:   "aŭstralia angla",
			language.EN_CA:   "kanada angla",
			language.EN_GB:   "brita angla",
			language.EN_US:   "usona angla",
			language.EO:      "esperanto",
			language.ES:      "hispana",
			language.ES_419:  "amerika hispana",
			language.ES_ES:   "eŭropa hispana",
			language.ES_MX:   "meksika hispana",
			language.ET:      "estona",
			language.EU:      "eŭska",
			language.FA:      "persa",
			language.FI:      "finna",
			language.FIL:     "filipina",
			language.FJ:      "fiĝia",
			language.FO:      "feroa",
			language.FR:      "franca",
			language.FR_CA:   "kanada franca",
			language.FR_CH:   "svisa franca",
			language.FY:      "frisa",
			language.GA:      "irlanda",
			language.GD:      "gaela",
			language.GL:      "galega",
			language.GN:      "gvarania",
			language.GU:      "guĝarata",
			language.HA:      "haŭsa",
			language.HAW:     "havaja",
			language.HE:      "hebrea",
			language.HI:      "hinda",
			language.HR:      "kroata",
			language.HT:      "haitia kreola",
			language.HU:      "hungara",
			language.HY:      "armena",
			language.IA:      "interlingvao",
			language.ID:      "indonezia",
			language.IE:      "okcidentalo",
			language.IK:      "eskima",
			language.IS:      "islanda",
			language.IT:      "itala",
			language.IU:      "inuita",
			language.JA:      "japana",
			language.JV:      "java",
			language.KA:      "kartvela",
			language.KK:      "kazaĥa",
			language.KL:      "gronlanda",
			language.KM:      "kmera",
			language.KN:      "kanara",
			language.KO:      "korea",
			language.KS:      "kaŝmira",
			language.KU:      "kurda",
			language.KY:      "kirgiza",
			language.LA:      "latino",
			language.LB:      "luksemburga",
			language.LN:      "lingala",
			language.LO:      "laŭa",
			language.LT:      "litova",
			language.LV:      "latva",
			language.MG:      "malagasa",
			language.MI:      "maoria",
			language.MK:      "makedona",
			language.ML:      "malajalama",
			language.MN:      "mongola",
			language.MR:      "marata",
			language.MS:      "malaja",
			language.MT:      "malta",
			language.MUL:     "pluraj lingvoj",
			language.MY:      "birma",
			language.NA:      "naura",
			language.NB:      "dannorvega",
			language.NE:      "nepala",
			language.NL:      "nederlanda",
			language.NL_BE:   "flandra",
			language.NN:      "novnorvega",
			language.NO:      "norvega",
			language.OC:      "okcitana",
			language.OM:      "oroma",
			language.OR:      "orijo",
			language.PA:      "panĝaba",
			language.PL:      "pola",
			language.PS:      "paŝtoa",
			language.PT:      "portugala",
			language.PT_BR:   "brazilportugala",
			language.PT_PT:   "eŭropportugala",
			language.QU:      "keĉua",
			language.RM:      "romanĉa",
			language.RN:      "burunda",
			language.RO:      "rumana",
			language.RU:      "rusa",
			language.RW:      "ruanda",
			language.SA:      "sanskrito",
			language.SD:      "sinda",
			language.SG:      "sangoa",
			language.SH:      "serbo-Kroata",
			language.SI:      "sinhala",
			language.SK:      "slovaka",
			language.SL:      "slovena",
			language.SM:      "samoa",
			language.SN:      "ŝona",
			language.SO:      "somala",
			language.SQ:      "albana",
			language.SR:      "serba",
			language.SS:      "svazia",
			language.ST:      "sota",
			language.SU:      "sunda",
			language.SV:      "sveda",
			language.SW:      "svahila",
			language.TA:      "tamila",
			language.TE:      "telugua",
			language.TG:      "taĝika",
			language.TH:      "taja",
			language.TI:      "tigraja",
			language.TK:      "turkmena",
			language.TL:      "tagaloga",
			language.TLH:     "klingona",
			language.TN:      "cvana",
			language.TO:      "tongaa",
			language.TR:      "turka",
			language.TS:      "conga",
			language.TT:      "tatara",
			language.UG:      "ujgura",
			language.UK:      "ukraina",
			language.UND:     "nekonata lingvo",
			language.UR:      "urduo",
			language.UZ:      "uzbeka",
			language.VI:      "vjetnama",
			language.VO:      "volapuko",
			language.WO:      "volofa",
			language.XH:      "ksosa",
			language.YI:      "jida",
			language.YO:      "joruba",
			language.ZA:      "ĝuanga",
			language.ZH:      "ĉina",
			language.ZH_HANS: "ĉina simpligita",
			language.ZH_HANT: "ĉina tradicia",
			language.ZU:      "zulua",
			language.ZXX:     "nelingvaĵo",
		},
		Territories: cldr.Territories{
			territory.V_001: "Mondo",
			territory.AD:    "Andoro",
			territory.AE:    "Unuiĝintaj Arabaj Emirlandoj",
			territory.AF:    "Afganujo",
			territory.AG:    "Antigvo-Barbudo",
			territory.AI:    "Angvilo",
			territory.AL:    "Albanujo",
			territory.AM:    "Armenujo",
			territory.AO:    "Angolo",
			territory.AQ:    "Antarkto",
			territory.AR:    "Argentino",
			territory.AT:    "Aŭstrujo",
			territory.AU:    "Aŭstralio",
			territory.AW:    "Arubo",
			territory.AZ:    "Azerbajĝano",
			territory.BA:    "Bosnio-Hercegovino",
			territory.BB:    "Barbado",
			territory.BD:    "Bangladeŝo",
			territory.BE:    "Belgujo",
			territory.BF:    "Burkino",
			territory.BG:    "Bulgarujo",
			territory.BH:    "Barejno",
			territory.BI:    "Burundo",
			territory.BJ:    "Benino",
			territory.BM:    "Bermudoj",
			territory.BN:    "Brunejo",
			territory.BO:    "Bolivio",
			territory.BR:    "Brazilo",
			territory.BS:    "Bahamoj",
			territory.BT:    "Butano",
			territory.BW:    "Bocvano",
			territory.BY:    "Belorusujo",
			territory.BZ:    "Belizo",
			territory.CA:    "Kanado",
			territory.CF:    "Centr-Afrika Respubliko",
			territory.CG:    "Kongolo",
			territory.CH:    "Svisujo",
			territory.CI:    "Ebur-Bordo",
			territory.CK:    "Kukinsuloj",
			territory.CL:    "Ĉilio",
			territory.CM:    "Kameruno",
			territory.CN:    "Ĉinujo",
			territory.CO:    "Kolombio",
			territory.CR:    "Kostariko",
			territory.CU:    "Kubo",
			territory.CV:    "Kabo-Verdo",
			territory.CY:    "Kipro",
			territory.CZ:    "Ĉeĥujo",
			territory.DE:    "Germanujo",
			territory.DJ:    "Ĝibutio",
			territory.DK:    "Danujo",
			territory.DM:    "Dominiko",
			territory.DO:    "Domingo",
			territory.DZ:    "Alĝerio",
			territory.EC:    "Ekvadoro",
			territory.EE:    "Estonujo",
			territory.EG:    "Egipto",
			territory.EH:    "Okcidenta Saharo",
			territory.ER:    "Eritreo",
			territory.ES:    "Hispanujo",
			territory.ET:    "Etiopujo",
			territory.FI:    "Finnlando",
			territory.FJ:    "Fiĝoj",
			territory.FM:    "Mikronezio",
			territory.FO:    "Ferooj",
			territory.FR:    "Francujo",
			territory.GA:    "Gabono",
			territory.GB:    "Unuiĝinta Reĝlando",
			territory.GD:    "Grenado",
			territory.GE:    "Kartvelujo",
			territory.GF:    "Franca Gviano",
			territory.GH:    "Ganao",
			territory.GI:    "Ĝibraltaro",
			territory.GL:    "Gronlando",
			territory.GM:    "Gambio",
			territory.GN:    "Gvineo",
			territory.GP:    "Gvadelupo",
			territory.GQ:    "Ekvatora Gvineo",
			territory.GR:    "Grekujo",
			territory.GS:    "Sud-Georgio kaj Sud-Sandviĉinsuloj",
			territory.GT:    "Gvatemalo",
			territory.GU:    "Gvamo",
			territory.GW:    "Gvineo-Bisaŭo",
			territory.GY:    "Gujano",
			territory.HK:    "Honkongo",
			territory.HM:    "Herda kaj Makdonaldaj Insuloj",
			territory.HN:    "Honduro",
			territory.HR:    "Kroatujo",
			territory.HT:    "Haitio",
			territory.HU:    "Hungarujo",
			territory.ID:    "Indonezio",
			territory.IE:    "Irlando",
			territory.IL:    "Israelo",
			territory.IN:    "Hindujo",
			territory.IO:    "Brita Hindoceana Teritorio",
			territory.IQ:    "Irako",
			territory.IR:    "Irano",
			territory.IS:    "Islando",
			territory.IT:    "Italujo",
			territory.JM:    "Jamajko",
			territory.JO:    "Jordanio",
			territory.JP:    "Japanujo",
			territory.KE:    "Kenjo",
			territory.KG:    "Kirgizistano",
			territory.KH:    "Kamboĝo",
			territory.KI:    "Kiribato",
			territory.KM:    "Komoroj",
			territory.KN:    "Sent-Kristofo kaj Neviso",
			territory.KP:    "Nord-Koreo",
			territory.KR:    "Sud-Koreo",
			territory.KW:    "Kuvajto",
			territory.KY:    "Kejmanoj",
			territory.KZ:    "Kazaĥstano",
			territory.LA:    "Laoso",
			territory.LB:    "Libano",
			territory.LC:    "Sent-Lucio",
			territory.LI:    "Liĥtenŝtejno",
			territory.LK:    "Sri-Lanko",
			territory.LR:    "Liberio",
			territory.LS:    "Lesoto",
			territory.LT:    "Litovujo",
			territory.LU:    "Luksemburgo",
			territory.LV:    "Latvujo",
			territory.LY:    "Libio",
			territory.MA:    "Maroko",
			territory.MC:    "Monako",
			territory.MD:    "Moldavujo",
			territory.MG:    "Madagaskaro",
			territory.MH:    "Marŝaloj",
			territory.ML:    "Malio",
			territory.MM:    "Mjanmao",
			territory.MN:    "Mongolujo",
			territory.MP:    "Nord-Marianoj",
			territory.MQ:    "Martiniko",
			territory.MR:    "Maŭritanujo",
			territory.MT:    "Malto",
			territory.MU:    "Maŭricio",
			territory.MV:    "Maldivoj",
			territory.MW:    "Malavio",
			territory.MX:    "Meksiko",
			territory.MY:    "Malajzio",
			territory.MZ:    "Mozambiko",
			territory.NA:    "Namibio",
			territory.NC:    "Nov-Kaledonio",
			territory.NE:    "Niĝero",
			territory.NF:    "Norfolkinsulo",
			territory.NG:    "Niĝerio",
			territory.NI:    "Nikaragvo",
			territory.NL:    "Nederlando",
			territory.NO:    "Norvegujo",
			territory.NP:    "Nepalo",
			territory.NR:    "Nauro",
			territory.NU:    "Niuo",
			territory.NZ:    "Nov-Zelando",
			territory.OM:    "Omano",
			territory.PA:    "Panamo",
			territory.PE:    "Peruo",
			territory.PF:    "Franca Polinezio",
			territory.PG:    "Papuo-Nov-Gvineo",
			territory.PH:    "Filipinoj",
			territory.PK:    "Pakistano",
			territory.PL:    "Pollando",
			territory.PM:    "Sent-Piero kaj Mikelono",
			territory.PN:    "Pitkarna Insulo",
			territory.PR:    "Puerto-Riko",
			territory.PT:    "Portugalujo",
			territory.PW:    "Belaŭo",
			territory.PY:    "Paragvajo",
			territory.QA:    "Kataro",
			territory.RE:    "Reunio",
			territory.RO:    "Rumanujo",
			territory.RU:    "Rusujo",
			territory.RW:    "Ruando",
			territory.SA:    "Saŭda Arabujo",
			territory.SB:    "Salomonoj",
			territory.SC:    "Sejŝeloj",
			territory.SD:    "Sudano",
			territory.SE:    "Svedujo",
			territory.SG:    "Singapuro",
			territory.SH:    "Sent-Heleno",
			territory.SI:    "Slovenujo",
			territory.SJ:    "Svalbardo kaj Jan-Majen-insulo",
			territory.SK:    "Slovakujo",
			territory.SL:    "Siera-Leono",
			territory.SM:    "San-Marino",
			territory.SN:    "Senegalo",
			territory.SO:    "Somalujo",
			territory.SR:    "Surinamo",
			territory.SS:    "Sud-Sudano",
			territory.ST:    "Sao-Tomeo kaj Principeo",
			territory.SV:    "Salvadoro",
			territory.SY:    "Sirio",
			territory.SZ:    "Svazilando",
			territory.TD:    "Ĉado",
			territory.TG:    "Togolo",
			territory.TH:    "Tajlando",
			territory.TJ:    "Taĝikujo",
			territory.TM:    "Turkmenujo",
			territory.TN:    "Tunizio",
			territory.TO:    "Tongo",
			territory.TR:    "Turkujo",
			territory.TT:    "Trinidado kaj Tobago",
			territory.TV:    "Tuvalo",
			territory.TW:    "Tajvano",
			territory.TZ:    "Tanzanio",
			territory.UA:    "Ukrajno",
			territory.UG:    "Ugando",
			territory.UM:    "Usonaj malgrandaj insuloj",
			territory.US:    "Usono",
			territory.UY:    "Urugvajo",
			territory.UZ:    "Uzbekujo",
			territory.VA:    "Vatikano",
			territory.VC:    "Sent-Vincento kaj la Grenadinoj",
			territory.VE:    "Venezuelo",
			territory.VG:    "Britaj Virgulininsuloj",
			territory.VI:    "Usonaj Virgulininsuloj",
			territory.VN:    "Vjetnamo",
			territory.VU:    "Vanuatuo",
			territory.WF:    "Valiso kaj Futuno",
			territory.WS:    "Samoo",
			territory.YE:    "Jemeno",
			territory.YT:    "Majoto",
			territory.ZA:    "Sud-Afriko",
			territory.ZM:    "Zambio",
			territory.ZW:    "Zimbabvo",
			territory.ZZ:    "nekonata regiono",
		},
	}
}
