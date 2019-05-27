package exif

import (
    "bytes"
    "fmt"
    "reflect"
    "testing"

    "github.com/dsoprea/go-logging"
)

func TestTagType_EncodeDecode_Byte(t *testing.T) {
    tt := NewTagType(TypeByte, TestDefaultByteOrder)

    data := []byte{0x11, 0x22, 0x33, 0x44, 0x55}

    encoded, err := tt.Encode(data)
    log.PanicIf(err)

    if bytes.Compare(encoded, data) != 0 {
        t.Fatalf("Data not encoded correctly.")
    }

    restored, err := tt.ParseBytes(encoded, uint32(len(data)))
    log.PanicIf(err)

    if bytes.Compare(restored, data) != 0 {
        t.Fatalf("Data not decoded correctly.")
    }
}

func TestTagType_EncodeDecode_Ascii(t *testing.T) {
    tt := NewTagType(TypeAscii, TestDefaultByteOrder)

    data := "hello"

    encoded, err := tt.Encode(data)
    log.PanicIf(err)

    if string(encoded) != fmt.Sprintf("%s\000", data) {
        t.Fatalf("Data not encoded correctly.")
    }

    restored, err := tt.ParseAscii(encoded, uint32(len(data)))
    log.PanicIf(err)

    if restored != data {
        t.Fatalf("Data not decoded correctly.")
    }
}

func TestTagType_EncodeDecode_Shorts(t *testing.T) {
    tt := NewTagType(TypeShort, TestDefaultByteOrder)

    data := []uint16{0x11, 0x22, 0x33}

    encoded, err := tt.Encode(data)
    log.PanicIf(err)

    if bytes.Compare(encoded, []byte{0x00, 0x11, 0x00, 0x22, 0x00, 0x33}) != 0 {
        t.Fatalf("Data not encoded correctly.")
    }

    restored, err := tt.ParseShorts(encoded, uint32(len(data)))
    log.PanicIf(err)

    if reflect.DeepEqual(restored, data) != true {
        t.Fatalf("Data not decoded correctly.")
    }
}

func TestTagType_EncodeDecode_Long(t *testing.T) {
    tt := NewTagType(TypeLong, TestDefaultByteOrder)

    data := []uint32{0x11, 0x22, 0x33}

    encoded, err := tt.Encode(data)
    log.PanicIf(err)

    if bytes.Compare(encoded, []byte{0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x33}) != 0 {
        t.Fatalf("Data not encoded correctly.")
    }

    restored, err := tt.ParseLongs(encoded, uint32(len(data)))
    log.PanicIf(err)

    if reflect.DeepEqual(restored, data) != true {
        t.Fatalf("Data not decoded correctly.")
    }
}

func TestTagType_EncodeDecode_Rational(t *testing.T) {
    tt := NewTagType(TypeRational, TestDefaultByteOrder)

    data := []Rational{
        Rational{Numerator: 0x11, Denominator: 0x22},
        Rational{Numerator: 0x33, Denominator: 0x44},
    }

    encoded, err := tt.Encode(data)
    log.PanicIf(err)

    if bytes.Compare(encoded, []byte{0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x44}) != 0 {
        t.Fatalf("Data not encoded correctly.")
    }

    restored, err := tt.ParseRationals(encoded, uint32(len(data)))
    log.PanicIf(err)

    if reflect.DeepEqual(restored, data) != true {
        t.Fatalf("Data not decoded correctly.")
    }
}

func TestTagType_EncodeDecode_SignedLong(t *testing.T) {
    tt := NewTagType(TypeSignedLong, TestDefaultByteOrder)

    data := []int32{0x11, 0x22, 0x33}

    encoded, err := tt.Encode(data)
    log.PanicIf(err)

    if bytes.Compare(encoded, []byte{0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x33}) != 0 {
        t.Fatalf("Data not encoded correctly.")
    }

    restored, err := tt.ParseSignedLongs(encoded, uint32(len(data)))
    log.PanicIf(err)

    if reflect.DeepEqual(restored, data) != true {
        t.Fatalf("Data not decoded correctly.")
    }
}

