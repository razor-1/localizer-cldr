// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_ce_RU() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ce_RU",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "янв", Feb: "фев", Mar: "мар", Apr: "апр", May: "май", Jun: "июн", Jul: "июл", Aug: "авг", Sep: "сен", Oct: "окт", Nov: "ноя", Dec: "дек"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "январь", Feb: "февраль", Mar: "март", Apr: "апрель", May: "май", Jun: "июнь", Jul: "июль", Aug: "август", Sep: "сентябрь", Oct: "октябрь", Nov: "ноябрь", Dec: "декабрь"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "кӀи", Mon: "ор", Tue: "ши", Wed: "кха", Thu: "еа", Fri: "пӀе", Sat: "шуо"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "кӀ", Mon: "о", Tue: "ш", Wed: "кх", Thu: "е", Fri: "пӀ", Sat: "ш"}, Short: cldr.CalendarDayFormatNameValue{Sun: "кӀи", Mon: "ор", Tue: "ши", Wed: "кха", Thu: "еа", Fri: "пӀе", Sat: "шуо"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "кӀира", Mon: "оршот", Tue: "шинара", Wed: "кхаара", Thu: "еара", Fri: "пӀераска", Sat: "шуот"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Дирхам ӀЦЭ", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "ОвхӀан-пачхьалкхан афгани", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Албанин лек", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Эрмалойчоьнан драм", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Нидерландин Антилин гульден", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Анголан кванза", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Аргентинан песо", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Австралин доллар", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Арубан флорин", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Азербайджанан манат", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Боснин а, Герцеговинан а хийцалун марка", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Барбадосан доллар", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Бангладешан така", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Болгарин лев", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Бахрейнан динар", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Бурундин франк", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Бермудан доллар", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Брунейн доллар", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Боливин боливиано", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Бразилин реал", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Багаман доллар", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Бутанан нгултрум", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Ботсванан пула", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Белоруссин сом", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Белоруссин сом (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Белизин доллар", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Канадан доллар", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Конголезин франк", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Швейцарин франк", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Чилин песо", Symbol: "CLP"},
				currency.CNY: cldr.Currency{DisplayName: "Китайн юань", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Колумбин песо", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Костарикан колон", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Кубан хийцалун песо", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Кубан песо", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Кабо-Верден эскудо", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Чехин крона", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Джибутин франк", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Данин крона", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Доминикан песо", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Алжиран динар", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Мисаран фунт", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Эритрейн накфа", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Эфиопин быр", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Евро", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Фиджин доллар", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Фолклендан гӀайренийн фунт", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Англин фунт", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Гуьржийчоьнан лари", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Ганан седи", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Гибралтаран фунт", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Гамбин даласи", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Гвинейн франк", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Гватемалан кетсаль", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Гайанан доллар", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Гонконган доллар", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Гондурасан лемпира", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Хорватин куна", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Гаитин гурд", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Венгрин форинт", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Индонезин рупи", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Израилан керла шекель", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Индин рупи", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Ӏиракъан динар", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ГӀажарийчоьнан риал", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Исландин крона", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Ямайн доллар", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Урданан динар", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Японин иена", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Кенин шиллинг", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Киргизин сом", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Камбоджан риель", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Коморийн гӀайренийн франк", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Къилбаседа Корейн вона", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Къилба Корейн вона", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Кувейтан динар", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Кайманийн гӀайренийн доллар", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Кхазакхстанан тенге", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Лаосан кип", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Ливанан фунт", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Шри-Ланкан рупи", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Либерин доллар", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Ливин динар", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Мароккон дирхам", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Молдавин лей", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Малагасийн ариари", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Македонин динар", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Мьянман кьят", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Монголин тугрик", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Макаон патака", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Мавританин уги (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Мавританин уги", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Маврикин рупи", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Мальдивийн руфи", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Малавин квача", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Мексикан песо", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Малайзин ринггит", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Мозамбикан метикал", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Намибин доллар", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Нигерин найра", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Никарагуан кордоба", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Норвегин крона", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Непалан рупи", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Керла Зеландин доллар", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Оманан риал", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Панаман бальбоа", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Перун соль", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Папуа — Керла Гвинейн кина", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Филиппинийн песо", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Пакистанан рупи", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Польшан злотый", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Парагвайн гуарани", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Катаран риал", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Румынин лей", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Сербин динар", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Российн сом", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Руандан франк", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "СаӀудийн Ӏаьрбийчоьнан риал", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Соломонан гӀайренийн доллар", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Сейшелан рупи", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Суданан фунт", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Швецин крона", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Сингапуран доллар", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Сийлахьчу Еленин гӀайрен фунт", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Леоне", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Сомалин шиллинг", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Суринаман доллар", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Къилба Суданан фунт", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Сан-Томен а, Принсипин а добра (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Сан-Томен а, Принсипин а добра", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Шеман фунт", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Свазилендан лилангени", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Таиландан бат", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "Таджикистанан сомони", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Туркменин керла манат", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Тунисан динар", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Тонганан паанга", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Туркойчоьнан лира", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Тринидадан а, Тобагон а доллар", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Тайванан керла доллар", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Танзанин шиллинг", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Украинан гривна", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Угандан шиллинг", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "АЦШн доллар", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Уругвайн песо", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Узбекистанан сом", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Венесуэлан боливар (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Венесуэлан боливар", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Вьетнаман донг", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Вануатун вату", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Самоанан тала", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Юккъерчу Африкан КФА франк", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Малхбален Карибийн доллар", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Малхбузен Африкан КФА франк", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Французийн Тийна океанан франк", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "девзаш доцу я лелаш доцу ахча", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Йеменан риал", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Къилба-Африкин рэнд", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Замбин квача", Symbol: "ZMW"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "афарийн",
			language.AB:      "абхазхойн",
			language.ACE:     "ачехийн",
			language.ADA:     "адангме",
			language.ADY:     "адигейн",
			language.AF:      "африкаанс",
			language.AGQ:     "агхӀем",
			language.AIN:     "айнийн",
			language.AK:      "акан",
			language.ALE:     "алеутийн",
			language.ALT:     "къилба алтайн",
			language.AM:      "амхаройн",
			language.AN:      "арагонойн",
			language.ANP:     "ангика",
			language.AR:      "Ӏаьрбийн",
			language.AR_001:  "ХӀинца болу стандартан Ӏаьрбийн",
			language.ARN:     "арауканхойн",
			language.ARP:     "арапахо",
			language.AS:      "ассамийн",
			language.ASA:     "асу",
			language.AST:     "астурийн",
			language.AV:      "суьйлийн",
			language.AWA:     "авадхи",
			language.AY:      "аймара",
			language.AZ:      "азербайджанийн",
			language.BA:      "башкирийн",
			language.BAN:     "балийн",
			language.BAS:     "баса",
			language.BE:      "белорусийн",
			language.BEM:     "бемба",
			language.BEZ:     "бена",
			language.BG:      "болгарийн",
			language.BGN:     "малхбузен-белуджийн",
			language.BHO:     "бходжпури",
			language.BI:      "бислама",
			language.BIN:     "бини",
			language.BLA:     "сиксика",
			language.BM:      "бамбара",
			language.BN:      "бенгалийн",
			language.BO:      "тибетхойн",
			language.BR:      "бретонийн",
			language.BRX:     "бодо",
			language.BS:      "боснийн",
			language.BUG:     "бугийн",
			language.BYN:     "билийн",
			language.CA:      "каталонийн",
			language.CE:      "нохчийн",
			language.CEB:     "себуано",
			language.CGG:     "чига",
			language.CH:      "чаморро",
			language.CHK:     "чукчийн",
			language.CHM:     "марийн",
			language.CHO:     "чоктавийн",
			language.CHR:     "чероки",
			language.CHY:     "шайенийн",
			language.CKB:     "юккъерчу курдийн",
			language.CO:      "корсиканийн",
			language.CRS:     "сейшелийн креолийн",
			language.CS:      "чехийн",
			language.CU:      "килсславянийн",
			language.CV:      "чувашийн",
			language.CY:      "валлийн",
			language.DA:      "датхойн",
			language.DAK:     "дакота",
			language.DAR:     "даьргӀойн",
			language.DAV:     "таита",
			language.DE:      "немцойн",
			language.DE_AT:   "австрин немцойн",
			language.DE_CH:   "швейцарин литературин немцойн",
			language.DGR:     "догриб",
			language.DJE:     "зарма",
			language.DSB:     "сорбийн",
			language.DUA:     "дуала",
			language.DV:      "мальдивийн",
			language.DYO:     "дьола-фоньи",
			language.DZ:      "дзонг-кэ",
			language.DZG:     "даза",
			language.EBU:     "эмбу",
			language.EE:      "эве",
			language.EFI:     "эфик",
			language.EKA:     "экаджук",
			language.EL:      "грекийн",
			language.EN:      "ингалсан",
			language.EN_AU:   "Австралин ингалсан",
			language.EN_CA:   "канадан ингалсан",
			language.EN_GB:   "ингалсан (Британи)",
			language.EN_US:   "американ ингалсан",
			language.EO:      "эсперанто",
			language.ES:      "испанхойн",
			language.ES_419:  "латинан американ испанхойн",
			language.ES_ES:   "европан испанхойн",
			language.ES_MX:   "мексикан испанхойн",
			language.ET:      "эстонийн",
			language.EU:      "баскийн",
			language.EWO:     "эвондо",
			language.FA:      "гӀажарийн",
			language.FF:      "фулах",
			language.FI:      "финнийн",
			language.FIL:     "филиппинийн",
			language.FJ:      "фиджи",
			language.FO:      "фарерийн",
			language.FON:     "фон",
			language.FR:      "французийн",
			language.FR_CA:   "канадан французийн",
			language.FR_CH:   "швейцарин французийн",
			language.FUR:     "фриулийн",
			language.FY:      "малхбузен-фризийн",
			language.GA:      "ирландхойн",
			language.GAA:     "га",
			language.GAG:     "гагаузийн",
			language.GD:      "гэлийн",
			language.GEZ:     "геэз",
			language.GIL:     "гильбертийн",
			language.GL:      "галисийн",
			language.GN:      "гуарани",
			language.GOR:     "горонтало",
			language.GSW:     "швейцарин немцойн",
			language.GU:      "гуджарати",
			language.GUZ:     "гусии",
			language.GV:      "мэнийн",
			language.GWI:     "гвичин",
			language.HA:      "хауса",
			language.HAW:     "гавайн",
			language.HE:      "жугтийн",
			language.HI:      "хӀинди",
			language.HIL:     "хилигайнон",
			language.HMN:     "хмонг",
			language.HR:      "хорватийн",
			language.HSB:     "лакхара сербийн",
			language.HT:      "гаитийн",
			language.HU:      "венгрийн",
			language.HUP:     "хупа",
			language.HY:      "эрмалойн",
			language.HZ:      "гереро",
			language.IA:      "интерлингва",
			language.IBA:     "ибанийн",
			language.IBB:     "ибибио",
			language.ID:      "индонезихойн",
			language.IG:      "игбо",
			language.II:      "сычуань",
			language.ILO:     "илоко",
			language.INH:     "гӀалгӀайн",
			language.IO:      "идо",
			language.IS:      "исландхойн",
			language.IT:      "итальянийн",
			language.IU:      "инуктитут",
			language.JA:      "японийн",
			language.JBO:     "ложбан",
			language.JGO:     "нгомба",
			language.JMC:     "мачаме",
			language.JV:      "яванийн",
			language.KA:      "гуьржийн",
			language.KAB:     "кабилийн",
			language.KAC:     "качинийн",
			language.KAJ:     "каджи",
			language.KAM:     "камба",
			language.KBD:     "гӀебартойн",
			language.KCG:     "тьяп",
			language.KDE:     "маконде",
			language.KEA:     "кабувердьяну",
			language.KFO:     "коро",
			language.KHA:     "кхаси",
			language.KHQ:     "койра чиини",
			language.KI:      "кикуйю",
			language.KJ:      "кунама",
			language.KK:      "кхазакхийн",
			language.KKJ:     "како",
			language.KL:      "гренландхойн",
			language.KLN:     "календжин",
			language.KM:      "кхмерийн",
			language.KMB:     "кимбунду",
			language.KN:      "каннада",
			language.KO:      "корейн",
			language.KOI:     "коми-пермякийн",
			language.KOK:     "конкани",
			language.KPE:     "кпелле",
			language.KR:      "канури",
			language.KRC:     "кхарачойн-балкхаройн",
			language.KRL:     "карелийн",
			language.KRU:     "курух",
			language.KS:      "кашмири",
			language.KSB:     "шамбала",
			language.KSF:     "бафиа",
			language.KSH:     "коьлнийн",
			language.KU:      "курдийн",
			language.KUM:     "гӀумкийн",
			language.KV:      "комийн",
			language.KW:      "корнуоллийн",
			language.KY:      "гӀиргӀизойн",
			language.LA:      "латинан",
			language.LAD:     "ладино",
			language.LAG:     "ланги",
			language.LB:      "люксембургхойн",
			language.LEZ:     "лаьзгийн",
			language.LG:      "ганда",
			language.LI:      "лимбургийн",
			language.LKT:     "лакота",
			language.LN:      "лингала",
			language.LO:      "лаоссийн",
			language.LOZ:     "лози",
			language.LRC:     "къилбаседа лури",
			language.LT:      "литвахойн",
			language.LU:      "луба-катанга",
			language.LUA:     "луба-лулуа",
			language.LUN:     "лунда",
			language.LUO:     "луо (Кени а, Танзани а)",
			language.LUS:     "лушей",
			language.LUY:     "лухья",
			language.LV:      "латышийн",
			language.MAD:     "мадурийн",
			language.MAG:     "магахи",
			language.MAI:     "майтхили",
			language.MAK:     "макасарийн",
			language.MAS:     "масаи",
			language.MDF:     "мокшанойн",
			language.MEN:     "менде",
			language.MER:     "меру",
			language.MFE:     "маврикин креолийн",
			language.MG:      "малагасийн",
			language.MGH:     "макуа-меетто",
			language.MGO:     "мета",
			language.MH:      "маршаллийн",
			language.MI:      "маори",
			language.MIC:     "микмак",
			language.MIN:     "минангкабау",
			language.MK:      "македонхойн",
			language.ML:      "малаялам",
			language.MN:      "монголийн",
			language.MNI:     "манипурийн",
			language.MOH:     "мохаук",
			language.MOS:     "моси",
			language.MR:      "маратхи",
			language.MS:      "малайн",
			language.MT:      "мальтойн",
			language.MUA:     "мунданг",
			language.MUL:     "тайп-тайпа доьзалан меттанаш",
			language.MUS:     "крик",
			language.MWL:     "мирандойн",
			language.MY:      "бирманийн",
			language.MYV:     "эрзянийн",
			language.MZN:     "мазандеранхойн",
			language.NA:      "науру",
			language.NAP:     "неаполитанойн",
			language.NAQ:     "нама",
			language.NB:      "норвегийн букмол",
			language.ND:      "къилбаседа ндебели",
			language.NDS:     "лахара германхойн",
			language.NDS_NL:  "лахара саксонийн",
			language.NE:      "непалхойн",
			language.NEW:     "неваройн",
			language.NG:      "ндонга",
			language.NIA:     "ниас",
			language.NIU:     "ниуэ",
			language.NL:      "голландхойн",
			language.NL_BE:   "фламандийн",
			language.NMG:     "квасио",
			language.NN:      "норвегийн нюнорск",
			language.NNH:     "нгиембунд",
			language.NOG:     "ногӀийн",
			language.NQO:     "нко",
			language.NR:      "къилба ндебеле",
			language.NSO:     "къилбаседа сото",
			language.NUS:     "нуэр",
			language.NV:      "навахо",
			language.NY:      "ньянджа",
			language.NYN:     "ньянколе",
			language.OC:      "окситанойн",
			language.OM:      "оромо",
			language.OR:      "ори",
			language.OS:      "хӀирийн",
			language.PA:      "панджаби",
			language.PAG:     "пангасинан",
			language.PAM:     "пампанга",
			language.PAP:     "папьяменто",
			language.PAU:     "палау",
			language.PCM:     "нигерийн-креолийн",
			language.PL:      "полякийн",
			language.PRG:     "пруссийн",
			language.PS:      "пушту",
			language.PT:      "португалихойн",
			language.PT_BR:   "бразилин португалихойн",
			language.PT_PT:   "европан португалихойн",
			language.QU:      "кечуа",
			language.QUC:     "киче",
			language.RAP:     "рапануйн",
			language.RAR:     "раротонга",
			language.RM:      "романшийн",
			language.RN:      "рунди",
			language.RO:      "румынийн",
			language.RO_MD:   "молдавийн",
			language.ROF:     "ромбо",
			language.ROOT:    "ораман мотт",
			language.RU:      "оьрсийн",
			language.RUP:     "аруминийн",
			language.RW:      "киньяруанда",
			language.RWK:     "руанда",
			language.SA:      "санскрит",
			language.SAD:     "сандаве",
			language.SAH:     "якутийн",
			language.SAQ:     "самбуру",
			language.SAT:     "сантали",
			language.SBA:     "нгамбайн",
			language.SBP:     "сангу",
			language.SC:      "сардинийн",
			language.SCN:     "сицилийн",
			language.SCO:     "шотландхойн",
			language.SD:      "синдхи",
			language.SE:      "къилбаседа саамийн",
			language.SEH:     "сена",
			language.SES:     "койраборо сенни",
			language.SG:      "санго",
			language.SHI:     "тахелхит",
			language.SHN:     "шанойн",
			language.SI:      "сингалхойн",
			language.SK:      "словакийн",
			language.SL:      "словенийн",
			language.SM:      "самоанойн",
			language.SMA:     "саамийн (къилба)",
			language.SMJ:     "луле-саамийн",
			language.SMN:     "инари-саамийн",
			language.SMS:     "скольт-саамийн",
			language.SN:      "шона",
			language.SNK:     "сонинке",
			language.SO:      "сомали",
			language.SQ:      "албанойн",
			language.SR:      "сербийн",
			language.SRN:     "сранан-тонго",
			language.SS:      "свази",
			language.SSY:     "сахо",
			language.ST:      "къилба сото",
			language.SU:      "сунданхойн",
			language.SUK:     "сукума",
			language.SV:      "шведийн",
			language.SW:      "суахили",
			language.SW_CD:   "суахили (Конго)",
			language.SWB:     "коморийн",
			language.SYR:     "шемахойн",
			language.TA:      "тамилхойн",
			language.TE:      "телугу",
			language.TEM:     "темне",
			language.TEO:     "тесо",
			language.TET:     "тетум",
			language.TG:      "таджикийн",
			language.TH:      "тайн",
			language.TI:      "тигринья",
			language.TIG:     "тигре",
			language.TK:      "туркменийн",
			language.TLH:     "клингонин",
			language.TN:      "тсвана",
			language.TO:      "тонганийн",
			language.TPI:     "ток-писин",
			language.TR:      "туркойн",
			language.TRV:     "седекойн",
			language.TS:      "тсонга",
			language.TT:      "гӀезалойн",
			language.TUM:     "тумбука",
			language.TVL:     "тувалу",
			language.TWQ:     "тасавак",
			language.TY:      "таитянойн",
			language.TYV:     "тувинийн",
			language.TZM:     "тамазигхтийн",
			language.UDM:     "удмуртийн",
			language.UG:      "уйгурийн",
			language.UK:      "украинийн",
			language.UMB:     "умбунду",
			language.UND:     "боьвзуш боцу мотт",
			language.UR:      "урду",
			language.UZ:      "узбекийн",
			language.VAI:     "ваи",
			language.VE:      "венда",
			language.VI:      "вьетнамхойн",
			language.VO:      "волапюк",
			language.VUN:     "вунджо",
			language.WA:      "валлонойн",
			language.WAE:     "валлисийн",
			language.WAL:     "воламо",
			language.WAR:     "варай",
			language.WBP:     "варлпири",
			language.WO:      "волоф",
			language.XAL:     "гӀалмакхойн",
			language.XH:      "коса",
			language.XOG:     "сога",
			language.YAV:     "янгбен",
			language.YBB:     "йемба",
			language.YI:      "идиш",
			language.YO:      "йоруба",
			language.YUE:     "кантонийн",
			language.ZGH:     "мороккон стандартан тамазигхтийн",
			language.ZH:      "цийн",
			language.ZH_HANS: "атта цийн",
			language.ZH_HANT: "ламастан цийн",
			language.ZU:      "зулу",
			language.ZUN:     "зуньи",
			language.ZXX:     "меттан чулацам боцуш",
			language.ZZA:     "заза",
		},
		Territories: cldr.Territories{
			territory.V_001: "Дерригдуьненан",
			territory.V_002: "Африка",
			territory.V_003: "Къилбаседа Америка",
			territory.V_005: "Къилба Америка",
			territory.V_009: "Океани",
			territory.V_011: "Малхбузен Африка",
			territory.V_013: "Юккъера Америка",
			territory.V_014: "Малхбален Африка",
			territory.V_015: "Къилбаседа Африка",
			territory.V_017: "Юккъера Африка",
			territory.V_018: "Къилба Африка",
			territory.V_019: "Къилбаседа а, къилба а Америка",
			territory.V_021: "Къилбаседа Америка – АЦШ а, Канада а",
			territory.V_029: "Карибаш",
			territory.V_030: "Юккъера Ази",
			territory.V_034: "Къилба Ази",
			territory.V_035: "Къилба-малхбален Ази",
			territory.V_039: "Къилба Европа",
			territory.V_053: "Австралази",
			territory.V_054: "Меланези",
			territory.V_057: "Микронези",
			territory.V_061: "Полинези",
			territory.V_142: "Ази",
			territory.V_143: "Юккъера Малхбале",
			territory.V_145: "Юккъера а, Гергара а Малхбале",
			territory.V_150: "Европа",
			territory.V_151: "Малхбален Европа",
			territory.V_154: "Къилбаседа Европа",
			territory.V_155: "Малхбузен Европа",
			territory.V_419: "Латинан Америка",
			territory.AC:    "Айъадаларан гӀайре",
			territory.AD:    "Андорра",
			territory.AE:    "Ӏарбийн Цхьанатоьхна Эмираташ",
			territory.AF:    "ОвхӀан мохк",
			territory.AG:    "Антигуа а, Барбуда а",
			territory.AI:    "Ангилья",
			territory.AL:    "Албани",
			territory.AM:    "Эрмалойчоь",
			territory.AO:    "Ангола",
			territory.AQ:    "Антарктида",
			territory.AR:    "Аргентина",
			territory.AS:    "Американ Самоа",
			territory.AT:    "Австри",
			territory.AU:    "Австрали",
			territory.AW:    "Аруба",
			territory.AX:    "Аландан гӀайренаш",
			territory.AZ:    "Азербайджан",
			territory.BA:    "Босни а, Герцеговина а",
			territory.BB:    "Барбадос",
			territory.BD:    "Бангладеш",
			territory.BE:    "Бельги",
			territory.BF:    "Буркина- Фасо",
			territory.BG:    "Болгари",
			territory.BH:    "Бахрейн",
			territory.BI:    "Бурунди",
			territory.BJ:    "Бенин",
			territory.BL:    "Сен-Бартельми",
			territory.BM:    "Бермудан гӀайренаш",
			territory.BN:    "Бруней-Даруссалам",
			territory.BO:    "Боливи",
			territory.BQ:    "Бонэйр, Синт-Эстатиус а, Саба а",
			territory.BR:    "Бразили",
			territory.BS:    "Багаман гӀайренаш",
			territory.BT:    "Бутан",
			territory.BV:    "Бувен гӀайре",
			territory.BW:    "Ботсвана",
			territory.BY:    "Белорусси",
			territory.BZ:    "Белиз",
			territory.CA:    "Канада",
			territory.CC:    "Кокосийн гӀайренаш",
			territory.CD:    "Демократин Республика Конго",
			territory.CF:    "Юккъерчу Африкин Республика",
			territory.CG:    "Конго - Браззавиль",
			territory.CH:    "Швейцари",
			territory.CI:    "Кот-Д’ивуар",
			territory.CK:    "Кукан гӀайренаш",
			territory.CL:    "Чили",
			territory.CM:    "Камерун",
			territory.CN:    "Цийчоь",
			territory.CO:    "Колумби",
			territory.CP:    "Клиппертон",
			territory.CR:    "Коста-Рика",
			territory.CU:    "Куба",
			territory.CV:    "Кабо-Верде",
			territory.CW:    "Кюрасао",
			territory.CX:    "ГӀайре ӏиса пайхӏамар вина де",
			territory.CY:    "Кипр",
			territory.CZ:    "Чехи",
			territory.DE:    "Германи",
			territory.DG:    "Диего-Гарси",
			territory.DJ:    "Джибути",
			territory.DK:    "Дани",
			territory.DM:    "Доминика",
			territory.DO:    "Доминикан Республика",
			territory.DZ:    "Алжир",
			territory.EA:    "Сеута а, Мелилья а",
			territory.EC:    "Эквадор",
			territory.EE:    "Эстони",
			territory.EG:    "Мисар",
			territory.EH:    "Малхбузен Саьхьара",
			territory.ER:    "Эритрей",
			territory.ES:    "Испани",
			territory.ET:    "Эфиопи",
			territory.EU:    "Евробарт",
			territory.EZ:    "еврозона",
			territory.FI:    "Финлянди",
			territory.FJ:    "Фиджи",
			territory.FK:    "Фолклендан гӀайренаш",
			territory.FM:    "Микронезин Федеративни штаташ",
			territory.FO:    "Фарерийн гӀайренаш",
			territory.FR:    "Франци",
			territory.GA:    "Габон",
			territory.GB:    "Йоккха Британи",
			territory.GD:    "Гренада",
			territory.GE:    "Гуьржийчоь",
			territory.GF:    "Французийн Гвиана",
			territory.GG:    "Гернси",
			territory.GH:    "Гана",
			territory.GI:    "Гибралтар",
			territory.GL:    "Гренланди",
			territory.GM:    "Гамби",
			territory.GN:    "Гвиней",
			territory.GP:    "Гваделупа",
			territory.GQ:    "Экваторан Гвиней",
			territory.GR:    "Греци",
			territory.GS:    "Къилба Джорджи а, Къилба Гавайн гӀайренаш а",
			territory.GT:    "Гватемала",
			territory.GU:    "Гуам",
			territory.GW:    "Гвиней-Бисау",
			territory.GY:    "Гайана",
			territory.HK:    "Гонконг (ша-къаьстина кӀошт)",
			territory.HM:    "Херд гӀайре а, Макдональд гӀайренаш а",
			territory.HN:    "Гондурас",
			territory.HR:    "Хорвати",
			territory.HT:    "Гаити",
			territory.HU:    "Венгри",
			territory.IC:    "Канаран гӀайренаш",
			territory.ID:    "Индонези",
			territory.IE:    "Ирланди",
			territory.IL:    "Израиль",
			territory.IM:    "Мэн гӀайре",
			territory.IN:    "ХӀинди",
			territory.IO:    "Британин латта Индин океанехь",
			territory.IQ:    "Ӏиракъ",
			territory.IR:    "ГӀажарийчоь",
			territory.IS:    "Исланди",
			territory.IT:    "Итали",
			territory.JE:    "Джерси",
			territory.JM:    "Ямайка",
			territory.JO:    "Урдан",
			territory.JP:    "Япони",
			territory.KE:    "Кени",
			territory.KG:    "Киргизи",
			territory.KH:    "Камбоджа",
			territory.KI:    "Кирибати",
			territory.KM:    "Комораш",
			territory.KN:    "Сент-Китс а, Невис а",
			territory.KP:    "Къилбаседа Корей",
			territory.KR:    "Къилба Корей",
			territory.KW:    "Кувейт",
			territory.KY:    "Кайман гӀайренаш",
			territory.KZ:    "Кхазакхстан",
			territory.LA:    "Лаос",
			territory.LB:    "Ливан",
			territory.LC:    "Сент-Люси",
			territory.LI:    "Лихтенштейн",
			territory.LK:    "Шри-Ланка",
			territory.LR:    "Либери",
			territory.LS:    "Лесото",
			territory.LT:    "Литва",
			territory.LU:    "Люксембург",
			territory.LV:    "Латви",
			territory.LY:    "Ливи",
			territory.MA:    "Марокко",
			territory.MC:    "Монако",
			territory.MD:    "Молдави",
			territory.ME:    "Ӏаьржаламанчоь",
			territory.MF:    "Сен-Мартен",
			territory.MG:    "Мадагаскар",
			territory.MH:    "Маршаллан гӀайренаш",
			territory.ML:    "Мали",
			territory.MM:    "Мьянма (Бирма)",
			territory.MN:    "Монголи",
			territory.MO:    "Макао (ша-къаьстина кӀошт)",
			territory.MP:    "Къилбаседа Марианан гӀайренаш",
			territory.MQ:    "Мартиника",
			territory.MR:    "Мавритани",
			territory.MS:    "Монтсеррат",
			territory.MT:    "Мальта",
			territory.MU:    "Маврики",
			territory.MV:    "Мальдиваш",
			territory.MW:    "Малави",
			territory.MX:    "Мексика",
			territory.MY:    "Малайзи",
			territory.MZ:    "Мозамбик",
			territory.NA:    "Намиби",
			territory.NC:    "Керла Каледони",
			territory.NE:    "Нигер",
			territory.NF:    "Норфолк гӀайре",
			territory.NG:    "Нигери",
			territory.NI:    "Никарагуа",
			territory.NL:    "Нидерландаш",
			territory.NO:    "Норвеги",
			territory.NP:    "Непал",
			territory.NR:    "Науру",
			territory.NU:    "Ниуэ",
			territory.NZ:    "Керла Зеланди",
			territory.OM:    "Ӏоман",
			territory.PA:    "Панама",
			territory.PE:    "Перу",
			territory.PF:    "Французийн Полинези",
			territory.PG:    "Папуа — Керла Гвиней",
			territory.PH:    "Филиппинаш",
			territory.PK:    "Пакистан",
			territory.PL:    "Польша",
			territory.PM:    "Сен-Пьер а, Микелон а",
			territory.PN:    "Питкэрн гӀайренаш",
			territory.PR:    "Пуэрто-Рико",
			territory.PS:    "ПалестӀинан латтанаш",
			territory.PT:    "Португали",
			territory.PW:    "Палау",
			territory.PY:    "Парагвай",
			territory.QA:    "Катар",
			territory.QO:    "Арахьара Океани",
			territory.RE:    "Реюньон",
			territory.RO:    "Румыни",
			territory.RS:    "Серби",
			territory.RU:    "Росси",
			territory.RW:    "Руанда",
			territory.SA:    "СаӀудийн Ӏаьрбийчоь",
			territory.SB:    "Соломонан гӀайренаш",
			territory.SC:    "Сейшелан гӀайренаш",
			territory.SD:    "Судан",
			territory.SE:    "Швеци",
			territory.SG:    "Сингапур",
			territory.SH:    "Сийлахьчу Еленин гӀайре",
			territory.SI:    "Словени",
			territory.SJ:    "Шпицберген а, Ян-Майен а",
			territory.SK:    "Словаки",
			territory.SL:    "Сьерра- Леоне",
			territory.SM:    "Сан-Марино",
			territory.SN:    "Сенегал",
			territory.SO:    "Сомали",
			territory.SR:    "Суринам",
			territory.SS:    "Къилба Судан",
			territory.ST:    "Сан-Томе а, Принсипи а",
			territory.SV:    "Сальвадор",
			territory.SX:    "Синт-Мартен",
			territory.SY:    "Шема",
			territory.SZ:    "Свазиленд",
			territory.TA:    "Тристан-да- Кунья",
			territory.TC:    "Тёркс а, Кайкос а гӀайренаш",
			territory.TD:    "Чад",
			territory.TF:    "Французийн къилба латтанаш",
			territory.TG:    "Того",
			territory.TH:    "Таиланд",
			territory.TJ:    "Таджикистан",
			territory.TK:    "Токелау",
			territory.TL:    "Малхбален Тимор",
			territory.TM:    "Туркмени",
			territory.TN:    "Тунис",
			territory.TO:    "Тонга",
			territory.TR:    "Туркойчоь",
			territory.TT:    "Тринидад а, Тобаго а",
			territory.TV:    "Тувалу",
			territory.TW:    "Тайвань",
			territory.TZ:    "Танзани",
			territory.UA:    "Украина",
			territory.UG:    "Уганда",
			territory.UM:    "АЦШн арахьара кегийн гӀайренаш",
			territory.UN:    "Вовшахкхетта Къаьмнийн Организаци",
			territory.US:    "Цхьанатоьхна Штаташ",
			territory.UY:    "Уругвай",
			territory.UZ:    "Узбекистан",
			territory.VA:    "Ватикан",
			territory.VC:    "Сент-Винсент а, Гренадинаш а",
			territory.VE:    "Венесуэла",
			territory.VG:    "Виргинийн гӀайренаш (Британи)",
			territory.VI:    "Виргинийн гӀайренаш (АЦШ)",
			territory.VN:    "Вьетнам",
			territory.VU:    "Вануату",
			territory.WF:    "Уоллис а, Футуна а",
			territory.WS:    "Самоа",
			territory.XK:    "Косово",
			territory.YE:    "Йемен",
			territory.YT:    "Майотта",
			territory.ZA:    "Къилба-Африкин Республика",
			territory.ZM:    "Замби",
			territory.ZW:    "Зимбабве",
			territory.ZZ:    "Йоьвзуш йоцу регион",
		},
	}
}
