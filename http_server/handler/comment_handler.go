package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
)

type Comment struct {
	Id      uuid.UUID `json:"id"`
	Content string    `json:"content"`
}

type CommentHandler struct {
	logger   *log.Logger
	comments []*Comment
}

func NewCommentHandler(logger *log.Logger) *CommentHandler {
	return &CommentHandler{
		logger: logger,
	}
}

func (h *CommentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	switch r.Method {
	case http.MethodGet:
		if err := h.handleGet(w, r); err != nil {
			http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		}
	case http.MethodPost:
		if err := h.handlePost(w, r); err != nil {
			http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		}
	default:
		http.Error(w, `{"error": "invalid HTTP method"}`, http.StatusBadRequest)
	}
}

func (h *CommentHandler) handleGet(w http.ResponseWriter, r *http.Request) error {
	return json.NewEncoder(w).Encode(h.comments)
}

func (h *CommentHandler) handlePost(w http.ResponseWriter, r *http.Request) error {
	comment := new(Comment)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, comment); err != nil {
		return err
	}

	h.comments = append(h.comments, comment)

	w.WriteHeader(http.StatusCreated)

	return nil
}
