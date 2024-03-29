// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_ia_001() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ia_001",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE 'le' d 'de' MMMM y", Long: "d 'de' MMMM y", Medium: "d MMM y", Short: "dd-MM-y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'a' {0}", Long: "{1} 'a' {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "mai", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "oct", Nov: "nov", Dec: "dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januario", Feb: "februario", Mar: "martio", Apr: "april", May: "maio", Jun: "junio", Jul: "julio", Aug: "augusto", Sep: "septembre", Oct: "octobre", Nov: "novembre", Dec: "decembre"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dom", Mon: "lun", Tue: "mar", Wed: "mer", Thu: "jov", Fri: "ven", Sat: "sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "d", Mon: "l", Tue: "m", Wed: "m", Thu: "j", Fri: "v", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "do", Mon: "lu", Tue: "ma", Wed: "me", Thu: "jo", Fri: "ve", Sat: "sa"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "dominica", Mon: "lunedi", Tue: "martedi", Wed: "mercuridi", Thu: "jovedi", Fri: "venerdi", Sat: "sabbato"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", CurrencyAccounting: "¤\u00a0#,##0.00;(¤\u00a0#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ALL: cldr.Currency{DisplayName: "lek albanese", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "florino antillan", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "kwanza angolan", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "peso argentin", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "dollar australian", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "florino aruban", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "marco convertibile de Bosnia-Herzegovina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "dollar barbadian", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "lev bulgare", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "franco burundese", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "dollar bermudan", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "boliviano bolivian", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "real brasilian", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "dollar bahamian", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "pula botswanese", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "rublo bielorusse", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "dollar belizan", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "dollar canadian", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "franco congolese", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "franco suisse", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "peso chilen", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "yuan chinese", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "peso colombian", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "colon costarican", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "peso cuban convertibile", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "peso cuban", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "escudo capoverdian", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "corona chec", Symbol: "Kč"},
				currency.DEM: cldr.Currency{DisplayName: "Marco geman", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "franco djibutian", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "corona danese", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "peso dominican", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "dinar algerin", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Corona estonian", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "libra egyptie", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "nakfa eritree", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "birr ethiope", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Marco finnese", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "dollar fijian", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "libra falklandese", Symbol: "£"},
				currency.FRF: cldr.Currency{DisplayName: "Franco francese", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "libra sterling", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "cedi ghanese", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "libra de Gibraltar", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi gambian", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "franco guinean", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "quetzal guatemaltec", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "dollar guyanese", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "lempira hondurese", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "kuna croate", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "gourde haitian", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "forint hungare", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.IEP: cldr.Currency{DisplayName: "Libra irlandese", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "rupia indian", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "corona islandese", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "dollar jamaican", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "yen japonese", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "shilling kenyan", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "franco comorian", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "dollar del Insulas Caiman", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "dollar liberian", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "dinar libyc", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "dirham marocchin", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "leu moldave", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "ariary malgache", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "denar macedonie", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "ouguiya mauritan (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ouguiya mauritan", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "rupia mauritian", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha malawian", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "peso mexican", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "metical mozambican", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "dollar namibian", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "naira nigerian", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "cordoba nicaraguan", Symbol: "C$"},
				currency.NLG: cldr.Currency{DisplayName: "Florino nederlandese", Symbol: "ƒ"},
				currency.NOK: cldr.Currency{DisplayName: "corona norvegian", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "dollar neozelandese", Symbol: "NZ$"},
				currency.PAB: cldr.Currency{DisplayName: "balboa panamen", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "sol peruvian", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "kina papuan", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "zloty polonese", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "guarani paraguayan", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "leu romanian", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "dinar serbe", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "rublo russe", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "franco ruandese", Symbol: "RF"},
				currency.SBD: cldr.Currency{DisplayName: "dollar del insulas Salomon", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "rupia seychellese", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "libra sudanese", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "corona svedese", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "libra de St. Helena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "leone sierraleonese", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "shilling somali", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "dollar surinamese", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "libra sud-sudanese", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "dobra de São Tomé e Príncipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni swazilandese", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "dinar tunisian", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "paʻanga tongan", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "dollar de Trinidad e Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "shilling tanzanian", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "hryvnia ukrainian", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "shilling ugandese", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "dollar statounitese", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "peso uruguayan", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "bolivar venezuelan (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "bolivar venezuelan", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "vatu vanuatuan", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "tala samoan", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "franco CFA de Africa Central", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "dollar del Caribes Oriental", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "franco CFA de Africa Occidental", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "franco CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "moneta incognite", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "rand sudafrican", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha zambian", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "afar",
			language.AB:      "abkhazo",
			language.ACE:     "acehnese",
			language.ADA:     "adangme",
			language.ADY:     "adygeano",
			language.AF:      "afrikaans",
			language.AGQ:     "aghem",
			language.AIN:     "ainu",
			language.AK:      "akan",
			language.ALE:     "aleuto",
			language.ALT:     "altai del sud",
			language.AM:      "amharico",
			language.AN:      "aragonese",
			language.ANP:     "angika",
			language.AR:      "arabe",
			language.AR_001:  "arabe standard moderne",
			language.ARN:     "mapuche",
			language.ARP:     "arapaho",
			language.AS:      "assamese",
			language.ASA:     "asu",
			language.AST:     "asturiano",
			language.AV:      "avaro",
			language.AWA:     "awadhi",
			language.AY:      "aymara",
			language.AZ:      "azeri",
			language.BA:      "bashkir",
			language.BAN:     "balinese",
			language.BAS:     "basaa",
			language.BE:      "bielorusso",
			language.BEM:     "bemba",
			language.BEZ:     "bena",
			language.BG:      "bulgaro",
			language.BHO:     "bhojpuri",
			language.BI:      "bislama",
			language.BIN:     "bini",
			language.BLA:     "siksika",
			language.BM:      "bambara",
			language.BN:      "bengalese",
			language.BO:      "tibetano",
			language.BR:      "breton",
			language.BRX:     "bodo",
			language.BS:      "bosniaco",
			language.BUG:     "buginese",
			language.BYN:     "blin",
			language.CA:      "catalano",
			language.CE:      "checheno",
			language.CEB:     "cebuano",
			language.CGG:     "chiga",
			language.CH:      "chamorro",
			language.CHK:     "chuukese",
			language.CHM:     "mari",
			language.CHO:     "choctaw",
			language.CHR:     "cherokee",
			language.CHY:     "cheyenne",
			language.CKB:     "kurdo central",
			language.CO:      "corso",
			language.CRS:     "creolo seychellese",
			language.CS:      "checo",
			language.CU:      "slavo ecclesiastic",
			language.CV:      "chuvash",
			language.CY:      "gallese",
			language.DA:      "danese",
			language.DAK:     "dakota",
			language.DAR:     "dargwa",
			language.DAV:     "taita",
			language.DE:      "germano",
			language.DE_AT:   "germano austriac",
			language.DE_CH:   "alte germano suisse",
			language.DGR:     "dogrib",
			language.DJE:     "zarma",
			language.DSB:     "basse sorabo",
			language.DUA:     "duala",
			language.DV:      "divehi",
			language.DYO:     "jola-fonyi",
			language.DZ:      "dzongkha",
			language.DZG:     "dazaga",
			language.EBU:     "embu",
			language.EE:      "ewe",
			language.EFI:     "efik",
			language.EKA:     "ekajuk",
			language.EL:      "greco",
			language.EN:      "anglese",
			language.EN_AU:   "anglese australian",
			language.EN_CA:   "anglese canadian",
			language.EN_GB:   "anglese britannic",
			language.EN_US:   "anglese american",
			language.EO:      "esperanto",
			language.ES:      "espaniol",
			language.ES_419:  "espaniol latinoamerican",
			language.ES_ES:   "espaniol europee",
			language.ES_MX:   "espaniol mexican",
			language.ET:      "estoniano",
			language.EU:      "basco",
			language.EWO:     "ewondo",
			language.FA:      "persa",
			language.FF:      "fula",
			language.FI:      "finnese",
			language.FIL:     "filipino",
			language.FJ:      "fijiano",
			language.FO:      "feroese",
			language.FON:     "fon",
			language.FR:      "francese",
			language.FR_CA:   "francese canadian",
			language.FR_CH:   "francese suisse",
			language.FUR:     "friulano",
			language.FY:      "frison occidental",
			language.GA:      "irlandese",
			language.GAA:     "ga",
			language.GD:      "gaelico scotese",
			language.GEZ:     "ge’ez",
			language.GIL:     "gilbertese",
			language.GL:      "galleco",
			language.GN:      "guarani",
			language.GOR:     "gorontalo",
			language.GSW:     "germano suisse",
			language.GU:      "gujarati",
			language.GUZ:     "gusii",
			language.GV:      "mannese",
			language.GWI:     "gwich’in",
			language.HA:      "hausa",
			language.HAW:     "hawaiano",
			language.HE:      "hebreo",
			language.HI:      "hindi",
			language.HIL:     "hiligaynon",
			language.HMN:     "hmong",
			language.HR:      "croato",
			language.HSB:     "alte sorabo",
			language.HT:      "creolo haitian",
			language.HU:      "hungaro",
			language.HUP:     "hupa",
			language.HY:      "armeniano",
			language.HZ:      "herero",
			language.IA:      "interlingua",
			language.IBA:     "iban",
			language.IBB:     "ibibio",
			language.ID:      "indonesiano",
			language.IE:      "Interlingue",
			language.IG:      "igbo",
			language.II:      "yi de Sichuan",
			language.ILO:     "ilocano",
			language.INH:     "ingush",
			language.IO:      "ido",
			language.IS:      "islandese",
			language.IT:      "italiano",
			language.IU:      "inuktitut",
			language.JA:      "japonese",
			language.JBO:     "lojban",
			language.JGO:     "ngomba",
			language.JMC:     "machame",
			language.JV:      "javanese",
			language.KA:      "georgiano",
			language.KAB:     "kabylo",
			language.KAC:     "kachin",
			language.KAJ:     "jju",
			language.KAM:     "kamba",
			language.KBD:     "cabardiano",
			language.KCG:     "tyap",
			language.KDE:     "makonde",
			language.KEA:     "capoverdiano",
			language.KFO:     "koro",
			language.KHA:     "khasi",
			language.KHQ:     "koyra chiini",
			language.KI:      "kikuyu",
			language.KJ:      "kuanyama",
			language.KK:      "kazakh",
			language.KKJ:     "kako",
			language.KL:      "groenlandese",
			language.KLN:     "kalenjin",
			language.KM:      "khmer",
			language.KMB:     "kimbundu",
			language.KN:      "kannada",
			language.KO:      "coreano",
			language.KOK:     "konkani",
			language.KPE:     "kpelle",
			language.KR:      "kanuri",
			language.KRC:     "karachay-balkaro",
			language.KRL:     "careliano",
			language.KRU:     "kurukh",
			language.KS:      "kashmiri",
			language.KSB:     "shambala",
			language.KSF:     "bafia",
			language.KSH:     "coloniese",
			language.KU:      "kurdo",
			language.KUM:     "kumyko",
			language.KV:      "komi",
			language.KW:      "cornico",
			language.KY:      "kirghizo",
			language.LA:      "latino",
			language.LAD:     "ladino",
			language.LAG:     "langi",
			language.LB:      "luxemburgese",
			language.LEZ:     "lezghiano",
			language.LG:      "luganda",
			language.LI:      "limburgese",
			language.LKT:     "lakota",
			language.LN:      "lingala",
			language.LO:      "laotiano",
			language.LOZ:     "lozi",
			language.LRC:     "luri del nord",
			language.LT:      "lithuano",
			language.LU:      "luba-katanga",
			language.LUA:     "luba-lulua",
			language.LUN:     "lunda",
			language.LUO:     "luo",
			language.LUS:     "mizo",
			language.LUY:     "luyia",
			language.LV:      "letton",
			language.MAD:     "madurese",
			language.MAG:     "magahi",
			language.MAI:     "maithili",
			language.MAK:     "macassarese",
			language.MAS:     "masai",
			language.MDF:     "moksha",
			language.MEN:     "mende",
			language.MER:     "meri",
			language.MFE:     "creolo mauritian",
			language.MG:      "malgache",
			language.MGH:     "makhuwa-meetto",
			language.MGO:     "metaʼ",
			language.MH:      "marshallese",
			language.MI:      "maori",
			language.MIC:     "micmac",
			language.MIN:     "minangkabau",
			language.MK:      "macedone",
			language.ML:      "malayalam",
			language.MN:      "mongol",
			language.MNI:     "manipuri",
			language.MOH:     "mohawk",
			language.MOS:     "mossi",
			language.MR:      "marathi",
			language.MS:      "malay",
			language.MT:      "maltese",
			language.MUA:     "mundang",
			language.MUL:     "plure linguas",
			language.MUS:     "creek",
			language.MWL:     "mirandese",
			language.MY:      "birmano",
			language.MYV:     "erzya",
			language.MZN:     "mazanderani",
			language.NA:      "nauru",
			language.NAP:     "napolitano",
			language.NAQ:     "nama",
			language.NB:      "norvegiano bokmål",
			language.ND:      "ndebele del nord",
			language.NE:      "nepalese",
			language.NEW:     "newari",
			language.NG:      "ndonga",
			language.NIA:     "nias",
			language.NIU:     "nieuano",
			language.NL:      "nederlandese",
			language.NL_BE:   "flamingo",
			language.NMG:     "kwasio",
			language.NN:      "norvegiano nynorsk",
			language.NNH:     "ngiemboon",
			language.NO:      "norvegiano",
			language.NOG:     "nogai",
			language.NQO:     "n’ko",
			language.NR:      "ndebele del sud",
			language.NSO:     "sotho del nord",
			language.NUS:     "nuer",
			language.NV:      "navajo",
			language.NY:      "nyanja",
			language.NYN:     "nyankole",
			language.OC:      "occitano",
			language.OM:      "oromo",
			language.OR:      "oriya",
			language.OS:      "osseto",
			language.PA:      "punjabi",
			language.PAG:     "pangasinan",
			language.PAM:     "pampanga",
			language.PAP:     "papiamento",
			language.PAU:     "palauano",
			language.PCM:     "pidgin nigerian",
			language.PL:      "polonese",
			language.PRG:     "prussiano",
			language.PS:      "pashto",
			language.PT:      "portugese",
			language.PT_BR:   "portugese de Brasil",
			language.PT_PT:   "portugese de Portugal",
			language.QU:      "quechua",
			language.QUC:     "kʼicheʼ",
			language.RAP:     "rapanui",
			language.RAR:     "rarotongano",
			language.RM:      "romanche",
			language.RN:      "rundi",
			language.RO:      "romaniano",
			language.RO_MD:   "moldavo",
			language.ROF:     "rombo",
			language.ROOT:    "radice",
			language.RU:      "russo",
			language.RUP:     "aromaniano",
			language.RW:      "kinyarwanda",
			language.RWK:     "rwa",
			language.SA:      "sanscrito",
			language.SAD:     "sandawe",
			language.SAH:     "yakuto",
			language.SAQ:     "samburu",
			language.SAT:     "santali",
			language.SBA:     "ngambay",
			language.SBP:     "sangu",
			language.SC:      "sardo",
			language.SCN:     "siciliano",
			language.SCO:     "scotese",
			language.SD:      "sindhi",
			language.SE:      "sami del nord",
			language.SEH:     "sena",
			language.SES:     "koyraboro senni",
			language.SG:      "sango",
			language.SH:      "serbocroate",
			language.SHI:     "tachelhit",
			language.SHN:     "shan",
			language.SI:      "cingalese",
			language.SK:      "slovaco",
			language.SL:      "sloveno",
			language.SM:      "samoano",
			language.SMA:     "sami del sud",
			language.SMJ:     "sami de Lule",
			language.SMN:     "sami de Inari",
			language.SMS:     "sami skolt",
			language.SN:      "shona",
			language.SNK:     "soninke",
			language.SO:      "somali",
			language.SQ:      "albanese",
			language.SR:      "serbo",
			language.SRN:     "sranan tongo",
			language.SS:      "swati",
			language.SSY:     "saho",
			language.ST:      "sotho del sud",
			language.SU:      "sundanese",
			language.SUK:     "sukuma",
			language.SV:      "svedese",
			language.SW:      "swahili",
			language.SW_CD:   "swahili del Congo",
			language.SWB:     "comoriano",
			language.SYR:     "syriaco",
			language.TA:      "tamil",
			language.TE:      "telugu",
			language.TEM:     "temne",
			language.TEO:     "teso",
			language.TET:     "tetum",
			language.TG:      "tajiko",
			language.TH:      "thai",
			language.TI:      "tigrinya",
			language.TIG:     "tigre",
			language.TK:      "turkmeno",
			language.TLH:     "klingon",
			language.TN:      "tswana",
			language.TO:      "tongano",
			language.TPI:     "tok pisin",
			language.TR:      "turco",
			language.TRV:     "taroko",
			language.TS:      "tsonga",
			language.TT:      "tataro",
			language.TUM:     "tumbuka",
			language.TVL:     "tuvaluano",
			language.TW:      "twi",
			language.TWQ:     "tasawaq",
			language.TY:      "tahitiano",
			language.TYV:     "tuvano",
			language.TZM:     "tamazight del Atlas Central",
			language.UDM:     "udmurto",
			language.UG:      "uighur",
			language.UK:      "ukrainiano",
			language.UMB:     "umbundu",
			language.UND:     "lingua incognite",
			language.UR:      "urdu",
			language.UZ:      "uzbeko",
			language.VAI:     "vai",
			language.VE:      "venda",
			language.VI:      "vietnamese",
			language.VO:      "volapük",
			language.VUN:     "vunjo",
			language.WA:      "wallon",
			language.WAE:     "walser",
			language.WAL:     "wolaytta",
			language.WAR:     "waray",
			language.WO:      "wolof",
			language.XAL:     "calmuco",
			language.XH:      "xhosa",
			language.XOG:     "soga",
			language.YAV:     "yangben",
			language.YBB:     "yemba",
			language.YI:      "yiddish",
			language.YO:      "yoruba",
			language.YUE:     "cantonese",
			language.ZGH:     "tamazight marocchin standard",
			language.ZH:      "chinese",
			language.ZH_HANS: "chinese simplificate",
			language.ZH_HANT: "chinese traditional",
			language.ZU:      "zulu",
			language.ZUN:     "zuni",
			language.ZXX:     "sin contento linguistic",
			language.ZZA:     "zaza",
		},
		Territories: cldr.Territories{
			territory.V_001: "Mundo",
			territory.V_002: "Africa",
			territory.V_003: "America del Nord",
			territory.V_005: "America del Sud",
			territory.V_009: "Oceania",
			territory.V_011: "Africa occidental",
			territory.V_013: "America central",
			territory.V_014: "Africa oriental",
			territory.V_015: "Africa septentrional",
			territory.V_017: "Africa central",
			territory.V_018: "Africa meridional",
			territory.V_019: "Americas",
			territory.V_021: "America septentrional",
			territory.V_029: "Caribes",
			territory.V_030: "Asia oriental",
			territory.V_034: "Asia meridional",
			territory.V_035: "Asia del sud-est",
			territory.V_039: "Europa meridional",
			territory.V_053: "Australasia",
			territory.V_054: "Melanesia",
			territory.V_057: "Region micronesian",
			territory.V_061: "Polynesia",
			territory.V_142: "Asia",
			territory.V_143: "Asia central",
			territory.V_145: "Asia occidental",
			territory.V_150: "Europa",
			territory.V_151: "Europa oriental",
			territory.V_154: "Europa septentrional",
			territory.V_155: "Europa occidental",
			territory.V_202: "Africa subsaharian",
			territory.V_419: "America latin",
			territory.AD:    "Andorra",
			territory.AE:    "Emiratos Arabe Unite",
			territory.AF:    "Afghanistan",
			territory.AG:    "Antigua e Barbuda",
			territory.AL:    "Albania",
			territory.AM:    "Armenia",
			territory.AO:    "Angola",
			territory.AQ:    "Antarctica",
			territory.AR:    "Argentina",
			territory.AS:    "Samoa american",
			territory.AT:    "Austria",
			territory.AU:    "Australia",
			territory.AX:    "Insulas Åland",
			territory.AZ:    "Azerbaidzhan",
			territory.BA:    "Bosnia e Herzegovina",
			territory.BD:    "Bangladesh",
			territory.BE:    "Belgica",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgaria",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BM:    "Bermuda",
			territory.BO:    "Bolivia",
			territory.BR:    "Brasil",
			territory.BS:    "Bahamas",
			territory.BT:    "Bhutan",
			territory.BV:    "Insula de Bouvet",
			territory.BW:    "Botswana",
			territory.BY:    "Bielorussia",
			territory.BZ:    "Belize",
			territory.CA:    "Canada",
			territory.CF:    "Republica African Central",
			territory.CG:    "Congo",
			territory.CH:    "Suissa",
			territory.CK:    "Insulas Cook",
			territory.CL:    "Chile",
			territory.CM:    "Camerun",
			territory.CN:    "China",
			territory.CO:    "Colombia",
			territory.CR:    "Costa Rica",
			territory.CU:    "Cuba",
			territory.CX:    "Insula de Natal",
			territory.CY:    "Cypro",
			territory.CZ:    "Chechia",
			territory.DE:    "Germania",
			territory.DK:    "Danmark",
			territory.DO:    "Republica Dominican",
			territory.DZ:    "Algeria",
			territory.EC:    "Ecuador",
			territory.EE:    "Estonia",
			territory.EG:    "Egypto",
			territory.EH:    "Sahara occidental",
			territory.ER:    "Eritrea",
			territory.ES:    "Espania",
			territory.ET:    "Ethiopia",
			territory.EU:    "Union Europee",
			territory.EZ:    "Zona euro",
			territory.FI:    "Finlandia",
			territory.FM:    "Micronesia",
			territory.FO:    "Insulas Feroe",
			territory.FR:    "Francia",
			territory.GA:    "Gabon",
			territory.GB:    "Regno Unite",
			territory.GE:    "Georgia",
			territory.GF:    "Guyana francese",
			territory.GG:    "Guernsey",
			territory.GH:    "Ghana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Groenlandia",
			territory.GM:    "Gambia",
			territory.GN:    "Guinea",
			territory.GQ:    "Guinea equatorial",
			territory.GR:    "Grecia",
			territory.GT:    "Guatemala",
			territory.GW:    "Guinea-Bissau",
			territory.HN:    "Honduras",
			territory.HR:    "Croatia",
			territory.HT:    "Haiti",
			territory.HU:    "Hungaria",
			territory.ID:    "Indonesia",
			territory.IE:    "Irlanda",
			territory.IL:    "Israel",
			territory.IM:    "Insula de Man",
			territory.IN:    "India",
			territory.IO:    "Territorio oceanic britanno-indian",
			territory.IQ:    "Irak",
			territory.IR:    "Iran",
			territory.IS:    "Islanda",
			territory.IT:    "Italia",
			territory.JE:    "Jersey",
			territory.JO:    "Jordania",
			territory.JP:    "Japon",
			territory.KE:    "Kenya",
			territory.KG:    "Kirghizistan",
			territory.KH:    "Cambodgia",
			territory.KI:    "Kiribati",
			territory.KM:    "Comoros",
			territory.KN:    "Sancte Christophoro e Nevis",
			territory.KP:    "Corea del Nord",
			territory.KR:    "Corea del Sud",
			territory.KY:    "Insulas de Caiman",
			territory.KZ:    "Kazakhstan",
			territory.LB:    "Libano",
			territory.LC:    "Sancte Lucia",
			territory.LI:    "Liechtenstein",
			territory.LK:    "Sri Lanka",
			territory.LR:    "Liberia",
			territory.LS:    "Lesotho",
			territory.LT:    "Lituania",
			territory.LU:    "Luxemburg",
			territory.LV:    "Lettonia",
			territory.LY:    "Libya",
			territory.MA:    "Marocco",
			territory.MC:    "Monaco",
			territory.MD:    "Moldavia",
			territory.ME:    "Montenegro",
			territory.MG:    "Madagascar",
			territory.MH:    "Insulas Marshall",
			territory.MK:    "Macedonia",
			territory.ML:    "Mali",
			territory.MM:    "Birmania/Myanmar",
			territory.MN:    "Mongolia",
			territory.MP:    "Insulas Marianna del Nord",
			territory.MR:    "Mauritania",
			territory.MT:    "Malta",
			territory.MW:    "Malawi",
			territory.MX:    "Mexico",
			territory.MY:    "Malaysia",
			territory.MZ:    "Mozambique",
			territory.NA:    "Namibia",
			territory.NC:    "Nove Caledonia",
			territory.NE:    "Niger",
			territory.NF:    "Insula Norfolk",
			territory.NG:    "Nigeria",
			territory.NI:    "Nicaragua",
			territory.NL:    "Nederlandia",
			territory.NO:    "Norvegia",
			territory.NP:    "Nepal",
			territory.NZ:    "Nove Zelanda",
			territory.OM:    "Oman",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PF:    "Polynesia francese",
			territory.PG:    "Papua Nove Guinea",
			territory.PH:    "Philippinas",
			territory.PK:    "Pakistan",
			territory.PL:    "Polonia",
			territory.PM:    "St. Pierre e Miquelon",
			territory.PT:    "Portugal",
			territory.PY:    "Paraguay",
			territory.QO:    "Oceania remote",
			territory.RO:    "Romania",
			territory.RS:    "Serbia",
			territory.RU:    "Russia",
			territory.RW:    "Ruanda",
			territory.SA:    "Arabia Saudita",
			territory.SB:    "Insulas Solomon",
			territory.SC:    "Seychelles",
			territory.SD:    "Sudan",
			territory.SE:    "Svedia",
			territory.SI:    "Slovenia",
			territory.SJ:    "Svalbard e Jan Mayen",
			territory.SK:    "Slovachia",
			territory.SL:    "Sierra Leone",
			territory.SM:    "San Marino",
			territory.SN:    "Senegal",
			territory.SO:    "Somalia",
			territory.SR:    "Suriname",
			territory.SS:    "Sudan del Sud",
			territory.SV:    "El Salvador",
			territory.SY:    "Syria",
			territory.SZ:    "Swazilandia",
			territory.TC:    "Insulas Turcos e Caicos",
			territory.TD:    "Tchad",
			territory.TF:    "Territorios meridional francese",
			territory.TG:    "Togo",
			territory.TH:    "Thailandia",
			territory.TJ:    "Tadzhikistan",
			territory.TK:    "Tokelau",
			territory.TL:    "Timor del Est",
			territory.TM:    "Turkmenistan",
			territory.TN:    "Tunisia",
			territory.TO:    "Tonga",
			territory.TR:    "Turchia",
			territory.TT:    "Trinidad e Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tanzania",
			territory.UA:    "Ukraina",
			territory.UG:    "Uganda",
			territory.UN:    "Nationes Unite",
			territory.US:    "Statos Unite",
			territory.UY:    "Uruguay",
			territory.UZ:    "Uzbekistan",
			territory.VA:    "Citate del Vaticano",
			territory.VC:    "Sancte Vincente e le Grenadinas",
			territory.VE:    "Venezuela",
			territory.VU:    "Vanuatu",
			territory.WS:    "Samoa",
			territory.XK:    "Kosovo",
			territory.YE:    "Yemen",
			territory.ZA:    "Sudafrica",
			territory.ZM:    "Zambia",
			territory.ZW:    "Zimbabwe",
			territory.ZZ:    "Region incognite",
		},
	}
}