func TestTagType_EncodeDecode_SignedRational(t *testing.T) {
    tt := NewTagType(TypeSignedRational, TestDefaultByteOrder)

    data := []SignedRational{
        SignedRational{Numerator: 0x11, Denominator: 0x22},
        SignedRational{Numerator: 0x33, Denominator: 0x44},
    }

    encoded, err := tt.Encode(data)
    log.PanicIf(err)

    if bytes.Compare(encoded, []byte{0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x22, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x44}) != 0 {
        t.Fatalf("Data not encoded correctly.")
    }

    restored, err := tt.ParseSignedRationals(encoded, uint32(len(data)))
    log.PanicIf(err)

    if reflect.DeepEqual(restored, data) != true {
        t.Fatalf("Data not decoded correctly.")
    }
}

func TestTagType_EncodeDecode_AsciiNoNul(t *testing.T) {
    tt := NewTagType(TypeAsciiNoNul, TestDefaultByteOrder)

    data := "hello"

    encoded, err := tt.Encode(data)
    log.PanicIf(err)

    if string(encoded) != data {
        t.Fatalf("Data not encoded correctly.")
    }

    restored, err := tt.ParseAsciiNoNul(encoded, uint32(len(data)))
    log.PanicIf(err)

    if restored != data {
        t.Fatalf("Data not decoded correctly.")
    }
}

// TODO(dustin): Add tests for TypeUndefined.

func TestTagType_FromString_Undefined(t *testing.T) {
    defer func() {
        if state := recover(); state != nil {
            err := log.Wrap(state.(error))
            log.PrintErrorf(err, "Test failure.")

            log.Panic(err)
        }
    }()

    tt := NewTagType(TypeUndefined, TestDefaultByteOrder)

    _, err := tt.FromString("")
    if err == nil {
        t.Fatalf("no error for undefined-type")
    } else if err.Error() != "undefined-type values are not supported" {
        fmt.Printf("[%s]\n", err.Error())
        log.Panic(err)
    }
}

func TestTagType_FromString_Byte(t *testing.T) {
    tt := NewTagType(TypeByte, TestDefaultByteOrder)

    value, err := tt.FromString("abc")
    log.PanicIf(err)

    if reflect.DeepEqual(value, []byte{'a', 'b', 'c'}) != true {
        t.Fatalf("byte value not correct")
    }
}

func TestTagType_FromString_Ascii(t *testing.T) {
    tt := NewTagType(TypeAscii, TestDefaultByteOrder)

    value, err := tt.FromString("abc")
    log.PanicIf(err)

    if reflect.DeepEqual(value, "abc") != true {
        t.Fatalf("ASCII value not correct: [%s]", value)
    }
}

func TestTagType_FromString_Short(t *testing.T) {
    tt := NewTagType(TypeShort, TestDefaultByteOrder)

    value, err := tt.FromString("55")
    log.PanicIf(err)

    if reflect.DeepEqual(value, uint16(55)) != true {
        t.Fatalf("short value not correct")
    }
}

func TestTagType_FromString_Long(t *testing.T) {
    tt := NewTagType(TypeLong, TestDefaultByteOrder)

    value, err := tt.FromString("66000")
    log.PanicIf(err)

    if reflect.DeepEqual(value, uint32(66000)) != true {
        t.Fatalf("long value not correct")
    }
}

func TestTagType_FromString_Rational(t *testing.T) {
    tt := NewTagType(TypeRational, TestDefaultByteOrder)

    value, err := tt.FromString("12/34")
    log.PanicIf(err)

    expected := Rational{
        Numerator:   12,
        Denominator: 34,
    }

    if reflect.DeepEqual(value, expected) != true {
        t.Fatalf("rational value not correct")
    }
}

func TestTagType_FromString_SignedLong(t *testing.T) {
    tt := NewTagType(TypeSignedLong, TestDefaultByteOrder)

    value, err := tt.FromString("-66000")
    log.PanicIf(err)

    if reflect.DeepEqual(value, int32(-66000)) != true {
        t.Fatalf("signed-long value not correct")
    }
}

