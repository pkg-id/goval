package govalregex

import "testing"

func TestExpr_RegExp(t *testing.T) {
	tests := []expr{
		Alpha,
		AlphaNumeric,
		AlphaUnicode,
		AlphaUnicodeNumeric,
		Numeric,
		Number,
		Hexadecimal,
		HexColor,
		RGB,
		RGBA,
		HSL,
		HSLA,
		Email,
		E164,
		Base64,
		Base64URL,
		Base64RawURL,
		ISBN10,
		ISBN13,
		UUID3,
		UUID4,
		UUID5,
		UUID,
		UUID3RFC4122,
		UUID4RFC4122,
		UUID5RFC4122,
		UUIDRFC4122,
		ULID,
		MD4,
		MD5,
		SHA256,
		SHA384,
		SHA512,
		RIPEMD128,
		RIPEMD160,
		Tiger128,
		Tiger160,
		Tiger192,
		ASCII,
		PrintableASCII,
		Multibyte,
		DataURI,
		Latitude,
		Longitude,
		SSN,
		HostnameRFC952,
		HostnameRFC1123,
		FQDNRFC1123,
		BTCAddress,
		BTCAddressUpperBech32,
		BTCAddressLowerBech32,
		ETHAddress,
		ETHAddressUpper,
		ETHAddressLower,
		URLEncoded,
		HTMLEncoded,
		HTML,
		JWT,
		SplitParams,
		Bic,
		Semver,
		DNSRFC1035Label,
		CVE,
		Mongodb,
		Cron,
	}

	for i, e := range tests {
		re := e.RegExp()
		if re == nil {
			t.Errorf("exoect expr at %d not nil", i)
		}
	}

	for i, e := range tests {
		re, ok := caches.Load(e)
		if !ok {
			t.Errorf("exoect expr at %d is cached", i)
		}
		if re == nil {
			t.Errorf("exoect expr at %d not nil", i)
		}
	}
}
