package mri

// TODO: большая проблема: то что рисуем на выход в сваггер != тому что туда реально уходит (уходит GRPC)
// TODO: сделать выход контуров через json nyy
import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	adapters "gateway/internal/adapters"
	"gateway/internal/domain"

	mriuploadpb "gateway/internal/generated/broker/produce/mriupload"
	mripb "gateway/internal/generated/grpc/client/mri"
	"gateway/internal/repository"

	"github.com/gorilla/mux"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	adapter adapters.Adapter
	dao     repository.DAO
}

func New(
	adapter adapters.Adapter,
	dao repository.DAO,
) *Handler {
	return &Handler{
		adapter: adapter,
		dao:     dao,
	}
}

// PostMri загружает узи на обработку
//
//	@Summary		Загружает узи на обработку
//	@Description	Загружает узи на обработку
//	@Tags			mri
//	@Produce		json
//	@Param			token		header		string	true	"access_token"
//	@Param			file		formData	file	true	"mri file. (обязательно с .tiff/.png)"
//	@Param			projection	formData	string	true	"проекция узи"
//	@Param			patient_id	formData	string	true	"id пациента"
//	@Param			device_id	formData	string	true	"id узи апапапапарата"
//	@Success		200			{string}	string	"molodec"
//	@Failure		500			{string}	string	"Internal Server Error"
//	@Router			/mri/mris [post]
func (h *Handler) PostMri(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projection := r.FormValue("projection")
	patientID := r.FormValue("patient_id")
	deviceID, _ := strconv.Atoi(r.FormValue("device_id"))

	mriResp, err := h.adapter.MriAdapter.CreateMri(ctx, &mripb.CreateMriIn{
		Projection: projection,
		PatientId:  patientID,
		DeviceId:   int64(deviceID),
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	file, meta, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
	defer file.Close()
	ext := filepath.Ext(meta.Filename)

	// TODO: заюзать библу
	mime, err := domain.ParseFormatFromExt(ext)
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	err = h.dao.NewFileRepo().LoadFile(ctx, filepath.Join(mriResp.Id, mriResp.Id), domain.File{Format: mime, Buf: file})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	// TODO: нужна тотальная сага тут
	if err := h.adapter.BrokerAdapter.SendMriUpload(&mriuploadpb.MriUpload{MriId: mriResp.Id}); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(mriResp.Id); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// TODO: проверить крайние случае, если что то не приходит например(неправильный id)
// TODO: убрать echographic из ответа на обновление
// PatchMri Обновляет узи
//
//	@Summary		Обновляет узи
//	@Description	Обновляет узи
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string		true	"access_token"
//	@Param			id		path		string		true	"mri_id"
//	@Param			body	body		PatchMriIn	true	"обновляемые значения"
//	@Success		200		{object}	PatchMriOut	"mri"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/mri/mris/{id} [patch]
func (h *Handler) PatchMri(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var req PatchMriIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	resp, err := h.adapter.MriAdapter.UpdateMri(ctx, &mripb.UpdateMriIn{
		Id:         id,
		Projection: req.Projection,
		Checked:    req.Checked,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp.Mri); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// PatchEchographics Обновляет эхографику
//
//	@Summary		Обновляет эхографику
//	@Description	Обновляет эхографику
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string					true	"access_token"
//	@Param			id		path		string					true	"mri_id"
//	@Param			body	body		PatchEchographicsIn		true	"обновляемые значения"
//	@Success		200		{object}	PatchEchographicsOut	"echographic"
//	@Failure		500		{string}	string					"Internal Server Error"
//	@Router			/mri/echographics/{id} [patch]
func (h *Handler) PatchEchographics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var req PatchEchographicsIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	resp, err := h.adapter.MriAdapter.UpdateEchographic(ctx, &mripb.UpdateEchographicIn{
		Echographic: &mripb.Echographic{
			Id:              id,
			LeftLobeLength:  req.LeftLobeLength,
			LeftLobeWidth:   req.LeftLobeWidth,
			LeftLobeThick:   req.LeftLobeThick,
			LeftLobeVolum:   req.LeftLobeVolum,
			RightLobeLength: req.RightLobeLength,
			RightLobeWidth:  req.RightLobeWidth,
			RightLobeThick:  req.RightLobeThick,
			RightLobeVolum:  req.RightLobeVolum,
			GlandVolum:      req.GlandVolum,
			Isthmus:         req.Isthmus,
			Struct:          req.Struct,
			Echogenicity:    req.Echogenicity,
			RegionalLymph:   req.RegionalLymph,
			Vascularization: req.Vascularization,
			Location:        req.Location,
			Additional:      req.Additional,
			Conclusion:      req.Conclusion,
		},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp.Echographic); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetMri получает mri
//
//	@Summary		получает mri
//	@Description	получает mri
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string		true	"access_token"
//	@Param			id		path		string		true	"mri_id"
//	@Success		200		{object}	GetMriOut	"mri"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/mri/mris/{id} [get]
func (h *Handler) GetMri(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.MriAdapter.GetMri(ctx, &mripb.GetMriIn{Id: id})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
	// TODO: понять почему тут узи возвращается без эхографикой, а тут с
	// TODO: подумать над content-tpye в ответе(посмотреть в каком порядке выставлять функции для ответа)
	if err := json.NewEncoder(w).Encode(resp.Mri); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetPatientMri Получить узи пациента
//
//	@Summary		Получить узи пациента
//	@Description	Получить узи пациента
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Param			id		path		string	true	"patient_id"
//	@Success		200		{object}	GetPatientMriOut
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/mri/patient/{id}/mris [get]
func (h *Handler) GetPatientMri(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	res, err := h.adapter.MriAdapter.GetPatientMris(ctx, &mripb.GetPatientMrisIn{
		PatientId: id,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	w.WriteHeader(200)
}

// GetEchographics получает mri
//
//	@Summary		получает эхографику mri
//	@Description	получает эхографику mri
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string				true	"access_token"
//	@Param			id		path		string				true	"mri_id"
//	@Success		200		{object}	GetEchographicsOut	"echographics"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/mri/echographics/{id} [get]
func (h *Handler) GetEchographics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.MriAdapter.GetEchographic(ctx, &mripb.GetEchographicIn{Id: id})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
	// TODO: понять почему тут узи возвращается без эхографикой, а тут с
	// TODO: подумать над content-tpye в ответе(посмотреть в каком порядке выставлять функции для ответа)
	if err := json.NewEncoder(w).Encode(resp.Echographic); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetMriImages получает id картинок mri
//
//	@Summary		получает списк id кадров mri
//	@Description	получает списк id кадров mri
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string			true	"access_token"
//	@Param			id		path		string			true	"mri_id"
//	@Success		200		{object}	GetMriImagesOut	"images"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/mri/mris/{id}/images [get]
func (h *Handler) GetMriImages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.MriAdapter.GetMriImages(ctx, &mripb.GetMriImagesIn{MriId: id})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp.Images); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetMriNodeSegments получит ноды и сегменты на указанном изображении
//
//	@Summary		получит ноды и сегменты на указанном изображении
//	@Description	получит ноды и сегменты на указанном изображении
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string					true	"access_token"
//	@Param			id		path		string					true	"image_id"
//	@Success		200		{object}	GetMriNodeSegmentsOut	"nodes&&segments"
//	@Failure		500		{string}	string					"Internal Server Error"
//	@Router			/mri/images/{id}/nodes-segments [get]
func (h *Handler) GetMriNodeSegments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.MriAdapter.GetImageSegmentsWithNodes(
		ctx,
		&mripb.GetImageSegmentsWithNodesIn{Id: id},
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetMriDevice получит список mri аппаратов
//
//	@Summary		получит список mri аппаратов
//	@Description	получит список mri аппаратов
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string			true	"access_token"
//	@Success		200		{object}	GetMriDeviceOut	"mri аппараты"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/mri/devices [get]
func (h *Handler) GetMriDevices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp, err := h.adapter.MriAdapter.GetDeviceList(ctx, &emptypb.Empty{})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// PostNodes добавить узел с сегментами
//
//	@Summary		добавить узел с сегментами
//	@Description	добавить узел с сегментами
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string		true	"access_token"
//	@Param			node	body		PostNodeIn	true	"узел с сегментами"
//	@Success		200		{object}	PostNodeOut	"id узла"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/mri/nodes [post]
func (h *Handler) PostNodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req PostNodeIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	segments := make([]*mripb.CreateNodeIn_NestedSegment, 0, len(req.Segments))
	for _, v := range req.Segments {
		segments = append(segments, &mripb.CreateNodeIn_NestedSegment{
			ImageId:   v.ImageID.String(),
			Contor:    v.Contor,
			Knosp_012: v.Knosp012,
			Knosp_3:   v.Knosp3,
			Knosp_4:   v.Knosp4,
		})
	}

	resp, err := h.adapter.MriAdapter.CreateNode(ctx, &mripb.CreateNodeIn{
		MriId:     req.MriID.String(),
		Segments:  segments,
		Knosp_012: req.Knosp012,
		Knosp_3:   req.Knosp3,
		Knosp_4:   req.Knosp4,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// GetAllNodes получить все узлы узи
//
//	@Summary		получить все узлы узи
//	@Description	получить все узлы узи
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string		true	"access_token"
//	@Param			id		path		string					true	"mri_id"
//	@Success		200		{object}	GetAllNodesOut	"id узла"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/mri/mris/{id}/nodes [get]
func (h *Handler) GetAllNodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	// TODO: в ответе пустые поля будут опущены, убрать теги omitempty.
	resp, err := h.adapter.MriAdapter.GetAllNodes(ctx, &mripb.GetAllNodesIn{
		MriId: id,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// DeleteNode удалит узел
//
//	@Summary		удалит узел
//	@Description	удалит узел
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Success		200		{string}	string	"molodec"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/mri/nodes/{id} [delete]
func (h *Handler) DeleteNode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.MriAdapter.DeleteNode(ctx, &mripb.DeleteNodeIn{Id: id})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// PatchNode обновит узел
//
//	@Summary		обновит узел
//	@Description	обновит узел
//	@Tags			mri
//	@Produce		json
//	@Param			node	body		PatchNodeIn		true	"узел с сегментами"
//	@Success		200		{object}	PatchNodeOut	"обновленный узел"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/mri/nodes/{id} [patch]
func (h *Handler) PatchNode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var req PatchNodeIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	resp, err := h.adapter.MriAdapter.UpdateNode(ctx, &mripb.UpdateNodeIn{
		Id:        id,
		Knosp_012: req.Knosp012,
		Knosp_3:   req.Knosp3,
		Knosp_4:   req.Knosp4,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp.Node); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// TODO: нет валидации что если ноды нет
// PostSegment добавит новый сегмент к указанному узлу
//
//	@Summary		добавит новый сегмент к указанному узлу
//	@Description	добавит новый сегмент к указанному узлу
//	@Tags			mri
//	@Produce		json
//	@Param			node	body		PostSegmentIn	true	"сегмент"
//	@Success		200		{object}	PostSegmentOut	"id узла"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/mri/segments [post]
func (h *Handler) PostSegment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req PostSegmentIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	resp, err := h.adapter.MriAdapter.CreateSegment(ctx, &mripb.CreateSegmentIn{
		ImageId:   req.ImageID.String(),
		NodeId:    req.NodeID.String(),
		Contor:    req.Contor,
		Knosp_012: req.Knosp012,
		Knosp_3:   req.Knosp3,
		Knosp_4:   req.Knosp4,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// DeleteSegment удалит сегмент
//
//	@Summary		удалит сегмент
//	@Description	удалит сегмент, ЕСЛИ У УЗЛА НЕ ОСТАНЕТСЯ СЕГМЕНТОВ, ОН ТОЖЕ БУДЕТ УДАЛЕН
//	@Tags			mri
//	@Produce		json
//	@Param			token	header		string	true	"access_token"
//	@Success		200		{string}	string	"molodec"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/mri/segments/{id} [delete]
func (h *Handler) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	resp, err := h.adapter.MriAdapter.DeleteSegment(ctx, &mripb.DeleteSegmentIn{Id: id})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// PatchSegment обновит сегмент
//
//	@Summary		обновит сегмент
//	@Description	обновит сегмент
//	@Tags			mri
//	@Produce		json
//	@Param			node	body		PatchSegmentIn	true	"узел с сегментами"
//	@Success		200		{object}	PatchSegmentOut	"обновленный узел"
//	@Failure		500		{string}	string			"Internal Server Error"
//	@Router			/mri/segments/{id} [patch]
func (h *Handler) PatchSegment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := mux.Vars(r)["id"]

	var req PatchSegmentIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	resp, err := h.adapter.MriAdapter.UpdateSegment(ctx, &mripb.UpdateSegmentIn{
		Id:        id,
		Knosp_012: req.Knosp012,
		Knosp_3:   req.Knosp3,
		Knosp_4:   req.Knosp4,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(resp.Segment); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}
