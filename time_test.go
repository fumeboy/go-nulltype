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
	var nt *Time

	want := ""
	got := fmt.Sprint(nt)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	nt = PtrTimeOf(time.Now())
	now := time.Now()
	want = now.Format("2006/01/02 15:04:05")
	nt.Set(now)
	got = fmt.Sprint(nt)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}

	nt = PtrTimeOf(now)
	got = fmt.Sprint(nt)
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func TestNullTimeMarshalJSON(t *testing.T) {
	var nt *Time

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
	nt = PtrTimeOf(time.Now())
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
	var nt *Time

	err := json.NewDecoder(strings.NewReader("null")).Decode(&nt)
	if err != nil {
		t.Fatal(err)
	}

	err = json.NewDecoder(strings.NewReader(`"2019-02-01T11:12:13Z"`)).Decode(&nt)
	if err != nil {
		t.Fatal(err)
	}

	if nt == nil {
		t.Fatalf("must not be null but got nil")
	}

	now, _ := time.Parse(time.RFC3339, "2019-02-01T11:12:13Z")
	want := now
	got := *nt
	if got != Time(want) {
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