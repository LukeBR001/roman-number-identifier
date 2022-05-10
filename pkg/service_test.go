package pkg

import (
	"github.com/google/go-cmp/cmp"
	"identifier/pkg/model"
	"testing"
)

func TestIdentifyCombinations(t *testing.T) {
	type args struct {
		textPayload model.TextPayload
	}

	validComb := model.TextPayload{Text: "XIVQWECQWEMMDCCXXIVQWELXQWELXXXQWEXLQWEMmDcCcxXIv"}
	//"XIV C MMDCCXXIV LX LXXX XL MmDcCcxXIv"

	justInvalidComb := model.TextPayload{Text: "XIVIQWEVCQWEMMDCCDXXIVQWELXLQWELXXXXQWEXLLQWEMmDcCcxXIvX"}
	randomComb := model.TextPayload{Text: "XIVQWEIIIIQWEMMDCCXXIVQWELLXQWELXXXQWEIXLQWEMmDcCcxXIvQWEICQWEXXL"}
	//"XIV MMDCCXXIV LXXX MMDCCCXXIV"

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "should return all valid combinations when successful",
			args: args{
				textPayload: validComb,
			},
			want:    []string{"XIV", "C", "MMDCCXXIV", "LX", "LXXX", "XL", "MMDCCCXXIV"},
			wantErr: false,
		},
		{
			name: "should return error when not found any valid roman combinations",
			args: args{
				textPayload: justInvalidComb,
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "should return just valid roman combinations",
			args: args{
				textPayload: randomComb,
			},
			want:    []string{"XIV", "MMDCCXXIV", "LXXX", "MMDCCCXXIV"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IdentifyCombinations(tt.args.textPayload)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentifyCombinations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Got error, Diff = %s", diff)
			}
		})
	}
}

func Test_isValidRomanSequence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true when sequence is a valid roman combination",
			args: args{
				s: "MMMDCCCLXXXIV",
			},
			want: true,
		},
		{
			name: "should return false when sequence not is an valid roman combination",
			args: args{
				s: "MMMDDCCCLXXXIV",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidRomanSequence(tt.args.s); got != tt.want {
				t.Errorf("isValidRomanSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdentifyBiggerNumber(t *testing.T) {
	type args struct {
		romanList []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return the bigger number in an roman sequence",
			args: args{
				romanList: []string{"XXX", "IV", "LX", "MMD", "MMCM", "MMCMX"},
			},
			// MMCMX = 2910
			want: 2910,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IdentifyBiggerNumber(tt.args.romanList); got != tt.want {
				t.Errorf("IdentifyBiggerNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRepeatable(t *testing.T) {
	type args struct {
		symbol string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true when is a roman number",
			args: args{
				symbol: "X",
			},
			want: true,
		},
		{
			name: "should return false when not is a roman number",
			args: args{
				symbol: "E",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRepeatable(tt.args.symbol); got != tt.want {
				t.Errorf("isRepeatable() = %v, want %v", got, tt.want)
			}
		})
	}
}
