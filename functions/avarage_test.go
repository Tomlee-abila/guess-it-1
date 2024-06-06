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
		{
			name: "average of positive numbers",
			args: args{
				sum:  15.0,
				data: []float64{3, 4, 5, 3},
			},
			want: 3.75,
		},
		{
			name: "average with zero sum",
			args: args{
				sum:  0.0,
				data: []float64{0, 0, 0, 0},
			},
			want: 0.0,
		},
		{
			name: "average of mixed numbers",
			args: args{
				sum:  1.0,
				data: []float64{-3, 4, -2, 2},
			},
			want: 0.25,
		},
		{
			name: "average with empty data",
			args: args{
				sum:  0.0,
				data: []float64{},
			},
			want: 0.0,
		},
		{
			name: "average of single element",
			args: args{
				sum:  5.0,
				data: []float64{5},
			},
			want: 5.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Avarage(tt.args.sum, tt.args.data); got != tt.want {
				t.Errorf("Avarage() = %v, want %v", got, tt.want)
			}
		})
	}
}
