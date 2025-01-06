package functions

import "testing"

func TestMedian(t *testing.T) {
	type args struct {
		intData []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Odd number of elements",
			args: args{intData: []float64{1, 3, 3, 6, 7, 8, 9}},
			want: 6,
		},
		{
			name: "Even number of elements",
			args: args{intData: []float64{1, 2, 3, 4, 5, 6, 8, 9}},
			want: 4.5,
		},
		{
			name: "Single element",
			args: args{intData: []float64{10}},
			want: 10,
		},
		{
			name: "Two elements",
			args: args{intData: []float64{1, 2}},
			want: 1.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.intData); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}
