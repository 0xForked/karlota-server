package utils_test

import (
	"fmt"
	"github.com/aasumitro/karlota/src/utils"
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
			got, err := h.Make(tt.args.s)

			if !tt.wantErr(t, err, fmt.Sprintf("Hash(%v)", tt.args.s)) {
				return
			}

			assert.Equalf(t, tt.want, h.Verify(tt.args.s, got), "Hash(%v)", tt.args.s)
		})
	}
}
