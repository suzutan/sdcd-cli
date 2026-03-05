package model

import (
	"encoding/json"
	"testing"
	"time"
)

func TestBuild_JSONUnmarshal(t *testing.T) {
	data := `{"id":1,"jobId":2,"eventId":3,"status":"SUCCESS","sha":"abc","number":5}`
	var b Build
	if err := json.Unmarshal([]byte(data), &b); err != nil {
		t.Fatal(err)
	}
	if b.ID != 1 || b.Status != "SUCCESS" {
		t.Errorf("unexpected build: %+v", b)
	}
	if b.CreateTime != nil {
		t.Error("expected nil CreateTime")
	}
}

func TestBuild_JSONUnmarshal_WithTime(t *testing.T) {
	data := `{"id":2,"createTime":"2024-01-15T12:00:00.000Z"}`
	var b Build
	if err := json.Unmarshal([]byte(data), &b); err != nil {
		t.Fatal(err)
	}
	if b.CreateTime == nil {
		t.Fatal("expected non-nil CreateTime")
	}
	expected := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	if !b.CreateTime.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, *b.CreateTime)
	}
}

func TestStep_JSONUnmarshal(t *testing.T) {
	data := `{"name":"install","code":0}`
	var s Step
	if err := json.Unmarshal([]byte(data), &s); err != nil {
		t.Fatal(err)
	}
	if s.Name != "install" {
		t.Errorf("expected name=install, got %q", s.Name)
	}
	if s.Code == nil || *s.Code != 0 {
		t.Errorf("expected code=0, got %v", s.Code)
	}
}

func TestLogLine_JSONUnmarshal(t *testing.T) {
	data := `[{"t":1000,"m":"hello world","n":0},{"t":2000,"m":"done","n":1}]`
	var lines []LogLine
	if err := json.Unmarshal([]byte(data), &lines); err != nil {
		t.Fatal(err)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
	if lines[0].M != "hello world" {
		t.Errorf("expected 'hello world', got %q", lines[0].M)
	}
}
