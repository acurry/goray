package point

import (
	"math"
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

func TestDot(t *testing.T) {
	type args struct {
		a *Point
		b *Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "all positive coords",
			args: args{
				a: &Point{
					1.0, 2.0, 3.0,
				},
				b: &Point{
					2.0, 3.0, 4.0,
				},
			},
			want: 20.0,
		},
		{
			name: "negative coords",
			args: args{
				a: &Point{
					-1.0, -2.0, -3.0,
				},
				b: &Point{
					2.0, 3.0, 4.0,
				},
			},
			want: -20.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dot(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Dot(t *testing.T) {
	type fields struct {
		x float64
		y float64
		z float64
	}
	type args struct {
		b *Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "all positive coords",
			fields: fields{
				1.0, 2.0, 3.0,
			},
			args: args{
				b: &Point{
					2.0, 3.0, 4.0,
				},
			},
			want: 20.0,
		},
		{
			name: "negative coords",
			fields: fields{
				-1.0, -2.0, -3.0,
			},
			args: args{
				b: &Point{
					2.0, 3.0, 4.0,
				},
			},
			want: -20.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Point{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			if got := a.Dot(tt.args.b); got != tt.want {
				t.Errorf("Point.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCross(t *testing.T) {
	type args struct {
		a *Point
		b *Point
	}
	tests := []struct {
		name string
		args args
		want *Point
	}{
		{
			name: "right-angle unit vectors",
			args: args{
				a: &Point{
					0.0, 1.0, 0.0,
				},
				b: &Point{
					-1.0, 0.0, 0.0,
				},
			},
			want: &Point{
				0.0, 0.0, 1.0,
			},
		},
		{
			name: "45-deg vectors",
			args: args{
				a: &Point{
					0.0, math.Sqrt(2), 0.0,
				},
				b: &Point{
					-1.0, 0.0, 0.0,
				},
			},
			want: &Point{
				0.0, 0.0, math.Sqrt(2),
			},
		},
		{
			name: "45-deg vectors flipped",
			args: args{
				a: &Point{
					math.Sqrt(2), 0.0, 0.0,
				},
				b: &Point{
					0.0, -1.0, 0.0,
				},
			},
			want: &Point{
				0.0, 0.0, -math.Sqrt(2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cross(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Cross(t *testing.T) {
	type fields struct {
		x float64
		y float64
		z float64
	}
	type args struct {
		b *Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Point{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			if got := a.Cross(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.Cross() = %v, want %v", got, tt.want)
			}
		})
	}
}
