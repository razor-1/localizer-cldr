// Code generated by make_resources.go; DO NOT EDIT.
package resources

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"

	"github.com/razor-1/localizer-cldr"
)

type pluralForms []plural.Form
type pluralFunc func(ops *cldr.PluralOperands) plural.Form

func cardinal_0() cldr.PluralData {
	// for locales: [bm bo dz id ig ii in ja jbo jv jw kde kea km ko lkt lo ms my nqo osa root sah ses sg su th to vi wo yo yue zh]
	forms := pluralForms{
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_1() cldr.PluralData {
	// for locales: [am as bn fa gu hi kn pcm zu]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// i = 0 or n = 1
		if intEqualsAny(ops.I, 0) ||
			ops.NEqualsAny(1) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_2() cldr.PluralData {
	// for locales: [ff fr hy kab]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// i = 0,1
		if intEqualsAny(ops.I, 0, 1) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_3() cldr.PluralData {
	// for locales: [pt]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// i = 0..1
		if intInRange(ops.I, 0, 1) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_4() cldr.PluralData {
	// for locales: [ast ca de en et fi fy gl ia io it ji nl pt_PT sc scn sv sw ur yi]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// i = 1 and v = 0
		if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_5() cldr.PluralData {
	// for locales: [si]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 0,1 or i = 0 and f = 1
		if ops.NEqualsAny(0, 1) ||
			intEqualsAny(ops.I, 0) && intEqualsAny(ops.F, 1) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_6() cldr.PluralData {
	// for locales: [ak bho guw ln mg nso pa ti wa]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 0..1
		if ops.NInRange(0, 1) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_7() cldr.PluralData {
	// for locales: [tzm]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 0..1 or n = 11..99
		if ops.NInRange(0, 1) ||
			ops.NInRange(11, 99) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_8() cldr.PluralData {
	// for locales: [af an asa az bem bez bg brx ce cgg chr ckb dv ee el eo es eu fo fur gsw ha haw hu jgo jmc ka kaj kcg kk kkj kl ks ksb ku ky lb lg mas mgo ml mn mr nah nb nd ne nn nnh no nr ny nyn om or os pap ps rm rof rwk saq sd sdh seh sn so sq ss ssy st syr ta te teo tig tk tn tr ts ug uz ve vo vun wae xh xog]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 1
		if ops.NEqualsAny(1) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_9() cldr.PluralData {
	// for locales: [da]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 1 or t != 0 and i = 0,1
		if ops.NEqualsAny(1) ||
			!intEqualsAny(ops.T, 0) && intEqualsAny(ops.I, 0, 1) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_10() cldr.PluralData {
	// for locales: [is]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// t = 0 and i % 10 = 1 and i % 100 != 11 or t != 0
		if intEqualsAny(ops.T, 0) && intEqualsAny(ops.I%10, 1) && !intEqualsAny(ops.I%100, 11) ||
			!intEqualsAny(ops.T, 0) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_11() cldr.PluralData {
	// for locales: [mk]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// v = 0 and i % 10 = 1 and i % 100 != 11 or f % 10 = 1 and f % 100 != 11
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 1) && !intEqualsAny(ops.I%100, 11) ||
			intEqualsAny(ops.F%10, 1) && !intEqualsAny(ops.F%100, 11) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_12() cldr.PluralData {
	// for locales: [ceb fil tl]
	forms := pluralForms{
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// v = 0 and i = 1,2,3 or v = 0 and i % 10 != 4,6,9 or v != 0 and f % 10 != 4,6,9
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I, 1, 2, 3) ||
			intEqualsAny(ops.V, 0) && !intEqualsAny(ops.I%10, 4, 6, 9) ||
			!intEqualsAny(ops.V, 0) && !intEqualsAny(ops.F%10, 4, 6, 9) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_13() cldr.PluralData {
	// for locales: [lv prg]
	forms := pluralForms{
		plural.Zero,
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n % 10 = 0 or n % 100 = 11..19 or v = 2 and f % 100 = 11..19
		if ops.NModEqualsAny(10, 0) ||
			ops.NModInRange(100, 11, 19) ||
			intEqualsAny(ops.V, 2) && intInRange(ops.F%100, 11, 19) {
			return plural.Zero
		}
		// n % 10 = 1 and n % 100 != 11 or v = 2 and f % 10 = 1 and f % 100 != 11 or v != 2 and f % 10 = 1
		if ops.NModEqualsAny(10, 1) && !ops.NModEqualsAny(100, 11) ||
			intEqualsAny(ops.V, 2) && intEqualsAny(ops.F%10, 1) && !intEqualsAny(ops.F%100, 11) ||
			!intEqualsAny(ops.V, 2) && intEqualsAny(ops.F%10, 1) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_14() cldr.PluralData {
	// for locales: [lag]
	forms := pluralForms{
		plural.Zero,
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 0
		if ops.NEqualsAny(0) {
			return plural.Zero
		}
		// i = 0,1 and n != 0
		if intEqualsAny(ops.I, 0, 1) && !ops.NEqualsAny(0) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_15() cldr.PluralData {
	// for locales: [ksh]
	forms := pluralForms{
		plural.Zero,
		plural.One,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 0
		if ops.NEqualsAny(0) {
			return plural.Zero
		}
		// n = 1
		if ops.NEqualsAny(1) {
			return plural.One
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_16() cldr.PluralData {
	// for locales: [iu naq sat se sma smi smj smn sms]
	forms := pluralForms{
		plural.One,
		plural.Two,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 1
		if ops.NEqualsAny(1) {
			return plural.One
		}
		// n = 2
		if ops.NEqualsAny(2) {
			return plural.Two
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_17() cldr.PluralData {
	// for locales: [shi]
	forms := pluralForms{
		plural.One,
		plural.Few,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// i = 0 or n = 1
		if intEqualsAny(ops.I, 0) ||
			ops.NEqualsAny(1) {
			return plural.One
		}
		// n = 2..10
		if ops.NInRange(2, 10) {
			return plural.Few
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_18() cldr.PluralData {
	// for locales: [mo ro]
	forms := pluralForms{
		plural.One,
		plural.Few,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// i = 1 and v = 0
		if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
			return plural.One
		}
		// v != 0 or n = 0 or n % 100 = 2..19
		if !intEqualsAny(ops.V, 0) ||
			ops.NEqualsAny(0) ||
			ops.NModInRange(100, 2, 19) {
			return plural.Few
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_19() cldr.PluralData {
	// for locales: [bs hr sh sr]
	forms := pluralForms{
		plural.One,
		plural.Few,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// v = 0 and i % 10 = 1 and i % 100 != 11 or f % 10 = 1 and f % 100 != 11
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 1) && !intEqualsAny(ops.I%100, 11) ||
			intEqualsAny(ops.F%10, 1) && !intEqualsAny(ops.F%100, 11) {
			return plural.One
		}
		// v = 0 and i % 10 = 2..4 and i % 100 != 12..14 or f % 10 = 2..4 and f % 100 != 12..14
		if intEqualsAny(ops.V, 0) && intInRange(ops.I%10, 2, 4) && !intInRange(ops.I%100, 12, 14) ||
			intInRange(ops.F%10, 2, 4) && !intInRange(ops.F%100, 12, 14) {
			return plural.Few
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_20() cldr.PluralData {
	// for locales: [gd]
	forms := pluralForms{
		plural.One,
		plural.Two,
		plural.Few,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 1,11
		if ops.NEqualsAny(1, 11) {
			return plural.One
		}
		// n = 2,12
		if ops.NEqualsAny(2, 12) {
			return plural.Two
		}
		// n = 3..10,13..19
		if ops.NInRange(3, 10) || ops.NInRange(13, 19) {
			return plural.Few
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_21() cldr.PluralData {
	// for locales: [sl]
	forms := pluralForms{
		plural.One,
		plural.Two,
		plural.Few,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// v = 0 and i % 100 = 1
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%100, 1) {
			return plural.One
		}
		// v = 0 and i % 100 = 2
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%100, 2) {
			return plural.Two
		}
		// v = 0 and i % 100 = 3..4 or v != 0
		if intEqualsAny(ops.V, 0) && intInRange(ops.I%100, 3, 4) ||
			!intEqualsAny(ops.V, 0) {
			return plural.Few
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_22() cldr.PluralData {
	// for locales: [dsb hsb]
	forms := pluralForms{
		plural.One,
		plural.Two,
		plural.Few,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// v = 0 and i % 100 = 1 or f % 100 = 1
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%100, 1) ||
			intEqualsAny(ops.F%100, 1) {
			return plural.One
		}
		// v = 0 and i % 100 = 2 or f % 100 = 2
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%100, 2) ||
			intEqualsAny(ops.F%100, 2) {
			return plural.Two
		}
		// v = 0 and i % 100 = 3..4 or f % 100 = 3..4
		if intEqualsAny(ops.V, 0) && intInRange(ops.I%100, 3, 4) ||
			intInRange(ops.F%100, 3, 4) {
			return plural.Few
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_23() cldr.PluralData {
	// for locales: [he iw]
	forms := pluralForms{
		plural.One,
		plural.Two,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// i = 1 and v = 0
		if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
			return plural.One
		}
		// i = 2 and v = 0
		if intEqualsAny(ops.I, 2) && intEqualsAny(ops.V, 0) {
			return plural.Two
		}
		// v = 0 and n != 0..10 and n % 10 = 0
		if intEqualsAny(ops.V, 0) && !ops.NInRange(0, 10) && ops.NModEqualsAny(10, 0) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_24() cldr.PluralData {
	// for locales: [cs sk]
	forms := pluralForms{
		plural.One,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// i = 1 and v = 0
		if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
			return plural.One
		}
		// i = 2..4 and v = 0
		if intInRange(ops.I, 2, 4) && intEqualsAny(ops.V, 0) {
			return plural.Few
		}
		// v != 0
		if !intEqualsAny(ops.V, 0) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_25() cldr.PluralData {
	// for locales: [pl]
	forms := pluralForms{
		plural.One,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// i = 1 and v = 0
		if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
			return plural.One
		}
		// v = 0 and i % 10 = 2..4 and i % 100 != 12..14
		if intEqualsAny(ops.V, 0) && intInRange(ops.I%10, 2, 4) && !intInRange(ops.I%100, 12, 14) {
			return plural.Few
		}
		// v = 0 and i != 1 and i % 10 = 0..1 or v = 0 and i % 10 = 5..9 or v = 0 and i % 100 = 12..14
		if intEqualsAny(ops.V, 0) && !intEqualsAny(ops.I, 1) && intInRange(ops.I%10, 0, 1) ||
			intEqualsAny(ops.V, 0) && intInRange(ops.I%10, 5, 9) ||
			intEqualsAny(ops.V, 0) && intInRange(ops.I%100, 12, 14) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_26() cldr.PluralData {
	// for locales: [be]
	forms := pluralForms{
		plural.One,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n % 10 = 1 and n % 100 != 11
		if ops.NModEqualsAny(10, 1) && !ops.NModEqualsAny(100, 11) {
			return plural.One
		}
		// n % 10 = 2..4 and n % 100 != 12..14
		if ops.NModInRange(10, 2, 4) && !ops.NModInRange(100, 12, 14) {
			return plural.Few
		}
		// n % 10 = 0 or n % 10 = 5..9 or n % 100 = 11..14
		if ops.NModEqualsAny(10, 0) ||
			ops.NModInRange(10, 5, 9) ||
			ops.NModInRange(100, 11, 14) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_27() cldr.PluralData {
	// for locales: [lt]
	forms := pluralForms{
		plural.One,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n % 10 = 1 and n % 100 != 11..19
		if ops.NModEqualsAny(10, 1) && !ops.NModInRange(100, 11, 19) {
			return plural.One
		}
		// n % 10 = 2..9 and n % 100 != 11..19
		if ops.NModInRange(10, 2, 9) && !ops.NModInRange(100, 11, 19) {
			return plural.Few
		}
		// f != 0
		if !intEqualsAny(ops.F, 0) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_28() cldr.PluralData {
	// for locales: [mt]
	forms := pluralForms{
		plural.One,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 1
		if ops.NEqualsAny(1) {
			return plural.One
		}
		// n = 0 or n % 100 = 2..10
		if ops.NEqualsAny(0) ||
			ops.NModInRange(100, 2, 10) {
			return plural.Few
		}
		// n % 100 = 11..19
		if ops.NModInRange(100, 11, 19) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_29() cldr.PluralData {
	// for locales: [ru uk]
	forms := pluralForms{
		plural.One,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// v = 0 and i % 10 = 1 and i % 100 != 11
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 1) && !intEqualsAny(ops.I%100, 11) {
			return plural.One
		}
		// v = 0 and i % 10 = 2..4 and i % 100 != 12..14
		if intEqualsAny(ops.V, 0) && intInRange(ops.I%10, 2, 4) && !intInRange(ops.I%100, 12, 14) {
			return plural.Few
		}
		// v = 0 and i % 10 = 0 or v = 0 and i % 10 = 5..9 or v = 0 and i % 100 = 11..14
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 0) ||
			intEqualsAny(ops.V, 0) && intInRange(ops.I%10, 5, 9) ||
			intEqualsAny(ops.V, 0) && intInRange(ops.I%100, 11, 14) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_30() cldr.PluralData {
	// for locales: [br]
	forms := pluralForms{
		plural.One,
		plural.Two,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n % 10 = 1 and n % 100 != 11,71,91
		if ops.NModEqualsAny(10, 1) && !ops.NModEqualsAny(100, 11, 71, 91) {
			return plural.One
		}
		// n % 10 = 2 and n % 100 != 12,72,92
		if ops.NModEqualsAny(10, 2) && !ops.NModEqualsAny(100, 12, 72, 92) {
			return plural.Two
		}
		// n % 10 = 3..4,9 and n % 100 != 10..19,70..79,90..99
		if (ops.NModInRange(10, 3, 4) || ops.NModEqualsAny(10, 9)) && !(ops.NModInRange(100, 10, 19) || ops.NModInRange(100, 70, 79) || ops.NModInRange(100, 90, 99)) {
			return plural.Few
		}
		// n != 0 and n % 1000000 = 0
		if !ops.NEqualsAny(0) && ops.NModEqualsAny(1000000, 0) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_31() cldr.PluralData {
	// for locales: [ga]
	forms := pluralForms{
		plural.One,
		plural.Two,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 1
		if ops.NEqualsAny(1) {
			return plural.One
		}
		// n = 2
		if ops.NEqualsAny(2) {
			return plural.Two
		}
		// n = 3..6
		if ops.NInRange(3, 6) {
			return plural.Few
		}
		// n = 7..10
		if ops.NInRange(7, 10) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_32() cldr.PluralData {
	// for locales: [gv]
	forms := pluralForms{
		plural.One,
		plural.Two,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// v = 0 and i % 10 = 1
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 1) {
			return plural.One
		}
		// v = 0 and i % 10 = 2
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 2) {
			return plural.Two
		}
		// v = 0 and i % 100 = 0,20,40,60,80
		if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%100, 0, 20, 40, 60, 80) {
			return plural.Few
		}
		// v != 0
		if !intEqualsAny(ops.V, 0) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_33() cldr.PluralData {
	// for locales: [kw]
	forms := pluralForms{
		plural.Zero,
		plural.One,
		plural.Two,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 0
		if ops.NEqualsAny(0) {
			return plural.Zero
		}
		// n = 1
		if ops.NEqualsAny(1) {
			return plural.One
		}
		// n % 100 = 2,22,42,62,82 or n % 1000 = 0 and n % 100000 = 1000..20000,40000,60000,80000 or n != 0 and n % 1000000 = 100000
		if ops.NModEqualsAny(100, 2, 22, 42, 62, 82) ||
			ops.NModEqualsAny(1000, 0) && (ops.NModInRange(100000, 1000, 20000) || ops.NModEqualsAny(100000, 40000, 60000, 80000)) ||
			!ops.NEqualsAny(0) && ops.NModEqualsAny(1000000, 100000) {
			return plural.Two
		}
		// n % 100 = 3,23,43,63,83
		if ops.NModEqualsAny(100, 3, 23, 43, 63, 83) {
			return plural.Few
		}
		// n != 1 and n % 100 = 1,21,41,61,81
		if !ops.NEqualsAny(1) && ops.NModEqualsAny(100, 1, 21, 41, 61, 81) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_34() cldr.PluralData {
	// for locales: [ar ars]
	forms := pluralForms{
		plural.Zero,
		plural.One,
		plural.Two,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 0
		if ops.NEqualsAny(0) {
			return plural.Zero
		}
		// n = 1
		if ops.NEqualsAny(1) {
			return plural.One
		}
		// n = 2
		if ops.NEqualsAny(2) {
			return plural.Two
		}
		// n % 100 = 3..10
		if ops.NModInRange(100, 3, 10) {
			return plural.Few
		}
		// n % 100 = 11..99
		if ops.NModInRange(100, 11, 99) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

func cardinal_35() cldr.PluralData {
	// for locales: [cy]
	forms := pluralForms{
		plural.Zero,
		plural.One,
		plural.Two,
		plural.Few,
		plural.Many,
		plural.Other,
	}
	plFunc := func(ops *cldr.PluralOperands) plural.Form {
		// n = 0
		if ops.NEqualsAny(0) {
			return plural.Zero
		}
		// n = 1
		if ops.NEqualsAny(1) {
			return plural.One
		}
		// n = 2
		if ops.NEqualsAny(2) {
			return plural.Two
		}
		// n = 3
		if ops.NEqualsAny(3) {
			return plural.Few
		}
		// n = 6
		if ops.NEqualsAny(6) {
			return plural.Many
		}
		return plural.Other
	}

	return cldr.PluralData{Forms: forms, Func: plFunc}
}

var cardinalPlural = map[language.Tag]func() cldr.PluralData{
	tag_af:    cardinal_8,
	tag_ak:    cardinal_6,
	tag_am:    cardinal_1,
	tag_an:    cardinal_8,
	tag_ar:    cardinal_34,
	tag_ars:   cardinal_34,
	tag_as:    cardinal_1,
	tag_asa:   cardinal_8,
	tag_ast:   cardinal_4,
	tag_az:    cardinal_8,
	tag_be:    cardinal_26,
	tag_bem:   cardinal_8,
	tag_bez:   cardinal_8,
	tag_bg:    cardinal_8,
	tag_bho:   cardinal_6,
	tag_bm:    cardinal_0,
	tag_bn:    cardinal_1,
	tag_bo:    cardinal_0,
	tag_br:    cardinal_30,
	tag_brx:   cardinal_8,
	tag_bs:    cardinal_19,
	tag_ca:    cardinal_4,
	tag_ce:    cardinal_8,
	tag_ceb:   cardinal_12,
	tag_cgg:   cardinal_8,
	tag_chr:   cardinal_8,
	tag_ckb:   cardinal_8,
	tag_cs:    cardinal_24,
	tag_cy:    cardinal_35,
	tag_da:    cardinal_9,
	tag_de:    cardinal_4,
	tag_dsb:   cardinal_22,
	tag_dv:    cardinal_8,
	tag_dz:    cardinal_0,
	tag_ee:    cardinal_8,
	tag_el:    cardinal_8,
	tag_en:    cardinal_4,
	tag_eo:    cardinal_8,
	tag_es:    cardinal_8,
	tag_et:    cardinal_4,
	tag_eu:    cardinal_8,
	tag_fa:    cardinal_1,
	tag_ff:    cardinal_2,
	tag_fi:    cardinal_4,
	tag_fil:   cardinal_12,
	tag_fo:    cardinal_8,
	tag_fr:    cardinal_2,
	tag_fur:   cardinal_8,
	tag_fy:    cardinal_4,
	tag_ga:    cardinal_31,
	tag_gd:    cardinal_20,
	tag_gl:    cardinal_4,
	tag_gsw:   cardinal_8,
	tag_gu:    cardinal_1,
	tag_guw:   cardinal_6,
	tag_gv:    cardinal_32,
	tag_ha:    cardinal_8,
	tag_haw:   cardinal_8,
	tag_he:    cardinal_23,
	tag_hi:    cardinal_1,
	tag_hr:    cardinal_19,
	tag_hsb:   cardinal_22,
	tag_hu:    cardinal_8,
	tag_hy:    cardinal_2,
	tag_ia:    cardinal_4,
	tag_id:    cardinal_0,
	tag_ig:    cardinal_0,
	tag_ii:    cardinal_0,
	tag_in:    cardinal_0,
	tag_io:    cardinal_4,
	tag_is:    cardinal_10,
	tag_it:    cardinal_4,
	tag_iu:    cardinal_16,
	tag_iw:    cardinal_23,
	tag_ja:    cardinal_0,
	tag_jbo:   cardinal_0,
	tag_jgo:   cardinal_8,
	tag_ji:    cardinal_4,
	tag_jmc:   cardinal_8,
	tag_jv:    cardinal_0,
	tag_jw:    cardinal_0,
	tag_ka:    cardinal_8,
	tag_kab:   cardinal_2,
	tag_kaj:   cardinal_8,
	tag_kcg:   cardinal_8,
	tag_kde:   cardinal_0,
	tag_kea:   cardinal_0,
	tag_kk:    cardinal_8,
	tag_kkj:   cardinal_8,
	tag_kl:    cardinal_8,
	tag_km:    cardinal_0,
	tag_kn:    cardinal_1,
	tag_ko:    cardinal_0,
	tag_ks:    cardinal_8,
	tag_ksb:   cardinal_8,
	tag_ksh:   cardinal_15,
	tag_ku:    cardinal_8,
	tag_kw:    cardinal_33,
	tag_ky:    cardinal_8,
	tag_lag:   cardinal_14,
	tag_lb:    cardinal_8,
	tag_lg:    cardinal_8,
	tag_lkt:   cardinal_0,
	tag_ln:    cardinal_6,
	tag_lo:    cardinal_0,
	tag_lt:    cardinal_27,
	tag_lv:    cardinal_13,
	tag_mas:   cardinal_8,
	tag_mg:    cardinal_6,
	tag_mgo:   cardinal_8,
	tag_mk:    cardinal_11,
	tag_ml:    cardinal_8,
	tag_mn:    cardinal_8,
	tag_mo:    cardinal_18,
	tag_mr:    cardinal_8,
	tag_ms:    cardinal_0,
	tag_mt:    cardinal_28,
	tag_my:    cardinal_0,
	tag_nah:   cardinal_8,
	tag_naq:   cardinal_16,
	tag_nb:    cardinal_8,
	tag_nd:    cardinal_8,
	tag_ne:    cardinal_8,
	tag_nl:    cardinal_4,
	tag_nn:    cardinal_8,
	tag_nnh:   cardinal_8,
	tag_no:    cardinal_8,
	tag_nqo:   cardinal_0,
	tag_nr:    cardinal_8,
	tag_nso:   cardinal_6,
	tag_ny:    cardinal_8,
	tag_nyn:   cardinal_8,
	tag_om:    cardinal_8,
	tag_or:    cardinal_8,
	tag_os:    cardinal_8,
	tag_osa:   cardinal_0,
	tag_pa:    cardinal_6,
	tag_pap:   cardinal_8,
	tag_pcm:   cardinal_1,
	tag_pl:    cardinal_25,
	tag_prg:   cardinal_13,
	tag_ps:    cardinal_8,
	tag_pt:    cardinal_3,
	tag_pt_PT: cardinal_4,
	tag_rm:    cardinal_8,
	tag_ro:    cardinal_18,
	tag_rof:   cardinal_8,
	tag_root:  cardinal_0,
	tag_ru:    cardinal_29,
	tag_rwk:   cardinal_8,
	tag_sah:   cardinal_0,
	tag_saq:   cardinal_8,
	tag_sat:   cardinal_16,
	tag_sc:    cardinal_4,
	tag_scn:   cardinal_4,
	tag_sd:    cardinal_8,
	tag_sdh:   cardinal_8,
	tag_se:    cardinal_16,
	tag_seh:   cardinal_8,
	tag_ses:   cardinal_0,
	tag_sg:    cardinal_0,
	tag_sh:    cardinal_19,
	tag_shi:   cardinal_17,
	tag_si:    cardinal_5,
	tag_sk:    cardinal_24,
	tag_sl:    cardinal_21,
	tag_sma:   cardinal_16,
	tag_smi:   cardinal_16,
	tag_smj:   cardinal_16,
	tag_smn:   cardinal_16,
	tag_sms:   cardinal_16,
	tag_sn:    cardinal_8,
	tag_so:    cardinal_8,
	tag_sq:    cardinal_8,
	tag_sr:    cardinal_19,
	tag_ss:    cardinal_8,
	tag_ssy:   cardinal_8,
	tag_st:    cardinal_8,
	tag_su:    cardinal_0,
	tag_sv:    cardinal_4,
	tag_sw:    cardinal_4,
	tag_syr:   cardinal_8,
	tag_ta:    cardinal_8,
	tag_te:    cardinal_8,
	tag_teo:   cardinal_8,
	tag_th:    cardinal_0,
	tag_ti:    cardinal_6,
	tag_tig:   cardinal_8,
	tag_tk:    cardinal_8,
	tag_tl:    cardinal_12,
	tag_tn:    cardinal_8,
	tag_to:    cardinal_0,
	tag_tr:    cardinal_8,
	tag_ts:    cardinal_8,
	tag_tzm:   cardinal_7,
	tag_ug:    cardinal_8,
	tag_uk:    cardinal_29,
	tag_ur:    cardinal_4,
	tag_uz:    cardinal_8,
	tag_ve:    cardinal_8,
	tag_vi:    cardinal_0,
	tag_vo:    cardinal_8,
	tag_vun:   cardinal_8,
	tag_wa:    cardinal_6,
	tag_wae:   cardinal_8,
	tag_wo:    cardinal_0,
	tag_xh:    cardinal_8,
	tag_xog:   cardinal_8,
	tag_yi:    cardinal_4,
	tag_yo:    cardinal_0,
	tag_yue:   cardinal_0,
	tag_zh:    cardinal_0,
	tag_zu:    cardinal_1,
}
