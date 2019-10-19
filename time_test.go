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

	type test struct {
		Bo *Time
		Aa int
	}
	data, _ := json.Marshal(test{Aa: 1})
	fmt.Println(string(data))
	data, _ = json.Marshal(test{Bo: PtrTimeOf(time.Now())})
	fmt.Println(string(data))

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
	type test struct {
		Bo *Time
		Aa int
	}
	var at = new(test)
	_ = json.Unmarshal([]byte(`{"Bo":"2019-10-19T18:28:52+08:00", "Aa":0}`), at)
	fmt.Println(at)
}