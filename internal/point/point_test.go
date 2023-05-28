package point

import (
	"reflect"
	"testing"
)

func TestClip(t *testing.T) {
	type args struct {
		point *Point
		min   float64
		max   float64
	}
	tests := []struct {
		name string
		args args
		want *Point
	}{
		{
			name: "All coords less than min",
			args: args{
				point: &Point{
					x: -1.0,
					y: -1.0,
					z: -1.0,
				},
				min: 0.0,
				max: 1.0,
			},
			want: &Point{
				x: 0.0,
				y: 0.0,
				z: 0.0,
			},
		},
		{
			name: "All coords greater than max",
			args: args{
				point: &Point{
					x: 2.0,
					y: 3.0,
					z: 4.0,
				},
				min: 0.0,
				max: 1.0,
			},
			want: &Point{
				x: 1.0,
				y: 1.0,
				z: 1.0,
			},
		},
		{
			name: "only x within bounds",
			args: args{
				point: &Point{
					x: 0.5,
					y: 3.0,
					z: 4.0,
				},
				min: 0.0,
				max: 1.0,
			},
			want: &Point{
				x: 0.5,
				y: 1.0,
				z: 1.0,
			},
		},
		{
			name: "only y within bounds",
			args: args{
				point: &Point{
					x: 2.5,
					y: 0.5,
					z: 4.0,
				},
				min: 0.0,
				max: 1.0,
			},
			want: &Point{
				x: 1.0,
				y: 0.5,
				z: 1.0,
			},
		},
		{
			name: "only z within bounds",
			args: args{
				point: &Point{
					x: 2.5,
					y: 3.5,
					z: 0.5,
				},
				min: 0.0,
				max: 1.0,
			},
			want: &Point{
				x: 1.0,
				y: 1.0,
				z: 0.5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clip(tt.args.point, tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clip() = %v, want %v", got, tt.want)
			}
		})
	}
}
