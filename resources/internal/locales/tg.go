// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_tg() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "tg",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Янв", Feb: "Фев", Mar: "Мар", Apr: "Апр", May: "Май", Jun: "Июн", Jul: "Июл", Aug: "Авг", Sep: "Сен", Oct: "Окт", Nov: "Ноя", Dec: "Дек"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Январ", Feb: "Феврал", Mar: "Март", Apr: "Апрел", May: "Май", Jun: "Июн", Jul: "Июл", Aug: "Август", Sep: "Сентябр", Oct: "Октябр", Nov: "Ноябр", Dec: "Декабр"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Яшб", Mon: "Дшб", Tue: "Сшб", Wed: "Чшб", Thu: "Пшб", Fri: "Ҷмъ", Sat: "Шнб"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Я", Mon: "Д", Tue: "С", Wed: "Ч", Thu: "П", Fri: "Ҷ", Sat: "Ш"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Яшб", Mon: "Дшб", Tue: "Сшб", Wed: "Чшб", Thu: "Пшб", Fri: "Ҷмъ", Sat: "Шнб"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Якшанбе", Mon: "Душанбе", Tue: "Сешанбе", Wed: "Чоршанбе", Thu: "Панҷшанбе", Fri: "Ҷумъа", Sat: "Шанбе"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0%"},
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
				currency.BRL: cldr.Currency{DisplayName: "Реали бразилиягӣ", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Иенаи хитоӣ", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "Евро", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Фунт стерлинги британӣ", Symbol: "£"},
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
				currency.INR: cldr.Currency{DisplayName: "Рупияи ҳиндустонӣ", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Иенаи японӣ", Symbol: "JP¥"},
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
				currency.RUB: cldr.Currency{DisplayName: "Рубли русӣ", Symbol: "RUB"},
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
				currency.TJS: cldr.Currency{DisplayName: "Сомонӣ", Symbol: "сом."},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "Доллари ИМА", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Асъори номаълум", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AF:      "африкаанс",
			language.AM:      "амҳарӣ",
			language.AR:      "арабӣ",
			language.AS:      "ассомӣ",
			language.AZ:      "озарбойҷонӣ",
			language.BA:      "бошқирдӣ",
			language.BAN:     "балинӣ",
			language.BE:      "белорусӣ",
			language.BEM:     "бемба",
			language.BG:      "булғорӣ",
			language.BN:      "бинғолӣ",
			language.BO:      "тибетӣ",
			language.BR:      "бретонӣ",
			language.BS:      "босниягӣ",
			language.CA:      "каталонӣ",
			language.CEB:     "себуано",
			language.CHM:     "марӣ",
			language.CHR:     "черокӣ",
			language.CKB:     "курдии марказӣ",
			language.CO:      "корсиканӣ",
			language.CS:      "чехӣ",
			language.CY:      "валлӣ",
			language.DA:      "даниягӣ",
			language.DE:      "немисӣ",
			language.DE_AT:   "немисии австриягӣ",
			language.DE_CH:   "немисии швейсарии болоӣ",
			language.DSB:     "сербии поёнӣ",
			language.DV:      "дивеҳӣ",
			language.DZ:      "дзонгха",
			language.EL:      "юнонӣ",
			language.EN:      "англисӣ",
			language.EN_AU:   "англисии австралиягӣ",
			language.EN_CA:   "англисии канадагӣ",
			language.EN_GB:   "англисӣ (ШМ)",
			language.EN_US:   "англисии америкоӣ",
			language.EO:      "эсперанто",
			language.ES:      "испанӣ",
			language.ES_419:  "испании америкоии лотинӣ",
			language.ES_ES:   "испании аврупоӣ",
			language.ES_MX:   "испании мексикоӣ",
			language.ET:      "эстонӣ",
			language.EU:      "баскӣ",
			language.FA:      "форсӣ",
			language.FF:      "фулаҳ",
			language.FI:      "финӣ",
			language.FIL:     "филиппинӣ",
			language.FO:      "фарерӣ",
			language.FR:      "франсузӣ",
			language.FR_CA:   "франсузии канадагӣ",
			language.FR_CH:   "франсузии швейсарӣ",
			language.FY:      "фризии ғарбӣ",
			language.GA:      "ирландӣ",
			language.GD:      "шотландии гэлӣ",
			language.GL:      "галисиягӣ",
			language.GN:      "гуаранӣ",
			language.GU:      "гуҷаротӣ",
			language.HA:      "ҳауса",
			language.HAW:     "ҳавайӣ",
			language.HE:      "ибронӣ",
			language.HI:      "ҳиндӣ",
			language.HIL:     "ҳилигайнон",
			language.HR:      "хорватӣ",
			language.HSB:     "сербии болоӣ",
			language.HT:      "гаитии креолӣ",
			language.HU:      "маҷорӣ",
			language.HY:      "арманӣ",
			language.HZ:      "ҳереро",
			language.IA:      "Байни забонӣ",
			language.IBB:     "ибибио",
			language.ID:      "индонезӣ",
			language.IG:      "игбо",
			language.IS:      "исландӣ",
			language.IT:      "италиявӣ",
			language.IU:      "инуктитутӣ",
			language.JA:      "японӣ",
			language.JV:      "Ҷаванизӣ",
			language.KA:      "гурҷӣ",
			language.KK:      "қазоқӣ",
			language.KM:      "кхмерӣ",
			language.KN:      "каннада",
			language.KO:      "кореягӣ",
			language.KOK:     "конканӣ",
			language.KR:      "канурӣ",
			language.KRU:     "курукс",
			language.KS:      "кашмирӣ",
			language.KU:      "курдӣ",
			language.KY:      "қирғизӣ",
			language.LA:      "лотинӣ",
			language.LB:      "люксембургӣ",
			language.LO:      "лаосӣ",
			language.LT:      "литвонӣ",
			language.LV:      "латишӣ",
			language.MEN:     "менде",
			language.MG:      "малагасӣ",
			language.MI:      "маорӣ",
			language.MK:      "мақдунӣ",
			language.ML:      "малаяламӣ",
			language.MN:      "муғулӣ",
			language.MNI:     "манипурӣ",
			language.MOH:     "моҳок",
			language.MR:      "маратҳӣ",
			language.MS:      "малайӣ",
			language.MT:      "малтӣ",
			language.MY:      "бирманӣ",
			language.NE:      "непалӣ",
			language.NIU:     "ниуэӣ",
			language.NL:      "голландӣ",
			language.NO:      "норвегӣ",
			language.NY:      "нянҷа",
			language.OC:      "окситанӣ",
			language.OM:      "оромо",
			language.OR:      "одия",
			language.PA:      "панҷобӣ",
			language.PAP:     "папиаменто",
			language.PL:      "лаҳистонӣ",
			language.PS:      "пушту",
			language.PT:      "португалӣ",
			language.PT_BR:   "португалии бразилиягӣ",
			language.PT_PT:   "португалии аврупоӣ",
			language.QU:      "кечуа",
			language.QUC:     "киче",
			language.RM:      "ретороманӣ",
			language.RO:      "руминӣ",
			language.RU:      "русӣ",
			language.RW:      "киняруанда",
			language.SA:      "санскрит",
			language.SAH:     "саха",
			language.SAT:     "санталӣ",
			language.SD:      "синдӣ",
			language.SE:      "самии шимолӣ",
			language.SI:      "сингалӣ",
			language.SK:      "словакӣ",
			language.SL:      "словенӣ",
			language.SMA:     "самии ҷанубӣ",
			language.SMJ:     "луле самӣ",
			language.SMN:     "инари самӣ",
			language.SMS:     "сколти самӣ",
			language.SO:      "сомалӣ",
			language.SQ:      "албанӣ",
			language.SR:      "сербӣ",
			language.SV:      "шведӣ",
			language.SYR:     "суриёнӣ",
			language.TA:      "тамилӣ",
			language.TE:      "телугу",
			language.TG:      "тоҷикӣ",
			language.TH:      "тайӣ",
			language.TI:      "тигриня",
			language.TK:      "туркманӣ",
			language.TO:      "тонганӣ",
			language.TR:      "туркӣ",
			language.TT:      "тоторӣ",
			language.TZM:     "тамазайти атласи марказӣ",
			language.UG:      "ӯйғурӣ",
			language.UK:      "украинӣ",
			language.UND:     "забони номаълум",
			language.UR:      "урду",
			language.UZ:      "ӯзбекӣ",
			language.VE:      "венда",
			language.VI:      "ветнамӣ",
			language.WO:      "волоф",
			language.YI:      "идиш",
			language.YO:      "йоруба",
			language.ZH:      "хитоӣ, мандаринӣ",
			language.ZH_HANS: "хитоии мандаринии осонфаҳм",
			language.ZH_HANT: "хитоии мандаринии анъанавӣ",
			language.ZU:      "Зулу",
		},
		Territories: cldr.Territories{
			territory.AC: "Асунсон",
			territory.AD: "Андорра",
			territory.AE: "Аморатҳои Муттаҳидаи Араб",
			territory.AF: "Афғонистон",
			territory.AG: "Антигуа ва Барбуда",
			territory.AI: "Ангилия",
			territory.AL: "Албания",
			territory.AM: "Арманистон",
			territory.AO: "Ангола",
			territory.AQ: "Антарктида",
			territory.AR: "Аргентина",
			territory.AS: "Самоаи Америка",
			territory.AT: "Австрия",
			territory.AU: "Австралия",
			territory.AW: "Аруба",
			territory.AX: "Ҷазираҳои Аланд",
			territory.AZ: "Озарбойҷон",
			territory.BA: "Босния ва Ҳерсеговина",
			territory.BB: "Барбадос",
			territory.BD: "Бангладеш",
			territory.BE: "Белгия",
			territory.BF: "Буркина-Фасо",
			territory.BG: "Булғория",
			territory.BH: "Баҳрайн",
			territory.BI: "Бурунди",
			territory.BJ: "Бенин",
			territory.BL: "Сент-Бартелми",
			territory.BM: "Бермуда",
			territory.BN: "Бруней",
			territory.BO: "Боливия",
			territory.BR: "Бразилия",
			territory.BS: "Багам",
			territory.BT: "Бутон",
			territory.BV: "Ҷазираи Буве",
			territory.BW: "Ботсвана",
			territory.BY: "Белорус",
			territory.BZ: "Белиз",
			territory.CA: "Канада",
			territory.CC: "Ҷазираҳои Кокос (Килинг)",
			territory.CF: "Ҷумҳурии Африқои Марказӣ",
			territory.CH: "Швейтсария",
			territory.CI: "Кот-д’Ивуар",
			territory.CK: "Ҷазираҳои Кук",
			territory.CL: "Чили",
			territory.CM: "Камерун",
			territory.CN: "Хитой",
			territory.CO: "Колумбия",
			territory.CR: "Коста-Рика",
			territory.CU: "Куба",
			territory.CV: "Кабо-Верде",
			territory.CW: "Кюрасао",
			territory.CX: "Ҷазираи Крисмас",
			territory.CY: "Кипр",
			territory.CZ: "Ҷумҳурии Чех",
			territory.DE: "Германия",
			territory.DJ: "Ҷибути",
			territory.DK: "Дания",
			territory.DM: "Доминика",
			territory.DO: "Ҷумҳурии Доминикан",
			territory.DZ: "Алҷазоир",
			territory.EC: "Эквадор",
			territory.EE: "Эстония",
			territory.EG: "Миср",
			territory.ER: "Эритрея",
			territory.ES: "Испания",
			territory.ET: "Эфиопия",
			territory.FI: "Финляндия",
			territory.FJ: "Фиҷи",
			territory.FK: "Ҷазираҳои Фолкленд",
			territory.FM: "Штатҳои Федеративии Микронезия",
			territory.FO: "Ҷазираҳои Фарер",
			territory.FR: "Фаронса",
			territory.GA: "Габон",
			territory.GB: "Шоҳигарии Муттаҳида",
			territory.GD: "Гренада",
			territory.GE: "Гурҷистон",
			territory.GF: "Гвианаи Фаронса",
			territory.GG: "Гернси",
			territory.GH: "Гана",
			territory.GI: "Гибралтар",
			territory.GL: "Гренландия",
			territory.GM: "Гамбия",
			territory.GN: "Гвинея",
			territory.GP: "Гваделупа",
			territory.GQ: "Гвинеяи Экваторӣ",
			territory.GR: "Юнон",
			territory.GS: "Ҷорҷияи Ҷанубӣ ва Ҷазираҳои Сандвич",
			territory.GT: "Гватемала",
			territory.GU: "Гуам",
			territory.GW: "Гвинея-Бисау",
			territory.GY: "Гайана",
			territory.HK: "Ҳонконг (МММ)",
			territory.HM: "Ҷазираи Ҳерд ва Ҷазираҳои Макдоналд",
			territory.HN: "Гондурас",
			territory.HR: "Хорватия",
			territory.HT: "Гаити",
			territory.HU: "Маҷористон",
			territory.ID: "Индонезия",
			territory.IE: "Ирландия",
			territory.IL: "Исроил",
			territory.IM: "Ҷазираи Мэн",
			territory.IN: "Ҳиндустон",
			territory.IO: "Қаламрави Британия дар уқёнуси Ҳинд",
			territory.IQ: "Ироқ",
			territory.IR: "Эрон",
			territory.IS: "Исландия",
			territory.IT: "Италия",
			territory.JE: "Ҷерси",
			territory.JM: "Ямайка",
			territory.JO: "Урдун",
			territory.JP: "Япония",
			territory.KE: "Кения",
			territory.KG: "Қирғизистон",
			territory.KH: "Камбоҷа",
			territory.KI: "Кирибати",
			territory.KM: "Комор",
			territory.KN: "Сент-Китс ва Невис",
			territory.KP: "Кореяи Шимолӣ",
			territory.KW: "Қувайт",
			territory.KY: "Ҷазираҳои Кайман",
			territory.KZ: "Қазоқистон",
			territory.LA: "Лаос",
			territory.LB: "Лубнон",
			territory.LC: "Сент-Люсия",
			territory.LI: "Лихтенштейн",
			territory.LK: "Шри-Ланка",
			territory.LR: "Либерия",
			territory.LS: "Лесото",
			territory.LT: "Литва",
			territory.LU: "Люксембург",
			territory.LV: "Латвия",
			territory.LY: "Либия",
			territory.MA: "Марокаш",
			territory.MC: "Монако",
			territory.MD: "Молдова",
			territory.ME: "Черногория",
			territory.MF: "Ҷазираи Сент-Мартин",
			territory.MG: "Мадагаскар",
			territory.MH: "Ҷазираҳои Маршалл",
			territory.MK: "Македонияи Шимолӣ",
			territory.ML: "Мали",
			territory.MM: "Мянма",
			territory.MN: "Муғулистон",
			territory.MO: "Макао (МММ)",
			territory.MP: "Ҷазираҳои Марианаи Шимолӣ",
			territory.MQ: "Мартиника",
			territory.MR: "Мавритания",
			territory.MS: "Монтсеррат",
			territory.MT: "Малта",
			territory.MU: "Маврикий",
			territory.MV: "Малдив",
			territory.MW: "Малави",
			territory.MX: "Мексика",
			territory.MY: "Малайзия",
			territory.MZ: "Мозамбик",
			territory.NA: "Намибия",
			territory.NC: "Каледонияи Нав",
			territory.NE: "Нигер",
			territory.NF: "Ҷазираи Норфолк",
			territory.NG: "Нигерия",
			territory.NI: "Никарагуа",
			territory.NL: "Нидерландия",
			territory.NO: "Норвегия",
			territory.NP: "Непал",
			territory.NR: "Науру",
			territory.NU: "Ниуэ",
			territory.NZ: "Зеландияи Нав",
			territory.OM: "Умон",
			territory.PA: "Панама",
			territory.PE: "Перу",
			territory.PF: "Полинезияи Фаронса",
			territory.PG: "Папуа Гвинеяи Нав",
			territory.PH: "Филиппин",
			territory.PK: "Покистон",
			territory.PL: "Лаҳистон",
			territory.PM: "Сент-Пер ва Микелон",
			territory.PN: "Ҷазираҳои Питкейрн",
			territory.PR: "Пуэрто-Рико",
			territory.PT: "Португалия",
			territory.PW: "Палау",
			territory.PY: "Парагвай",
			territory.QA: "Қатар",
			territory.RE: "Реюнион",
			territory.RO: "Руминия",
			territory.RS: "Сербия",
			territory.RU: "Русия",
			territory.RW: "Руанда",
			territory.SA: "Арабистони Саудӣ",
			territory.SB: "Ҷазираҳои Соломон",
			territory.SC: "Сейшел",
			territory.SD: "Судон",
			territory.SE: "Шветсия",
			territory.SG: "Сингапур",
			territory.SH: "Сент Елена",
			territory.SI: "Словения",
			territory.SJ: "Шпитсберген ва Ян Майен",
			territory.SK: "Словакия",
			territory.SL: "Сиерра-Леоне",
			territory.SM: "Сан-Марино",
			territory.SN: "Сенегал",
			territory.SO: "Сомалӣ",
			territory.SR: "Суринам",
			territory.SS: "Судони Ҷанубӣ",
			territory.ST: "Сан Томе ва Принсипи",
			territory.SV: "Эл-Салвадор",
			territory.SX: "Синт-Маартен",
			territory.SY: "Сурия",
			territory.SZ: "Свазиленд",
			territory.TA: "Тристан-да-Куня",
			territory.TC: "Ҷазираҳои Теркс ва Кайкос",
			territory.TD: "Чад",
			territory.TF: "Минтақаҳои Ҷанубии Фаронса",
			territory.TG: "Того",
			territory.TH: "Таиланд",
			territory.TJ: "Тоҷикистон",
			territory.TK: "Токелау",
			territory.TL: "Тимор-Лесте",
			territory.TM: "Туркманистон",
			territory.TN: "Тунис",
			territory.TO: "Тонга",
			territory.TR: "Туркия",
			territory.TT: "Тринидад ва Тобаго",
			territory.TV: "Тувалу",
			territory.TW: "Тайван",
			territory.TZ: "Танзания",
			territory.UA: "Украина",
			territory.UG: "Уганда",
			territory.UM: "Ҷазираҳои Хурди Дурдасти ИМА",
			territory.US: "Иёлоти Муттаҳида",
			territory.UY: "Уругвай",
			territory.UZ: "Ӯзбекистон",
			territory.VA: "Шаҳри Вотикон",
			territory.VC: "Сент-Винсент ва Гренадина",
			territory.VE: "Венесуэла",
			territory.VG: "Ҷазираҳои Виргини Британия",
			territory.VI: "Ҷазираҳои Виргини ИМА",
			territory.VN: "Ветнам",
			territory.VU: "Вануату",
			territory.WF: "Уоллис ва Футуна",
			territory.WS: "Самоа",
			territory.XK: "Косово",
			territory.YE: "Яман",
			territory.YT: "Майотта",
			territory.ZA: "Африкаи Ҷанубӣ",
			territory.ZM: "Замбия",
			territory.ZW: "Зимбабве",
			territory.ZZ: "Минтақаи номаълум",
		},
	}
}