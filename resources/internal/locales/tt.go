// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_tt() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "tt",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "d MMMM, y 'ел', EEEE", Long: "d MMMM, y 'ел'", Medium: "d MMM, y 'ел'", Short: "dd.MM.y"}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "гыйн.", Feb: "фев.", Mar: "мар.", Apr: "апр.", May: "май", Jun: "июнь", Jul: "июль", Aug: "авг.", Sep: "сент.", Oct: "окт.", Nov: "нояб.", Dec: "дек."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "гыйнвар", Feb: "февраль", Mar: "март", Apr: "апрель", May: "май", Jun: "июнь", Jul: "июль", Aug: "август", Sep: "сентябрь", Oct: "октябрь", Nov: "ноябрь", Dec: "декабрь"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "якш.", Mon: "дүш.", Tue: "сиш.", Wed: "чәр.", Thu: "пәнҗ.", Fri: "җом.", Sat: "шим."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Я", Mon: "Д", Tue: "С", Wed: "Ч", Thu: "П", Fri: "Җ", Sat: "Ш"}, Short: cldr.CalendarDayFormatNameValue{Sun: "якш.", Mon: "дүш.", Tue: "сиш.", Wed: "чәр.", Thu: "пәнҗ.", Fri: "җом.", Sat: "шим."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "якшәмбе", Mon: "дүшәмбе", Tue: "сишәмбе", Wed: "чәршәмбе", Thu: "пәнҗешәмбе", Fri: "җомга", Sat: "шимбә"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00¤", CurrencyAccounting: "", Percent: "#,##0\u00a0%"},
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
				currency.BRL: cldr.Currency{DisplayName: "Бразилия реалы", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "кытай юане", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "евро", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "фунт стерлинг", Symbol: "£"},
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
				currency.INR: cldr.Currency{DisplayName: "Индия рупиясе", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "япон иенасы", Symbol: "JP¥"},
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
				currency.RUB: cldr.Currency{DisplayName: "Россия сумы", Symbol: "₽"},
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
				currency.USD: cldr.Currency{DisplayName: "АКШ доллары", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "билгесез валюта", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AF:      "африкаанс",
			language.AM:      "амхар",
			language.AR:      "гарәп",
			language.ARN:     "мапуче",
			language.AS:      "ассам",
			language.AZ:      "әзәрбайҗан",
			language.BA:      "башкорт",
			language.BAN:     "бали",
			language.BE:      "белорус",
			language.BEM:     "бемба",
			language.BG:      "болгар",
			language.BN:      "бенгали",
			language.BO:      "тибет",
			language.BR:      "бретон",
			language.BS:      "босния",
			language.CA:      "каталан",
			language.CEB:     "себуано",
			language.CHM:     "мари",
			language.CHR:     "чероки",
			language.CKB:     "үзәк көрд",
			language.CO:      "корсика",
			language.CS:      "чех",
			language.CY:      "уэльс",
			language.DA:      "дания",
			language.DE:      "алман",
			language.DE_CH:   "югары алман (Швейцария)",
			language.DSB:     "түбән сорб",
			language.DV:      "мальдив",
			language.DZ:      "дзонг-кха",
			language.EL:      "грек",
			language.EN:      "инглиз",
			language.EN_GB:   "Британия инглизчәсе",
			language.EN_US:   "Америка инглизчәсе",
			language.EO:      "эсперанто",
			language.ES:      "испан",
			language.ES_419:  "испан (Латин Америкасы)",
			language.ES_ES:   "испан (Европа)",
			language.ET:      "эстон",
			language.EU:      "баск",
			language.FA:      "фарсы",
			language.FF:      "фула",
			language.FI:      "фин",
			language.FIL:     "филиппин",
			language.FO:      "фарер",
			language.FR:      "француз",
			language.GA:      "ирланд",
			language.GD:      "шотланд гэль",
			language.GL:      "галисия",
			language.GN:      "гуарани",
			language.GU:      "гуҗарати",
			language.HA:      "хауса",
			language.HAW:     "гавайи",
			language.HE:      "яһүд",
			language.HI:      "һинд",
			language.HIL:     "хилигайнон",
			language.HR:      "хорват",
			language.HSB:     "югары сорб",
			language.HT:      "гаити креол",
			language.HU:      "венгр",
			language.HY:      "әрмән",
			language.HZ:      "гереро",
			language.IBB:     "ибибио",
			language.ID:      "индонезия",
			language.IG:      "игбо",
			language.IS:      "исланд",
			language.IT:      "итальян",
			language.IU:      "инуктикут",
			language.JA:      "япон",
			language.KA:      "грузин",
			language.KK:      "казакъ",
			language.KM:      "кхмер",
			language.KN:      "каннада",
			language.KO:      "корея",
			language.KOK:     "конкани",
			language.KR:      "канури",
			language.KRU:     "курух",
			language.KS:      "кашмири",
			language.KU:      "көрд",
			language.KY:      "кыргыз",
			language.LA:      "латин",
			language.LB:      "люксембург",
			language.LO:      "лаос",
			language.LT:      "литва",
			language.LV:      "латыш",
			language.MEN:     "менде",
			language.MG:      "малагаси",
			language.MI:      "маори",
			language.MK:      "македон",
			language.ML:      "малаялам",
			language.MN:      "монгол",
			language.MNI:     "манипури",
			language.MOH:     "могаук",
			language.MR:      "маратхи",
			language.MS:      "малай",
			language.MT:      "мальта",
			language.MY:      "бирма",
			language.NE:      "непали",
			language.NIU:     "ниуэ",
			language.NL:      "голланд",
			language.NY:      "ньянҗа",
			language.OC:      "окситан",
			language.OM:      "оромо",
			language.OR:      "ория",
			language.PA:      "пәнҗаби",
			language.PAP:     "папьяменто",
			language.PL:      "поляк",
			language.PS:      "пушту",
			language.PT:      "португал",
			language.PT_PT:   "португал (Европа)",
			language.QU:      "кечуа",
			language.QUC:     "киче",
			language.RM:      "ретороман",
			language.RO:      "румын",
			language.RU:      "рус",
			language.RW:      "руанда",
			language.SA:      "санскрит",
			language.SAH:     "саха",
			language.SAT:     "сантали",
			language.SD:      "синдһи",
			language.SE:      "төньяк саам",
			language.SI:      "сингал",
			language.SK:      "словак",
			language.SL:      "словен",
			language.SMA:     "көньяк саам",
			language.SMJ:     "луле-саам",
			language.SMN:     "инари-саам",
			language.SMS:     "колтта-саам",
			language.SO:      "сомали",
			language.SQ:      "албан",
			language.SR:      "серб",
			language.SV:      "швед",
			language.SYR:     "сүрия",
			language.TA:      "тамил",
			language.TE:      "телугу",
			language.TG:      "таҗик",
			language.TH:      "тай",
			language.TI:      "тигринья",
			language.TK:      "төрекмән",
			language.TO:      "тонга",
			language.TR:      "төрек",
			language.TT:      "татар",
			language.TZM:     "үзәк атлас тамазигт",
			language.UG:      "уйгыр",
			language.UK:      "украин",
			language.UND:     "билгесез тел",
			language.UR:      "урду",
			language.UZ:      "үзбәк",
			language.VE:      "венда",
			language.VI:      "вьетнам",
			language.WO:      "волоф",
			language.YI:      "идиш",
			language.YO:      "йоруба",
			language.ZH:      "мандарин кытайчасы",
			language.ZH_HANS: "гадиләштерелгән мандарин кытайчасы",
			language.ZH_HANT: "традицион мандарин кытайчасы",
		},
		Territories: cldr.Territories{
			territory.AD: "Андорра",
			territory.AE: "Берләшкән Гарәп Әмирлекләре",
			territory.AF: "Әфганстан",
			territory.AG: "Антигуа һәм Барбуда",
			territory.AI: "Ангилья",
			territory.AL: "Албания",
			territory.AM: "Әрмәнстан",
			territory.AO: "Ангола",
			territory.AQ: "Антарктика",
			territory.AR: "Аргентина",
			territory.AS: "Америка Самоасы",
			territory.AT: "Австрия",
			territory.AU: "Австралия",
			territory.AW: "Аруба",
			territory.AX: "Аланд утраулары",
			territory.AZ: "Әзәрбайҗан",
			territory.BA: "Босния һәм Герцеговина",
			territory.BB: "Барбадос",
			territory.BD: "Бангладеш",
			territory.BE: "Бельгия",
			territory.BF: "Буркина-Фасо",
			territory.BG: "Болгария",
			territory.BH: "Бәхрәйн",
			territory.BI: "Бурунди",
			territory.BJ: "Бенин",
			territory.BL: "Сен-Бартельми",
			territory.BM: "Бермуд утраулары",
			territory.BN: "Бруней",
			territory.BO: "Боливия",
			territory.BR: "Бразилия",
			territory.BS: "Багам утраулары",
			territory.BT: "Бутан",
			territory.BV: "Буве утравы",
			territory.BW: "Ботсвана",
			territory.BY: "Беларусь",
			territory.BZ: "Белиз",
			territory.CA: "Канада",
			territory.CC: "Кокос (Килинг) утраулары",
			territory.CF: "Үзәк Африка Республикасы",
			territory.CH: "Швейцария",
			territory.CI: "Кот-д’Ивуар",
			territory.CK: "Кук утраулары",
			territory.CL: "Чили",
			territory.CM: "Камерун",
			territory.CN: "Кытай",
			territory.CO: "Колумбия",
			territory.CR: "Коста-Рика",
			territory.CU: "Куба",
			territory.CV: "Кабо-Верде",
			territory.CW: "Кюрасао",
			territory.CX: "Раштуа утравы",
			territory.CY: "Кипр",
			territory.CZ: "Чехия Республикасы",
			territory.DE: "Германия",
			territory.DJ: "Җибүти",
			territory.DK: "Дания",
			territory.DM: "Доминика",
			territory.DO: "Доминикана Республикасы",
			territory.DZ: "Алжир",
			territory.EC: "Эквадор",
			territory.EE: "Эстония",
			territory.EG: "Мисыр",
			territory.ER: "Эритрея",
			territory.ES: "Испания",
			territory.ET: "Эфиопия",
			territory.FI: "Финляндия",
			territory.FJ: "Фиджи",
			territory.FK: "Фолкленд утраулары",
			territory.FM: "Микронезия",
			territory.FO: "Фарер утраулары",
			territory.FR: "Франция",
			territory.GA: "Габон",
			territory.GB: "Берләшкән Корольлек",
			territory.GD: "Гренада",
			territory.GE: "Грузия",
			territory.GF: "Француз Гвианасы",
			territory.GG: "Гернси",
			territory.GH: "Гана",
			territory.GI: "Гибралтар",
			territory.GL: "Гренландия",
			territory.GM: "Гамбия",
			territory.GN: "Гвинея",
			territory.GP: "Гваделупа",
			territory.GQ: "Экваториаль Гвинея",
			territory.GR: "Греция",
			territory.GS: "Көньяк Георгия һәм Көньяк Сандвич утраулары",
			territory.GT: "Гватемала",
			territory.GU: "Гуам",
			territory.GW: "Гвинея-Бисау",
			territory.GY: "Гайана",
			territory.HK: "Гонконг Махсус Идарәле Төбәге",
			territory.HM: "Херд утравы һәм Макдональд утраулары",
			territory.HN: "Гондурас",
			territory.HR: "Хорватия",
			territory.HT: "Гаити",
			territory.HU: "Венгрия",
			territory.ID: "Индонезия",
			territory.IE: "Ирландия",
			territory.IL: "Израиль",
			territory.IM: "Мэн утравы",
			territory.IN: "Индия",
			territory.IO: "Британиянең Һинд Океанындагы Территориясе",
			territory.IQ: "Гыйрак",
			territory.IR: "Иран",
			territory.IS: "Исландия",
			territory.IT: "Италия",
			territory.JE: "Джерси",
			territory.JM: "Ямайка",
			territory.JO: "Иордания",
			territory.JP: "Япония",
			territory.KE: "Кения",
			territory.KG: "Кыргызстан",
			territory.KH: "Камбоджа",
			territory.KI: "Кирибати",
			territory.KM: "Комор утраулары",
			territory.KN: "Сент-Китс һәм Невис",
			territory.KP: "Төньяк Корея",
			territory.KW: "Күвәйт",
			territory.KY: "Кайман утраулары",
			territory.KZ: "Казахстан",
			territory.LA: "Лаос",
			territory.LB: "Ливан",
			territory.LC: "Сент-Люсия",
			territory.LI: "Лихтенштейн",
			territory.LK: "Шри-Ланка",
			territory.LR: "Либерия",
			territory.LS: "Лесото",
			territory.LT: "Литва",
			territory.LU: "Люксембург",
			territory.LV: "Латвия",
			territory.LY: "Ливия",
			territory.MA: "Марокко",
			territory.MC: "Монако",
			territory.MD: "Молдова",
			territory.ME: "Черногория",
			territory.MF: "Сент-Мартин",
			territory.MG: "Мадагаскар",
			territory.MH: "Маршалл утраулары",
			territory.MK: "Төньяк Македония",
			territory.ML: "Мали",
			territory.MN: "Монголия",
			territory.MO: "Макао Махсус Идарәле Төбәге",
			territory.MP: "Төньяк Мариана утраулары",
			territory.MQ: "Мартиника",
			territory.MR: "Мавритания",
			territory.MS: "Монтсеррат",
			territory.MT: "Мальта",
			territory.MU: "Маврикий",
			territory.MV: "Мальдив утраулары",
			territory.MW: "Малави",
			territory.MX: "Мексика",
			territory.MY: "Малайзия",
			territory.MZ: "Мозамбик",
			territory.NA: "Намибия",
			territory.NC: "Яңа Каледония",
			territory.NE: "Нигер",
			territory.NF: "Норфолк утравы",
			territory.NG: "Нигерия",
			territory.NI: "Никарагуа",
			territory.NL: "Нидерланд",
			territory.NO: "Норвегия",
			territory.NP: "Непал",
			territory.NR: "Науру",
			territory.NU: "Ниуэ",
			territory.NZ: "Яңа Зеландия",
			territory.OM: "Оман",
			territory.PA: "Панама",
			territory.PE: "Перу",
			territory.PF: "Француз Полинезиясе",
			territory.PG: "Папуа - Яңа Гвинея",
			territory.PH: "Филиппин",
			territory.PK: "Пакистан",
			territory.PL: "Польша",
			territory.PM: "Сен-Пьер һәм Микелон",
			territory.PN: "Питкэрн утраулары",
			territory.PR: "Пуэрто-Рико",
			territory.PT: "Португалия",
			territory.PW: "Палау",
			territory.PY: "Парагвай",
			territory.QA: "Катар",
			territory.RE: "Реюньон",
			territory.RO: "Румыния",
			territory.RS: "Сербия",
			territory.RU: "Россия",
			territory.RW: "Руанда",
			territory.SA: "Согуд Гарәбстаны",
			territory.SB: "Сөләйман утраулары",
			territory.SC: "Сейшел утраулары",
			territory.SD: "Судан",
			territory.SE: "Швеция",
			territory.SG: "Сингапур",
			territory.SI: "Словения",
			territory.SJ: "Шпицберген һәм Ян-Майен",
			territory.SK: "Словакия",
			territory.SL: "Сьерра-Леоне",
			territory.SM: "Сан-Марино",
			territory.SN: "Сенегал",
			territory.SO: "Сомали",
			territory.SR: "Суринам",
			territory.SS: "Көньяк Судан",
			territory.ST: "Сан-Томе һәм Принсипи",
			territory.SV: "Сальвадор",
			territory.SX: "Синт-Мартен",
			territory.SY: "Сүрия",
			territory.SZ: "Свазиленд",
			territory.TC: "Теркс һәм Кайкос утраулары",
			territory.TD: "Чад",
			territory.TF: "Франциянең Көньяк Территорияләре",
			territory.TG: "Того",
			territory.TH: "Тайланд",
			territory.TJ: "Таҗикстан",
			territory.TK: "Токелау",
			territory.TL: "Тимор-Лесте",
			territory.TM: "Төркмәнстан",
			territory.TN: "Тунис",
			territory.TO: "Тонга",
			territory.TR: "Төркия",
			territory.TT: "Тринидад һәм Тобаго",
			territory.TV: "Тувалу",
			territory.TW: "Тайвань",
			territory.TZ: "Танзания",
			territory.UA: "Украина",
			territory.UG: "Уганда",
			territory.UM: "АКШ Кече Читтәге утраулары",
			territory.US: "АКШ",
			territory.UY: "Уругвай",
			territory.UZ: "Үзбәкстан",
			territory.VC: "Сент-Винсент һәм Гренадин",
			territory.VE: "Венесуэла",
			territory.VG: "Британия Виргин утраулары",
			territory.VI: "АКШ Виргин утраулары",
			territory.VN: "Вьетнам",
			territory.VU: "Вануату",
			territory.WF: "Уоллис һәм Футуна",
			territory.WS: "Самоа",
			territory.XK: "Косово",
			territory.YE: "Йәмән",
			territory.YT: "Майотта",
			territory.ZA: "Көньяк Африка",
			territory.ZM: "Замбия",
			territory.ZW: "Зимбабве",
			territory.ZZ: "билгесез төбәк",
		},
	}
}
