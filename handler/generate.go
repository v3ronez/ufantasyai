package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/v3ronez/ufantasyai/db"
	"github.com/v3ronez/ufantasyai/types"
	"github.com/v3ronez/ufantasyai/view/generate"
)

func HandlerGenerateIndex(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	images, err := db.GetImagesFromUserId(user.ID)
	if err != nil {
		return err
	}
	viewData := generate.ViewData{Images: images}
	return RenderComponent(w, r, generate.Index(viewData))
}
func HandlerGenerateImageCreate(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	image := types.Image{
		ID:            2,
		UserId:        user.ID,
		Prompt:        "beautiful womem",
		Status:        types.ImageStatusPending,
		ImageLocation: "https://img.freepik.com/premium-photo/cute-girl-with-pretty-face-creative-ai_634423-2810.jpg",
	}

	if err := db.CreateImage(&image); err != nil {
		return err
	}
	return RenderComponent(w, r, generate.GalleryImage(image))
}

func HandlerGenerateImageStatus(w http.ResponseWriter, r *http.Request) error {
	imageID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return err
	}
	//fetch from db
	image, err := db.GetImageForID(imageID)
	if err != nil {
		return err
	}
	return RenderComponent(w, r, generate.GalleryImage(image))
}
