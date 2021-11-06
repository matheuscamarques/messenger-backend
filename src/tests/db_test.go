package test

import (
	"reflect"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/matheuscamarques/messenger-backend/main"
)

func TestDBConnection(t *testing.T) {
	tests := []struct {
		name string
		want *pg.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main.DBConnection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DBConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
