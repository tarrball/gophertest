package expects

import (
	"testing"

	"github.com/tarrball/go-expects/mocks"
)

func TestToBeGreaterThanFloatGreaterPasses(t *testing.T) {
	actual, expected := 2.1, 1.2

	expectation := ExpectationFloat{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanFloatLessFails(t *testing.T) {
	actual, expected := 1.2, 2.1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be greater than '%f'", actual, expected)

	expectation := ExpectationFloat{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanFloatEqualFails(t *testing.T) {
	actual, expected := 1.1, 1.1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be greater than '%f'", actual, expected)

	expectation := ExpectationFloat{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanFloat32GreaterPasses(t *testing.T) {
	actual, expected := float32(2.1), float32(1.2)

	expectation := ExpectationFloat32{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanFloat32LessFails(t *testing.T) {
	actual, expected := float32(1.2), float32(2.1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be greater than '%f'", actual, expected)

	expectation := ExpectationFloat32{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanFloat32EqualFails(t *testing.T) {
	actual, expected := float32(1.1), float32(1.1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be greater than '%f'", actual, expected)

	expectation := ExpectationFloat32{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanIntGreaterPasses(t *testing.T) {
	actual, expected := 2, 1

	expectation := ExpectationInt{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanIntLessFails(t *testing.T) {
	actual, expected := 1, 2
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationInt{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanIntEqualFails(t *testing.T) {
	actual, expected := 1, 1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationInt{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt8GreaterPasses(t *testing.T) {
	actual, expected := int8(2), int8(1)

	expectation := ExpectationInt8{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt8LessFails(t *testing.T) {
	actual, expected := int8(1), int8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationInt8{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt8EqualFails(t *testing.T) {
	actual, expected := int8(1), int8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationInt8{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt16GreaterPasses(t *testing.T) {
	actual, expected := int16(2), int16(1)

	expectation := ExpectationInt16{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt16LessFails(t *testing.T) {
	actual, expected := int16(1), int16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationInt16{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt16EqualFails(t *testing.T) {
	actual, expected := int16(1), int16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationInt16{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt32GreaterPasses(t *testing.T) {
	actual, expected := int32(2), int32(1)

	expectation := ExpectationInt32{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt32LessFails(t *testing.T) {
	actual, expected := int32(1), int32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationInt32{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt32EqualFails(t *testing.T) {
	actual, expected := int32(1), int32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationInt32{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt64GreaterPasses(t *testing.T) {
	actual, expected := int64(2), int64(1)

	expectation := ExpectationInt64{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt64LessFails(t *testing.T) {
	actual, expected := int64(1), int64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationInt64{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanInt64EqualFails(t *testing.T) {
	actual, expected := int64(1), int64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationInt64{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanStringGreaterPasses(t *testing.T) {
	actual, expected := "def", "abc"

	expectation := ExpectationString{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanStringLessFails(t *testing.T) {
	actual, expected := "abc", "def"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to be greater than '%s'", actual, expected)

	expectation := ExpectationString{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanStringEqualFails(t *testing.T) {
	actual, expected := "abc", "abc"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to be greater than '%s'", actual, expected)

	expectation := ExpectationString{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUintGreaterPasses(t *testing.T) {
	actual, expected := uint(2), uint(1)

	expectation := ExpectationUint{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUintLessFails(t *testing.T) {
	actual, expected := uint(1), uint(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationUint{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUintEqualFails(t *testing.T) {
	actual, expected := uint(1), uint(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationUint{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint8GreaterPasses(t *testing.T) {
	actual, expected := uint8(2), uint8(1)

	expectation := ExpectationUint8{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint8LessFails(t *testing.T) {
	actual, expected := uint8(1), uint8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationUint8{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint8EqualFails(t *testing.T) {
	actual, expected := uint8(1), uint8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationUint8{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint16GreaterPasses(t *testing.T) {
	actual, expected := uint16(2), uint16(1)

	expectation := ExpectationUint16{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint16LessFails(t *testing.T) {
	actual, expected := uint16(1), uint16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationUint16{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint16EqualFails(t *testing.T) {
	actual, expected := uint16(1), uint16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationUint16{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint32GreaterPasses(t *testing.T) {
	actual, expected := uint32(2), uint32(1)

	expectation := ExpectationUint32{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint32LessFails(t *testing.T) {
	actual, expected := uint32(1), uint32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationUint32{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint32EqualFails(t *testing.T) {
	actual, expected := uint32(1), uint32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationUint32{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint64GreaterPasses(t *testing.T) {
	actual, expected := uint64(2), uint64(1)

	expectation := ExpectationUint64{actual, Expectation{t}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint64LessFails(t *testing.T) {
	actual, expected := uint64(1), uint64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationUint64{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}

func TestToBeGreaterThanUint64EqualFails(t *testing.T) {
	actual, expected := uint64(1), uint64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'", actual, expected)

	expectation := ExpectationUint64{actual, Expectation{mock}}

	expectation.ToBeGreaterThan(expected)
}
