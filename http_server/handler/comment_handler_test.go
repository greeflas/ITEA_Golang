package handler

import (
	"bytes"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCommentHandlerGet(t *testing.T) {
	req := &http.Request{
		Method: http.MethodGet,
	}
	rw := httptest.NewRecorder()

	handler := NewCommentHandler(nil)
	handler.comments = []*Comment{
		{
			Id:      uuid.MustParse("6a4019b2-8791-46e1-8cf8-c270c5d1bdcf"),
			Content: "test",
		},
		{
			Id:      uuid.MustParse("a4c873d4-8866-4980-9b7f-3abe0e038ed3"),
			Content: "test 2",
		},
	}

	handler.ServeHTTP(rw, req)

	expectedBody := `[{"id":"6a4019b2-8791-46e1-8cf8-c270c5d1bdcf","content":"test"},{"id":"a4c873d4-8866-4980-9b7f-3abe0e038ed3","content":"test 2"}]`

	if strings.Trim(rw.Body.String(), "\n") != expectedBody {
		t.Errorf("invalid response body: got: %s, want: %s", rw.Body, expectedBody)
	}
}

func TestCommentHandlerPost(t *testing.T) {
	req := &http.Request{
		Method: http.MethodPost,
		Body:   io.NopCloser(bytes.NewBufferString(`{"id":"5360e351-82d6-4055-9b07-abfb71498a0f","content":"test"}`)),
	}
	rw := httptest.NewRecorder()

	handler := NewCommentHandler(nil)
	handler.ServeHTTP(rw, req)

	if rw.Code != http.StatusCreated {
		t.Errorf("invalid response status code: got: %d, want: %d", rw.Code, http.StatusCreated)
	}
}
