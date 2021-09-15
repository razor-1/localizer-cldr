// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_lag_TZ() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "lag_TZ",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Fúngatɨ", Feb: "Naanɨ", Mar: "Keenda", Apr: "Ikúmi", May: "Inyambala", Jun: "Idwaata", Jul: "Mʉʉnchɨ", Aug: "Vɨɨrɨ", Sep: "Saatʉ", Oct: "Inyi", Nov: "Saano", Dec: "Sasatʉ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "F", Feb: "N", Mar: "K", Apr: "I", May: "I", Jun: "I", Jul: "M", Aug: "V", Sep: "S", Oct: "I", Nov: "S", Dec: "S"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Kʉfúngatɨ", Feb: "Kʉnaanɨ", Mar: "Kʉkeenda", Apr: "Kwiikumi", May: "Kwiinyambála", Jun: "Kwiidwaata", Jul: "Kʉmʉʉnchɨ", Aug: "Kʉvɨɨrɨ", Sep: "Kʉsaatʉ", Oct: "Kwiinyi", Nov: "Kʉsaano", Dec: "Kʉsasatʉ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Píili", Mon: "Táatu", Tue: "Íne", Wed: "Táano", Thu: "Alh", Fri: "Ijm", Sat: "Móosi"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "P", Mon: "T", Tue: "E", Wed: "O", Thu: "A", Fri: "I", Sat: "M"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Jumapíiri", Mon: "Jumatátu", Tue: "Jumaíne", Wed: "Jumatáano", Thu: "Alamíisi", Fri: "Ijumáa", Sat: "Jumamóosi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "TOO", PM: "MUU"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "TOO", PM: "MUU"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "¤\u00a00K", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Diriháamu ya Ʉtemi wa Kɨaráabu", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwáanza ya Angóola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dóola ya Ausitereelía", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dináari ya Baharéeni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Faráanga ya Burúundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Púula ya Botiswáana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dóola ya Kánada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faráanga ya Kóongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faráaka ya Uswíisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yúani Renimínibi ya Chíina", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Esikúudo ya Kepuvéede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faráanga ya Jibóuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dináairi ya Alijéria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Páundi ya Mísiri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nákɨfa ya Eriterea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Bíiri ya Ʉhabéeshi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yúuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Páundi ya Ʉɨngɨréesa", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Séedi ya Gáana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Daláasi ya Gámbia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Faráanga ya Gíine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupía ya Índia", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yéeni ya Japáani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilíingi ya Kéenya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Faráanga ya Komóoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dóola ya Libéria", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lóoti ya Lesóoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dináari ya Líbia", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Diriháamu ya Moróoko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Mpía ya bukini", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ungwíiya ya Moritánia (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ungwíiya ya Moritánia", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupía ya Moríisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwáacha ya Maláawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikáali ya Musumbíiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dóola ya Namíbia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naíira ya Niijéria", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Faráanga ya Rwáanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyáali ya Saudía", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupía ya Shelishéeli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Páundi ya Sudáani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Páundi ya Mʉtakatíifu Heléena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leóoni", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilíingi ya Somália", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dóbura ya SaoTóome na Pirínsipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dóbura ya SaoTóome na Pirínsipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilengéeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dináari ya Tunísia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilíingi ya Taansanía", Symbol: "TSh"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Shilíingi ya Ugáanda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dóola ya Amerɨ́ka", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Faráanga ya CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Faráanga ya CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Ráandi ya Afɨrɨka ya Saame", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwácha ya Sámbia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwácha ya Sámbia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dóola ya Simbáabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Kɨakáani",
			language.AM:  "Kɨmʉháari",
			language.AR:  "Kɨaráabu",
			language.BE:  "Kɨberalúusi",
			language.BG:  "Kɨbulugária",
			language.BN:  "Kɨbangála",
			language.CS:  "Kɨchéeki",
			language.DE:  "Kɨjerʉmáani",
			language.EL:  "Kɨgiríki",
			language.EN:  "Kɨɨngeréesa",
			language.ES:  "Kɨhispánia",
			language.FA:  "Kɨajéemi",
			language.FR:  "Kɨfaráansa",
			language.HA:  "Kɨhaúusa",
			language.HI:  "Kɨhíindi",
			language.HU:  "Kɨhungári",
			language.ID:  "Kɨɨndonésia",
			language.IG:  "Kiígibo",
			language.IT:  "Kɨtaliáano",
			language.JA:  "Kɨjapáani",
			language.JV:  "Kɨjáava",
			language.KM:  "Kɨkambódia",
			language.KO:  "Kɨkoréa",
			language.LAG: "Kɨlaangi",
			language.MS:  "Kɨmelésia",
			language.MY:  "Kɨbáama",
			language.NE:  "Kɨnepáali",
			language.NL:  "Kɨholáanzi",
			language.PA:  "Kɨpúnjabi",
			language.PL:  "Kɨpólandi",
			language.PT:  "Kɨréeno",
			language.RO:  "Kɨromanía",
			language.RU:  "Kɨrúusi",
			language.RW:  "Kɨnyarwáanda",
			language.SO:  "Kɨsómáali",
			language.SV:  "Kɨswíidi",
			language.TA:  "Kɨtamíili",
			language.TH:  "Kɨtáilandi",
			language.TR:  "Kɨturúuki",
			language.UK:  "Kɨukɨranía",
			language.UR:  "Kɨúrdu",
			language.VI:  "Kɨvietináamu",
			language.YO:  "Kɨyorúuba",
			language.ZH:  "Kɨchíina",
			language.ZU:  "Kɨzúulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andóra",
			territory.AE: "Ʉtemi wa Kɨaráabu",
			territory.AF: "Afuganisitáani",
			territory.AG: "Antigúua na Baribúuda",
			territory.AI: "Anguíila",
			territory.AL: "Alubánia",
			territory.AM: "Ariménia",
			territory.AO: "Angóola",
			territory.AR: "Ajentíina",
			territory.AS: "Samóoa ya Amerɨ́ka",
			territory.AT: "Áusitiria",
			territory.AU: "Ausiteréelia",
			territory.AW: "Arúuba",
			territory.AZ: "Azabajáani",
			territory.BA: "Bósinia",
			territory.BB: "Babadóosi",
			territory.BD: "Bangaladéeshi",
			territory.BE: "Ʉbeligíiji",
			territory.BF: "Bukinafáaso",
			territory.BG: "Buligaría",
			territory.BH: "Baharéeni",
			territory.BI: "Burúundi",
			territory.BJ: "Beníini",
			territory.BM: "Berimúuda",
			territory.BN: "Burunéei",
			territory.BO: "Bolívia",
			territory.BR: "Brasíili",
			territory.BS: "Baháama",
			territory.BT: "Butáani",
			territory.BW: "Botiswáana",
			territory.BY: "Belarúusi",
			territory.BZ: "Belíise",
			territory.CA: "Kánada",
			territory.CD: "Jamuhúuri ya Kɨdemokurasía ya Kóongo",
			territory.CF: "Juhúuri ya Afɨrɨka ya katɨ katɨ",
			territory.CG: "Kóongo",
			territory.CH: "Uswíisi",
			territory.CI: "Ivori Kositi",
			territory.CK: "Visíiwa vya Kúuku",
			territory.CL: "Chíile",
			territory.CM: "Kamerúuni",
			territory.CN: "Chíina",
			territory.CO: "Kolómbia",
			territory.CR: "Kósita Rɨ́ɨka",
			territory.CU: "Kyúuba",
			territory.CV: "Kepuvéede",
			territory.CY: "Kupuróosi",
			territory.CZ: "Jamuhúuri ya Chéeki",
			territory.DE: "Ʉjerumáani",
			territory.DJ: "Jibúuti",
			territory.DK: "Denimaki",
			territory.DM: "Domínɨka",
			territory.DO: "Jamuhúuri ya Dominɨka",
			territory.DZ: "Alijéria",
			territory.EC: "Íkwado",
			territory.EE: "Estonía",
			territory.EG: "Mísiri",
			territory.ER: "Eriterea",
			territory.ES: "Hisipánia",
			territory.ET: "Ʉhabéeshi",
			territory.FI: "Ufíini",
			territory.FJ: "Fíiji",
			territory.FK: "Visíiwa vya Fakulandi",
			territory.FM: "Mikironésia",
			territory.FR: "Ʉfaráansa",
			territory.GA: "Gabóoni",
			territory.GB: "Ʉɨngeréesa",
			territory.GD: "Girenáada",
			territory.GE: "Jójia",
			territory.GF: "Gwiyáana yʉ Ʉfaráansa",
			territory.GH: "Gáana",
			territory.GI: "Jiburálita",
			territory.GL: "Giriniláandi",
			territory.GM: "Gámbia",
			territory.GN: "Gíine",
			territory.GP: "Gwadelúupe",
			territory.GQ: "Gíine Ikwéeta",
			territory.GR: "Ugiríki",
			territory.GT: "Gwatemáala",
			territory.GU: "Gwani",
			territory.GW: "Gíine Bisáau",
			territory.GY: "Guyáana",
			territory.HN: "Honduráasi",
			territory.HR: "Koréshia",
			territory.HT: "Haíiti",
			territory.HU: "Hungária",
			territory.ID: "Indonésia",
			territory.IE: "Ayaláandi",
			territory.IL: "Isiraéeli",
			territory.IN: "Índia",
			territory.IO: "Ɨsɨ yʉ Ʉɨngeréesa irivii ra Híindi",
			territory.IQ: "Iráaki",
			territory.IR: "Ʉajéemi",
			territory.IS: "Aisiláandi",
			territory.IT: "Itália",
			territory.JM: "Jamáika",
			territory.JO: "Jódani",
			territory.JP: "Japáani",
			territory.KE: "Kéenya",
			territory.KG: "Kirigisitáani",
			territory.KH: "Kambódia",
			territory.KI: "Kiribáati",
			territory.KM: "Komóoro",
			territory.KN: "Mʉtakatíifu kitisi na Nevíisi",
			territory.KP: "Koréa yʉ ʉtʉrʉko",
			territory.KR: "Koréa ya Saame",
			territory.KW: "Kʉwáiti",
			territory.KY: "Visíiwa vya Kayimani",
			territory.KZ: "Kazakasitáani",
			territory.LA: "Laóosi",
			territory.LB: "Lebanóoni",
			territory.LC: "Mʉtakatíifu Lusíia",
			territory.LI: "Lishentéeni",
			territory.LK: "Siriláanka",
			territory.LR: "Liibéria",
			territory.LS: "Lesóoto",
			territory.LT: "Lisuánia",
			territory.LU: "Lasembáagi",
			territory.LV: "Lativia",
			territory.LY: "Líbia",
			territory.MA: "Moróoko",
			territory.MC: "Monáako",
			territory.MD: "Molidóova",
			territory.MG: "Bukíini",
			territory.MH: "Visíiwa vya Marisháali",
			territory.ML: "Máali",
			territory.MM: "Miáama",
			territory.MN: "Mongólia",
			territory.MP: "Visiwa vya Mariana vya Kaskazini",
			territory.MQ: "Maritiníiki",
			territory.MR: "Moritánia",
			territory.MS: "Monteráati",
			territory.MT: "Málita",
			territory.MU: "Moríisi",
			territory.MV: "Modíivu",
			territory.MW: "Maláawi",
			territory.MX: "Mekisiko",
			territory.MY: "Maleísia",
			territory.MZ: "Musumbíiji",
			territory.NA: "Namíbia",
			territory.NC: "Kaledónia Ifya",
			territory.NE: "Níija",
			territory.NF: "Kisíiwa cha Nofifóoki",
			territory.NG: "Niijéria",
			territory.NI: "Nikarágʉa",
			territory.NL: "Ʉholáanzi",
			territory.NO: "Norwe",
			territory.NP: "Nepáali",
			territory.NR: "Naúuru",
			territory.NU: "Niúue",
			territory.NZ: "Nyuzílandi",
			territory.OM: "Ómani",
			territory.PA: "Panáama",
			territory.PE: "Péeru",
			territory.PF: "Polinésia yʉ Ʉfaráansa",
			territory.PG: "Papúua",
			territory.PH: "Ufilipíino",
			territory.PK: "Pakisitáani",
			territory.PL: "Pólandi",
			territory.PM: "Mʉtakatíifu Peéteri na Mɨkaéeli",
			territory.PN: "Patikaírini",
			territory.PR: "Pwetorɨ́ɨko",
			territory.PS: "Mweemberera wa kʉmweeri wa Gáaza",
			territory.PT: "Ʉréeno",
			territory.PW: "Paláau",
			territory.PY: "Paraguáai",
			territory.QA: "Katáari",
			territory.RE: "Reyunióoni",
			territory.RO: "Romaníia",
			territory.RU: "Urúusi",
			territory.RW: "Rwáanda",
			territory.SA: "Saudíia Arabíia",
			territory.SB: "Visíiwa vya Solomóoni",
			territory.SC: "Shelishéeli",
			territory.SD: "Sudáani",
			territory.SE: "Uswíidi",
			territory.SG: "Singapoo",
			territory.SH: "Mʉtakatíifu Heléena",
			territory.SI: "Sulovénia",
			territory.SK: "Sulováakia",
			territory.SL: "Seraleóoni",
			territory.SM: "Samaríino",
			territory.SN: "Senegáali",
			territory.SO: "Somália",
			territory.SR: "Surináamu",
			territory.ST: "Sao Tóome na Pirinsipe",
			territory.SV: "Elisalivado",
			territory.SY: "Síria",
			territory.SZ: "Ʉswáazi",
			territory.TC: "Visíiwa vya Turíiki na Kaíiko",
			territory.TD: "Cháadi",
			territory.TG: "Tóogo",
			territory.TH: "Táilandi",
			territory.TJ: "Tajikisitáani",
			territory.TK: "Tokeláau",
			territory.TL: "Timóori yi Itʉʉmba",
			territory.TM: "Uturukimenisitáani",
			territory.TN: "Tunísia",
			territory.TO: "Tóonga",
			territory.TR: "Uturúuki",
			territory.TT: "Tiriníida ya Tobáago",
			territory.TV: "Tuváalu",
			territory.TW: "Taiwáani",
			territory.TZ: "Taansanía",
			territory.UA: "Ʉkɨréeni",
			territory.UG: "Ʉgáanda",
			territory.US: "Amerɨka",
			territory.UY: "Uruguáai",
			territory.UZ: "Usibekisitáani",
			territory.VA: "Vatikáani",
			territory.VC: "Mʉtakatíifu Viséenti na Gernadíini",
			territory.VE: "Venezuéela",
			territory.VG: "Visíiwa vya Vigíini vya Ʉɨngeréesa",
			territory.VI: "Visíiwa vya Vigíini vya Amerɨ́ka",
			territory.VN: "Vietináamu",
			territory.VU: "Vanuáatu",
			territory.WF: "Walíisi na Futúuna",
			territory.WS: "Samóoa",
			territory.YE: "Yémeni",
			territory.YT: "Mayóote",
			territory.ZA: "Afɨrɨka ya Saame",
			territory.ZM: "Sámbia",
			territory.ZW: "Simbáabwe",
		},
	}
}
