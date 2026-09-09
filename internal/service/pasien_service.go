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
		ID:          id,
		NIK:         nik,
		NamaLengkap: nama,
		AlergiObat:  alergi,
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
func (s *PasienService) ListPasien(page int, pageSize int) ([]*storage.Pasien, int, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	// ambil data dari storage
	data, err := s.storage.GetAllPasienTx(pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	// hitung total data untuk info paginasi
	total, err := s.storage.CountPasien()
	if err != nil {
		return nil, 0, err
	}
	return data, total, err
}

func (s *PasienService) SoftDeletePasien(id string) (*storage.Pasien, error) {
	tx, err := s.storage.StartTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	deletePasien := &storage.Pasien{
		ID: id,
	}
	err = s.storage.SoftDeleteTx(tx, deletePasien.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return deletePasien, nil
}

func (s *PasienService) GetById(id string) (*storage.Pasien, error) {

	hasil, err := s.storage.GetByID(id)
	if err != nil {
		return nil, err
	}
	return hasil, nil

}
