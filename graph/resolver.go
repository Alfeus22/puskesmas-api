package graph

import "github.com/Alfeus22/puskesmas-api/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	PasienService *service.PasienService
}
