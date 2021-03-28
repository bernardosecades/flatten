package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	flatten "github.com/bernardosecades/flatten/internal"
)

const LimitRows = 100

type Handler interface {
	Flatten(w http.ResponseWriter, r *http.Request)
	History(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	repository flatten.Repository
}

func NewHandler(h flatten.Repository) Handler {
	return &handler{repository: h}
}

func (h *handler) History(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	histories, err := h.repository.GetHistoryByLimit(LimitRows)
	if err != nil {
		encodeCustomError(ErrorResponse{StatusCode: http.StatusInternalServerError, Err: err.Error()}, w)
		return
	}

	response, err := json.Marshal(histories)
	if err != nil {
		encodeCustomError(ErrorResponse{StatusCode: http.StatusInternalServerError, Err: err.Error()}, w)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		encodeCustomError(ErrorResponse{StatusCode: http.StatusInternalServerError, Err: err.Error()}, w)
	}
}

func (h *handler) Flatten(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		encodeCustomError(ErrorResponse{StatusCode: http.StatusInternalServerError, Err: err.Error()}, w)
		return
	}

	var request Request
	_ = json.Unmarshal(body, &request)

	if err := request.Validate(); err != nil {
		encodeCustomError(ErrorResponse{StatusCode: http.StatusBadRequest, Err: err.Error()}, w)
		return
	}

	result, depth := flatten.ApplyFlatten(request.Input)

	response, err := json.Marshal(Response{Output: result, Depth: depth})
	if err != nil {
		encodeCustomError(ErrorResponse{StatusCode: http.StatusInternalServerError, Err: err.Error()}, w)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		encodeCustomError(ErrorResponse{StatusCode: http.StatusInternalServerError, Err: err.Error()}, w)
		return
	}

	bytesInput, err := json.Marshal(request.Input)
	if err != nil {
		log.Println(err)
		return
	}

	bytesOutput, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		return
	}

	if err = h.repository.SaveHistory(bytesInput, bytesOutput, depth); err != nil {
		log.Println(err)
	}
}

func encodeCustomError(customError ErrorResponse, w http.ResponseWriter) {

	log.Println(customError.Err)
	w.WriteHeader(customError.StatusCode)
	response, err := json.Marshal(customError)

	if err != nil {
		log.Println("Error to encode customError")
	}
	_, err = w.Write(response)

	if err != nil {
		log.Println("Error to write customError", err)
	}
}
