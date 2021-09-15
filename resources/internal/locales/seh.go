// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_seh() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "seh",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Fev", Mar: "Mar", Apr: "Abr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Set", Oct: "Otu", Nov: "Nov", Dec: "Dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Janeiro", Feb: "Fevreiro", Mar: "Marco", Apr: "Abril", May: "Maio", Jun: "Junho", Jul: "Julho", Aug: "Augusto", Sep: "Setembro", Oct: "Otubro", Nov: "Novembro", Dec: "Decembro"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dim", Mon: "Pos", Tue: "Pir", Wed: "Tat", Thu: "Nai", Fri: "Sha", Sat: "Sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "P", Tue: "C", Wed: "T", Thu: "N", Fri: "S", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Dimingu", Mon: "Chiposi", Tue: "Chipiri", Wed: "Chitatu", Thu: "Chinai", Fri: "Chishanu", Sat: "Sabudu"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "#,##0.00¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirém dos Emirados Árabes Unidos", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Cuanza angolano", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dólar australiano", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dinar bareinita", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Franco do Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pula botsuanesa", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dólar canadense", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Franco congolês", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Franco suíço", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Renminbi chinês", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo cabo-verdiano", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Franco do Djibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar argelino", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Libra egípcia", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa da Eritréia", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr etíope", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Libra britânica", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Cedi de Gana (1979–2007)", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi de Gâmbia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Syli da Guiné", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rúpia indiana", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Iene japonês", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Xelim queniano", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Franco de Comores", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dólar liberiano", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti do Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinar líbio", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirém marroquino", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Franco de Madagascar", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya da Mauritânia (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya da Mauritânia", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupia de Maurício", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Cuacha do Maláui", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metical antigo de Moçambique", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Metical de Moçambique", Symbol: "MTn"},
				currency.NAD: cldr.Currency{DisplayName: "Dólar da Namíbia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira nigeriana", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Franco ruandês", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Rial saudita", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia das Seychelles", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dinar sudanês", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Libra sudanesa antiga", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Libra de Santa Helena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leone de Serra Leoa", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Xelim somali", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra de São Tomé e Príncipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra de São Tomé e Príncipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni da Suazilândia", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinar tunisiano", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Xelim da Tanzânia", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Xelim ugandense (1966–1987)", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dólar norte-americano", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Franco CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Franco CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Rand sul-africano", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Cuacha zambiano (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Cuacha zambiano", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dólar do Zimbábue", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "akan",
			language.AM:  "amárico",
			language.AR:  "árabe",
			language.BE:  "bielo-russo",
			language.BG:  "búlgaro",
			language.BN:  "bengali",
			language.CS:  "tcheco",
			language.DE:  "alemão",
			language.EL:  "grego",
			language.EN:  "inglês",
			language.ES:  "espanhol",
			language.FA:  "persa",
			language.FR:  "francês",
			language.HA:  "hausa",
			language.HI:  "hindi",
			language.HU:  "húngaro",
			language.ID:  "indonésio",
			language.IG:  "ibo",
			language.IT:  "italiano",
			language.JA:  "japonês",
			language.JV:  "javanês",
			language.KM:  "cmer",
			language.KO:  "coreano",
			language.MS:  "malaio",
			language.MY:  "birmanês",
			language.NE:  "nepalês",
			language.NL:  "holandês",
			language.PA:  "panjabi",
			language.PL:  "polonês",
			language.PT:  "português",
			language.RO:  "romeno",
			language.RU:  "russo",
			language.RW:  "kinyarwanda",
			language.SEH: "sena",
			language.SO:  "somali",
			language.SV:  "sueco",
			language.TA:  "tâmil",
			language.TH:  "tailandês",
			language.TR:  "turco",
			language.UK:  "ucraniano",
			language.UR:  "urdu",
			language.VI:  "vietnamita",
			language.YO:  "iorubá",
			language.ZH:  "chinês",
			language.ZU:  "zulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andorra",
			territory.AE: "Emirados Árabes Unidos",
			territory.AF: "Afeganistão",
			territory.AG: "Antígua e Barbuda",
			territory.AI: "Anguilla",
			territory.AL: "Albânia",
			territory.AM: "Armênia",
			territory.AO: "Angola",
			territory.AR: "Argentina",
			territory.AS: "Samoa Americana",
			territory.AT: "Áustria",
			territory.AU: "Austrália",
			territory.AW: "Aruba",
			territory.AZ: "Azerbaijão",
			territory.BA: "Bósnia-Herzegovina",
			territory.BB: "Barbados",
			territory.BD: "Bangladesh",
			territory.BE: "Bélgica",
			territory.BF: "Burquina Faso",
			territory.BG: "Bulgária",
			territory.BH: "Bahrain",
			territory.BI: "Burundi",
			territory.BJ: "Benin",
			territory.BM: "Bermudas",
			territory.BN: "Brunei",
			territory.BO: "Bolívia",
			territory.BR: "Brasil",
			territory.BS: "Bahamas",
			territory.BT: "Butão",
			territory.BW: "Botsuana",
			territory.BY: "Belarus",
			territory.BZ: "Belize",
			territory.CA: "Canadá",
			territory.CD: "Congo-Kinshasa",
			territory.CF: "República Centro-Africana",
			territory.CG: "Congo",
			territory.CH: "Suíça",
			territory.CI: "Costa do Marfim",
			territory.CK: "Ilhas Cook",
			territory.CL: "Chile",
			territory.CM: "República dos Camarões",
			territory.CN: "China",
			territory.CO: "Colômbia",
			territory.CR: "Costa Rica",
			territory.CU: "Cuba",
			territory.CV: "Cabo Verde",
			territory.CY: "Chipre",
			territory.CZ: "República Tcheca",
			territory.DE: "Alemanha",
			territory.DJ: "Djibuti",
			territory.DK: "Dinamarca",
			territory.DM: "Dominica",
			territory.DO: "República Dominicana",
			territory.DZ: "Argélia",
			territory.EC: "Equador",
			territory.EE: "Estônia",
			territory.EG: "Egito",
			territory.ER: "Eritréia",
			territory.ES: "Espanha",
			territory.ET: "Etiópia",
			territory.FI: "Finlândia",
			territory.FJ: "Fiji",
			territory.FK: "Ilhas Malvinas",
			territory.FM: "Micronésia",
			territory.FR: "França",
			territory.GA: "Gabão",
			territory.GB: "Reino Unido",
			territory.GD: "Granada",
			territory.GE: "Geórgia",
			territory.GF: "Guiana Francesa",
			territory.GH: "Gana",
			territory.GI: "Gibraltar",
			territory.GL: "Groênlandia",
			territory.GM: "Gâmbia",
			territory.GN: "Guiné",
			territory.GP: "Guadalupe",
			territory.GQ: "Guiné Equatorial",
			territory.GR: "Grécia",
			territory.GT: "Guatemala",
			territory.GU: "Guam",
			territory.GW: "Guiné Bissau",
			territory.GY: "Guiana",
			territory.HN: "Honduras",
			territory.HR: "Croácia",
			territory.HT: "Haiti",
			territory.HU: "Hungria",
			territory.ID: "Indonésia",
			territory.IE: "Irlanda",
			territory.IL: "Israel",
			territory.IN: "Índia",
			territory.IO: "Território Britânico do Oceano Índico",
			territory.IQ: "Iraque",
			territory.IR: "Irã",
			territory.IS: "Islândia",
			territory.IT: "Itália",
			territory.JM: "Jamaica",
			territory.JO: "Jordânia",
			territory.JP: "Japão",
			territory.KE: "Quênia",
			territory.KG: "Quirguistão",
			territory.KH: "Camboja",
			territory.KI: "Quiribati",
			territory.KM: "Comores",
			territory.KN: "São Cristovão e Nevis",
			territory.KP: "Coréia do Norte",
			territory.KR: "Coréia do Sul",
			territory.KW: "Kuwait",
			territory.KY: "Ilhas Caiman",
			territory.KZ: "Casaquistão",
			territory.LA: "Laos",
			territory.LB: "Líbano",
			territory.LC: "Santa Lúcia",
			territory.LI: "Liechtenstein",
			territory.LK: "Sri Lanka",
			territory.LR: "Libéria",
			territory.LS: "Lesoto",
			territory.LT: "Lituânia",
			territory.LU: "Luxemburgo",
			territory.LV: "Letônia",
			territory.LY: "Líbia",
			territory.MA: "Marrocos",
			territory.MC: "Mônaco",
			territory.MD: "Moldávia",
			territory.MG: "Madagascar",
			territory.MH: "Ilhas Marshall",
			territory.ML: "Mali",
			territory.MM: "Mianmar",
			territory.MN: "Mongólia",
			territory.MP: "Ilhas Marianas do Norte",
			territory.MQ: "Martinica",
			territory.MR: "Mauritânia",
			territory.MS: "Montserrat",
			territory.MT: "Malta",
			territory.MU: "Maurício",
			territory.MV: "Maldivas",
			territory.MW: "Malawi",
			territory.MX: "México",
			territory.MY: "Malásia",
			territory.MZ: "Moçambique",
			territory.NA: "Namíbia",
			territory.NC: "Nova Caledônia",
			territory.NE: "Níger",
			territory.NF: "Ilhas Norfolk",
			territory.NG: "Nigéria",
			territory.NI: "Nicarágua",
			territory.NL: "Holanda",
			territory.NO: "Noruega",
			territory.NP: "Nepal",
			territory.NR: "Nauru",
			territory.NU: "Niue",
			territory.NZ: "Nova Zelândia",
			territory.OM: "Omã",
			territory.PA: "Panamá",
			territory.PE: "Peru",
			territory.PF: "Polinésia Francesa",
			territory.PG: "Papua-Nova Guiné",
			territory.PH: "Filipinas",
			territory.PK: "Paquistão",
			territory.PL: "Polônia",
			territory.PM: "Saint Pierre e Miquelon",
			territory.PN: "Pitcairn",
			territory.PR: "Porto Rico",
			territory.PS: "Território da Palestina",
			territory.PT: "Portugal",
			territory.PW: "Palau",
			territory.PY: "Paraguai",
			territory.QA: "Catar",
			territory.RE: "Reunião",
			territory.RO: "Romênia",
			territory.RU: "Rússia",
			territory.RW: "Ruanda",
			territory.SA: "Arábia Saudita",
			territory.SB: "Ilhas Salomão",
			territory.SC: "Seychelles",
			territory.SD: "Sudão",
			territory.SE: "Suécia",
			territory.SG: "Cingapura",
			territory.SH: "Santa Helena",
			territory.SI: "Eslovênia",
			territory.SK: "Eslováquia",
			territory.SL: "Serra Leoa",
			territory.SM: "San Marino",
			territory.SN: "Senegal",
			territory.SO: "Somália",
			territory.SR: "Suriname",
			territory.ST: "São Tomé e Príncipe",
			territory.SV: "El Salvador",
			territory.SY: "Síria",
			territory.SZ: "Suazilândia",
			territory.TC: "Ilhas Turks e Caicos",
			territory.TD: "Chade",
			territory.TG: "Togo",
			territory.TH: "Tailândia",
			territory.TJ: "Tadjiquistão",
			territory.TK: "Tokelau",
			territory.TL: "Timor Leste",
			territory.TM: "Turcomenistão",
			territory.TN: "Tunísia",
			territory.TO: "Tonga",
			territory.TR: "Turquia",
			territory.TT: "Trinidad e Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwan",
			territory.UA: "Ucrânia",
			territory.UG: "Uganda",
			territory.US: "Estados Unidos",
			territory.UY: "Uruguai",
			territory.UZ: "Uzbequistão",
			territory.VA: "Vaticano",
			territory.VC: "São Vicente e Granadinas",
			territory.VE: "Venezuela",
			territory.VG: "Ilhas Virgens Britânicas",
			territory.VI: "Ilhas Virgens dos EUA",
			territory.VN: "Vietnã",
			territory.VU: "Vanuatu",
			territory.WF: "Wallis e Futuna",
			territory.WS: "Samoa",
			territory.YE: "Iêmen",
			territory.YT: "Mayotte",
			territory.ZA: "África do Sul",
			territory.ZM: "Zâmbia",
			territory.ZW: "Zimbábue",
		},
	}
}
