package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// "github.com/smartystreets/assertions/assert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=tomsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Fatalf("expected status code %d got %d", http.StatusOK, status)
	}

	t.Logf("expectes status code %d. Succeed", http.StatusOK)
}

func TestMainHandlerMissingCount(t *testing.T) {
	expectedMessage := "count missing"
	req := httptest.NewRequest("GET", "/cafe?city=krasnodar", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Fatalf("expected status code %d got %d", http.StatusBadRequest, status)
	}

	t.Logf("expected status code %d. Succed", http.StatusBadRequest)

	if responseRecorder.Body.String() != expectedMessage {
		t.Errorf("expected body: %s, got: %s", expectedMessage, responseRecorder.Body.String())
	}
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	total := 3

	req := httptest.NewRequest("GET", "/cafe?count=5&city=krasnodar", nil)

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(r, req)

	require.Equal(t, r.Code, http.StatusOK)
	// if r.Code != http.StatusOK {
	// 	t.Fatalf("expected status code %d got %d", http.StatusOK, r.Code)
	// }
	assert.NotEmpty(t, r.Body)

	answer := strings.Split(r.Body.String(), ",")
	assert.Len(t, answer, total)

	// assert.Equal(t, total, answer)
	// // if total != answer {
	// 	t.Errorf("Expected count cafe %d got %d", total, answer)
	// 	return
	// }
	t.Log("everything is ok")

}
