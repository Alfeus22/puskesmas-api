package storage

import (
	"database/sql"

	"errors"

	customErr "github.com/Alfeus22/puskesmas-api/pkg/errors"
	"github.com/jmoiron/sqlx"
)

// DTO model
type Pasien struct {
	ID          string  `db:"id"`
	NIK         string  `db:"nik"`
	NamaLengkap string  `db:"nama_lengkap"`
	AlergiObat  *string `db:"alergi_obat"` //pointer agar bisa null
	IsDeleted   int     `db:"is_deleted"`
}

type PasienStorage struct {
	db *sqlx.DB
}

func NewPasienStorage(db *sqlx.DB) *PasienStorage {
	return &PasienStorage{db: db}
}

// fungsi read yg hanya mengambil yg is_deleted = 0
func (s *PasienStorage) GetByID(id string) (*Pasien, error) {
	var pasien Pasien

	query := "SELECT id,nik,nama_lengkap,alergi_obat FROM pasien where id = ? and is_deleted = 0"
	err := s.db.Get(&pasien, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErr.ErrNotFound
		}
		return nil, err
	}
	return &pasien, nil
}

// fungsi delete diubah menjadi soft delete
func (s *PasienStorage) SoftDeleteTx(tx *sqlx.Tx, id string) error {
	query := "UPDATE pasien SET is_deleted = 1 WHERE id = ?"
	_, err := tx.Exec(query, id)
	if err != nil {
		return errors.New(err.Error())
	}
	return err
}

// start transaction
func (s *PasienStorage) StartTransaction() (*sqlx.Tx, error) {
	return s.db.Beginx()
}

// finish transaction akan otomstis roolback jika ada error, atau commit jika sukses
func (s *PasienStorage) FinishTransaction(tx *sqlx.Tx, err error) error {
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// fungsi create
func (s *PasienStorage) CreatePasienTx(tx *sqlx.Tx, pasien *Pasien) error {
	query := `INSERT into pasien (id,nik,nama_lengkap,alergi_obat) VALUES (:id,:nik,:nama_lengkap,:alergi_obat)`

	_, err := tx.NamedExec(query, pasien)
	return err
}

// fungsi update/ modifikasi pasien
func (s *PasienStorage) UpdatePasienTx(tx *sqlx.Tx, id string, pasien *Pasien) error {
	query := `UPDATE pasien SET nik = ?, nama_Lengkap = ?, alergi_obat = ? WHERE id = ?`
	_, err := tx.Exec(query, pasien.NIK, pasien.NamaLengkap, pasien.AlergiObat, id)
	return err
}
