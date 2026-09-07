package service

import (
	"github.com/Alfeus22/puskesmas-api/internal/storage"
	"github.com/google/uuid"
)

type PasienService struct {
	storage *storage.PasienStorage
}

func NewPasienService(storage *storage.PasienStorage) *PasienService {
	return &PasienService{storage: storage}
}

// (err error) ini namanay return named
func (s *PasienService) DaftarPasienBaru(nik string, nama string, alergi *string) (*storage.Pasien, error) {
	// mulai transaksi
	tx, err := s.storage.StartTransaction()
	if err != nil {
		return nil, err
	}
	// selesaikan transaksi
	defer tx.Rollback()

	// persiapkan DTO pasien
	pasienBaru := &storage.Pasien{
		ID:          uuid.NewString(), //generate uuid otomatis
		NIK:         nik,
		NamaLengkap: nama,
		// jika alergi == "" maka akan berubah jadi 'nil' dan di databas eakna jadi NULL
		AlergiObat: alergi,
	}

	// eksekusi ke gudang
	err = s.storage.CreatePasienTx(tx, pasienBaru)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return pasienBaru, nil
}

func (s *PasienService) UpdatePasien(id string, nik string, nama string, alergi *string) (*storage.Pasien, error) {
	tx, err := s.storage.StartTransaction()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	// dto update pasien
	updatePasien := &storage.Pasien{
		NIK:         nik,
		NamaLengkap: nama,
		AlergiObat:  *&alergi,
	}
	err = s.storage.UpdatePasienTx(tx, id, updatePasien)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return updatePasien, nil

}
