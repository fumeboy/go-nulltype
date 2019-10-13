package nulltype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestNullTimeStringer(t *testing.T) {
	var nt Time

	want := ""
	got := fmt.Sprint(nt)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	now := time.Now()
	want = now.Format("2006/01/02 15:04:05")
	nt.Set(now)
	got = fmt.Sprint(nt)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	nt = TimeOf(now)
	got = fmt.Sprint(nt)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	want = ""
	nt.Reset()
	got = fmt.Sprint(nt)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullTimeMarshalJSON(t *testing.T) {
	var nt Time

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(nt)
	if err != nil {
		t.Fatal(err)
	}

	want := "null"
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	buf.Reset()

	now := time.Now()
	nt.Set(now)
	err = json.NewEncoder(&buf).Encode(nt)
	if err != nil {
		t.Fatal(err)
	}

	want = `"` + now.Format(time.RFC3339) + `"`
	got = strings.TrimSpace(buf.String())
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullTimeUnmarshalJSON(t *testing.T) {
	var nt Time

	err := json.NewDecoder(strings.NewReader("null")).Decode(&nt)
	if err != nil {
		t.Fatal(err)
	}

	if nt.Valid() {
		t.Fatalf("must be null but got %v", nt)
	}

	err = json.NewDecoder(strings.NewReader(`"2019-02-01T11:12:13Z"`)).Decode(&nt)
	if err != nil {
		t.Fatal(err)
	}

	if !nt.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	now, _ := time.Parse(time.RFC3339, "2019-02-01T11:12:13Z")
	want := now
	got := nt.TimeValue()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	err = json.NewDecoder(strings.NewReader("{}")).Decode(&nt)
	if err == nil {
		t.Fatal("should be fail")
	}

	err = json.NewDecoder(strings.NewReader(`"2019-02-01"`)).Decode(&nt)
	if err == nil {
		t.Fatal("should be fail")
	}
}

func TestNullTimeValueConverter(t *testing.T) {
	var nt Time

	now := time.Now()
	err := nt.Scan(now)
	if err != nil {
		t.Fatal(err)
	}

	if !nt.Valid() {
		t.Fatalf("must not be null but got nil")
	}

	want := now
	got := nt.TimeValue()
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	gotv, err := nt.Value()
	if err != nil {
		t.Fatal(err)
	}
	if gotv != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	nt.Reset()

	gotv, err = nt.Value()
	if err != nil {
		t.Fatal(err)
	}
	if gotv != nil {
		t.Fatalf("must be null but got %v", gotv)
	}

	err = nt.Scan("2019-02-19 17:53:00")
	if err != nil {
		t.Fatal(err)
	}
	wants := "2019-02-19 17:53:00"
	gots := nt.TimeValue().Format("2006-01-02 15:04:05")
	if gots != wants {
		t.Fatalf("want %v, but %v:", wants, gots)
	}

	err = nt.Scan("2019+02+19")
	if err != nil {
		t.Fatal(err)
	}

	err = nt.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}

	var v complex64
	err = nt.Scan(v)
	if err == nil {
		t.Fatal("should be fail")
	}
}
