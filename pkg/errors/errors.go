package errors

import (
	"errors"
)

var (
	ErrNotFound = errors.New("data tidak ditemukan")
	ErrConflict = errors.New("data sudah ada atau bentrok")
	ErrInternal = errors.New("terjadi kesalahn pada server")
)
