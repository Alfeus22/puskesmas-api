package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Alfeus22/puskesmas-api/internal/service"
)

type PasienController struct {
	service *service.PasienService
}

func NewPasienController(service *service.PasienService) *PasienController {
	return &PasienController{service: service}
}

// struktur json yg diharapkan dari postman/fe

type RegisterPasienRequest struct {
	NIK         string `json:"nik"`
	NamaLengkap string `json:"nama_lengkap"`
	AlergiObat  string `json:"alergi_obat"`
}

func (c *PasienController) RegisterPasien(w http.ResponseWriter, r *http.Request) {
	// buka paket json
	var req RegisterPasienRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Format JSON tidak valid"})
		return
	}

	// lempar ke service
	var alergiPtr *string

	if req.AlergiObat != "" {
		tmp := req.AlergiObat
		alergiPtr = &tmp
	}
	result, err := c.service.DaftarPasienBaru(req.NIK, req.NamaLengkap, alergiPtr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	// balasan sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Pasien berhasil didaftarkan", "data": result})
}
