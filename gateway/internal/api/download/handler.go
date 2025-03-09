package download

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"gateway/internal/repository"

	"github.com/gorilla/mux"
)

type Handler struct {
	dao repository.DAO
}

func New(
	dao repository.DAO,
) *Handler {
	return &Handler{
		dao: dao,
	}
}

// GetUzi Получение мрт
//
//	@Summary		Получение мрт
//	@Description	Получение мрт
//	@Tags			download
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Param			mri_id	path		string	true	"id мрт"
//	@Success		200		{file}		File	"Изображение МРТ"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/download/mri/{id} [get]
func (h *Handler) GetUzi(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	file, err := h.dao.NewFileRepo().GetFile(ctx, filepath.Join(id, id))
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	// TODO: переписать dao, Добавить content-type
	w.Header().Set("Content-Type", "image/tiff")
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, fmt.Sprintf("не удалось вернуть изображение: %v", err), 500)
		return
	}
}

// GetImage Получение image mri
//
//	@Summary		Получение image mri
//	@Description	Получение image mri
//	@Tags			download
//	@Produce		json
//	@Param			token		header		string	true	"access_token"
//	@Param			mri_id		path		string	true	"id мрт"
//	@Param			image_id	path		string	true	"id image"
//	@Success		200			{file}		File	"Изображение кадра Узи"
//	@Failure		500			{string}	string	"Internal Server Error"
//	@Router			/download/mri/{mri_id}/image/{image_id} [get]
func (h *Handler) GetImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uziID := mux.Vars(r)["mri_id"]
	imageID := mux.Vars(r)["image_id"]

	file, err := h.dao.NewFileRepo().GetFile(
		ctx,
		filepath.Join(
			uziID,
			imageID,
			imageID,
		),
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	// TODO: переписать dao, Добавить content-type
	w.Header().Set("Content-Type", "image/png")
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, fmt.Sprintf("не удалось вернуть изображение: %v", err), 500)
		return
	}
}