func TestTagType_FromString_SignedRational(t *testing.T) {
    tt := NewTagType(TypeSignedRational, TestDefaultByteOrder)

    value, err := tt.FromString("-12/34")
    log.PanicIf(err)

    expected := SignedRational{
        Numerator:   -12,
        Denominator: 34,
    }

    if reflect.DeepEqual(value, expected) != true {
        t.Fatalf("signd-rational value not correct")
    }
}

func TestTagType_FromString_AsciiNoNul(t *testing.T) {
    tt := NewTagType(TypeAsciiNoNul, TestDefaultByteOrder)

    value, err := tt.FromString("abc")
    log.PanicIf(err)

    if reflect.DeepEqual(value, "abc") != true {
        t.Fatalf("ASCII-no-nul value not correct")
    }
}

func TestParseValue_TypeByte(t *testing.T) {
    tt := NewTagType(TypeByte, TestDefaultByteOrder)

    value := []byte{1, 2, 3, 4}

    encodedBytes, err := tt.Encode(value)
    log.PanicIf(err)

    wrib := WrappedRawInputBytes{
        data:      encodedBytes,
        unitCount: 4,
    }

    wov, err := ParseValue(TypeByte, wrib, TestDefaultByteOrder)
    log.PanicIf(err)

    if wov.tagType != TypeByte {
        t.Fatalf("WOV tag-type not correct: (%d)", wov.tagType)
    } else if wov.unitCount != 4 {
        t.Fatalf("WOV unit-count not correct: (%d)", wov.unitCount)
    }

    recoveredValue := wov.Raw().([]byte)

    if reflect.DeepEqual(recoveredValue, value) != true {
        t.Fatalf("recovered value does not equal expected value: [%v] != [%v]", recoveredValue, value)
    }

    s := wov.GetString(true)
    if s != "0x01" {
        t.Fatalf("recovered first value not correct: [%s]", s)
    }

    s = wov.String()
    if s != "01 02 03 04" {
        t.Fatalf("recovered value not correct: [%s]", s)
    }
}

func TestParseValue_TypeAscii(t *testing.T) {
    tt := NewTagType(TypeAscii, TestDefaultByteOrder)

    value := "testing"

    encodedBytes, err := tt.Encode(value)
    log.PanicIf(err)

    wrib := WrappedRawInputBytes{
        data:      encodedBytes,
        unitCount: uint32(len(value)),
    }

    wov, err := ParseValue(TypeAscii, wrib, TestDefaultByteOrder)
    log.PanicIf(err)

    if wov.tagType != TypeAscii {
        t.Fatalf("WOV tag-type not correct: (%d)", wov.tagType)
    } else if wov.unitCount != uint32(len(value)) {
        t.Fatalf("WOV unit-count not correct: (%d)", wov.unitCount)
    }

    recoveredValue := wov.Raw().(string)

    if reflect.DeepEqual(recoveredValue, value) != true {
        t.Fatalf("recovered value does not equal expected value: [%v] != [%v]", recoveredValue, value)
    }

    s := wov.GetString(true)
    if s != value {
        t.Fatalf("recovered first value not correct: [%s]", s)
    }

    s = wov.String()
    if s != value {
        t.Fatalf("recovered value not correct: [%s]", s)
    }
}

func TestParseValue_TypeAsciiNoNul(t *testing.T) {
    tt := NewTagType(TypeAsciiNoNul, TestDefaultByteOrder)

    value := "testing"

    encodedBytes, err := tt.Encode(value)
    log.PanicIf(err)

    wrib := WrappedRawInputBytes{
        data:      encodedBytes,
        unitCount: uint32(len(value)),
    }

    wov, err := ParseValue(TypeAsciiNoNul, wrib, TestDefaultByteOrder)
    log.PanicIf(err)

    if wov.tagType != TypeAsciiNoNul {
        t.Fatalf("WOV tag-type not correct: (%d)", wov.tagType)
    } else if wov.unitCount != uint32(len(value)) {
        t.Fatalf("WOV unit-count not correct: (%d)", wov.unitCount)
    }

    recoveredValue := wov.Raw().(string)

    if reflect.DeepEqual(recoveredValue, value) != true {
        t.Fatalf("recovered value does not equal expected value: [%v] != [%v]", recoveredValue, value)
    }

    s := wov.GetString(true)
    if s != value {
        t.Fatalf("recovered first value not correct: [%s]", s)
    }

    s = wov.String()
    if s != value {
        t.Fatalf("recovered value not correct: [%s]", s)
    }
}

