// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_fo_DK() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "fo_DK",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'kl'. {0}", Long: "{1} 'kl'. {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "mai", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "mars", Apr: "apríl", May: "mai", Jun: "juni", Jul: "juli", Aug: "august", Sep: "september", Oct: "oktober", Nov: "november", Dec: "desember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sun", Mon: "mán", Tue: "týs", Wed: "mik", Thu: "hós", Fri: "frí", Sat: "ley"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "M", Thu: "H", Fri: "F", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "su", Mon: "má", Tue: "tý", Wed: "mi", Thu: "hó", Fri: "fr", Sat: "le"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "sunnudagur", Mon: "mánadagur", Tue: "týsdagur", Wed: "mikudagur", Thu: "hósdagur", Fri: "fríggjadagur", Sat: "leygardagur"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤;(#,##0.00\u00a0¤)", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Sameindu Emirríkini dirham", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afganistan afghani", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Albania lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Armenia dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Niðurlonds Karibia gyllin", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angola kwanza", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Argentina peso", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Avstralskur dollari", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruba florin", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Aserbadjan manat", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Bosnia-Hersegovina mark (kann vekslast)", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados dollari", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladesj taka", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Bulgaria lev", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Barein dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundi frankur", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda dollari", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Brunei dollari", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Bolivia boliviano", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Brasilianskur real", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Bahamaoyggjar dollari", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Butan ngultrum", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Botsvana pula", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Hvítarussland ruble", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Hvítarussland ruble (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Belis dollari", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanada dollari", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongo frankur", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "sveisiskur frankur", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Kili peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "kinesiskur yuan (úr landi)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "kinesiskur yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolombia peso", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Kosta Rika colón", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Kuba peso (sum kann vekslast)", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kuba peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Grønhøvdaoyggjar escudo", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Kekkia koruna", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Djibuti frankur", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "donsk króna", Symbol: "kr."},
				currency.DOP: cldr.Currency{DisplayName: "Dominika peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Algeria dinar", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Egyptaland pund", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrea nakfa", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Etiopia birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Evra", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Fiji dollari", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Falklandsoyggjar pund", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "bretsk pund", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Georgia lari", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Gana cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltar pund", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambia dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Guinea frankur", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemala quetzal", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Gujana dollari", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hong Kong dollari", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Honduras lempira", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Kroatia kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haiti gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Ungarn forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesia rupiah", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Ísrael new shekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "indiskir rupis", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Irak dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "iranskir rials", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "íslendsk króna", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaika dollari", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Jordan dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "japanskur yen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "kenjanskur skillingur", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kirgisia som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kambodja riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Komoroyggjar frankur", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Norðurkorea won", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Suðurkorea won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuvait dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Caymanoyggjar dollari", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Kasakstan tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laos kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Libanon pund", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lanka rupi", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberia dollari", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Libya dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Marokko dirham", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Moldova leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Madagaskar ariary", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Makedónia denar", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Myanmar (Burma) kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongolia tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Makao pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Móritania ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Móritania ouguiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Móritius rupi", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Maldivoyggjar rufiyaa", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malavi kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Meksiko peso", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Malaisia ringgit", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Mosambik metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibia dollari", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigeria naira", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Nikaragua córdoba", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "norsk króna", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepal rupi", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Nýsæland dollari", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Oman rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panama balboa", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Peru sol", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Papua Nýguinea kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Filipsoyggjar peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistan rupi", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Pólland zloty", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Paraguai guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Katar rial", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Rumenia leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Serbia dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Russland ruble", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ruanda frankur", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudiarabia riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Salomonoyggjar dollari", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seyskelloyggjar rupi", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Sudan pund", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "svensk króna", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapor dollari", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "St. Helena pund", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leona leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somalia skillingur", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Surinam dollari", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Suðursudan pund", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Sao Tome & Prinsipi dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Sao Tome & Prinsipi dobra", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Sýria pund", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Svasiland lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Tailand baht", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "Tadsjikistan somoni", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Turkmenistan manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tunesia dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tonga paʻanga", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Turkaland liri", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad & Tobago dollari", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Taivan new dollari", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tansania skillingur", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukraina hryvnia", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Uganda skillingur", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "US dollari", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Uruguai peso", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Usbekistan som", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Venesuela bolívar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Venesuela bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Vjetnam dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatu vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoa tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Miðafrika CFA frankur", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "unse sølv", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "unse guld", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Eystur Karibia dollari", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Vesturafrika CFA frankur", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "unse palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP frankur", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "unse platin", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "ókent gjaldoyra", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Jemen rial", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Suðurafrika rand", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Sambia kwacha", Symbol: "ZMW"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "afar",
			language.AB:      "abkhasiskt",
			language.ACE:     "achinese",
			language.ADA:     "adangme",
			language.ADY:     "adyghe",
			language.AF:      "afrikaans",
			language.AGQ:     "aghem",
			language.AIN:     "ainu",
			language.AK:      "akan",
			language.ALE:     "aleut",
			language.ALT:     "suður altai",
			language.AM:      "amhariskt",
			language.AN:      "aragoniskt",
			language.ANP:     "angika",
			language.AR:      "arabiskt",
			language.AR_001:  "nútíðar vanligt arabiskt",
			language.ARN:     "mapuche",
			language.ARP:     "arapaho",
			language.AS:      "assamesiskt",
			language.ASA:     "asu",
			language.AST:     "asturianskt",
			language.AV:      "avariskt",
			language.AWA:     "awadhi",
			language.AY:      "aymara",
			language.AZ:      "azeri",
			language.BA:      "bashkir",
			language.BAN:     "balinesiskt",
			language.BAS:     "basaa",
			language.BE:      "hvitarussiskt",
			language.BEM:     "bemba",
			language.BEZ:     "bena",
			language.BG:      "bulgarskt",
			language.BGN:     "vestur balochi",
			language.BHO:     "bhojpuri",
			language.BI:      "bislama",
			language.BIN:     "bini",
			language.BLA:     "siksika",
			language.BM:      "bambara",
			language.BN:      "bangla",
			language.BO:      "tibetskt",
			language.BR:      "bretonskt",
			language.BRX:     "bodo",
			language.BS:      "bosniskt",
			language.BSS:     "bakossi",
			language.BUG:     "buginesiskt",
			language.BYN:     "blin",
			language.CA:      "katalani",
			language.CCP:     "khakma",
			language.CE:      "tjetjenskt",
			language.CEB:     "cebuano",
			language.CGG:     "chiga",
			language.CH:      "chamorro",
			language.CHK:     "chuukese",
			language.CHM:     "mari",
			language.CHO:     "choctaw",
			language.CHR:     "cherokee",
			language.CHY:     "cheyenne",
			language.CKB:     "miðkurdiskt",
			language.CO:      "korsikanskt",
			language.CRS:     "seselwa creole franskt",
			language.CS:      "kekkiskt",
			language.CU:      "kirkju sláviskt",
			language.CV:      "chuvash",
			language.CY:      "walisiskt",
			language.DA:      "danskt",
			language.DAK:     "dakota",
			language.DAR:     "dargwa",
			language.DAV:     "taita",
			language.DE:      "týskt",
			language.DE_CH:   "høgt týskt (Sveis)",
			language.DGR:     "dogrib",
			language.DJE:     "sarma",
			language.DSB:     "lágt sorbian",
			language.DUA:     "duala",
			language.DV:      "divehi",
			language.DYO:     "jola-fonyi",
			language.DZ:      "dzongkha",
			language.DZG:     "dazaga",
			language.EBU:     "embu",
			language.EE:      "ewe",
			language.EFI:     "efik",
			language.EKA:     "ekajuk",
			language.EL:      "grikskt",
			language.EN:      "enskt",
			language.EO:      "esperanto",
			language.ES:      "spanskt",
			language.ET:      "estiskt",
			language.EU:      "baskiskt",
			language.EWO:     "ewondo",
			language.FA:      "persiskt",
			language.FF:      "fulah",
			language.FI:      "finskt",
			language.FIL:     "filipiniskt",
			language.FJ:      "fijimál",
			language.FO:      "føroyskt",
			language.FON:     "fon",
			language.FR:      "franskt",
			language.FUR:     "friuliskt",
			language.FY:      "vestur frísiskt",
			language.GA:      "írskt",
			language.GAA:     "ga",
			language.GAG:     "gagauz",
			language.GAN:     "gan kinesiskt",
			language.GD:      "skotskt gæliskt",
			language.GEZ:     "geez",
			language.GIL:     "kiribatiskt",
			language.GL:      "galisiskt",
			language.GN:      "guarani",
			language.GOR:     "gorontalo",
			language.GSW:     "týskt (Sveis)",
			language.GU:      "gujarati",
			language.GUZ:     "gusii",
			language.GV:      "manx",
			language.GWI:     "gwich’in",
			language.HA:      "hausa",
			language.HAK:     "hakka kinesiskt",
			language.HAW:     "hawaiianskt",
			language.HE:      "hebraiskt",
			language.HI:      "hindi",
			language.HIL:     "hiligaynon",
			language.HMN:     "hmong",
			language.HR:      "kroatiskt",
			language.HSB:     "ovara sorbian",
			language.HSN:     "xiang kinesiskt",
			language.HT:      "haitiskt creole",
			language.HU:      "ungarskt",
			language.HUP:     "hupa",
			language.HY:      "armenskt",
			language.HZ:      "herero",
			language.IA:      "interlingua",
			language.IBA:     "iban",
			language.IBB:     "ibibio",
			language.ID:      "indonesiskt",
			language.IE:      "interlingue",
			language.IG:      "igbo",
			language.II:      "sichuan yi",
			language.ILO:     "iloko",
			language.INH:     "inguish",
			language.IO:      "ido",
			language.IS:      "íslendskt",
			language.IT:      "italskt",
			language.IU:      "inuktitut",
			language.JA:      "japanskt",
			language.JBO:     "lojban",
			language.JGO:     "ngomba",
			language.JMC:     "machame",
			language.JV:      "javanskt",
			language.KA:      "georgiskt",
			language.KAB:     "kabyle",
			language.KAC:     "kachin",
			language.KAJ:     "jju",
			language.KAM:     "kamba",
			language.KBD:     "kabardinskt",
			language.KCG:     "tyap",
			language.KDE:     "makonde",
			language.KEA:     "grønhøvdaoyggjarskt",
			language.KFO:     "koro",
			language.KHA:     "khasi",
			language.KHQ:     "koyra chiini",
			language.KI:      "kikuyu",
			language.KJ:      "kuanyama",
			language.KK:      "kazakh",
			language.KKJ:     "kako",
			language.KL:      "kalaallisut",
			language.KLN:     "kalenjin",
			language.KM:      "khmer",
			language.KMB:     "kimbundu",
			language.KN:      "kannada",
			language.KO:      "koreanskt",
			language.KOI:     "komi-permyak",
			language.KOK:     "konkani",
			language.KPE:     "kpelle",
			language.KR:      "kanuri",
			language.KRC:     "karachay-balkar",
			language.KRL:     "karelskt",
			language.KRU:     "kurukh",
			language.KS:      "kashmiri",
			language.KSB:     "shambala",
			language.KSF:     "bafia",
			language.KSH:     "kølnskt",
			language.KU:      "kurdiskt",
			language.KUM:     "kumyk",
			language.KV:      "komi",
			language.KW:      "corniskt",
			language.KY:      "kyrgyz",
			language.LA:      "latín",
			language.LAD:     "ladino",
			language.LAG:     "langi",
			language.LAH:     "lahnda",
			language.LB:      "luksemborgskt",
			language.LEZ:     "lezghian",
			language.LG:      "ganda",
			language.LI:      "limburgiskt",
			language.LKT:     "lakota",
			language.LN:      "lingala",
			language.LO:      "laoskt",
			language.LOZ:     "lozi",
			language.LRC:     "norður luri",
			language.LT:      "litaviskt",
			language.LU:      "luba-katanga",
			language.LUA:     "luba-lulua",
			language.LUN:     "lunda",
			language.LUO:     "luo",
			language.LUS:     "mizo",
			language.LUY:     "luyia",
			language.LV:      "lettiskt",
			language.MAD:     "maduresiskt",
			language.MAG:     "magahi",
			language.MAI:     "maithili",
			language.MAK:     "makasar",
			language.MAS:     "masai",
			language.MDF:     "moksha",
			language.MEN:     "mende",
			language.MER:     "meru",
			language.MFE:     "morisyen",
			language.MG:      "malagassiskt",
			language.MGH:     "makhuwa-meetto",
			language.MGO:     "metaʼ",
			language.MH:      "marshallesiskt",
			language.MI:      "maori",
			language.MIC:     "micmac",
			language.MIN:     "minangkabau",
			language.MK:      "makedónskt",
			language.ML:      "malayalam",
			language.MN:      "mongolskt",
			language.MNI:     "manupuri",
			language.MOH:     "mohawk",
			language.MOS:     "mossi",
			language.MR:      "marathi",
			language.MS:      "malaiiskt",
			language.MT:      "maltiskt",
			language.MUA:     "mundang",
			language.MUL:     "ymisk mál",
			language.MUS:     "creek",
			language.MWL:     "mirandesiskt",
			language.MY:      "burmesiskt",
			language.MYV:     "erzya",
			language.MZN:     "mazanderani",
			language.NA:      "nauru",
			language.NAN:     "min nan kinesiskt",
			language.NAP:     "napolitanskt",
			language.NAQ:     "nama",
			language.NB:      "norskt bókmál",
			language.ND:      "norður ndebele",
			language.NDS:     "lágt týskt",
			language.NDS_NL:  "lágt saksiskt",
			language.NE:      "nepalskt",
			language.NEW:     "newari",
			language.NG:      "ndonga",
			language.NIA:     "nias",
			language.NIU:     "niuean",
			language.NL:      "hálendskt",
			language.NL_BE:   "flamskt",
			language.NMG:     "kwasio",
			language.NN:      "nýnorskt",
			language.NNH:     "ngiemboon",
			language.NO:      "norskt",
			language.NOG:     "nogai",
			language.NQO:     "nʼko",
			language.NR:      "suður ndebele",
			language.NSO:     "norður sotho",
			language.NUS:     "nuer",
			language.NV:      "navajo",
			language.NY:      "nyanja",
			language.NYN:     "nyankole",
			language.OC:      "occitanskt",
			language.OM:      "oromo",
			language.OR:      "odia",
			language.OS:      "ossetiskt",
			language.PA:      "punjabi",
			language.PAG:     "pangasinan",
			language.PAM:     "pampanga",
			language.PAP:     "papiamento",
			language.PAU:     "palauan",
			language.PCM:     "nigeriskt pidgin",
			language.PL:      "pólskt",
			language.PRG:     "prusslanskt",
			language.PS:      "pashto",
			language.PT:      "portugiskiskt",
			language.PT_BR:   "portugiskiskt (Brasilia)",
			language.PT_PT:   "portugiskiskt (Evropa)",
			language.QU:      "quechua",
			language.QUC:     "kʼicheʼ",
			language.RAP:     "rapanui",
			language.RAR:     "rarotongiskt",
			language.RM:      "retoromanskt",
			language.RN:      "rundi",
			language.RO:      "rumenskt",
			language.RO_MD:   "moldaviskt",
			language.ROF:     "rombo",
			language.RU:      "russiskt",
			language.RUP:     "aromenskt",
			language.RW:      "kinyarwanda",
			language.RWK:     "rwa",
			language.SA:      "sanskrit",
			language.SAD:     "sandawe",
			language.SAH:     "sakha",
			language.SAQ:     "samburu",
			language.SAT:     "santali",
			language.SBA:     "ngambay",
			language.SBP:     "sangu",
			language.SC:      "sardiskt",
			language.SCN:     "sisilanskt",
			language.SCO:     "skotskt",
			language.SD:      "sindhi",
			language.SDH:     "suður kurdiskt",
			language.SE:      "norður sámiskt",
			language.SEH:     "sena",
			language.SES:     "koyraboro senni",
			language.SG:      "sango",
			language.SH:      "serbokroatiskt",
			language.SHI:     "tachelhit",
			language.SHN:     "shan",
			language.SI:      "singalesiskt",
			language.SK:      "slovakiskt",
			language.SL:      "slovenskt",
			language.SM:      "sámoiskt",
			language.SMA:     "suður sámiskt",
			language.SMJ:     "lule sámiskt",
			language.SMN:     "inari sami",
			language.SMS:     "skolt sámiskt",
			language.SN:      "shona",
			language.SNK:     "soninke",
			language.SO:      "somaliskt",
			language.SQ:      "albanskt",
			language.SR:      "serbiskt",
			language.SRN:     "sranan tongo",
			language.SS:      "swatiskt",
			language.SSY:     "saho",
			language.ST:      "sesotho",
			language.SU:      "sundanesiskt",
			language.SUK:     "sukuma",
			language.SV:      "svenskt",
			language.SW:      "swahili",
			language.SW_CD:   "kongo svahili",
			language.SWB:     "komoriskt",
			language.SYR:     "syriac",
			language.TA:      "tamilskt",
			language.TE:      "telugu",
			language.TEM:     "timne",
			language.TEO:     "teso",
			language.TET:     "tetum",
			language.TG:      "tajik",
			language.TH:      "tailendskt",
			language.TI:      "tigrinya",
			language.TIG:     "tigre",
			language.TK:      "turkmenskt",
			language.TL:      "tagalog",
			language.TLH:     "klingonskt",
			language.TN:      "tswana",
			language.TO:      "tonganskt",
			language.TPI:     "tok pisin",
			language.TR:      "turkiskt",
			language.TRV:     "taroko",
			language.TS:      "tsonga",
			language.TT:      "tatar",
			language.TUM:     "tumbuka",
			language.TVL:     "tuvalu",
			language.TW:      "twi",
			language.TWQ:     "tasawaq",
			language.TY:      "tahitiskt",
			language.TYV:     "tuvinian",
			language.TZM:     "miðatlasfjøll tamazight",
			language.UDM:     "udmurt",
			language.UG:      "uyghur",
			language.UK:      "ukrainskt",
			language.UMB:     "umbundu",
			language.UND:     "ókent mál",
			language.UR:      "urdu",
			language.UZ:      "usbekiskt",
			language.VAI:     "vai",
			language.VE:      "venda",
			language.VI:      "vjetnamesiskt",
			language.VO:      "volapykk",
			language.VUN:     "vunjo",
			language.WA:      "walloon",
			language.WAE:     "walser",
			language.WAL:     "wolaytta",
			language.WAR:     "waray",
			language.WBP:     "warlpiri",
			language.WO:      "wolof",
			language.WUU:     "wu kinesiskt",
			language.XAL:     "kalmyk",
			language.XH:      "xhosa",
			language.XOG:     "soga",
			language.YAV:     "yangben",
			language.YBB:     "yemba",
			language.YI:      "jiddiskt",
			language.YO:      "yoruba",
			language.YUE:     "kinesiskt, kantonesiskt",
			language.ZGH:     "vanligt marokanskt tamazight",
			language.ZH:      "kinesiskt, mandarin",
			language.ZH_HANS: "mandarin kinesiskt (einkult)",
			language.ZH_HANT: "mandarin kinesiskt (vanligt)",
			language.ZU:      "sulu",
			language.ZUN:     "zuni",
			language.ZXX:     "einki málsligt innihald",
			language.ZZA:     "zaza",
		},
		Territories: cldr.Territories{
			territory.V_001: "heimur",
			territory.V_002: "Afrika",
			territory.V_003: "Norðuramerika",
			territory.V_005: "Suðuramerika",
			territory.V_009: "Osiania",
			territory.V_011: "Vesturafrika",
			territory.V_013: "Miðamerika",
			territory.V_014: "Eysturafrika",
			territory.V_015: "Norðurafrika",
			territory.V_017: "Miðafrika",
			territory.V_018: "sunnari partur av Afrika",
			territory.V_019: "Amerika",
			territory.V_021: "Amerika norðanfyri Meksiko",
			territory.V_029: "Karibia",
			territory.V_030: "Eysturasia",
			territory.V_034: "Suðurasia",
			territory.V_035: "Útsynningsasia",
			territory.V_039: "Suðurevropa",
			territory.V_053: "Avstralasia",
			territory.V_054: "Melanesia",
			territory.V_057: "Mikronesi øki",
			territory.V_061: "Polynesia",
			territory.V_142: "Asia",
			territory.V_143: "Miðasia",
			territory.V_145: "Vesturasia",
			territory.V_150: "Evropa",
			territory.V_151: "Eysturevropa",
			territory.V_154: "Norðurevropa",
			territory.V_155: "Vesturevropa",
			territory.V_202: "Afrika sunnanfyri Sahara",
			territory.V_419: "Latínamerika",
			territory.AC:    "Ascension",
			territory.AD:    "Andorra",
			territory.AE:    "Sameindu Emirríkini",
			territory.AF:    "Afganistan",
			territory.AG:    "Antigua & Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Albania",
			territory.AM:    "Armenia",
			territory.AO:    "Angola",
			territory.AQ:    "Antarktis",
			territory.AR:    "Argentina",
			territory.AS:    "Amerikanska Samoa",
			territory.AT:    "Eysturríki",
			territory.AU:    "Avstralia",
			territory.AW:    "Aruba",
			territory.AX:    "Áland",
			territory.AZ:    "Aserbadjan",
			territory.BA:    "Bosnia-Hersegovina",
			territory.BB:    "Barbados",
			territory.BD:    "Bangladesj",
			territory.BE:    "Belgia",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgaria",
			territory.BH:    "Barein",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BL:    "St. Barthélemy",
			territory.BM:    "Bermuda",
			territory.BN:    "Brunei",
			territory.BO:    "Bolivia",
			territory.BQ:    "Niðurlonds Karibia",
			territory.BR:    "Brasil",
			territory.BS:    "Bahamaoyggjar",
			territory.BT:    "Butan",
			territory.BV:    "Bouvetoyggj",
			territory.BW:    "Botsvana",
			territory.BY:    "Hvítarussland",
			territory.BZ:    "Belis",
			territory.CA:    "Kanada",
			territory.CC:    "Kokosoyggjar",
			territory.CD:    "Kongo, Dem. Lýðveldið",
			territory.CF:    "Miðafrikalýðveldið",
			territory.CG:    "Kongo",
			territory.CH:    "Sveis",
			territory.CI:    "Fílabeinsstrondin",
			territory.CK:    "Cooksoyggjar",
			territory.CL:    "Kili",
			territory.CM:    "Kamerun",
			territory.CN:    "Kina",
			territory.CO:    "Kolombia",
			territory.CP:    "Clipperton",
			territory.CR:    "Kosta Rika",
			territory.CU:    "Kuba",
			territory.CV:    "Grønhøvdaoyggjar",
			territory.CW:    "Curaçao",
			territory.CX:    "Jólaoyggjin",
			territory.CY:    "Kýpros",
			territory.CZ:    "Kekkia",
			territory.DE:    "Týskland",
			territory.DG:    "Diego Garcia",
			territory.DJ:    "Djibuti",
			territory.DK:    "Danmark",
			territory.DM:    "Dominika",
			territory.DO:    "Dominikalýðveldið",
			territory.DZ:    "Algeria",
			territory.EA:    "Ceuta & Melilla",
			territory.EC:    "Ekvador",
			territory.EE:    "Estland",
			territory.EG:    "Egyptaland",
			territory.EH:    "Vestursahara",
			territory.ER:    "Eritrea",
			territory.ES:    "Spania",
			territory.ET:    "Etiopia",
			territory.EU:    "Evropasamveldið",
			territory.EZ:    "Evrasona",
			territory.FI:    "Finnland",
			territory.FJ:    "Fiji",
			territory.FK:    "Falklandsoyggjar",
			territory.FM:    "Mikronesiasamveldið",
			territory.FO:    "Føroyar",
			territory.FR:    "Frakland",
			territory.GA:    "Gabon",
			territory.GB:    "Stórabretland",
			territory.GD:    "Grenada",
			territory.GE:    "Georgia",
			territory.GF:    "Franska Gujana",
			territory.GG:    "Guernsey",
			territory.GH:    "Gana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Grønland",
			territory.GM:    "Gambia",
			territory.GN:    "Guinea",
			territory.GP:    "Guadeloupe",
			territory.GQ:    "Ekvatorguinea",
			territory.GR:    "Grikkaland",
			territory.GS:    "Suðurgeorgia og Suðursandwichoyggjar",
			territory.GT:    "Guatemala",
			territory.GU:    "Guam",
			territory.GW:    "Guinea-Bissau",
			territory.GY:    "Gujana",
			territory.HK:    "Hong Kong SAR Kina",
			territory.HM:    "Heard og McDonaldoyggjar",
			territory.HN:    "Honduras",
			territory.HR:    "Kroatia",
			territory.HT:    "Haiti",
			territory.HU:    "Ungarn",
			territory.IC:    "Kanariuoyggjar",
			territory.ID:    "Indonesia",
			territory.IE:    "Írland",
			territory.IL:    "Ísrael",
			territory.IM:    "Isle of Man",
			territory.IN:    "India",
			territory.IO:    "Stóra Bretlands Indiahavoyggjar",
			territory.IQ:    "Irak",
			territory.IR:    "Iran",
			territory.IS:    "Ísland",
			territory.IT:    "Italia",
			territory.JE:    "Jersey",
			territory.JM:    "Jamaika",
			territory.JO:    "Jordan",
			territory.JP:    "Japan",
			territory.KE:    "Kenja",
			territory.KG:    "Kirgisia",
			territory.KH:    "Kambodja",
			territory.KI:    "Kiribati",
			territory.KM:    "Komoroyggjar",
			territory.KN:    "St. Kitts & Nevis",
			territory.KP:    "Norðurkorea",
			territory.KR:    "Suðurkorea",
			territory.KW:    "Kuvait",
			territory.KY:    "Caymanoyggjar",
			territory.KZ:    "Kasakstan",
			territory.LA:    "Laos",
			territory.LB:    "Libanon",
			territory.LC:    "St. Lusia",
			territory.LI:    "Liktinstein",
			territory.LK:    "Sri Lanka",
			territory.LR:    "Liberia",
			territory.LS:    "Lesoto",
			territory.LT:    "Litava",
			territory.LU:    "Luksemborg",
			territory.LV:    "Lettland",
			territory.LY:    "Libya",
			territory.MA:    "Marokko",
			territory.MC:    "Monako",
			territory.MD:    "Moldova",
			territory.ME:    "Montenegro",
			territory.MF:    "St-Martin",
			territory.MG:    "Madagaskar",
			territory.MH:    "Marshalloyggjar",
			territory.ML:    "Mali",
			territory.MM:    "Myanmar (Burma)",
			territory.MN:    "Mongolia",
			territory.MO:    "Makao SAR Kina",
			territory.MP:    "Norðaru Mariuoyggjar",
			territory.MQ:    "Martinique",
			territory.MR:    "Móritania",
			territory.MS:    "Montserrat",
			territory.MT:    "Malta",
			territory.MU:    "Móritius",
			territory.MV:    "Maldivoyggjar",
			territory.MW:    "Malavi",
			territory.MX:    "Meksiko",
			territory.MY:    "Malaisia",
			territory.MZ:    "Mosambik",
			territory.NA:    "Namibia",
			territory.NC:    "Nýkaledónia",
			territory.NE:    "Niger",
			territory.NF:    "Norfolksoyggj",
			territory.NG:    "Nigeria",
			territory.NI:    "Nikaragua",
			territory.NL:    "Niðurlond",
			territory.NO:    "Noreg",
			territory.NP:    "Nepal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Nýsæland",
			territory.OM:    "Oman",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PF:    "Franska Polynesia",
			territory.PG:    "Papua Nýguinea",
			territory.PH:    "Filipsoyggjar",
			territory.PK:    "Pakistan",
			territory.PL:    "Pólland",
			territory.PM:    "Saint Pierre & Miquelon",
			territory.PN:    "Pitcairnoyggjar",
			territory.PR:    "Puerto Riko",
			territory.PS:    "Palestinskt landøki",
			territory.PT:    "Portugal",
			territory.PW:    "Palau",
			territory.PY:    "Paraguai",
			territory.QA:    "Katar",
			territory.QO:    "fjarskoti Osiania",
			territory.RE:    "Réunion",
			territory.RO:    "Rumenia",
			territory.RS:    "Serbia",
			territory.RU:    "Russland",
			territory.RW:    "Ruanda",
			territory.SA:    "Saudiarabia",
			territory.SB:    "Salomonoyggjar",
			territory.SC:    "Seyskelloyggjar",
			territory.SD:    "Sudan",
			territory.SE:    "Svøríki",
			territory.SG:    "Singapor",
			territory.SH:    "St. Helena",
			territory.SI:    "Slovenia",
			territory.SJ:    "Svalbard & Jan Mayen",
			territory.SK:    "Slovakia",
			territory.SL:    "Sierra Leona",
			territory.SM:    "San Marino",
			territory.SN:    "Senegal",
			territory.SO:    "Somalia",
			territory.SR:    "Surinam",
			territory.SS:    "Suðursudan",
			territory.ST:    "Sao Tome & Prinsipi",
			territory.SV:    "El Salvador",
			territory.SX:    "Sint Maarten",
			territory.SY:    "Sýria",
			territory.SZ:    "Esvatini",
			territory.TA:    "Tristan da Cunha",
			territory.TC:    "Turks- og Caicosoyggjar",
			territory.TD:    "Kjad",
			territory.TF:    "Fronsku sunnaru landaøki",
			territory.TG:    "Togo",
			territory.TH:    "Tailand",
			territory.TJ:    "Tadsjikistan",
			territory.TK:    "Tokelau",
			territory.TL:    "Eysturtimor",
			territory.TM:    "Turkmenistan",
			territory.TN:    "Tunesia",
			territory.TO:    "Tonga",
			territory.TR:    "Turkaland",
			territory.TT:    "Trinidad & Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taivan",
			territory.TZ:    "Tansania",
			territory.UA:    "Ukraina",
			territory.UG:    "Uganda",
			territory.UM:    "Sambandsríki Amerikas fjarskotnu oyggjar",
			territory.UN:    "Sameindu Tjóðir",
			territory.US:    "Sambandsríki Amerika",
			territory.UY:    "Uruguai",
			territory.UZ:    "Usbekistan",
			territory.VA:    "Vatikanbýur",
			territory.VC:    "St. Vinsent & Grenadinoyggjar",
			territory.VE:    "Venesuela",
			territory.VG:    "Stóra Bretlands Jomfrúoyggjar",
			territory.VI:    "Sambandsríki Amerikas Jomfrúoyggjar",
			territory.VN:    "Vjetnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Wallis- og Futunaoyggjar",
			territory.WS:    "Samoa",
			territory.XA:    "óekta tónalag",
			territory.XB:    "óektaður BIDI tekstur",
			territory.XK:    "Kosovo",
			territory.YE:    "Jemen",
			territory.YT:    "Mayotte",
			territory.ZA:    "Suðurafrika",
			territory.ZM:    "Sambia",
			territory.ZW:    "Simbabvi",
			territory.ZZ:    "ókent øki",
		},
	}
}
