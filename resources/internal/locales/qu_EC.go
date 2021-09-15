// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_qu_EC() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "qu_EC",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{0} {1}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ene", Feb: "Feb", Mar: "Mar", Apr: "Abr", May: "May", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Set", Oct: "Oct", Nov: "Nov", Dec: "Dic"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Enero", Feb: "Febrero", Mar: "Marzo", Apr: "Abril", May: "Mayo", Jun: "Junio", Jul: "Julio", Aug: "Agosto", Sep: "Setiembre", Oct: "Octubre", Nov: "Noviembre", Dec: "Diciembre"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dom", Mon: "Lun", Tue: "Mar", Wed: "Mié", Thu: "Jue", Fri: "Vie", Sat: "Sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "X", Thu: "J", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Dom", Mon: "Lun", Tue: "Mar", Wed: "Mié", Thu: "Jue", Fri: "Vie", Sat: "Sab"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Domingo", Mon: "Lunes", Tue: "Martes", Wed: "Miércoles", Thu: "Jueves", Fri: "Viernes", Sat: "Sábado"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", CurrencyAccounting: "¤\u00a0#,##0.00", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afgani Afgano", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Lek albanés", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Florín Antillano Neerlandés", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza Angoleño", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Peso Argentino", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Dólar Australiano", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Florín Arubeño", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Marco Bosnioherzegovino", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Dólar de Barbados", Symbol: "BBG"},
				currency.BDT: cldr.Currency{DisplayName: "Taka Bangladesí", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Lev", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Dinar Bareiní", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Franco Burundés", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Dólar Bermudeño", Symbol: "DBM"},
				currency.BND: cldr.Currency{DisplayName: "Dólar de Brunéi", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Real Brasileño", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dólar Bahameño", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum Butanés", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Pula Botswano", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Nuevo Rublo Bielorruso", Symbol: "BYN"},
				currency.BZD: cldr.Currency{DisplayName: "Dólar Beliceño", Symbol: "DBZ"},
				currency.CAD: cldr.Currency{DisplayName: "Dólar Canadiense", Symbol: "$CA"},
				currency.CDF: cldr.Currency{DisplayName: "Franco Congoleño", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Franco Suizo", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Peso Chileno", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan Chino (offshore)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Chino", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso Colombiano", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Colón Costarricense", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Peso Cubano Convertible", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Peso Cubano", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo Caboverdiano", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Corona Checa", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Franco Yibutiano", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Corona Danesa", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Peso Dominicano", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar Argelino", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Libra Egipcia", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa Eritreano", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr Etíope", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dólar Fiyiano", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Libra Malvinense", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Libra Esterlina", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Lari Georgiano", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Cedi Ganés", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Libra Gibraltareña", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Franco Guineano", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal Guatemalteco", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Dólar Guyanés", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Dólar de Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira Hondureño", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna Croata", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde Haitiano", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Florín Húngaro", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Rupia Indonesia", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Nuevo Séquel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupia India", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinar Iraquí", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Rial Iraní", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Corona Islandesa", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Dólar Jamaiquino", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Dinar Jordano", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Yen Japonés", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Chelín Keniano", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Som Kirguís", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Riel Camboyano", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Franco Comorense", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Won Norcoreano", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Won Surcoreano", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinar Kuwaití", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Dólar de las Islas Caimán", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge Kazajo", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Kip Laosiano", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Libra Libanesa", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Rupia de Sri Lanka", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Dólar Liberiano", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinar Libio", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Dírham Marroquí", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Leu Moldavo", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Ariary Malgache", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Dinar Macedonio", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Kyat Birmano", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik Mongol", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca Macaense", Symbol: "MOP"},
				currency.MRU: cldr.Currency{DisplayName: "Uguiya Mauritano", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "Rupia de Mauricio", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Rupia de Maldivas", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha Malauí", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Peso Mexicano", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit Malayo", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Metical Mozambiqueño", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Dólar Namibio", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naira Nigeriano", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Córdova Nicaragüense", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Corona Noruega", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Rupia Nepalí", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Dólar Neozelandés", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial Omaní", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Balboa Panameño", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Sol Peruano", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Kina Papuano", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Peso Filipino", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Rupia Pakistaní", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Guaraní Paraguayo", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Riyal Catarí", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Leu Rumano", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar Serbio", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rublo Ruso", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Franco Ruandés", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal Saudí", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Dólar de las Islas Salomón", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia de Seychelles", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Libra Sudanesa", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Corona Sueca", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Dólar de Singapur", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Libra de Santa Helena", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Leone de Sierra Leona", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Chelín Somalí", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Dólar Surinamés", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Libra Sursudanesa", Symbol: "SSP"},
				currency.STN: cldr.Currency{DisplayName: "Dobra Santotomense", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Libra Siria", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni Swazi", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Baht Tailandés", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "Somoni Tayiko", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Manat Turcomano", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Dinar Tunecino", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Paʻanga Tongano", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Lira Turca", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Dólar de Trinidad y Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Nuevo Dólar Taiwanés", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Chelín Tanzano", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Grivna", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Chelín Ugandés", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Dólar Americano", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "Peso Uruguayo", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Som Ubzeko", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Bolívar Venezolano", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Dong Vietnamita", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Tala Samoano", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Franco CFA de África Central", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dólar del Caribe Oriental", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Franco CFA de África Occidental", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Franco CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Mana riqsisqa Qullqi", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Rial Yemení", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Rand Sudafricano", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha Zambiano", Symbol: "ZMW"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AF:      "Afrikaans Simi",
			language.AGQ:     "Aghem Simi",
			language.AK:      "Akan Simi",
			language.AM:      "Amarico Simi",
			language.AR:      "Arabe Simi",
			language.ARN:     "Mapuche Simi",
			language.AS:      "Asames Simi",
			language.ASA:     "Asu Simi",
			language.AST:     "Asturiano Simi",
			language.AY:      "Aymara Simi",
			language.AZ:      "Azerí Simi",
			language.BA:      "Baskir Simi",
			language.BAS:     "Basaa Simi",
			language.BE:      "Bielorruso Simi",
			language.BEM:     "Bemba Simi",
			language.BEZ:     "Bena Simi",
			language.BG:      "Bulgaro Simi",
			language.BM:      "Bambara Simi",
			language.BN:      "Bangla Simi",
			language.BO:      "Tibetano Simi",
			language.BR:      "Breton Simi",
			language.BRX:     "Bodo Simi",
			language.BS:      "Bosnio Simi",
			language.CA:      "Catalan Simi",
			language.CCP:     "Chakma Simi",
			language.CE:      "Checheno Simi",
			language.CEB:     "Cebuano Simi",
			language.CGG:     "Kiga Simi",
			language.CHR:     "Cheroqui Simi",
			language.CKB:     "Chawpi Kurdo Simi",
			language.CO:      "Corso Simi",
			language.CS:      "Checo Simi",
			language.CU:      "Eslavo Eclesiástico Simi",
			language.CY:      "Gales Simi",
			language.DA:      "Danes Simi",
			language.DAV:     "Taita Simi",
			language.DE:      "Aleman Simi",
			language.DJE:     "Zarma Simi",
			language.DSB:     "Bajo Sorbio Simi",
			language.DUA:     "Duala Simi",
			language.DV:      "Divehi Simi",
			language.DYO:     "Jola-Fonyi Simi",
			language.DZ:      "Butanés Simi",
			language.EBU:     "Embu Simi",
			language.EE:      "Ewé Simi",
			language.EL:      "Griego Simi",
			language.EN:      "Ingles Simi",
			language.EN_GB:   "Ingles Simi (GB)",
			language.EN_US:   "Ingles Simi (US)",
			language.EO:      "Esperanto Simi",
			language.ES:      "Español Simi",
			language.ES_419:  "Español Simi (Latino América)",
			language.ET:      "Estonio Simi",
			language.EU:      "Euskera Simi",
			language.EWO:     "Ewondo Simi",
			language.FA:      "Persa Simi",
			language.FF:      "Fulah Simi",
			language.FI:      "Fines Simi",
			language.FIL:     "Filipino Simi",
			language.FO:      "Feroes Simi",
			language.FR:      "Frances Simi",
			language.FUR:     "Friulano Simi",
			language.FY:      "Frison Simi",
			language.GA:      "Irlandes Simi",
			language.GD:      "Gaelico Escoces Simi",
			language.GL:      "Gallego Simi",
			language.GSW:     "Alsaciano Simi",
			language.GU:      "Gujarati Simi",
			language.GUZ:     "Guzí Simi",
			language.GV:      "Manés Simi",
			language.HA:      "Hausa Simi",
			language.HAW:     "Hawaiano Simi",
			language.HE:      "Hebreo Simi",
			language.HI:      "Hindi Simi",
			language.HMN:     "Hmong Daw Simi",
			language.HR:      "Croata Simi",
			language.HSB:     "Alto Sorbio Simi",
			language.HT:      "Haitiano Criollo Simi",
			language.HU:      "Hungaro Simi",
			language.HY:      "Armenio Simi",
			language.IA:      "Interlingua Simi",
			language.ID:      "Indonesio Simi",
			language.IG:      "Igbo Simi",
			language.II:      "Yi Simi",
			language.IS:      "Islandes Simi",
			language.IT:      "Italiano Simi",
			language.IU:      "Inuktitut Simi",
			language.JA:      "Japones Simi",
			language.JGO:     "Ngomba Simi",
			language.JMC:     "Machame Simi",
			language.JV:      "Javanés Simi",
			language.KA:      "Georgiano Simi",
			language.KAB:     "Cabilio Simi",
			language.KAM:     "Kamba Simi",
			language.KDE:     "Makonde Simi",
			language.KEA:     "Caboverdiano Simi",
			language.KHQ:     "Koyra Chiini Simi",
			language.KI:      "Kikuyu Simi",
			language.KK:      "Kazajo Simi",
			language.KKJ:     "Kako Simi",
			language.KL:      "Groenlandes Simi",
			language.KLN:     "Kalenjin Simi",
			language.KM:      "Khmer Simi",
			language.KN:      "Kannada Simi",
			language.KO:      "Coreano Simi",
			language.KOK:     "Konkani Simi",
			language.KS:      "Cachemir Simi",
			language.KSB:     "Shambala Simi",
			language.KSF:     "Bafia Simi",
			language.KSH:     "Kölsch Simi",
			language.KU:      "Kurdo Simi",
			language.KW:      "Córnico Simi",
			language.KY:      "Kirghiz Simi",
			language.LA:      "Latín Simi",
			language.LAG:     "Langi Simi",
			language.LB:      "Luxemburgues Simi",
			language.LG:      "Luganda Simi",
			language.LKT:     "Lakota Simi",
			language.LN:      "Lingala Simi",
			language.LO:      "Lao Simi",
			language.LRC:     "Luri septentrional Simi",
			language.LT:      "Lituano Simi",
			language.LU:      "Luba-Katanga Simi",
			language.LUO:     "Luo Simi",
			language.LUY:     "Luyia Simi",
			language.LV:      "Leton Simi",
			language.MAS:     "Masai Simi",
			language.MER:     "Meru Simi",
			language.MFE:     "Mauriciano Simi",
			language.MG:      "Malgache Simi",
			language.MGH:     "Makhuwa-Meetto Simi",
			language.MGO:     "Metaʼ Simi",
			language.MI:      "Maori Simi",
			language.MK:      "Macedonio Simi",
			language.ML:      "Malayalam Simi",
			language.MN:      "Mongol Simi",
			language.MOH:     "Mohawk Simi",
			language.MR:      "Marathi Simi",
			language.MS:      "Malayo Simi",
			language.MT:      "Maltes Simi",
			language.MUA:     "Mundang Simi",
			language.MUL:     "Idiomas M´últiples Simi",
			language.MY:      "Birmano Simi",
			language.MZN:     "Mazandaraní Simi",
			language.NAQ:     "Nama Simi",
			language.NB:      "Noruego Bokmål Simi",
			language.ND:      "Ndebele septentrional Simi",
			language.NDS:     "Bajo Alemán Simi",
			language.NE:      "Nepali Simi",
			language.NL:      "Neerlandes Simi",
			language.NL_BE:   "Flamenco Simi",
			language.NMG:     "Kwasio Ngumba Simi",
			language.NN:      "Noruego Nynorsk Simi",
			language.NNH:     "Ngiemboon Simi",
			language.NO:      "Noruego Simi",
			language.NSO:     "Sesotho Sa Leboa Simi",
			language.NUS:     "Nuer Simi",
			language.NY:      "Nyanja Simi",
			language.NYN:     "Nyankole Simi",
			language.OC:      "Occitano Simi",
			language.OM:      "Oromo Simi",
			language.OR:      "Odia Simi",
			language.OS:      "Osetio Simi",
			language.PA:      "Punyabi Simi",
			language.PAP:     "Papiamento Simi",
			language.PL:      "Polaco Simi",
			language.PRG:     "Prusiano Simi",
			language.PS:      "Pashto Simi",
			language.PT:      "Portugues Simi",
			language.QU:      "Runasimi",
			language.QUC:     "Kʼicheʼ Simi",
			language.RM:      "Romanche Simi",
			language.RN:      "Rundi Simi",
			language.RO:      "Rumano Simi",
			language.ROF:     "Rombo Simi",
			language.RU:      "Ruso Simi",
			language.RW:      "Kinyarwanda Simi",
			language.RWK:     "Rwa Simi",
			language.SA:      "Sanscrito Simi",
			language.SAH:     "Sakha Simi",
			language.SAQ:     "Samburu Simi",
			language.SBP:     "Sangu Simi",
			language.SD:      "Sindhi Simi",
			language.SE:      "Chincha Sami Simi",
			language.SEH:     "Sena Simi",
			language.SES:     "Koyraboro Senni Simi",
			language.SG:      "Sango Simi",
			language.SHI:     "Tashelhit Simi",
			language.SI:      "Cingales Simi",
			language.SK:      "Eslovaco Simi",
			language.SL:      "Esloveno Simi",
			language.SM:      "Samoano Simi",
			language.SMA:     "Qulla Sami Simi",
			language.SMJ:     "Sami Lule Simi",
			language.SMN:     "Sami Inari Simi",
			language.SMS:     "Sami Skolt Simi",
			language.SN:      "Shona Simi",
			language.SO:      "Somali Simi",
			language.SQ:      "Albanes Simi",
			language.SR:      "Serbio Simi",
			language.ST:      "Soto Meridional Simi",
			language.SU:      "Sundanés Simi",
			language.SV:      "Sueco Simi",
			language.SW:      "Suajili Simi",
			language.SW_CD:   "Suajili Simi (Congo (RDC))",
			language.SYR:     "Siriaco Simi",
			language.TA:      "Tamil Simi",
			language.TE:      "Telugu Simi",
			language.TEO:     "Teso Simi",
			language.TG:      "Tayiko Simi",
			language.TH:      "Tailandes Simi",
			language.TI:      "Tigriña Simi",
			language.TK:      "Turcomano Simi",
			language.TN:      "Setsuana Simi",
			language.TO:      "Tongano Simi",
			language.TR:      "Turco Simi",
			language.TT:      "Tartaro Simi",
			language.TWQ:     "Tasawaq Simi",
			language.TZM:     "Tamazight Simi",
			language.UG:      "Uigur Simi",
			language.UK:      "Ucraniano Simi",
			language.UND:     "Mana Riqsisqa Simi",
			language.UR:      "Urdu Simi",
			language.UZ:      "Uzbeko Simi",
			language.VAI:     "Vai Simi",
			language.VI:      "Vietnamita Simi",
			language.VO:      "Volapük Simi",
			language.VUN:     "Vunjo Simi",
			language.WAE:     "Walser Simi",
			language.WO:      "Wolof Simi",
			language.XH:      "Isixhosa Simi",
			language.XOG:     "Soga Simi",
			language.YAV:     "Yangben Simi",
			language.YI:      "Yiddish Simi",
			language.YO:      "Yoruba Simi",
			language.YUE:     "Chino Cantonés Simi",
			language.ZGH:     "Bereber Marroquí Estándar Simi",
			language.ZH:      "Chino Mandarín Simi",
			language.ZH_HANS: "Chino Mandarín Simplificado Simi",
			language.ZH_HANT: "Chino Mandarín Tradicional Simi",
			language.ZU:      "Isizulu Simi",
			language.ZXX:     "Sin contenido linguístico",
		},
		Territories: cldr.Territories{
			territory.AC: "Islas Ascensión",
			territory.AD: "Andorra",
			territory.AE: "Emiratos Árabes Unidos",
			territory.AF: "Afganistán",
			territory.AG: "Antigua y Barbuda",
			territory.AI: "Anguila",
			territory.AL: "Albania",
			territory.AM: "Armenia",
			territory.AO: "Angola",
			territory.AQ: "Antártida",
			territory.AR: "Argentina",
			territory.AS: "Samoa Americana",
			territory.AT: "Austria",
			territory.AU: "Australia",
			territory.AW: "Aruba",
			territory.AX: "Islas Åland",
			territory.AZ: "Azerbaiyán",
			territory.BA: "Bosnia y Herzegovina",
			territory.BB: "Barbados",
			territory.BD: "Bangladesh",
			territory.BE: "Bélgica",
			territory.BF: "Burkina Faso",
			territory.BG: "Bulgaria",
			territory.BH: "Baréin",
			territory.BI: "Burundi",
			territory.BJ: "Benín",
			territory.BL: "San Bartolomé",
			territory.BM: "Bermudas",
			territory.BN: "Brunéi",
			territory.BO: "Bolivia",
			territory.BQ: "Bonaire",
			territory.BR: "Brasil",
			territory.BS: "Bahamas",
			territory.BT: "Bután",
			territory.BV: "Isla Bouvet",
			territory.BW: "Botsuana",
			territory.BY: "Belarús",
			territory.BZ: "Belice",
			territory.CA: "Canadá",
			territory.CC: "Islas Cocos",
			territory.CD: "Congo (RDC)",
			territory.CF: "República Centroafricana",
			territory.CG: "Congo",
			territory.CH: "Suiza",
			territory.CI: "Côte d’Ivoire",
			territory.CK: "Islas Cook",
			territory.CL: "Chile",
			territory.CM: "Camerún",
			territory.CN: "China",
			territory.CO: "Colombia",
			territory.CP: "Isla Clipperton",
			territory.CR: "Costa Rica",
			territory.CU: "Cuba",
			territory.CV: "Cabo Verde",
			territory.CW: "Curazao",
			territory.CX: "Isla Christmas",
			territory.CY: "Chipre",
			territory.CZ: "Chequia",
			territory.DE: "Alemania",
			territory.DG: "Diego García",
			territory.DJ: "Yibuti",
			territory.DK: "Dinamarca",
			territory.DM: "Dominica",
			territory.DO: "República Dominicana",
			territory.DZ: "Argelia",
			territory.EA: "Ceuta y Melilla",
			territory.EC: "Ecuador",
			territory.EE: "Estonia",
			territory.EG: "Egipto",
			territory.EH: "Sahara Occidental",
			territory.ER: "Eritrea",
			territory.ES: "España",
			territory.ET: "Etiopía",
			territory.FI: "Finlandia",
			territory.FJ: "Fiyi",
			territory.FK: "Islas Malvinas",
			territory.FM: "Micronesia",
			territory.FO: "Islas Feroe",
			territory.FR: "Francia",
			territory.GA: "Gabón",
			territory.GB: "Reino Unido",
			territory.GD: "Granada",
			territory.GE: "Georgia",
			territory.GF: "Guayana Francesa",
			territory.GG: "Guernesey",
			territory.GH: "Ghana",
			territory.GI: "Gibraltar",
			territory.GL: "Groenlandia",
			territory.GM: "Gambia",
			territory.GN: "Guinea",
			territory.GP: "Guadalupe",
			territory.GQ: "Guinea Ecuatorial",
			territory.GR: "Grecia",
			territory.GS: "Georgia del Sur e Islas Sandwich del Sur",
			territory.GT: "Guatemala",
			territory.GU: "Guam",
			territory.GW: "Guinea-Bisáu",
			territory.GY: "Guyana",
			territory.HK: "Hong Kong (RAE)",
			territory.HM: "Islas Heard y McDonald",
			territory.HN: "Honduras",
			territory.HR: "Croacia",
			territory.HT: "Haití",
			territory.HU: "Hungría",
			territory.IC: "Islas Canarias",
			territory.ID: "Indonesia",
			territory.IE: "Irlanda",
			territory.IL: "Israel",
			territory.IM: "Isla de Man",
			territory.IN: "India",
			territory.IO: "Territorio Británico del Océano Índico",
			territory.IQ: "Irak",
			territory.IR: "Irán",
			territory.IS: "Islandia",
			territory.IT: "Italia",
			territory.JE: "Jersey",
			territory.JM: "Jamaica",
			territory.JO: "Jordania",
			territory.JP: "Japón",
			territory.KE: "Kenia",
			territory.KG: "Kirguistán",
			territory.KH: "Camboya",
			territory.KI: "Kiribati",
			territory.KM: "Comoras",
			territory.KN: "San Cristóbal y Nieves",
			territory.KP: "Corea del Norte",
			territory.KR: "Corea del Sur",
			territory.KW: "Kuwait",
			territory.KY: "Islas Caimán",
			territory.KZ: "Kazajistán",
			territory.LA: "Laos",
			territory.LB: "Líbano",
			territory.LC: "Santa Lucia",
			territory.LI: "Liechtenstein",
			territory.LK: "Sri Lanka",
			territory.LR: "Liberia",
			territory.LS: "Lesoto",
			territory.LT: "Lituania",
			territory.LU: "Luxemburgo",
			territory.LV: "Letonia",
			territory.LY: "Libia",
			territory.MA: "Marruecos",
			territory.MC: "Mónaco",
			territory.MD: "Moldova",
			territory.ME: "Montenegro",
			territory.MF: "San Martín",
			territory.MG: "Madagascar",
			territory.MH: "Islas Marshall",
			territory.MK: "Macedonia del Norte",
			territory.ML: "Malí",
			territory.MM: "Myanmar",
			territory.MN: "Mongolia",
			territory.MO: "Macao RAE",
			territory.MP: "Islas Marianas del Norte",
			territory.MQ: "Martinica",
			territory.MR: "Mauritania",
			territory.MS: "Montserrat",
			territory.MT: "Malta",
			territory.MU: "Mauricio",
			territory.MV: "Maldivas",
			territory.MW: "Malawi",
			territory.MX: "México",
			territory.MY: "Malasia",
			territory.MZ: "Mozambique",
			territory.NA: "Namibia",
			territory.NC: "Nueva Caledonia",
			territory.NE: "Níger",
			territory.NF: "Isla Norfolk",
			territory.NG: "Nigeria",
			territory.NI: "Nicaragua",
			territory.NL: "Países Bajos",
			territory.NO: "Noruega",
			territory.NP: "Nepal",
			territory.NR: "Nauru",
			territory.NU: "Niue",
			territory.NZ: "Nueva Zelanda",
			territory.OM: "Omán",
			territory.PA: "Panamá",
			territory.PE: "Perú",
			territory.PF: "Polinesia Francesa",
			territory.PG: "Papúa Nueva Guinea",
			territory.PH: "Filipinas",
			territory.PK: "Pakistán",
			territory.PL: "Polonia",
			territory.PM: "San Pedro y Miquelón",
			territory.PN: "Islas Pitcairn",
			territory.PR: "Puerto Rico",
			territory.PS: "Palestina Kamachikuq",
			territory.PT: "Portugal",
			territory.PW: "Palaos",
			territory.PY: "Paraguay",
			territory.QA: "Qatar",
			territory.RE: "Reunión",
			territory.RO: "Rumania",
			territory.RS: "Serbia",
			territory.RU: "Rusia",
			territory.RW: "Ruanda",
			territory.SA: "Arabia Saudí",
			territory.SB: "Islas Salomón",
			territory.SC: "Seychelles",
			territory.SD: "Sudán",
			territory.SE: "Suecia",
			territory.SG: "Singapur",
			territory.SH: "Santa Elena",
			territory.SI: "Eslovenia",
			territory.SJ: "Svalbard y Jan Mayen",
			territory.SK: "Eslovaquia",
			territory.SL: "Sierra Leona",
			territory.SM: "San Marino",
			territory.SN: "Senegal",
			territory.SO: "Somalia",
			territory.SR: "Surinam",
			territory.SS: "Sudán del Sur",
			territory.ST: "Santo Tomé y Príncipe",
			territory.SV: "El Salvador",
			territory.SX: "Sint Maarten",
			territory.SY: "Siria",
			territory.SZ: "Suazilandia",
			territory.TA: "Tristán de Acuña",
			territory.TC: "Islas Turcas y Caicos",
			territory.TD: "Chad",
			territory.TF: "Territorios Australes Franceses",
			territory.TG: "Togo",
			territory.TH: "Tailandia",
			territory.TJ: "Tayikistán",
			territory.TK: "Tokelau",
			territory.TL: "Timor-Leste",
			territory.TM: "Turkmenistán",
			territory.TN: "Túnez",
			territory.TO: "Tonga",
			territory.TR: "Turquía",
			territory.TT: "Trinidad y Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwán",
			territory.TZ: "Tanzania",
			territory.UA: "Ucrania",
			territory.UG: "Uganda",
			territory.UM: "Islas menores alejadas de los EE.UU.",
			territory.US: "Estados Unidos",
			territory.UY: "Uruguay",
			territory.UZ: "Uzbekistán",
			territory.VA: "Santa Sede (Ciudad del Vaticano)",
			territory.VC: "San Vicente y las Granadinas",
			territory.VE: "Venezuela",
			territory.VG: "Islas Vírgenes Británicas",
			territory.VI: "EE.UU. Islas Vírgenes",
			territory.VN: "Vietnam",
			territory.VU: "Vanuatu",
			territory.WF: "Wallis y Futuna",
			territory.WS: "Samoa",
			territory.XK: "Kosovo",
			territory.YE: "Yemen",
			territory.YT: "Mayotte",
			territory.ZA: "Sudáfrica",
			territory.ZM: "Zambia",
			territory.ZW: "Zimbabue",
		},
	}
}
