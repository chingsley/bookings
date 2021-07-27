package main

import (
	"fmt"
	"testing"

	"github.com/chingsley/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing; test passed

	default:
		t.Error(fmt.Sprintf("expected type *chi.Mux, but got type %T", v))
	}
}