func TestParseValue_TypeShort(t *testing.T) {
    tt := NewTagType(TypeShort, TestDefaultByteOrder)

    value := []uint16{11, 22, 33, 44}

    encodedBytes, err := tt.Encode(value)
    log.PanicIf(err)

    wrib := WrappedRawInputBytes{
        data:      encodedBytes,
        unitCount: uint32(len(value)),
    }

    wov, err := ParseValue(TypeShort, wrib, TestDefaultByteOrder)
    log.PanicIf(err)

    if wov.tagType != TypeShort {
        t.Fatalf("WOV tag-type not correct: (%d)", wov.tagType)
    } else if wov.unitCount != uint32(len(value)) {
        t.Fatalf("WOV unit-count not correct: (%d)", wov.unitCount)
    }

    recoveredValue := wov.Raw().([]uint16)

    if reflect.DeepEqual(recoveredValue, value) != true {
        t.Fatalf("recovered value does not equal expected value: [%v] != [%v]", recoveredValue, value)
    }

    s := wov.GetString(true)
    if s != "11" {
        t.Fatalf("recovered first value not correct: [%s]", s)
    }

    s = wov.String()
    if s != "[11 22 33 44]" {
        t.Fatalf("recovered value not correct: [%s]", s)
    }
}

func TestParseValue_TypeLong(t *testing.T) {
    tt := NewTagType(TypeLong, TestDefaultByteOrder)

    value := []uint32{11223344}

    encodedBytes, err := tt.Encode(value)
    log.PanicIf(err)

    wrib := WrappedRawInputBytes{
        data:      encodedBytes,
        unitCount: 1,
    }

    wov, err := ParseValue(TypeLong, wrib, TestDefaultByteOrder)
    log.PanicIf(err)

    if wov.tagType != TypeLong {
        t.Fatalf("WOV tag-type not correct: (%d)", wov.tagType)
    } else if wov.unitCount != 1 {
        t.Fatalf("WOV unit-count not correct: (%d)", wov.unitCount)
    }

    recoveredValue := wov.Raw().([]uint32)

    if reflect.DeepEqual(recoveredValue, value) != true {
        t.Fatalf("recovered value does not equal expected value: [%v] != [%v]", recoveredValue, value)
    }

    s := wov.GetString(true)
    if s != "11223344" {
        t.Fatalf("recovered first value not correct: [%s]", s)
    }

    s = wov.String()
    if s != "[11223344]" {
        t.Fatalf("recovered value not correct: [%s]", s)
    }
}

func TestParseValue_TypeRational(t *testing.T) {
    tt := NewTagType(TypeRational, TestDefaultByteOrder)

    value := []Rational{
        Rational{Numerator: 0x11, Denominator: 0x22},
    }

    encodedBytes, err := tt.Encode(value)
    log.PanicIf(err)

    wrib := WrappedRawInputBytes{
        data:      encodedBytes,
        unitCount: 1,
    }

    wov, err := ParseValue(TypeRational, wrib, TestDefaultByteOrder)
    log.PanicIf(err)

    if wov.tagType != TypeRational {
        t.Fatalf("WOV tag-type not correct: (%d)", wov.tagType)
    } else if wov.unitCount != 1 {
        t.Fatalf("WOV unit-count not correct: (%d)", wov.unitCount)
    }

    recoveredValue := wov.Raw().([]Rational)

    if reflect.DeepEqual(recoveredValue, value) != true {
        t.Fatalf("recovered value does not equal expected value: [%v] != [%v]", recoveredValue, value)
    }

    s := wov.GetString(true)
    if s != "17/34" {
        t.Fatalf("recovered first value not correct: [%s]", s)
    }

    s = wov.String()
    if s != "[17/34]" {
        t.Fatalf("recovered value not correct: [%s]", s)
    }
}

