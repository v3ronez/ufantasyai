package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/v3ronez/ufantasyai/types"
	"github.com/v3ronez/ufantasyai/view/generate"
)

func HandlerGenerateIndex(w http.ResponseWriter, r *http.Request) error {
	images := []types.Image{}
	viewData := generate.ViewData{Images: images}
	return RenderComponent(w, r, generate.Index(viewData))
}
func HandlerGenerateImageCreate(w http.ResponseWriter, r *http.Request) error {
	return RenderComponent(w, r, generate.GalleryImage(types.Image{Status: types.ImageStatusPending}))
}

func HandlerGenerateImageStatus(w http.ResponseWriter, r *http.Request) error {
	imageID := chi.URLParam(r, "id")
	//fetch from db
	slog.Info("checking image status", "id", imageID)
	image := types.Image{Status: types.ImageStatusPending}
	return RenderComponent(w, r, generate.GalleryImage(image))
}
