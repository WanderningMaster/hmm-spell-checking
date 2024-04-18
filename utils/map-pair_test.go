package utils

import (
	"testing"
)

func compareTupleSlice(got []Tuple, tuple []Tuple) bool {
	for tup := range got {
		if got[tup] != tuple[tup] {
			return false
		}
	}

	return true
}

func TestMapWordPair(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    []Tuple
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				line: "foo bar",
			},
			want: []Tuple{
				{first: 'f', second: 'b'},
				{first: 'o', second: 'a'},
				{first: 'o', second: 'r'},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				line: "foo fo",
			},
			want: []Tuple{
				{first: 'f', second: 'f'},
				{first: 'o', second: 'o'},
				{first: 'o', second: ' '},
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				line: "fo foo",
			},
			want: []Tuple{
				{first: 'f', second: 'f'},
				{first: 'o', second: 'o'},
				{first: ' ', second: 'o'},
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				line: "foo bar baz",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MapWordPair(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapWordPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !compareTupleSlice(got, tt.want) {
				t.Errorf("MapWordPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
