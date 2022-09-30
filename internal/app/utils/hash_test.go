package utils_test

import (
	"github.com/aasumitro/karlota/internal/app/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "test hash make and verify",
			args:    args{s: "test"},
			want:    true,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := utils.Hash{}
			got := h.Make(tt.args.s)
			assert.Equalf(t, tt.want, h.Verify(tt.args.s, got), "Hash(%v)", tt.args.s)
		})
	}
}
