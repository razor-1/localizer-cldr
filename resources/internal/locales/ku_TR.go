// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_ku_TR() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ku_TR",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "rêb", Feb: "reş", Mar: "ada", Apr: "avr", May: "gul", Jun: "pûş", Jul: "tîr", Aug: "gel", Sep: "rez", Oct: "kew", Nov: "ser", Dec: "ber"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "R", Feb: "R", Mar: "A", Apr: "A", May: "G", Jun: "P", Jul: "T", Aug: "G", Sep: "R", Oct: "K", Nov: "S", Dec: "B"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "rêbendan", Feb: "reşemî", Mar: "adar", Apr: "avrêl", May: "gulan", Jun: "pûşper", Jul: "tîrmeh", Aug: "gelawêj", Sep: "rezber", Oct: "kewçêr", Nov: "sermawez", Dec: "berfanbar"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "yş", Mon: "dş", Tue: "sş", Wed: "çş", Thu: "pş", Fri: "în", Sat: "ş"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Y", Mon: "D", Tue: "S", Wed: "Ç", Thu: "P", Fri: "Î", Sat: "Ş"}, Short: cldr.CalendarDayFormatNameValue{Sun: "yş", Mon: "dş", Tue: "sş", Wed: "çş", Thu: "pş", Fri: "în", Sat: "ş"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "yekşem", Mon: "duşem", Tue: "sêşem", Wed: "çarşem", Thu: "pêncşem", Fri: "în", Sat: "şemî"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "BN", PM: "PN"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "BN", PM: "PN"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤;(#,##0.00\u00a0¤)", Percent: "%#,##0"},
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
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
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
				currency.EUR: cldr.Currency{DisplayName: "ewro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
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
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
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
				currency.USD: cldr.Currency{DisplayName: "", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "afarî",
			language.AB:      "abxazî",
			language.ACE:     "açehî",
			language.ADY:     "adîgeyî",
			language.AF:      "afrîkansî",
			language.AIN:     "aynuyî",
			language.ALE:     "alêwîtî",
			language.AM:      "amharî",
			language.AN:      "aragonî",
			language.AR:      "erebî",
			language.AR_001:  "erebiya standard",
			language.AS:      "asamî",
			language.AST:     "astûrî",
			language.AV:      "avarî",
			language.AY:      "aymarayî",
			language.AZ:      "azerî",
			language.BA:      "başkîrî",
			language.BAN:     "balînî",
			language.BE:      "belarusî",
			language.BEM:     "bembayî",
			language.BG:      "bulgarî",
			language.BHO:     "bojpûrî",
			language.BI:      "bîslamayî",
			language.BLA:     "blakfotî",
			language.BM:      "bambarayî",
			language.BN:      "bengalî",
			language.BO:      "tîbetî",
			language.BR:      "bretonî",
			language.BS:      "bosnî",
			language.BUG:     "bugî",
			language.CA:      "katalanî",
			language.CE:      "çeçenî",
			language.CEB:     "sebwanoyî",
			language.CH:      "çamoroyî",
			language.CHK:     "çûkî",
			language.CHM:     "marî",
			language.CHR:     "çerokî",
			language.CHY:     "çeyenî",
			language.CKB:     "soranî",
			language.CO:      "korsîkayî",
			language.CS:      "çekî",
			language.CV:      "çuvaşî",
			language.CY:      "weylsî",
			language.DA:      "danmarkî",
			language.DE:      "elmanî",
			language.DSB:     "sorbiya jêrîn",
			language.DUA:     "diwalayî",
			language.DV:      "divehî",
			language.DZ:      "conxayî",
			language.EE:      "eweyî",
			language.EL:      "yewnanî",
			language.EN:      "îngilîzî",
			language.EO:      "esperantoyî",
			language.ES:      "spanî",
			language.ET:      "estonî",
			language.EU:      "baskî",
			language.FA:      "farisî",
			language.FF:      "fulahî",
			language.FI:      "fînî",
			language.FIL:     "fîlîpînoyî",
			language.FJ:      "fîjî",
			language.FO:      "ferî",
			language.FR:      "frensî",
			language.FUR:     "friyolî",
			language.FY:      "frîsî",
			language.GA:      "îrî",
			language.GD:      "gaelîka skotî",
			language.GIL:     "kîrîbatî",
			language.GL:      "galîsî",
			language.GN:      "guwaranî",
			language.GOR:     "gorontaloyî",
			language.GSW:     "elmanîşî",
			language.GU:      "gujaratî",
			language.GV:      "manksî",
			language.HA:      "hawsayî",
			language.HAW:     "hawayî",
			language.HE:      "îbranî",
			language.HI:      "hindî",
			language.HIL:     "hîlîgaynonî",
			language.HR:      "xirwatî",
			language.HSB:     "sorbiya jorîn",
			language.HT:      "haîtî",
			language.HU:      "mecarî",
			language.HY:      "ermenî",
			language.HZ:      "hereroyî",
			language.IA:      "interlingua",
			language.ID:      "indonezî",
			language.IG:      "îgboyî",
			language.ILO:     "îlokanoyî",
			language.INH:     "îngûşî",
			language.IO:      "îdoyî",
			language.IS:      "îzlendî",
			language.IT:      "îtalî",
			language.IU:      "înuîtî",
			language.JA:      "japonî",
			language.JBO:     "lojbanî",
			language.JV:      "javayî",
			language.KA:      "gurcî",
			language.KAB:     "kabîlî",
			language.KEA:     "kapverdî",
			language.KK:      "qazaxî",
			language.KL:      "kalalîsûtî",
			language.KM:      "ximêrî",
			language.KN:      "kannadayî",
			language.KO:      "koreyî",
			language.KOK:     "konkanî",
			language.KS:      "keşmîrî",
			language.KSH:     "rîpwarî",
			language.KU:      "kurdî",
			language.KV:      "komî",
			language.KW:      "kornî",
			language.KY:      "kirgizî",
			language.LAD:     "ladînoyî",
			language.LB:      "luksembûrgî",
			language.LEZ:     "lezgînî",
			language.LG:      "lugandayî",
			language.LI:      "lîmbûrgî",
			language.LKT:     "lakotayî",
			language.LN:      "lingalayî",
			language.LO:      "lawsî",
			language.LRC:     "luriya bakur",
			language.LT:      "lîtwanî",
			language.LV:      "latviyayî",
			language.MAD:     "madurayî",
			language.MAS:     "masayî",
			language.MDF:     "mokşayî",
			language.MG:      "malagasî",
			language.MH:      "marşalî",
			language.MI:      "maorî",
			language.MIC:     "mîkmakî",
			language.MIN:     "mînangkabawî",
			language.MK:      "makedonî",
			language.ML:      "malayalamî",
			language.MN:      "mongolî",
			language.MOH:     "mohawkî",
			language.MR:      "maratî",
			language.MS:      "malezî",
			language.MT:      "maltayî",
			language.MY:      "burmayî",
			language.MYV:     "erzayî",
			language.MZN:     "mazenderanî",
			language.NA:      "nawrûyî",
			language.NAP:     "napolîtanî",
			language.NB:      "norwecî (bokmål)",
			language.NE:      "nepalî",
			language.NIU:     "nîwî",
			language.NL:      "holendî",
			language.NL_BE:   "flamî",
			language.NN:      "norwecî (nynorsk)",
			language.NSO:     "sotoyiya bakur",
			language.NV:      "navajoyî",
			language.OC:      "oksîtanî",
			language.OM:      "oromoyî",
			language.OR:      "oriyayî",
			language.OS:      "osetî",
			language.PA:      "puncabî",
			language.PAM:     "kapampanganî",
			language.PAP:     "papyamentoyî",
			language.PAU:     "palawî",
			language.PL:      "polonî",
			language.PRG:     "prûsyayî",
			language.PS:      "peştûyî",
			language.PT:      "portugalî",
			language.QU:      "keçwayî",
			language.RAP:     "rapanuyî",
			language.RAR:     "rarotongî",
			language.RM:      "romancî",
			language.RO:      "romanî",
			language.RU:      "rusî",
			language.RUP:     "aromanî",
			language.RW:      "kînyariwandayî",
			language.SA:      "sanskrîtî",
			language.SC:      "sardînî",
			language.SCN:     "sicîlî",
			language.SCO:     "skotî",
			language.SD:      "sindhî",
			language.SE:      "samiya bakur",
			language.SI:      "kîngalî",
			language.SK:      "slovakî",
			language.SL:      "slovenî",
			language.SM:      "samoayî",
			language.SMN:     "samiya înarî",
			language.SN:      "şonayî",
			language.SO:      "somalî",
			language.SQ:      "elbanî",
			language.SR:      "sirbî",
			language.SRN:     "sirananî",
			language.SS:      "swazî",
			language.ST:      "sotoyiya başûr",
			language.SU:      "sundanî",
			language.SV:      "swêdî",
			language.SW:      "swahîlî",
			language.SWB:     "komorî",
			language.SYR:     "siryanî",
			language.TA:      "tamîlî",
			language.TE:      "telûgûyî",
			language.TET:     "tetûmî",
			language.TG:      "tacikî",
			language.TH:      "tayî",
			language.TI:      "tigrînî",
			language.TK:      "tirkmenî",
			language.TLH:     "klîngonî",
			language.TN:      "tswanayî",
			language.TO:      "tongî",
			language.TPI:     "tokpisinî",
			language.TR:      "tirkî",
			language.TRV:     "tarokoyî",
			language.TS:      "tsongayî",
			language.TT:      "teterî",
			language.TUM:     "tumbukayî",
			language.TVL:     "tuvalûyî",
			language.TY:      "tahîtî",
			language.TZM:     "temazîxtî",
			language.UDM:     "udmurtî",
			language.UG:      "oygurî",
			language.UK:      "ukraynî",
			language.UND:     "zimanê nenas",
			language.UR:      "urdûyî",
			language.UZ:      "ozbekî",
			language.VI:      "viyetnamî",
			language.VO:      "volapûkî",
			language.WA:      "walonî",
			language.WAR:     "warayî",
			language.WO:      "wolofî",
			language.XH:      "xosayî",
			language.YI:      "yidîşî",
			language.YO:      "yorubayî",
			language.YUE:     "kantonî",
			language.ZH:      "çînî, mandarînî",
			language.ZH_HANS: "çîniya mandarînî ya sadekirî",
			language.ZH_HANT: "çîniya mandarînî ya kevneşopî",
			language.ZU:      "zuluyî",
			language.ZZA:     "zazakî",
		},
		Territories: cldr.Territories{
			territory.V_001: "Cîhan",
			territory.V_002: "Afrîka",
			territory.V_003: "Amerîkaya Bakur",
			territory.V_005: "Amerîkaya Başûr",
			territory.V_009: "Okyanûsya",
			territory.V_011: "Afrîkaya Rojava",
			territory.V_013: "Amerîkaya Navîn",
			territory.V_014: "Afrîkaya Rojhilat",
			territory.V_015: "Afrîkaya Bakur",
			territory.V_017: "Afrîkaya Navîn",
			territory.V_019: "Amerîka",
			territory.V_029: "Karîb",
			territory.V_053: "Awistralasya",
			territory.V_054: "Melanezya",
			territory.V_057: "Herêma Mîkronezya",
			territory.V_061: "Polînezya",
			territory.V_142: "Asya",
			territory.V_150: "Ewropa",
			territory.V_151: "Ewropaya Rojhilat",
			territory.V_155: "Ewropaya Rojava",
			territory.V_419: "Amerîkaya Latînî",
			territory.AD:    "Andorra",
			territory.AE:    "Emîrtiyên Erebî yên Yekbûyî",
			territory.AF:    "Efxanistan",
			territory.AG:    "Antîgua û Berbûda",
			territory.AL:    "Albanya",
			territory.AM:    "Ermenistan",
			territory.AO:    "Angola",
			territory.AQ:    "Antarktîka",
			territory.AR:    "Arjentîn",
			territory.AS:    "Samoaya Amerîkanî",
			territory.AT:    "Awistirya",
			territory.AU:    "Awistralya",
			territory.AW:    "Arûba",
			territory.AZ:    "Azerbaycan",
			territory.BA:    "Bosniya û Herzegovîna",
			territory.BB:    "Barbados",
			territory.BD:    "Bangladeş",
			territory.BE:    "Belçîka",
			territory.BF:    "Burkîna Faso",
			territory.BG:    "Bulgaristan",
			territory.BH:    "Behreyn",
			territory.BI:    "Burundî",
			territory.BJ:    "Bênîn",
			territory.BL:    "Saint-Barthélemy",
			territory.BM:    "Bermûda",
			territory.BN:    "Brûney",
			territory.BO:    "Bolîvya",
			territory.BR:    "Brazîl",
			territory.BS:    "Bahama",
			territory.BT:    "Bûtan",
			territory.BW:    "Botswana",
			territory.BY:    "Belarûs",
			territory.BZ:    "Belîze",
			territory.CA:    "Kanada",
			territory.CD:    "Kongo - Kînşasa",
			territory.CF:    "Komara Afrîkaya Navend",
			territory.CG:    "Kongo - Brazzaville",
			territory.CH:    "Swîsre",
			territory.CI:    "Peravê Diranfîl",
			territory.CK:    "Giravên Cook",
			territory.CL:    "Şîle",
			territory.CM:    "Kamerûn",
			territory.CN:    "Çîn",
			territory.CO:    "Kolombiya",
			territory.CR:    "Kosta Rîka",
			territory.CU:    "Kûba",
			territory.CV:    "Kap Verde",
			territory.CY:    "Kîpros",
			territory.CZ:    "Çekya",
			territory.DE:    "Almanya",
			territory.DJ:    "Cîbûtî",
			territory.DK:    "Danîmarka",
			territory.DM:    "Domînîka",
			territory.DO:    "Komara Domînîk",
			territory.DZ:    "Cezayir",
			territory.EC:    "Ekuador",
			territory.EE:    "Estonya",
			territory.EG:    "Misir",
			territory.EH:    "Sahraya Rojava",
			territory.ER:    "Erîtrea",
			territory.ES:    "Spanya",
			territory.ET:    "Etiyopya",
			territory.EU:    "Yekîtiya Ewropayê",
			territory.FI:    "Fînlenda",
			territory.FJ:    "Fîjî",
			territory.FK:    "Giravên Malvîn",
			territory.FM:    "Mîkronezya",
			territory.FO:    "Giravên Feroe",
			territory.FR:    "Fransa",
			territory.GA:    "Gabon",
			territory.GB:    "Keyaniya Yekbûyî",
			territory.GD:    "Grenada",
			territory.GE:    "Gurcistan",
			territory.GF:    "Guyanaya Fransî",
			territory.GH:    "Gana",
			territory.GI:    "Cîbraltar",
			territory.GL:    "Grînlenda",
			territory.GM:    "Gambiya",
			territory.GN:    "Gîne",
			territory.GP:    "Guadeloupe",
			territory.GQ:    "Gîneya Rojbendî",
			territory.GR:    "Yewnanistan",
			territory.GT:    "Guatemala",
			territory.GU:    "Guam",
			territory.GW:    "Gîne-Bissau",
			territory.GY:    "Guyana",
			territory.HN:    "Hondûras",
			territory.HR:    "Kroatya",
			territory.HT:    "Haîtî",
			territory.HU:    "Macaristan",
			territory.IC:    "Giravên Qenariyê",
			territory.ID:    "Îndonezya",
			territory.IE:    "Îrlenda",
			territory.IL:    "Îsraêl",
			territory.IM:    "Girava Man",
			territory.IN:    "Hindistan",
			territory.IQ:    "Iraq",
			territory.IR:    "Îran",
			territory.IS:    "Îslenda",
			territory.IT:    "Îtalya",
			territory.JM:    "Jamaîka",
			territory.JO:    "Urdun",
			territory.JP:    "Japon",
			territory.KE:    "Kenya",
			territory.KG:    "Qirgizistan",
			territory.KH:    "Kamboca",
			territory.KI:    "Kirîbatî",
			territory.KM:    "Komor",
			territory.KN:    "Saint Kitts û Nevîs",
			territory.KP:    "Korêya Bakur",
			territory.KR:    "Korêya Başûr",
			territory.KW:    "Kuweyt",
			territory.KY:    "Giravên Kaymanê",
			territory.KZ:    "Qazaxistan",
			territory.LA:    "Laos",
			territory.LB:    "Libnan",
			territory.LC:    "Saint Lucia",
			territory.LI:    "Liechtenstein",
			territory.LK:    "Srî Lanka",
			territory.LR:    "Lîberya",
			territory.LS:    "Lesoto",
			territory.LT:    "Lîtvanya",
			territory.LU:    "Lûksembûrg",
			territory.LV:    "Letonya",
			territory.LY:    "Lîbya",
			territory.MA:    "Maroko",
			territory.MC:    "Monako",
			territory.MD:    "Moldova",
			territory.ME:    "Montenegro",
			territory.MG:    "Madagaskar",
			territory.MH:    "Giravên Marşal",
			territory.MK:    "Makedonya",
			territory.ML:    "Malî",
			territory.MM:    "Myanmar (Birmanya)",
			territory.MN:    "Mongolya",
			territory.MP:    "Giravên Bakurê Marianan",
			territory.MQ:    "Martinique",
			territory.MR:    "Morîtanya",
			territory.MT:    "Malta",
			territory.MU:    "Maurîtius",
			territory.MV:    "Maldîv",
			territory.MW:    "Malawî",
			territory.MX:    "Meksîk",
			territory.MY:    "Malezya",
			territory.MZ:    "Mozambîk",
			territory.NA:    "Namîbya",
			territory.NC:    "Kaledonyaya Nû",
			territory.NE:    "Nîjer",
			territory.NF:    "Girava Norfolk",
			territory.NG:    "Nîjerya",
			territory.NI:    "Nîkaragua",
			territory.NL:    "Holenda",
			territory.NO:    "Norwêc",
			territory.NP:    "Nepal",
			territory.NR:    "Naûrû",
			territory.NU:    "Niûe",
			territory.NZ:    "Nû Zelenda",
			territory.OM:    "Oman",
			territory.PA:    "Panama",
			territory.PE:    "Perû",
			territory.PF:    "Polînezyaya Fransî",
			territory.PG:    "Papua Gîneya Nû",
			territory.PH:    "Filîpîn",
			territory.PK:    "Pakistan",
			territory.PL:    "Polonya",
			territory.PM:    "Saint-Pierre û Miquelon",
			territory.PN:    "Giravên Pitcairn",
			territory.PR:    "Porto Rîko",
			territory.PS:    "Xakên filistînî",
			territory.PT:    "Portûgal",
			territory.PW:    "Palau",
			territory.PY:    "Paraguay",
			territory.QA:    "Qeter",
			territory.RE:    "Réunion",
			territory.RO:    "Romanya",
			territory.RS:    "Serbistan",
			territory.RU:    "Rûsya",
			territory.RW:    "Rwanda",
			territory.SA:    "Erebistana Siyûdî",
			territory.SB:    "Giravên Salomon",
			territory.SC:    "Seyşel",
			territory.SD:    "Sûdan",
			territory.SE:    "Swêd",
			territory.SG:    "Singapûr",
			territory.SI:    "Slovenya",
			territory.SK:    "Slovakya",
			territory.SL:    "Sierra Leone",
			territory.SM:    "San Marîno",
			territory.SN:    "Senegal",
			territory.SO:    "Somalya",
			territory.SR:    "Sûrînam",
			territory.SS:    "Sûdana Başûr",
			territory.ST:    "Sao Tome û Prînsîpe",
			territory.SV:    "El Salvador",
			territory.SY:    "Sûrî",
			territory.SZ:    "Swazîlenda",
			territory.TC:    "Giravên Turk û Kaîkos",
			territory.TD:    "Çad",
			territory.TG:    "Togo",
			territory.TH:    "Taylenda",
			territory.TJ:    "Tacîkistan",
			territory.TK:    "Tokelau",
			territory.TL:    "Tîmora-Leste",
			territory.TM:    "Tirkmenistan",
			territory.TN:    "Tûnis",
			territory.TO:    "Tonga",
			territory.TR:    "Tirkiye",
			territory.TT:    "Trînîdad û Tobago",
			territory.TV:    "Tûvalû",
			territory.TW:    "Taywan",
			territory.TZ:    "Tanzanya",
			territory.UA:    "Ûkrayna",
			territory.UG:    "Ûganda",
			territory.UN:    "Neteweyên Yekbûyî",
			territory.US:    "Dewletên Yekbûyî yên Amerîkayê",
			territory.UY:    "Ûrûguay",
			territory.UZ:    "Ûzbêkistan",
			territory.VA:    "Vatîkan",
			territory.VC:    "Saint Vincent û Giravên Grenadîn",
			territory.VE:    "Venezuela",
			territory.VN:    "Viyetnam",
			territory.VU:    "Vanûatû",
			territory.WF:    "Wallis û Futuna",
			territory.WS:    "Samoa",
			territory.XK:    "Kosovo",
			territory.YE:    "Yemen",
			territory.ZA:    "Afrîkaya Başûr",
			territory.ZM:    "Zambiya",
			territory.ZW:    "Zîmbabwe",
		},
	}
}
