package functions

import "testing"

func TestStandardDeviation(t *testing.T) {
	type args struct {
		variance float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive variance",
			args: args{variance: 4},
			want: 2,
		},
		{
			name: "Zero variance",
			args: args{variance: 0},
			want: 0,
		},
		{
			name: "Small variance",
			args: args{variance: 0.25},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardDeviation(tt.args.variance); got != tt.want {
				t.Errorf("StandardDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
