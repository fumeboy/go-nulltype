package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNullInt64Stringer(t *testing.T) {
	var i *Int64

	want := ""
	got := fmt.Sprint(i)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	i = PtrInt64Of(0)
	want = "3"
	i.Set(3)
	got = fmt.Sprint(i)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "5"
	i = PtrInt64Of(5)
	got = fmt.Sprint(i)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullInt64MarshalJSON(t *testing.T) {
	var i Int64

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(i)
	if err != nil {
		t.Fatal(err)
	}

	want := "null"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()

	i.Set(3)
	err = json.NewEncoder(&buf).Encode(i)
	if err != nil {
		t.Fatal(err)
	}

	want = "3"
	got = strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullInt64UnmarshalJSON(t *testing.T) {
	var i *Int64

	err := json.NewDecoder(strings.NewReader("null")).Decode(&i)
	if err != nil {
		t.Fatal(err)
	}

	if i == nil {
		t.Fatalf("must be null but got %v", i)
	}

	err = json.NewDecoder(strings.NewReader(`3`)).Decode(&i)
	if err != nil {
		t.Fatal(err)
	}

	if i == nil {
		t.Fatalf("must not be null but got nil")
	}

	want := Int64(3)
	got := *i
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	err = json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&i)
	if err == nil {
		t.Fatal("should be fail")
	}
}