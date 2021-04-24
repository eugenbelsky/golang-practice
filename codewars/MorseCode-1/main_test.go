package main

import "testing"

func TestDecodeMorse(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"... --- ..."}, "SOS"},
		{"test2", args{"...---..."}, "SOS"},
		{"test3", args{".-. --- --. . .-."}, "ROGER"},
		{"test4", args{".---."}, "ROGER"},
		{"test5", args{"-.. ---   -.-- --- ..-   .-.. .. -.- .   .-- .... .- -   -.-- --- ..-   ... . . ..--.. -.-.--"}, "DO YOU LIKE WHAT YOU SEE?!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeMorse(tt.args.sentence); got != tt.want {
				t.Errorf("DecodeMorse() = %v, want %v", got, tt.want)
			}
		})
	}
}
