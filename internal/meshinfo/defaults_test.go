package meshinfo

import (
	"reflect"
	"testing"
)

func TestNewNodeColorInfo(t *testing.T) {
	type args struct {
		currentFaceCount uint32
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{"base", args{1}, reflect.TypeOf((*NodeColor)(nil))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNodeColorInfo(tt.args.currentFaceCount).InfoType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNodeColorInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTextureCoordsInfo(t *testing.T) {
	type args struct {
		currentFaceCount uint32
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{"base", args{1}, reflect.TypeOf((*TextureCoords)(nil))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTextureCoordsInfo(tt.args.currentFaceCount).InfoType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTextureCoordsInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBaseMaterialInfo(t *testing.T) {
	type args struct {
		currentFaceCount uint32
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{"base", args{1}, reflect.TypeOf((*BaseMaterial)(nil))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBaseMaterialInfo(tt.args.currentFaceCount).InfoType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseMaterialInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
