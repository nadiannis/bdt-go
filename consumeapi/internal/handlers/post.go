package handlers

import (
	"consumeapi/internal/domain"
	"consumeapi/internal/utils"
	"encoding/json"
	"net/http"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
	defer res.Body.Close()

	var posts []domain.Post
	err = json.NewDecoder(res.Body).Decode(&posts)
	if err != nil {
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, posts)
	if err != nil {
		utils.ServerErrorResponse(w, err)
		return
	}
}
