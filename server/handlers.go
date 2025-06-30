package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/exp/slices"
)

func (s *Server) Emojis(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	emojis := s.Store.GetAll()
	json, _ := json.Marshal(emojis)

	w.Write(json)
}

func (s *Server) EmojisByCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	category := strings.Replace(ps.ByName("category"), "-", " ", -1)
	if !slices.Contains(s.Store.Categories, category) {
		HandleError(w, http.StatusNotFound, "emojis with this category do not exist")
		return
	}

	emojis := s.Store.GetAllByCategory(category)
	json, _ := json.Marshal(emojis)

	w.Write(json)
}

func (s *Server) EmojisByGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	group := strings.Replace(ps.ByName("group"), "-", " ", -1)
	if !slices.Contains(s.Store.Groups, group) {
		HandleError(w, http.StatusNotFound, "emojis with this group do not exist")
		return
	}

	emojis := s.Store.GetAllByGroup(group)
	json, _ := json.Marshal(emojis)

	w.Write(json)
}

func (s *Server) RandomEmoji(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	emoji := s.Store.GetRandom()
	json, _ := json.Marshal(emoji)

	w.Write(json)
}

func (s *Server) RandomEmojiByCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	category := strings.Replace(ps.ByName("category"), "-", " ", -1)
	if !slices.Contains(s.Store.Categories, category) {
		HandleError(w, http.StatusNotFound, "emojis with this category do not exist")
		return
	}

	emoji := s.Store.GetRandomByCategory(category)
	json, _ := json.Marshal(emoji)

	w.Write(json)
}

func (s *Server) RandomEmojiByGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	group := strings.Replace(ps.ByName("group"), "-", " ", -1)
	if !slices.Contains(s.Store.Groups, group) {
		HandleError(w, http.StatusNotFound, "emojis with this group do not exist")
		return
	}

	emoji := s.Store.GetRandomByGroup(group)
	json, _ := json.Marshal(emoji)

	w.Write(json)
}

func (s *Server) Categories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json, _ := json.Marshal(s.Store.Categories)

	w.Write(json)
}

func (s *Server) Groups(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json, _ := json.Marshal(s.Store.Groups)

	w.Write(json)
}

func (s *Server) Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		HandleError(w, http.StatusBadRequest, "query parameter 'q' is required")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	emojis := s.Store.Search(query)
	json, _ := json.Marshal(emojis)

	w.Write(json)
}

func (s *Server) Similar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	name := ps.ByName("name")
	if name == "" {
		HandleError(w, http.StatusBadRequest, "name parameter is required")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	emojis := s.Store.GetSimilar(name)
	json, _ := json.Marshal(emojis)

	w.Write(json)
}
