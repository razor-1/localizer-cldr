// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_shi_Latn() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "shi_Latn",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "inn", Feb: "bṛa", Mar: "maṛ", Apr: "ibr", May: "may", Jun: "yun", Jul: "yul", Aug: "ɣuc", Sep: "cut", Oct: "ktu", Nov: "nuw", Dec: "duj"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "i", Feb: "b", Mar: "m", Apr: "i", May: "m", Jun: "y", Jul: "y", Aug: "ɣ", Sep: "c", Oct: "k", Nov: "n", Dec: "d"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "innayr", Feb: "bṛayṛ", Mar: "maṛṣ", Apr: "ibrir", May: "mayyu", Jun: "yunyu", Jul: "yulyuz", Aug: "ɣuct", Sep: "cutanbir", Oct: "ktubr", Nov: "nuwanbir", Dec: "dujanbir"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "asa", Mon: "ayn", Tue: "asi", Wed: "akṛ", Thu: "akw", Fri: "asim", Sat: "asiḍ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "asamas", Mon: "aynas", Tue: "asinas", Wed: "akṛas", Thu: "akwas", Fri: "asimwas", Sat: "asiḍyas"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "tifawt", PM: "tadggʷat"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "tifawt", PM: "tadggʷat"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "#,##0.00¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "adrim n limarat", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "kwanza n angula", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "adular n ustralya", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "adinar n bḥrayn", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "frank n burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "abula n butswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "adular n kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "frank n kungu", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "afrank n swisra", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "ayan n ccinwa", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "iskudu n kabbirdi", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "frank n djibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "adinar n dzayr", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "ajnih n miṣṛ", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "nafka n iritirya", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "bir n ityubya", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "uru", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "ajnih astrlini n nngliz", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "sidi n ɣana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi n gambya", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "frank n ɣinya", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "arubi n lhind", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "ayan n lyaban", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "acilin n kinya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "frank n qumuṛ", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "adular n libirya", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "luti n liṣuṭu", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "adinar n libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "adrim n lmɣrib", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "frank n madaɣacqar", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "uqiyya n muṛiṭanya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "uqiyya n muṛiṭanya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "arubi n muris", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "kwaca n malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "amitikl n muznbiq", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "adular n namibya", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "nayra n nijirya", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "afrank n rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "aryal n ssaɛudiya", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "arubi n ssicil", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "adinar n ssudan", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "ajnih n ssudan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "ajnih n santilin", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "liyun", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "acilin n ṣṣumal", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "adubra n sanṭumi (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "adubra n sanṭumi", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "lilanjini", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "adinar n tuns", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "acilin n ṭanẓanya", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "acilin n uɣanda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "adular n iwunak imunn", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "frank ṣifa", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "frank ṣifa bisaw", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "arand n afriqya n iffus", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "akwaca n zambya (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "akwaca n zambya", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "adular n zimbabwi", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Takant",
			language.AM:  "Tamharit",
			language.AR:  "Taɛrabt",
			language.BE:  "Tabilarust",
			language.BG:  "Tablɣarit",
			language.BN:  "Tabnɣalit",
			language.CS:  "Tatcikit",
			language.DE:  "Talimant",
			language.EL:  "Tagrigit",
			language.EN:  "Tanglizt",
			language.ES:  "Tasbnyulit",
			language.FA:  "Tafursit",
			language.FR:  "Tafransist",
			language.HA:  "Tahawsat",
			language.HI:  "Tahindit",
			language.HU:  "Tahnɣarit",
			language.ID:  "Tandunisit",
			language.IG:  "Tigbut",
			language.IT:  "Taṭalyant",
			language.JA:  "Tajabbunit",
			language.JV:  "Tajavanit",
			language.KM:  "Taxmirt",
			language.KO:  "Takurit",
			language.MS:  "Tamalawit",
			language.MY:  "Tabirmanit",
			language.NE:  "Tanibalit",
			language.NL:  "Tahulandit",
			language.PA:  "Tabnjabit",
			language.PL:  "Tabulunit",
			language.PT:  "Tabṛṭqizt",
			language.RO:  "Tarumanit",
			language.RU:  "Tarusit",
			language.RW:  "Taruwandit",
			language.SHI: "Tashelḥiyt",
			language.SO:  "Tasumalit",
			language.SV:  "Taswidit",
			language.TA:  "Tatamilt",
			language.TH:  "Tataylandit",
			language.TR:  "Taturkit",
			language.UK:  "Tukranit",
			language.UR:  "Turdut",
			language.VI:  "Tafitnamit",
			language.YO:  "Tayrubat",
			language.ZH:  "Tacinwit",
			language.ZU:  "Tazulut",
		},
		Territories: cldr.Territories{
			territory.AD: "andura",
			territory.AE: "limarat",
			territory.AF: "afɣanistan",
			territory.AG: "antiga d brbuda",
			territory.AI: "angila",
			territory.AL: "albanya",
			territory.AM: "arminya",
			territory.AO: "angula",
			territory.AR: "arjantin",
			territory.AS: "samwa tamirikanit",
			territory.AT: "nnmsa",
			territory.AU: "ustralya",
			territory.AW: "aruba",
			territory.AZ: "adrabijan",
			territory.BA: "busna d hirsik",
			territory.BB: "barbad",
			territory.BD: "bangladic",
			territory.BE: "bljika",
			territory.BF: "burkina fasu",
			territory.BG: "blɣara",
			territory.BH: "bḥrayn",
			territory.BI: "burundi",
			territory.BJ: "binin",
			territory.BM: "brmuda",
			territory.BN: "bruni",
			territory.BO: "bulibya",
			territory.BR: "brazil",
			territory.BS: "bahamas",
			territory.BT: "bhutan",
			territory.BW: "butswana",
			territory.BY: "bilarusya",
			territory.BZ: "biliz",
			territory.CA: "kanada",
			territory.CD: "tagdudant tadimukratit n Kongo",
			territory.CF: "tagdudant tanammast n ifriqya",
			territory.CG: "kungu",
			territory.CH: "swisra",
			territory.CI: "kut difwar",
			territory.CK: "tigzirin n kuk",
			territory.CL: "ccili",
			territory.CM: "kamirun",
			territory.CN: "ccinwa",
			territory.CO: "culumbya",
			territory.CR: "kusta rika",
			territory.CU: "kuba",
			territory.CV: "tigzirin n kabbirdi",
			territory.CY: "qubrus",
			territory.CZ: "tagdudant tatcikit",
			territory.DE: "almanya",
			territory.DJ: "djibuti",
			territory.DK: "danmark",
			territory.DM: "duminik",
			territory.DO: "tagdudant taduminikt",
			territory.DZ: "dzayr",
			territory.EC: "ikwadur",
			territory.EE: "istunya",
			territory.EG: "miṣṛ",
			territory.ER: "iritirya",
			territory.ES: "sbanya",
			territory.ET: "ityubya",
			territory.FI: "fillanda",
			territory.FJ: "fidji",
			territory.FK: "tigzirin n malawi",
			territory.FM: "mikrunizya",
			territory.FR: "fransa",
			territory.GA: "gabun",
			territory.GB: "tagldit imunn",
			territory.GD: "ɣrnaṭa",
			territory.GE: "jurjya",
			territory.GF: "gwiyan tafransist",
			territory.GH: "ɣana",
			territory.GI: "adrar n ṭaṛiq",
			territory.GL: "griland",
			territory.GM: "gambya",
			territory.GN: "ɣinya",
			territory.GP: "gwadalub",
			territory.GQ: "ɣinya n ikwadur",
			territory.GR: "lyunan",
			territory.GT: "gwatimala",
			territory.GU: "gwam",
			territory.GW: "ɣinya bisaw",
			territory.GY: "gwiyana",
			territory.HN: "hunduras",
			territory.HR: "krwatya",
			territory.HT: "hayti",
			territory.HU: "hnɣarya",
			territory.ID: "andunisya",
			territory.IE: "irlanda",
			territory.IL: "israyil",
			territory.IN: "lhind",
			territory.IO: "tamnaḍt tanglizit n ugaru ahindi",
			territory.IQ: "lɛiraq",
			territory.IR: "iran",
			territory.IS: "island",
			territory.IT: "iṭalya",
			territory.JM: "jamayka",
			territory.JO: "lurdun",
			territory.JP: "lyaban",
			territory.KE: "kinya",
			territory.KG: "kirɣizistan",
			territory.KH: "kambudya",
			territory.KI: "kiribati",
			territory.KM: "cumur",
			territory.KN: "sankris d nifis",
			territory.KP: "kurya n iẓẓlmḍ",
			territory.KR: "kurya n iffus",
			territory.KW: "lkwit",
			territory.KY: "tigzirin n kayman",
			territory.KZ: "kazaxstan",
			territory.LA: "laws",
			territory.LB: "lubnan",
			territory.LC: "santlusi",
			territory.LI: "likinctayn",
			territory.LK: "srilanka",
			territory.LR: "libirya",
			territory.LS: "liṣuṭu",
			territory.LT: "litwanya",
			territory.LU: "luksanburg",
			territory.LV: "latfya",
			territory.LY: "libya",
			territory.MA: "lmɣrib",
			territory.MC: "munaku",
			territory.MD: "muldufya",
			territory.MG: "madaɣacqar",
			territory.MH: "tigzirin n marcal",
			territory.ML: "mali",
			territory.MM: "myanmar",
			territory.MN: "mnɣulya",
			territory.MP: "tigzirin n maryan n iẓẓlmḍ",
			territory.MQ: "martinik",
			territory.MR: "muṛiṭanya",
			territory.MS: "munsirat",
			territory.MT: "malṭa",
			territory.MU: "muris",
			territory.MV: "maldif",
			territory.MW: "malawi",
			territory.MX: "miksik",
			territory.MY: "malizya",
			territory.MZ: "muznbiq",
			territory.NA: "namibya",
			territory.NC: "kalidunya tamaynut",
			territory.NE: "nnijir",
			territory.NF: "tigzirin n nurfulk",
			territory.NG: "nijirya",
			territory.NI: "nikaragwa",
			territory.NL: "hulanda",
			territory.NO: "nnrwij",
			territory.NP: "nibal",
			territory.NR: "nawru",
			territory.NU: "niwi",
			territory.NZ: "nyuzilanda",
			territory.OM: "ɛuman",
			territory.PA: "banama",
			territory.PE: "biru",
			territory.PF: "bulinizya tafransist",
			territory.PG: "babwa ɣinya tamaynut",
			territory.PH: "filibbin",
			territory.PK: "bakistan",
			territory.PL: "bulunya",
			territory.PM: "sanbyir d miklun",
			territory.PN: "bitkayrn",
			territory.PR: "burtu riku",
			territory.PS: "agmmaḍ n tagut d ɣzza",
			territory.PT: "bṛṭqiz",
			territory.PW: "balaw",
			territory.PY: "baragway",
			territory.QA: "qatar",
			territory.RE: "riyunyun",
			territory.RO: "rumanya",
			territory.RU: "rusya",
			territory.RW: "rwanda",
			territory.SA: "ssaɛudiya",
			territory.SB: "tigzirin n saluman",
			territory.SC: "ssicil",
			territory.SD: "ssudan",
			territory.SE: "sswid",
			territory.SG: "snɣafura",
			territory.SH: "santilin",
			territory.SI: "slufinya",
			territory.SK: "slufakya",
			territory.SL: "ssiralyun",
			territory.SM: "sanmarinu",
			territory.SN: "ssinigal",
			territory.SO: "ṣṣumal",
			territory.SR: "surinam",
			territory.ST: "sawṭumi d bransib",
			territory.SV: "salfadur",
			territory.SY: "surya",
			territory.SZ: "swazilanda",
			territory.TC: "tigzirin n turkya d kayk",
			territory.TD: "tcad",
			territory.TG: "ṭugu",
			territory.TH: "ṭayland",
			territory.TJ: "tadjakistan",
			territory.TK: "ṭuklaw",
			territory.TL: "timur n lqblt",
			territory.TM: "turkmanstan",
			territory.TN: "tuns",
			territory.TO: "ṭunga",
			territory.TR: "turkya",
			territory.TT: "trinidad d ṭubagu",
			territory.TV: "tufalu",
			territory.TW: "ṭaywan",
			territory.TZ: "ṭanẓanya",
			territory.UA: "ukranya",
			territory.UG: "uɣanda",
			territory.US: "iwunak munnin n mirikan",
			territory.UY: "urugway",
			territory.UZ: "uzbakistan",
			territory.VA: "awank n fatikan",
			territory.VC: "sanfansan d grinadin",
			territory.VE: "finzwila",
			territory.VG: "tigzirin timgad n nngliz",
			territory.VI: "tigzirin timgad n iwunak munnin",
			territory.VN: "fitnam",
			territory.VU: "fanwaṭu",
			territory.WF: "walis d futuna",
			territory.WS: "samwa",
			territory.YE: "yaman",
			territory.YT: "mayuṭ",
			territory.ZA: "afriqya n iffus",
			territory.ZM: "zambya",
			territory.ZW: "zimbabwi",
		},
	}
}
