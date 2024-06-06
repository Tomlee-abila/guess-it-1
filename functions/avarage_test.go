package functions

import "testing"

func TestAvarage(t *testing.T) {
	type args struct {
		sum  float64
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Avarage(tt.args.sum, tt.args.data); got != tt.want {
				t.Errorf("Avarage() = %v, want %v", got, tt.want)
			}
		})
	}
}
