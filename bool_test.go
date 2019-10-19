package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestNullBoolStringer(t *testing.T) {
	var b *Bool = nil

	want := ""
	got := fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
	b = PtrBoolOf(false)
	want = "true"
	b.Set(true)
	got = fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = "false"
	b = PtrBoolOf(false)
	got = fmt.Sprint(b)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullBoolMarshalJSON(t *testing.T) {
	var b *Bool

	type test struct {
		Bo *Bool
		Aa int
	}
	data, _ := json.Marshal(test{Aa: 1})
	fmt.Println(string(data))
	data, _ = json.Marshal(test{Bo: PtrBoolOf(true)})
	fmt.Println(string(data))

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(b)
	if err != nil {
		t.Fatal(err)
	}

	want := "null"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()
	b = PtrBoolOf(false)
	b.Set(true)
	err = json.NewEncoder(&buf).Encode(b)
	if err != nil {
		t.Fatal(err)
	}

	want = "true"
	got = strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()

	b.Set(false)
	err = json.NewEncoder(&buf).Encode(b)
	if err != nil {
		t.Fatal(err)
	}

	want = "false"
	got = strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullBoolUnmarshalJSON(t *testing.T) {
	var b *Bool

	type test struct {
		Bo *Bool
		Aa int
	}

	fmt.Println(json.Unmarshal(json.Marshal(test{})))
	fmt.Println(test{Bo: PtrBoolOf(true)})

	err := json.NewDecoder(strings.NewReader("null")).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	err = json.NewDecoder(strings.NewReader("true")).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	if b == nil {
		t.Fatalf("must not be null but got nil")
	}

	want := Bool(true)
	got := *b
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	err = json.NewDecoder(strings.NewReader(`"foo"`)).Decode(&b)
	if err == nil {
		t.Fatal("should be fail")
	}
}
