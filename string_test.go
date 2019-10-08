package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNullStringStringer(t *testing.T) {
	var s NullString

	want := ""
	got := fmt.Sprint(s)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "foo"
	s.Set("foo")
	got = fmt.Sprint(s)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "bar"
	s = NullStringOf("bar")
	got = fmt.Sprint(s)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = ""
	s.Reset()
	got = fmt.Sprint(s)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullStringMarshalJSON(t *testing.T) {
	var s NullString

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(s)
	if err != nil {
		t.Fatal(err)
	}

	want := "null"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()

	s.Set("foo")
	err = json.NewEncoder(&buf).Encode(s)
	if err != nil {
		t.Fatal(err)
	}

	want = `"foo"`
	got = strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullStringUnmarshalJSON(t *testing.T) {
	var s NullString

	err := json.NewDecoder(strings.NewReader("null")).Decode(&s)
	if err != nil {
		t.Fatal(err)
	}

	if s.Valid() {
		t.Fatalf("must be null but got %v", s)
	}

	err = json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&s)
	if err != nil {
		t.Fatal(err)
	}

	if !s.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := "foo"
	got := s.StringValue()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	err = json.NewDecoder(strings.NewReader("{}")).Decode(&s)
	if err == nil {
		t.Fatal("should be fail")
	}
}

func TestNullStringValueConverter(t *testing.T) {
	var s NullString

	err := s.Scan("1")
	if err != nil {
		t.Fatal(err)
	}

	if !s.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := "1"
	got := s.StringValue()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	gotv, err := s.Value()
	if err != nil {
		t.Fatal(err)
	}
	if gotv != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	s.Reset()

	gotv, err = s.Value()
	if err != nil {
		t.Fatal(err)
	}
	if gotv != nil {
		t.Fatalf("must be null but got %v", gotv)
	}
}

func TestTest(t *testing.T) {
	type test2 struct {
		S string
		A int
	}
	type test struct {
		S *NullString
		T2 *test2
	}
	var tt = test{}
	var ss = NullStringOf("")
	var tt2 = test2{"a", 1, }
	tt2.S = "a"
	tt.T2 = &tt2
	tt.S = &ss
	d, _ := json.Marshal(&tt)
	fmt.Println(string(d), 1, tt.S == nil)
}