// TODO(dustin): Add more tests for multiple values.

func TestParseValue_TypeRational_Multiple(t *testing.T) {
    tt := NewTagType(TypeRational, TestDefaultByteOrder)

    value := []Rational{
        Rational{Numerator: 0x11, Denominator: 0x22},
        Rational{Numerator: 0x33, Denominator: 0x44},
    }

    encodedBytes, err := tt.Encode(value)
    log.PanicIf(err)

    wrib := WrappedRawInputBytes{
        data:      encodedBytes,
        unitCount: 2,
    }

    wov, err := ParseValue(TypeRational, wrib, TestDefaultByteOrder)
    log.PanicIf(err)

    if wov.tagType != TypeRational {
        t.Fatalf("WOV tag-type not correct: (%d)", wov.tagType)
    } else if wov.unitCount != 2 {
        t.Fatalf("WOV unit-count not correct: (%d)", wov.unitCount)
    }

    recoveredValue := wov.Raw().([]Rational)

    if reflect.DeepEqual(recoveredValue, value) != true {
        t.Fatalf("recovered value does not equal expected value: [%v] != [%v]", recoveredValue, value)
    }

    s := wov.GetString(true)
    if s != "17/34" {
        t.Fatalf("recovered first value not correct: [%s]", s)
    }

    s = wov.String()
    if s != "[17/34 51/68]" {
        t.Fatalf("recovered value not correct: [%s]", s)
    }
}

func TestParseValue_TypeSignedLong(t *testing.T) {
    tt := NewTagType(TypeSignedLong, TestDefaultByteOrder)

    value := []int32{0x11}

    encodedBytes, err := tt.Encode(value)
    log.PanicIf(err)

    wrib := WrappedRawInputBytes{
        data:      encodedBytes,
        unitCount: uint32(len(value)),
    }

    wov, err := ParseValue(TypeSignedLong, wrib, TestDefaultByteOrder)
    log.PanicIf(err)

    if wov.tagType != TypeSignedLong {
        t.Fatalf("WOV tag-type not correct: (%d)", wov.tagType)
    } else if wov.unitCount != uint32(len(value)) {
        t.Fatalf("WOV unit-count not correct: (%d)", wov.unitCount)
    }

    recoveredValue := wov.Raw().([]int32)

    if reflect.DeepEqual(recoveredValue, value) != true {
        t.Fatalf("recovered value does not equal expected value: [%v] != [%v]", recoveredValue, value)
    }

    s := wov.GetString(true)
    if s != "17" {
        t.Fatalf("recovered first value not correct: [%s]", s)
    }

    s = wov.String()
    if s != "[17]" {
        t.Fatalf("recovered value not correct: [%s]", s)
    }
}

func TestParseValue_TypeSignedRational(t *testing.T) {
    tt := NewTagType(TypeSignedRational, TestDefaultByteOrder)

    value := []SignedRational{
        SignedRational{Numerator: 0x11, Denominator: 0x22},
    }

    encodedBytes, err := tt.Encode(value)
    log.PanicIf(err)

    wrib := WrappedRawInputBytes{
        data:      encodedBytes,
        unitCount: 1,
    }

    wov, err := ParseValue(TypeSignedRational, wrib, TestDefaultByteOrder)
    log.PanicIf(err)

    if wov.tagType != TypeSignedRational {
        t.Fatalf("WOV tag-type not correct: (%d)", wov.tagType)
    } else if wov.unitCount != 1 {
        t.Fatalf("WOV unit-count not correct: (%d)", wov.unitCount)
    }

    recoveredValue := wov.Raw().([]SignedRational)

    if reflect.DeepEqual(recoveredValue, value) != true {
        t.Fatalf("recovered value does not equal expected value: [%v] != [%v]", recoveredValue, value)
    }

    s := wov.GetString(true)
    if s != "17/34" {
        t.Fatalf("recovered first value not correct: [%s]", s)
    }

    s = wov.String()
    if s != "[17/34]" {
        t.Fatalf("recovered value not correct: [%s]", s)
    }
}
