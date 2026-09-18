package service

import (
	"database/sql"
	"errors"

	"github.com/Alfeus22/puskesmas-api/internal/storage"
	errCustom "github.com/Alfeus22/puskesmas-api/pkg/errors"
	"github.com/google/uuid"
	// "github.com/google/uuid"
)

type DokterService struct {
	storage *storage.DokterStorage
}

func NewDokterService(storage *storage.DokterStorage) *DokterService {
	return &DokterService{storage: storage}
}

func (s *DokterService) DaftarDokterBaru(nid string, nama string) (*storage.Dokter, error) {
	tx, err := s.storage.StartTransactionTx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dokterBaru := &storage.Dokter{
		ID:   uuid.NewString(),
		NID:  nid,
		Nama: nama,
	}

	err = s.storage.CreateDokterTx(tx, dokterBaru)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return dokterBaru, nil
}

func (s *DokterService) GetAllDokter(page int, pageSize int) ([]*storage.Dokter, int, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * pageSize

	data, err := s.storage.GetAllDokter(pageSize, offset)
	if err != nil {
		return nil, 0, err
	}

	// hitung total untuk pagiantion
	total, err := s.storage.CountDokter()
	if err != nil {
		return nil, 0, err
	}
	return data, total, err
}

func (s *DokterService) GetDokterById(id string) (*storage.Dokter, error) {
	hasil, err := s.GetDokterById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errCustom.ErrNotFound
		}
	}

	return hasil, nil
}
