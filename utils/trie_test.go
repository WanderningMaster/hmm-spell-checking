package utils

import "testing"

func TestTrie(t *testing.T) {
	words := []string{
		"cat",
		"orchestra",
	}
	trie := NewTrie()

	for _, w := range words {
		trie.Insert(w)
	}

	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				word: "cat",
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				word: "Cat",
			},
			want: false,
		},
		{
			name: "Test 3",
			args: args{
				word: "orc",
			},
			want: false,
		},
		{
			name: "Test 4",
			args: args{
				word: "orchestra",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trie.Search(tt.args.word)
			if got != tt.want {
				t.Errorf("Trie.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
