// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/localizer-cldr"
	"github.com/razor-1/localizer-cldr/resources/currency"
	"github.com/razor-1/localizer-cldr/resources/language"
	"github.com/razor-1/localizer-cldr/resources/territory"
)

func Get_kok_IN() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "kok_IN",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "dd-MM-y", Short: "d-M-yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "जाने", Feb: "फेब्रु", Mar: "मार्च", Apr: "एप्री", May: "मे", Jun: "जून", Jul: "जुल", Aug: "ऑग", Sep: "सप्टें", Oct: "ऑक्टो", Nov: "नो", Dec: "डिसे"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "जानेवारी", Feb: "फेब्रुवारी", Mar: "मार्च", Apr: "एप्रिल", May: "मे", Jun: "जून", Jul: "जुलय", Aug: "ऑगस्ट", Sep: "सप्टेंबर", Oct: "ऑक्टोबर", Nov: "नोव्हेंबर", Dec: "डिसेंबर"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "आयतार", Mon: "सोमार", Tue: "मंगळार", Wed: "बुधवार", Thu: "बिरेस्तार", Fri: "शुक्रार", Sat: "शेनवार"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "आ", Mon: "सो", Tue: "मं", Wed: "बु", Thu: "ब", Fri: "शु", Sat: "शे"}, Short: cldr.CalendarDayFormatNameValue{Sun: "आय", Mon: "सोम", Tue: "मंगळ", Wed: "बुध", Thu: "बिरे", Fri: "शुक्र", Sat: "शेन"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "आयतार", Mon: "सोमार", Tue: "मंगळार", Wed: "बुधवार", Thu: "बिरेस्तार", Fri: "शुक्रार", Sat: "शेनवार"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "युनाइटेड अरब इमीरॅट्स दिरहम", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "अफगाण अफगाणी", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "अल्बेनियन लेक", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "अर्मेनियन ड्राम", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "नॅदरलँड अँटिलियन गिल्डर", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "अंगोलन क्वॉन्ज", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "अर्जेंटिना पेसो", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ऑस्ट्रेलियाई डॉलर", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "अरुबान फ्लोरिन", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "अज़रबैजानी मनात", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "बोस्निया-हेर्जेगोविना रुपांतरीत मार्क", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "बार्बाडियान डॉलर", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "बांगलादेशी टाका", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "बल्गेरियन लेव", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "बहरिनी डिनार", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "बुरुंडी फ्रँक", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "बरमुदान डॉलर", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ब्रूनेई डॉलर", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "बोलिव्हियन बोलिवियानो", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "ब्राझिलियन रियाल", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "बहामियन डॉलर", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "भुतानीज नागल्ट्रम", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "बोत्सवाना पुला", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "बैलोरुसियन् रूबल", Symbol: "BYN"},
				currency.BZD: cldr.Currency{DisplayName: "बेलिझ डॉलर", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "कॅनाडियन डॉलर", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "काँगोलिसी फ्रँक", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "स्विस फ्रँक", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "चिली पेसो", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "चिनी युआन (ऑफशोर)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "चिनी युआन", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "कोलंबियन पेसो", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "कोस्ता रिका कॉलॉन", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "क्युबान रुपांतरीत पेसो", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "क्युबान पेसो", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "केप वर्दे एस्कुडो", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "चेक कोरुना", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "जिबूती फ्रँक", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "डॅनिश क्रोन", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "डोमिनिकन पेसो", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "अल्जेरियाई डिनार", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "ईजिप्ती पावंड", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "इरिट्रियन नाक्फा", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "इथियोपियाई बिरर", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "युरो", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "फिजी डॉलर", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "फ़ॉकलैंड आइलैंड्स पावंड", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "ब्रिटिश पावंड", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "जॉर्जियन लारी", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "घानाई सेडी", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "जिब्राल्टर पावंड", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "गॅम्बियन दलासी", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "गिनीन फ्रँक", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ग्वाटेमाला कुएट्झल", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "गयाना डॉलर", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "हाँग काँग डॉलर", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "होंडुरान लेम्पिरा", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "क्रोयेषियन् कुना", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "हैतीयन गौर्डे", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "हंगेरियन फोरिंट", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "इंडोनेशियन रुपिया", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "इस्त्रायली न्यु शेकेल", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "भारतीय रुपया", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "इराकी डिनार", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ईरानी रियाल", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "आईस्लान्डिक क्रोना", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "जमैकन डॉलर", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "जॉर्डनियन डिनार", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "जपानी येन", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "केनयाई शिलिंग", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "किरगिझस्तान सोम", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "कंबोडियन रियाल", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "कोमोरियन फ्रँक", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "उत्तर कोरियन वॉन", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "दक्षिण कोरियन वॉन", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "कुवेती डिनार", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "कैमेन आइलैंड्स डॉलर", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "कझाकस्तानी टेंग", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "लाओ किप", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "लिबानेस पावंड", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "श्री लंका रुपया", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "लायबेरियन डॉलर", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "लीबियान डिनार", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "मोरक्कन दिरहम", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "मोल्दोवान लियू", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "मलागासी एरियारी", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "मसीडोनियन् डिनर", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "म्यानमार क्यात", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "मंगोलियन तुगरिक", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "मकानेसे पटका", Symbol: "MOP"},
				currency.MRU: cldr.Currency{DisplayName: "मॉरिटानिया उगिया", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "मॉरिशस रुपी", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "मालदिवी रुफिया", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "मलावियन क्वाचा", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "मेक्सिकन पेसो", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "मलेशियाई रिंग्गित", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "मोझांबिकन मेटिकल", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "नामीबिया डॉलर", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "नायजेरियन नायरा", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "निकारागुआन कॉर्डोबा", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "नॉर्वेगन क्रोन", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "नेपाळी रुपया", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "न्युझीलॅन्ड डॉलर", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ओमानी रियाल", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "पानामानियन बाल्बोआ", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "पेरिवियन सोल", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "पापुआ न्यु गिनी किना", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "फिलिपिनी पेसो", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "पाकिस्तानी रुपया", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "पोलिष झ्लोटी", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "पराग्वेन गौरानी", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "कतारी रियाल", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "रोमानियन् लियू", Symbol: "रॉन"},
				currency.RSD: cldr.Currency{DisplayName: "सर्बियन डिनार", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "रुसी रुबल", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "रवांडा फ्रँक", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "सौदी रियाल", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "सोलोमन आयलँड्स डॉलर", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "सेशेल्लोइस रुपी", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "सुदानी पावंड", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "स्वीदीष क्रोन", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "सिंगापूरी डॉलर", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "सेंट हेलिना पावंड", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "सिएरा लियॉनी लियॉन", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "सोमाली शिलिंग", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "सुरीनामी डॉलर", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "दक्षिण सुडानी पावंड", Symbol: "SSP"},
				currency.STN: cldr.Currency{DisplayName: "साओ टोम आनी प्रिन्सिप डोब्रा", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "सिरियन पावंड", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "स्वाजी लिलांगेनी", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "थाई बाट", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "ताजिकिस्तानी सोमोनी", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "तुर्कमेनिस्तानी मनत", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "ट्यूनीशियन डिनार", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "टोंगन पांगा", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "तुर्किश लायरा", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ट्रिनीडाड आनी टोबॅगो डॉलर", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "न्यू तायवान डॉलर", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "तंजानिया शिलिंग", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "युक्रेनियन् रिव्निया", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "युगांडा शिलिंग", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "युएस डॉलर", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "उरुग्वेन पेसो", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "उज़्बेकिस्तानी सोम", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "विनेझुएला बोलिव्हर", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "वियतनामी डोंग", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "वानूआतू वातू", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "समोआई टाला", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "मध्य अफ्रीकी सीएफए फ्रँक", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "उदेंत कॅरिबियन डॉलर", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "अस्तंत आफ्रिकी सीएफए फ्रँक", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "सीएफपी फ्रँक", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "अज्ञात चलन", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "येमेनी रियाल", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "दक्षिण आफ्रिकन रँड", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "झांबियन क्वाचा", Symbol: "ZMW"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "अफार",
			language.AB:      "अबखेज़ियन",
			language.ACE:     "अचायनीज",
			language.ADA:     "अडांग्मे",
			language.ADY:     "अडिघे",
			language.AF:      "अफ्रिकान्स",
			language.AGQ:     "अघेम",
			language.AIN:     "आयनू",
			language.AK:      "अकान",
			language.ALE:     "आलिट",
			language.ALT:     "दक्षिणी अल्टाय",
			language.AM:      "अमहारिक्",
			language.AN:      "आरागोनिस",
			language.ANP:     "अंगिका",
			language.AR:      "अरेबिक",
			language.AR_001:  "आधुनिक प्रमाणित अरेबिक",
			language.ARN:     "मापुचे",
			language.ARP:     "अरापाहो",
			language.AS:      "आसामी",
			language.ASA:     "असु",
			language.AST:     "अस्टुरियान",
			language.AV:      "अवारिक",
			language.AWA:     "अवधी",
			language.AY:      "ऐमरा",
			language.AZ:      "अझेरी",
			language.BA:      "बष्किर",
			language.BAN:     "बालिनिस",
			language.BAS:     "बस्सा",
			language.BE:      "बैलोरुसियन्",
			language.BEM:     "बेम्बा",
			language.BEZ:     "बेना",
			language.BG:      "बल्गेरियन",
			language.BHO:     "भोजपुरी",
			language.BI:      "बिसलमा",
			language.BIN:     "बिनी",
			language.BLA:     "सिकसिका",
			language.BM:      "बंबारा",
			language.BN:      "बांग्ला",
			language.BO:      "तिबेटी",
			language.BR:      "ब्रेटन",
			language.BRX:     "बोडो",
			language.BS:      "बोस्नियन",
			language.BUG:     "बगिनिस",
			language.BYN:     "ब्लीन",
			language.CA:      "कटलान",
			language.CCP:     "चक्मा",
			language.CE:      "चिचेन",
			language.CEB:     "सिबौना",
			language.CGG:     "चिगा",
			language.CH:      "चामोर्रो",
			language.CHK:     "छुनिस",
			language.CHM:     "मारी",
			language.CHO:     "चोताव",
			language.CHR:     "चॅरोकी",
			language.CHY:     "चेयनी",
			language.CKB:     "मध्य खुर्दीश",
			language.CO:      "कोर्शियन",
			language.CRS:     "सेसेल्वा क्रयॉल फ्रेन्च",
			language.CS:      "चेक",
			language.CU:      "चर्च स्लेव्हीक",
			language.CV:      "छुवास",
			language.CY:      "वेळ्ष्",
			language.DA:      "डॅनिश",
			language.DAK:     "डाकोटा",
			language.DAR:     "दार्ग्वा",
			language.DAV:     "तायता",
			language.DE:      "जर्मन",
			language.DE_AT:   "ऑस्ट्रियन जर्मन",
			language.DE_CH:   "स्विझ हाय जर्मन",
			language.DGR:     "डोगरीब",
			language.DJE:     "झर्मा",
			language.DSB:     "लोवर सोर्बियन",
			language.DUA:     "डौल",
			language.DV:      "दिवेही",
			language.DYO:     "जोला-फोनी",
			language.DZ:      "झोंग्खा",
			language.DZG:     "डाझागा",
			language.EBU:     "एम्बु",
			language.EE:      "एव",
			language.EFI:     "एफीक",
			language.EKA:     "एकाजुक",
			language.EL:      "ग्रीक्",
			language.EN:      "इंग्लीश",
			language.EN_AU:   "ऑस्ट्रेलियन इंग्लीश",
			language.EN_CA:   "कॅनडियन इंग्लीश",
			language.EN_GB:   "यूके इंग्लीश",
			language.EN_US:   "यूएस इंग्लीश",
			language.EO:      "इस्परान्टो",
			language.ES:      "स्पॅनीश",
			language.ES_419:  "लातीं अमेरिकन स्पॅनीश",
			language.ES_ES:   "युरोपियन स्पॅनीश",
			language.ES_MX:   "मॅक्सिकन स्पॅनीश",
			language.ET:      "इस्टोनियन्",
			language.EU:      "बास्क",
			language.EWO:     "एवोंडो",
			language.FA:      "पर्शियन",
			language.FA_AF:   "दारी",
			language.FF:      "फुला",
			language.FI:      "फिन्निष्",
			language.FIL:     "फिलिपिनो",
			language.FJ:      "फिजी",
			language.FO:      "फेरोस्",
			language.FON:     "फोन",
			language.FR:      "फ्रेंच",
			language.FR_CA:   "कॅनडियन फ्रेंच",
			language.FR_CH:   "स्विझ फ्रेंच",
			language.FUR:     "फ्रिलियन",
			language.FY:      "पश्चिमी फ्रिशियन",
			language.GA:      "ऐरिष",
			language.GAA:     "गा",
			language.GD:      "स्कॉटीश गॅलीक",
			language.GEZ:     "गेझ",
			language.GIL:     "गिलबर्टीस",
			language.GL:      "गेलीशियन",
			language.GN:      "गौरानी",
			language.GOR:     "गोरोंटालो",
			language.GSW:     "स्विज जर्मन",
			language.GU:      "गुजराती",
			language.GUZ:     "गुसी",
			language.GV:      "मॅन्स",
			language.GWI:     "ग्विच",
			language.HA:      "हौसा",
			language.HAW:     "हवायियान",
			language.HE:      "हिब्रू",
			language.HI:      "हिन्दी",
			language.HIL:     "हिलीगायनॉन",
			language.HMN:     "मोंग",
			language.HR:      "क्रोयेषियन्",
			language.HSB:     "अपर सोर्बियन",
			language.HT:      "हैतियन क्रेयॉल",
			language.HU:      "हंगेरियन्",
			language.HUP:     "हुपा",
			language.HY:      "आर्मेनियन",
			language.HZ:      "हिरिरो",
			language.IA:      "इन्टरलिंग्वा",
			language.IBA:     "आयबन",
			language.IBB:     "ईबिबियो",
			language.ID:      "इंडोनेशियन",
			language.IE:      "इन्टरलिंग्",
			language.IG:      "इग्बो",
			language.II:      "सिच्युआन यी",
			language.IK:      "इनूपेयाक्",
			language.ILO:     "लोको",
			language.INH:     "इंगूश",
			language.IO:      "इदो",
			language.IS:      "आईस्लान्डिक",
			language.IT:      "इटालियन",
			language.IU:      "इन्युकट्ट",
			language.JA:      "जपानी",
			language.JBO:     "लोबजान",
			language.JGO:     "न्गोंबा",
			language.JMC:     "मचामे",
			language.JV:      "जावनीस्",
			language.KA:      "जार्जियन्",
			language.KAB:     "काबायले",
			language.KAC:     "काचीन",
			language.KAJ:     "जु",
			language.KAM:     "कंबा",
			language.KBD:     "काबार्डियन",
			language.KCG:     "त्याप",
			language.KDE:     "माकोंडे",
			language.KEA:     "काबुवर्डियनु",
			language.KFO:     "कोरो",
			language.KHA:     "खासी",
			language.KHQ:     "कोयरा छिनी",
			language.KI:      "किकुयु",
			language.KJ:      "कुयांमा",
			language.KK:      "कज़ख्",
			language.KKJ:     "काको",
			language.KL:      "कालाल्लिसुट",
			language.KLN:     "कालेंजीन",
			language.KM:      "कंबोडियन",
			language.KMB:     "किंबुंडु",
			language.KN:      "कन्नडा",
			language.KO:      "कोरियन्",
			language.KOK:     "कोंकणी",
			language.KPE:     "पेल्ले",
			language.KR:      "कानुरी",
			language.KRC:     "कराची-बाल्कर",
			language.KRL:     "कारेलियन",
			language.KRU:     "कुरुख",
			language.KS:      "कश्मीरी",
			language.KSB:     "शांबाला",
			language.KSF:     "बाफिया",
			language.KSH:     "कोलोनियन",
			language.KU:      "कुर्दिष",
			language.KUM:     "कुमयक",
			language.KV:      "कोमी",
			language.KW:      "कोर्निश",
			language.KY:      "किर्गिज़",
			language.LA:      "लाटिन",
			language.LAD:     "लाडिनो",
			language.LAG:     "लांगी",
			language.LB:      "लक्सेमबर्गीश",
			language.LEZ:     "लेझघियान",
			language.LG:      "गांडा",
			language.LI:      "लिंबुर्ग",
			language.LKT:     "लाकोटा",
			language.LN:      "लिंगाला",
			language.LO:      "लाओ",
			language.LOZ:     "लोझीं",
			language.LRC:     "उत्तरीय लुरी",
			language.LT:      "लिथुआनियन्",
			language.LU:      "लुबा-काटांगा",
			language.LUA:     "लुबा-लुलुआ",
			language.LUN:     "लुंडा",
			language.LUO:     "लुओ",
			language.LUS:     "मिझो",
			language.LUY:     "लुय",
			language.LV:      "लाट्वियन् (लेट्टिष्)",
			language.MAD:     "मादुरेसे",
			language.MAG:     "मगाही",
			language.MAI:     "मैथिली",
			language.MAK:     "माक",
			language.MAS:     "मसाई",
			language.MDF:     "मोक्ष",
			language.MEN:     "मेंडे",
			language.MER:     "मेरू",
			language.MFE:     "मोरिसेन",
			language.MG:      "मलागसी",
			language.MGH:     "माखुवा-मिट्टो",
			language.MGO:     "मेटा",
			language.MH:      "मार्शली",
			language.MI:      "मुरी",
			language.MIC:     "मिक्माक",
			language.MIN:     "मिनाग्काबौ",
			language.MK:      "मसीडोनियन्",
			language.ML:      "मळियाळम",
			language.MN:      "मंगोलियन",
			language.MNI:     "मणिपुरी",
			language.MOH:     "मोहाक",
			language.MOS:     "मोस्सी",
			language.MR:      "मराठी",
			language.MS:      "मलय",
			language.MT:      "मालतीस्",
			language.MUA:     "मुडांग",
			language.MUL:     "साबार भाशा",
			language.MUS:     "क्रिक",
			language.MWL:     "मिरांडीस",
			language.MY:      "बर्मीज़्",
			language.MYV:     "एरझिया",
			language.MZN:     "मझांडेराणी",
			language.NA:      "नौरो",
			language.NAP:     "नेपोलिटन",
			language.NAQ:     "नामा",
			language.NB:      "नोर्वेजियन बोकमाल",
			language.ND:      "उत्तर न्डेबेले",
			language.NDS:     "निम्न जर्मन",
			language.NE:      "नेपाळी",
			language.NEW:     "नेवरी",
			language.NG:      "डोंगा",
			language.NIA:     "नियास",
			language.NIU:     "नियुन",
			language.NL:      "डच्",
			language.NL_BE:   "फ्लेमिश",
			language.NMG:     "ख्वासी",
			language.NN:      "नोर्वोजियन नायनोर्स्क",
			language.NNH:     "न्गेबून",
			language.NO:      "नोर्वेजियन",
			language.NOG:     "नोगाय",
			language.NQO:     "नको",
			language.NR:      "दक्षिण डेबेले",
			language.NSO:     "उत्तरीय सोथो",
			language.NUS:     "न्युयर",
			language.NV:      "नावाजो",
			language.NY:      "नांन्जा",
			language.NYN:     "नानकोले",
			language.OC:      "ओसिटान्",
			language.OM:      "ओरोमो",
			language.OR:      "ओडिया",
			language.OS:      "ओसेटिक",
			language.PA:      "पंजाबी",
			language.PAG:     "पांगासियान",
			language.PAM:     "पांपान्गा",
			language.PAP:     "पापिमेंटो",
			language.PAU:     "पालुयान",
			language.PCM:     "नायझेरियन पिडगीन",
			language.PL:      "पॉलीश",
			language.PRG:     "प्रुसियन",
			language.PS:      "पाष्टो",
			language.PT:      "पोर्तुगीज",
			language.PT_BR:   "ब्राझिलियन पोर्तुगीज",
			language.PT_PT:   "युरोपियन पोर्तुगीज",
			language.QU:      "क्वेच्वा",
			language.QUC:     "किचे",
			language.RAP:     "रापान्यु",
			language.RAR:     "रारोटोंगान",
			language.RM:      "रोमान्श",
			language.RN:      "रुंदी",
			language.RO:      "रोमानियन",
			language.RO_MD:   "मोल्डावियन्",
			language.ROF:     "रोम्बो",
			language.ROOT:    "रूट",
			language.RU:      "रशियन",
			language.RUP:     "आरोमेनियन",
			language.RW:      "किन्यार्वान्डा",
			language.RWK:     "रवा",
			language.SA:      "संस्कृत",
			language.SAD:     "संडावे",
			language.SAH:     "सखा",
			language.SAQ:     "साम्बरू",
			language.SAT:     "संथाली",
			language.SBA:     "गांबे",
			language.SBP:     "सांगू",
			language.SC:      "सार्डिनियान",
			language.SCN:     "सिसिलियान",
			language.SCO:     "स्कॉट्स",
			language.SD:      "सिंधी",
			language.SE:      "उत्तरीय सामी",
			language.SEH:     "सेना",
			language.SES:     "कोयराबोरो सेन्नी",
			language.SG:      "सांगो",
			language.SH:      "सेर्बो-क्रोयेषियन्",
			language.SHI:     "ताछेहीट",
			language.SHN:     "शान",
			language.SI:      "सिनहालीस",
			language.SK:      "स्लोवाक",
			language.SL:      "स्लोवानियन",
			language.SM:      "सामोअन",
			language.SMA:     "दक्षिणी सामी",
			language.SMJ:     "लुले सामी",
			language.SMN:     "ईनारी सामी",
			language.SMS:     "स्कोल्ट सामी",
			language.SN:      "शोना",
			language.SNK:     "सोनिके",
			language.SO:      "सोमाली",
			language.SQ:      "आल्बेनियन्",
			language.SR:      "सर्बियन",
			language.SRN:     "श्रानन टोंगो",
			language.SS:      "स्वाती",
			language.SSY:     "साहो",
			language.ST:      "सावथर्न सोथो",
			language.SU:      "सुंदनीस",
			language.SUK:     "सुकुमा",
			language.SV:      "स्विडीश",
			language.SW:      "स्वाहिली",
			language.SW_CD:   "काँगो स्वाहिली",
			language.SWB:     "कोमोरियन",
			language.SYR:     "सिरियाक",
			language.TA:      "तमिळ",
			language.TE:      "तेलुगू",
			language.TEM:     "तिम्ने",
			language.TEO:     "तेसो",
			language.TET:     "तेतम",
			language.TG:      "ताजिक",
			language.TH:      "थाई",
			language.TI:      "तिग्रिन्या",
			language.TIG:     "टिग्रे",
			language.TK:      "तुर्कमेन",
			language.TL:      "तगालोग",
			language.TLH:     "क्लिंगन",
			language.TN:      "सेत्स्वाना",
			language.TO:      "तोंगा",
			language.TPI:     "तोक पिसीन",
			language.TR:      "तुर्किष",
			language.TRV:     "तारोको",
			language.TS:      "त्सोगा",
			language.TT:      "तटार",
			language.TUM:     "तुंबुका",
			language.TVL:     "तुवालू",
			language.TW:      "त्वि",
			language.TWQ:     "तासावाक",
			language.TY:      "ताहीशियन",
			language.TYV:     "तुविनियन",
			language.TZM:     "केंद्रीय अटलास तामाझायट",
			language.UDM:     "उडमुर्त",
			language.UG:      "उयघूर",
			language.UK:      "युक्रेनियन्",
			language.UMB:     "यमबुंडु",
			language.UND:     "अज्ञात भास",
			language.UR:      "उर्दू",
			language.UZ:      "उझबेक",
			language.VAI:     "वाई",
			language.VE:      "वेंदा",
			language.VI:      "वियत्नामीज़",
			language.VO:      "ओलापुक",
			language.VUN:     "वुंजो",
			language.WA:      "वालून",
			language.WAE:     "वाल्सर",
			language.WAL:     "वोलायटा",
			language.WAR:     "वरय",
			language.WO:      "उलोफ़",
			language.XAL:     "कालमायक",
			language.XH:      "झ़ौसा",
			language.XOG:     "सोगा",
			language.YAV:     "यांगबेन",
			language.YBB:     "येम्बा",
			language.YI:      "इद्दिष्",
			language.YO:      "यूरुबा",
			language.YUE:     "चिनी, कॅण्टोनीस",
			language.ZA:      "झ्हुन्ग",
			language.ZGH:     "प्रमाणीत मॉरोक्कन तमाझीट",
			language.ZH:      "चिनी, मंडारीन",
			language.ZH_HANS: "सोंपी मंडारीन चिनी",
			language.ZH_HANT: "पारंपारीक मंडारीन चिनी",
			language.ZU:      "जुलू",
			language.ZUN:     "झून",
			language.ZXX:     "अणकार सामुग्री ना",
			language.ZZA:     "झाझा",
		},
		Territories: cldr.Territories{
			territory.V_001: "जग",
			territory.V_002: "आफ्रिका",
			territory.V_003: "उत्तर अमेरिका",
			territory.V_005: "दक्षिण अमेरिका",
			territory.V_009: "ओसेनिया",
			territory.V_011: "अस्तंत आफ्रिका",
			territory.V_013: "मध्य अमेरिका",
			territory.V_014: "उदेंत आफ्रिका",
			territory.V_015: "उत्तरीय आफ्रिका",
			territory.V_017: "मध्य आफ्रिका",
			territory.V_018: "दक्षिण आफ्रिका",
			territory.V_019: "अमेरिकास",
			territory.V_021: "उत्तरीय अमेरिका",
			territory.V_029: "कॅरिबियन",
			territory.V_030: "उदेंत आशिया",
			territory.V_034: "दक्षिण आशिया",
			territory.V_035: "आग्नेय आशिया",
			territory.V_039: "दक्षिण येवरोप",
			territory.V_053: "ऑस्ट्रेलेसिया",
			territory.V_054: "मेलानेसिया",
			territory.V_057: "मायक्रोनेशियन प्रांत",
			territory.V_061: "पोलिनेशिया",
			territory.V_142: "आशिया",
			territory.V_143: "मध्य आशिया",
			territory.V_145: "अस्तंत आशिया",
			territory.V_150: "येवरोप",
			territory.V_151: "उदेंत येवरोप",
			territory.V_154: "उत्तर येवरोप",
			territory.V_155: "अस्तंत येवरोप",
			territory.V_202: "उप-सहाराई आफ्रिका",
			territory.V_419: "लॅटीन अमेरिका",
			territory.AC:    "असेशन आयलँड",
			territory.AD:    "अंडोरा",
			territory.AE:    "युनाइटेड अरब इमीरॅट्स",
			territory.AF:    "अफगानिस्तान",
			territory.AG:    "एँटिगुआ आनी बारबुडा",
			territory.AI:    "अंगुला",
			territory.AL:    "अल्बानीया",
			territory.AM:    "आर्मीनीया",
			territory.AO:    "अंगोला",
			territory.AQ:    "अंटार्क्टिका",
			territory.AR:    "अर्जेंटिना",
			territory.AS:    "अमेरिकी सामोआ",
			territory.AT:    "ऑस्ट्रिया",
			territory.AU:    "ऑस्ट्रेलीया",
			territory.AW:    "अरुबा",
			territory.AX:    "अलांड जुवे",
			territory.AZ:    "अजरबैजान",
			territory.BA:    "बोस्निया आनी हेर्जेगोविना",
			territory.BB:    "बारबाडोस",
			territory.BD:    "बांगलादेश",
			territory.BE:    "बेल्जियम",
			territory.BF:    "बुर्किना फॅसो",
			territory.BG:    "बल्गेरीया",
			territory.BH:    "बेहरेन",
			territory.BI:    "बुरुंडी",
			territory.BJ:    "बेनीन",
			territory.BL:    "सॅंट बार्थेल्मी",
			territory.BM:    "बर्मुडा",
			territory.BN:    "ब्रूनेई",
			territory.BO:    "बोलिव्हिया",
			territory.BQ:    "कॅरिबियन निदरलँड",
			territory.BR:    "ब्राझील",
			territory.BS:    "बहामास",
			territory.BT:    "भूतान",
			territory.BV:    "बोवट आयलँड",
			territory.BW:    "बोत्सवाना",
			territory.BY:    "बेलारूस",
			territory.BZ:    "बेलिझ",
			territory.CA:    "कॅनडा",
			territory.CC:    "कोकोस (कीलिंग) आयलँड",
			territory.CD:    "कोंगो - किंशासा",
			territory.CF:    "मध्य अफ्रीकी लोकसत्तकराज्य",
			territory.CG:    "काँगो - ब्राझाविला",
			territory.CH:    "स्विट्ज़रलैंड",
			territory.CI:    "कोत द’ईवोआर",
			territory.CK:    "कुक आयलँड्स",
			territory.CL:    "चिली",
			territory.CM:    "कॅमेरून",
			territory.CN:    "चीन",
			territory.CO:    "कोलंबिया",
			territory.CP:    "क्लिपरटॉन आयलँड",
			territory.CR:    "कोस्ता रिका",
			territory.CU:    "क्युबा",
			territory.CV:    "केप वर्दी",
			territory.CW:    "कुरसावो",
			territory.CX:    "क्रिसमस आयलँड",
			territory.CY:    "सायप्रस",
			territory.CZ:    "चेकिया",
			territory.DE:    "जर्मनी",
			territory.DG:    "दिगो गार्सिया",
			territory.DJ:    "जिबूती",
			territory.DK:    "डेनमार्क",
			territory.DM:    "डोमिनीका",
			territory.DO:    "डोमिनिकन प्रजासत्ताक",
			territory.DZ:    "अल्जेरिया",
			territory.EA:    "सिटा आनी मेलिल्ला",
			territory.EC:    "इक्वाडोर",
			territory.EE:    "एस्टोनिया",
			territory.EG:    "ईजिप्त",
			territory.EH:    "अस्तंत सहारा",
			territory.ER:    "इरिट्रिया",
			territory.ES:    "स्पेन",
			territory.ET:    "इथियोपिया",
			territory.EU:    "युरोपियन युनियन",
			territory.EZ:    "युरोझोन",
			territory.FI:    "फिनलँड",
			territory.FJ:    "फिजी",
			territory.FK:    "फ़ॉकलैंड आइलैंड्स",
			territory.FM:    "मायक्रोनेशिया",
			territory.FO:    "फैरो आयलँड्स",
			territory.FR:    "फ्रान्स",
			territory.GA:    "गॅबोन",
			territory.GB:    "युनायटेड किंगडम",
			territory.GD:    "ग्रेनॅडा",
			territory.GE:    "जॉर्जिया",
			territory.GF:    "फ्रेन्च गयाना",
			territory.GG:    "गर्नसी",
			territory.GH:    "घाना",
			territory.GI:    "जिब्राल्टर",
			territory.GL:    "ग्रीनलँड",
			territory.GM:    "गॅम्बिया",
			territory.GN:    "गुएनिया",
			territory.GP:    "ग्वाडेलोप",
			territory.GQ:    "इक्वेटोरियल गुएनिया",
			territory.GR:    "ग्रीस",
			territory.GS:    "दक्षिण जोर्जिया आनी दक्षिण सॅण्डविच आयलँड्स",
			territory.GT:    "ग्वाटेमाला",
			territory.GU:    "गुआम",
			territory.GW:    "गुअनिया-बिसाउ",
			territory.GY:    "गयाना",
			territory.HK:    "हाँग काँग SAR चीन",
			territory.HM:    "हर्ड आयलँड्स ऍंड मॅक्डोनाल्ड आयलँड्स",
			territory.HN:    "हॉनडुरस",
			territory.HR:    "क्रोयेशीया",
			territory.HT:    "हैती",
			territory.HU:    "हंगेरी",
			territory.IC:    "कॅनरी आयलैंड्स",
			territory.ID:    "इंडोनेशीया",
			territory.IE:    "आयरलँड",
			territory.IL:    "इज़राइल",
			territory.IM:    "इसले ऑफ मॅन",
			territory.IN:    "भारत",
			territory.IO:    "ब्रिटिश हिंद महासागरीय क्षेत्र",
			territory.IQ:    "इराक",
			territory.IR:    "इरान",
			territory.IS:    "आइसलैंड",
			territory.IT:    "इटली",
			territory.JE:    "जर्सी",
			territory.JM:    "जमैका",
			territory.JO:    "जॉर्डन",
			territory.JP:    "जपान",
			territory.KE:    "केनया",
			territory.KG:    "किर्गिज़स्तान",
			territory.KH:    "कंबोडिया",
			territory.KI:    "किरिबाती",
			territory.KM:    "कोमोरोस",
			territory.KN:    "सेंट किट्स आनी नेविस",
			territory.KP:    "उत्तर कोरिया",
			territory.KR:    "दक्षिण कोरिया",
			territory.KW:    "कुवेत",
			territory.KY:    "कैमेन आइलैंड्स",
			territory.KZ:    "कझाकस्तान",
			territory.LA:    "लाओस",
			territory.LB:    "लेबनान",
			territory.LC:    "सँट लुसिया",
			territory.LI:    "लिचेंस्टीन",
			territory.LK:    "श्री लंका",
			territory.LR:    "लायबेरीया",
			territory.LS:    "लिसोथो",
			territory.LT:    "लिथुआनिया",
			territory.LU:    "लक्सेमबर्ग",
			territory.LV:    "लॅटविया",
			territory.LY:    "लीबिया",
			territory.MA:    "मोरोक्को",
			territory.MC:    "मोनॅको",
			territory.MD:    "माल्डोवा",
			territory.ME:    "मॉन्टॅनग्रो",
			territory.MF:    "सॅंट मार्टिन",
			territory.MG:    "माडागास्कर",
			territory.MH:    "मार्शल आयलँड्स",
			territory.MK:    "उत्तर मॅसिडोनिया",
			territory.ML:    "माली",
			territory.MM:    "म्यानमार (बर्मा)",
			territory.MN:    "मंगोलिया",
			territory.MO:    "मकाव SAR चीन",
			territory.MP:    "उत्तरी मरिना आयसलैण्ड",
			territory.MQ:    "मार्टीनिक",
			territory.MR:    "मॉरिटानिया",
			territory.MS:    "मॉन्टसेराट",
			territory.MT:    "माल्टा",
			territory.MU:    "मॉरिशस",
			territory.MV:    "मालदीव",
			territory.MW:    "मलावी",
			territory.MX:    "मेक्सिको",
			territory.MY:    "मलेशिया",
			territory.MZ:    "मॉझांबीक",
			territory.NA:    "नामीबिया",
			territory.NC:    "न्यू कॅलिडोनिया",
			territory.NE:    "नायजर",
			territory.NF:    "नॉरफॉक आयलँड",
			territory.NG:    "नायजेरिया",
			territory.NI:    "निकारगुवा",
			territory.NL:    "नॅदरलँड",
			territory.NO:    "नॉर्वे",
			territory.NP:    "नेपाळ",
			territory.NR:    "नावरू",
			territory.NU:    "नीयू",
			territory.NZ:    "न्युझीलॅन्ड",
			territory.OM:    "ओमान",
			territory.PA:    "पनामा",
			territory.PE:    "पेरू",
			territory.PF:    "फ्रेन्च पोलिनेसिया",
			territory.PG:    "पापुआ न्यु गिनी",
			territory.PH:    "फिलीपिन्झ",
			territory.PK:    "पाकिस्तान",
			territory.PL:    "पोलंड",
			territory.PM:    "सँ. पायरे आनी मिकेलन",
			territory.PN:    "पिटकॅरन आयलँड्स",
			territory.PR:    "पिर्टो रिको",
			territory.PS:    "पेलेस्टीनियन प्रांत",
			territory.PT:    "पुर्तगाल",
			territory.PW:    "पलाऊ",
			territory.PY:    "पैराग्वे",
			territory.QA:    "कतार",
			territory.QO:    "आवटलायींग ओशेनिया",
			territory.RE:    "रीयूनियन",
			territory.RO:    "रोमानीया",
			territory.RS:    "सर्बिया",
			territory.RU:    "रूस",
			territory.RW:    "रवांडा",
			territory.SA:    "सऊदी अरेबिया",
			territory.SB:    "सोलोमन आइलँड्स",
			territory.SC:    "सेशेल्स",
			territory.SD:    "सूडान",
			territory.SE:    "स्वीडन",
			territory.SG:    "सिंगापूर",
			territory.SH:    "सेंट हेलिना",
			territory.SI:    "स्लोवेनिया",
			territory.SJ:    "स्वालबार्ड आनी जान मेयन",
			territory.SK:    "स्लोवाकिया",
			territory.SL:    "सिएरा लियॉन",
			territory.SM:    "सॅन मारीनो",
			territory.SN:    "सिनिगल",
			territory.SO:    "सोमालिया",
			territory.SR:    "सुरीनाम",
			territory.SS:    "दक्षिण सुडान",
			territory.ST:    "साओ टोम आनी प्रिन्सिप",
			territory.SV:    "एल साल्वाडोर",
			territory.SX:    "सिंट मार्टेन",
			territory.SY:    "सिरिया",
			territory.SZ:    "इस्वातिनी",
			territory.TA:    "त्रिस्तान दा कुन्हा",
			territory.TC:    "तुर्क्स आनी कॅकोज आयलँड्स",
			territory.TD:    "चाड",
			territory.TF:    "फ्रेंच दक्षिणी प्रांत",
			territory.TG:    "टोगो",
			territory.TH:    "थायलँड",
			territory.TJ:    "तजीकिस्तान",
			territory.TK:    "टोकलाऊ",
			territory.TL:    "तिमोर-लेस्ते",
			territory.TM:    "तुर्कमेनिस्तान",
			territory.TN:    "ट्यूनीशिया",
			territory.TO:    "टोंगा",
			territory.TR:    "तुर्की",
			territory.TT:    "ट्रिनीडाड आनी टोबॅगो",
			territory.TV:    "टुवालू",
			territory.TW:    "तायवान",
			territory.TZ:    "तांझानिया",
			territory.UA:    "युक्रेन",
			territory.UG:    "युगांडा",
			territory.UM:    "यु. एस. मायनर आवटलायींग आयलँड्\u200dस",
			territory.UN:    "युनायटेड नेशन्स",
			territory.US:    "युनायटेड स्टेट्स",
			territory.UY:    "उरूग्वे",
			territory.UZ:    "उज़्बेकिस्तान",
			territory.VA:    "वॅटिकन सिटी",
			territory.VC:    "सेंट विंसेंट ऐंड द ग्रेनेडाइंस",
			territory.VE:    "विनेझुएला",
			territory.VG:    "ब्रिटिश वर्जिन आयलँड्स",
			territory.VI:    "यु. एस. वर्जिन आयलँड्\u200dस",
			territory.VN:    "व्हिएतनाम",
			territory.VU:    "वनातू",
			territory.WF:    "वालिस आनी फ्यूचूना",
			territory.WS:    "सामोआ",
			territory.XA:    "स्युडो-ऍक्सेंट",
			territory.XB:    "स्युडो-बिडी",
			territory.XK:    "कोसोवो",
			territory.YE:    "येमेन",
			territory.YT:    "मेयोट",
			territory.ZA:    "दक्षिण आफ्रीका",
			territory.ZM:    "झांबिया",
			territory.ZW:    "जिम्बाब्वे",
			territory.ZZ:    "अज्ञात प्रांत",
		},
	}
}
