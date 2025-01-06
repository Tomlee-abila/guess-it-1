package functions

import "testing"

func TestVariance(t *testing.T) {
	type args struct {
		avarage float64
		data    []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Typical case",
			args: args{avarage: 3, data: []float64{1, 2, 3, 4, 5}},
			want: 2,
		},
		{
			name: "Zero variance",
			args: args{avarage: 5, data: []float64{5, 5, 5, 5, 5}},
			want: 0,
		},
		{
			name: "Single element",
			args: args{avarage: 4, data: []float64{4}},
			want: 0,
		},
		{
			name: "Negative numbers",
			args: args{avarage: -3, data: []float64{-1, -2, -3, -4, -5}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.avarage, tt.args.data); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}
