package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/status", nil)
	w := httptest.NewRecorder()
	statusHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 got %d", w.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}
	if resp["status"] != "ok" {
		t.Errorf("expected status 'ok' got %s", resp["status"])
	}
}

func TestWeekHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/game/week", nil)
	w := httptest.NewRecorder()
	weekHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 got %d", w.Code)
	}

	var resp map[string][]string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	letters := resp["letters"]
	if len(letters) != 7 {
		t.Errorf("expected 7 letters got %d", len(letters))
	}
}

func TestWordHandler(t *testing.T) {
	body := bytes.NewBufferString(`{"word":"test"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/game/word", body)
	w := httptest.NewRecorder()
	wordHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 got %d", w.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}
	expected := "Received word test"
	if resp["message"] != expected {
		t.Errorf("expected %q got %q", expected, resp["message"])
	}
}
