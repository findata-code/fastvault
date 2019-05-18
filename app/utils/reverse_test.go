package utils

import (
	"reflect"
	"testing"
)

func TestReversePalindrome(t *testing.T) {
	v := []byte("civic")
	if !reflect.DeepEqual(v, Reverse(v)){
		t.Error("expect", v, "actual", Reverse(v))
	}
}

func TestReverse(t *testing.T) {
	v := []byte("0123456789")
	if !reflect.DeepEqual([]byte("9876543210"), Reverse(v)){
		t.Error("expect", []byte("9876543210"), "actual", Reverse(v))
	}
}

func TestReverse2(t *testing.T) {
	v := []byte("01")
	if !reflect.DeepEqual([]byte("10"), Reverse(v)){
		t.Error("expect", []byte("10"), "actual", Reverse(v))
	}
}

func TestReverse3(t *testing.T) {
	v := []byte("1")
	if !reflect.DeepEqual([]byte("1"), Reverse(v)){
		t.Error("expect", []byte("1"), "actual", Reverse(v))
	}
}