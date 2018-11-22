package slice

import (
	"reflect"
	"testing"
)

func TestStringHas(t *testing.T) {
	islice := NewStringSlice([]string{"key1"})
	if !islice.Has("key1") {
		t.Fatalf("key 'key1' exist but false returned")
	}
}

func TestStringHasNot(t *testing.T) {
	islice := NewStringSlice([]string{"key1"})
	if islice.Has("key2") {
		t.Fatalf("key 'key2' not exists but true")
	}
}

func TestStringAppend(t *testing.T) {
	expected := []string{"key1", "key2"}
	islice := NewStringSlice([]string{"key1"})
	islice.Append("key2")
	if !islice.Has("key2") {
		t.Fatalf("key 'key2' exist but false returned")
	}

	if !reflect.DeepEqual(expected, islice.Values()) {
		t.Fatalf("should be '%v' given '%v'", expected, islice.Values())
	}
}

func TestIntHas(t *testing.T) {
	islice := NewIntSlice([]int{1})
	if !islice.Has(1) {
		t.Fatalf("key '1' exist but false returned")
	}
}

func TestIntHasNot(t *testing.T) {
	islice := NewIntSlice([]int{1})
	if islice.Has(2) {
		t.Fatalf("key '2' not exists but true")
	}
}

func TestIntAppend(t *testing.T) {
	expected := []int{1, 2}
	islice := NewIntSlice([]int{1})
	islice.Append(2)
	if !islice.Has(2) {
		t.Fatalf("key '2' exist but false returned")
	}

	if !reflect.DeepEqual(expected, islice.Values()) {
		t.Fatalf("should be '%v' given '%v'", expected, islice.Values())
	}
}


func TestFloatHas(t *testing.T) {
	islice := NewFloatSlice([]float64{1.1})
	if !islice.Has(1.1) {
		t.Fatalf("key '1.1' exist but false returned")
	}
}

func TestFloatHasNot(t *testing.T) {
	islice := NewFloatSlice([]float64{1.1})
	if islice.Has(1.2) {
		t.Fatalf("key '1.2' not exists but true")
	}
}

func TestFloatAppend(t *testing.T) {
	expected := []float64{1.1, 1.2}
	islice := NewFloatSlice([]float64{1.1})
	islice.Append(1.2)
	if !islice.Has(1.2) {
		t.Fatalf("key '1.2' exist but false returned")
	}

	if !reflect.DeepEqual(expected, islice.Values()) {
		t.Fatalf("should be '%v' given '%v'", expected, islice.Values())
	}
}