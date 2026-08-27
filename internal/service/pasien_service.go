package service

import (
	"github.com/Alfeus22/puskesmas-api/internal/storage"
	"github.com/Alfeus22/puskesmas-api/pkg/utils"
	"github.com/google/uuid"
)

type PasienService struct {
	storage *storage.PasienStorage
}

func NewPasienService(storage *storage.PasienStorage) *PasienService {
	return &PasienService{storage: storage}
}

// (err error) ini namanay return named
func (s *PasienService) DaftarPasienBaru(nik string, nama string, alergi string) (err error) {
	// mulai transaksi
	tx, err := s.storage.StartTransaction()
	if err != nil {
		return err
	}
	// selesaikan transaksi
	defer func() {
		err = s.storage.FinishTransaction(tx, err)
	}()
	// persiapkan DTO pasien
	pasienBaru := storage.Pasien{
		ID:          uuid.NewString(), //generate uuid otomatis
		NIK:         nik,
		NamaLengkap: nama,
		// jika alergi == "" maka akan berubah jadi 'nil' dan di databas eakna jadi NULL
		AlergiObat: utils.NullableString(alergi),
	}

	// eksekusi ke gudang
	err = s.storage.CreatePasienTx(tx, &pasienBaru)
	if err != nil {
		return err
	}
	return nil
}
