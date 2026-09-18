package storage

import (
	"database/sql"
	"errors"

	customErr "github.com/Alfeus22/puskesmas-api/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type Dokter struct {
	ID   string `db:"id"`
	NID  string `db:"nid"`
	Nama string `db:"nama"`
}

type DokterStorage struct {
	db *sqlx.DB
}

func NewDokterStorage(db *sqlx.DB) *DokterStorage {
	return &DokterStorage{db: db}
}

func (s *DokterStorage) StartTransactionTx() (*sqlx.Tx, error) {
	return s.db.Beginx()
}
func (s *DokterStorage) FinishTransactionTx(tx *sqlx.Tx, err error) error {
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *DokterStorage) GetAllDokter(limit, offset int) ([]*Dokter, error) {
	var dokter []*Dokter
	query := "SELECT id,nid,nama FROM dokter WHERE is_deleted = 0 LIMIT ? OFFSET ?"
	err := s.db.Select(&dokter, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return dokter, nil
}

func (s *DokterStorage) CreateDokterTx(tx *sqlx.Tx, Dokter *Dokter) error {
	query := "INSERT into dokter (id,nid,nama) VALUES(:id,:nid,:nama)"
	_, err := tx.NamedExec(query, Dokter)
	return err
}

func (s *DokterStorage) GetDokterById(id string) (*Dokter, error) {
	var dokter Dokter
	query := "SELECT id,nid,nama FROM dokter where id = ? and is_deleted = 0 "
	err := s.db.Get(&dokter, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErr.ErrNotFound
		}
		return nil, err
	}
	return &dokter, nil
}

func (s *DokterStorage) CountDokter() (int, error) {
	var jumlahDokter int
	query := "SELECT COUNT(*) FROM dokter WHERE is_deleted = 0"
	err := s.db.Select(&jumlahDokter, query)
	if err != nil {
		return 0, err
	}
	return jumlahDokter, nil
}
