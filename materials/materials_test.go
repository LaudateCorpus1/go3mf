package materials

import (
	"reflect"
	"testing"
)

func TestTexture2DResource_Identify(t *testing.T) {
	tests := []struct {
		name  string
		t     *Texture2DResource
		want  string
		want1 uint32
	}{
		{"base", &Texture2DResource{ID: 1, ModelPath: "3d/3dmodel.model"}, "3d/3dmodel.model", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.t.Identify()
			if got != tt.want {
				t.Errorf("Texture2DResource.Identify() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Texture2DResource.Identify() got = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTextureCoord_U(t *testing.T) {
	tests := []struct {
		name string
		t    TextureCoord
		want float32
	}{
		{"base", TextureCoord{1, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.U(); got != tt.want {
				t.Errorf("TextureCoord.U() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextureCoord_V(t *testing.T) {
	tests := []struct {
		name string
		t    TextureCoord
		want float32
	}{
		{"base", TextureCoord{1, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.V(); got != tt.want {
				t.Errorf("TextureCoord.V() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTexture2DGroupResource_Identify(t *testing.T) {
	tests := []struct {
		name  string
		t     *Texture2DGroupResource
		want  string
		want1 uint32
	}{
		{"base", &Texture2DGroupResource{ID: 1, ModelPath: "3d/3dmodel"}, "3d/3dmodel", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.t.Identify()
			if got != tt.want {
				t.Errorf("Texture2DGroupResource.Identify() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Texture2DGroupResource.Identify() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestColorGroupResource_Identify(t *testing.T) {
	tests := []struct {
		name  string
		c     *ColorGroupResource
		want  string
		want1 uint32
	}{
		{"base", &ColorGroupResource{ID: 1, ModelPath: "3d/3dmodel"}, "3d/3dmodel", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.Identify()
			if got != tt.want {
				t.Errorf("ColorGroupResource.Identify() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ColorGroupResource.Identify() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCompositeMaterialsResource_Identify(t *testing.T) {
	tests := []struct {
		name  string
		c     *CompositeMaterialsResource
		want  string
		want1 uint32
	}{
		{"base", &CompositeMaterialsResource{ID: 1, ModelPath: "3d/3dmodel.model"}, "3d/3dmodel.model", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.Identify()
			if got != tt.want {
				t.Errorf("CompositeMaterialsResource.Identify() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CompositeMaterialsResource.Identify() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTexture2DType_String(t *testing.T) {
	tests := []struct {
		name string
		t    Texture2DType
	}{
		{"image/png", TextureTypePNG},
		{"image/jpeg", TextureTypeJPEG},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.name {
				t.Errorf("Texture2DType.String() = %v, want %v", got, tt.name)
			}
		})
	}
}

func TestBlendMethod_String(t *testing.T) {
	tests := []struct {
		name string
		b    BlendMethod
	}{
		{"mix", BlendMix},
		{"multiply", BlendMultiply},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.name {
				t.Errorf("BlendMethod.String() = %v, want %v", got, tt.name)
			}
		})
	}
}

func TestTileStyle_String(t *testing.T) {
	tests := []struct {
		name string
		t    TileStyle
	}{
		{"wrap", TileWrap},
		{"mirror", TileMirror},
		{"clamp", TileClamp},
		{"none", TileNone},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.name {
				t.Errorf("TileStyle.String() = %v, want %v", got, tt.name)
			}
		})
	}
}

func TestTextureFilter_String(t *testing.T) {
	tests := []struct {
		name string
		t    TextureFilter
	}{
		{"auto", TextureFilterAuto},
		{"linear", TextureFilterLinear},
		{"nearest", TextureFilterNearest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.name {
				t.Errorf("TextureFilter.String() = %v, want %v", got, tt.name)
			}
		})
	}
}

func TestMultiPropertiesResource_Identify(t *testing.T) {
	tests := []struct {
		name  string
		c     *MultiPropertiesResource
		want  string
		want1 uint32
	}{
		{"base", &MultiPropertiesResource{ID: 1, ModelPath: "3d/3dmodel.model"}, "3d/3dmodel.model", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.Identify()
			if got != tt.want {
				t.Errorf("MultiPropertiesResource.Identify() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MultiPropertiesResource.Identify() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_newBlendMethod(t *testing.T) {
	tests := []struct {
		name   string
		wantB  BlendMethod
		wantOk bool
	}{
		{"mix", BlendMix, true},
		{"multiply", BlendMultiply, true},
		{"empty", BlendMix, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, gotOk := newBlendMethod(tt.name)
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("newBlendMethod() gotB = %v, want %v", gotB, tt.wantB)
			}
			if gotOk != tt.wantOk {
				t.Errorf("newBlendMethod() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_newTextureFilter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		want   TextureFilter
		wantOk bool
	}{
		{"auto", TextureFilterAuto, true},
		{"linear", TextureFilterLinear, true},
		{"nearest", TextureFilterNearest, true},
		{"empty", TextureFilterAuto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := newTextureFilter(tt.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTextureFilter() got = %v, want %v", got, tt.want)
			}
			if got != tt.want {
				t.Errorf("newTextureFilter() got1 = %v, want %v", got1, tt.want)
			}
		})
	}
}

func Test_newTileStyle(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		want  TileStyle
		want1 bool
	}{
		{"wrap", TileWrap, true},
		{"mirror", TileMirror, true},
		{"clamp", TileClamp, true},
		{"none", TileNone, true},
		{"empty", TileWrap, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := newTileStyle(tt.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTileStyle() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("newTileStyle() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_newTexture2DType(t *testing.T) {
	tests := []struct {
		name  string
		want  Texture2DType
		want1 bool
	}{
		{"image/png", TextureTypePNG, true},
		{"image/jpeg", TextureTypeJPEG, true},
		{"", Texture2DType(0), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := newTexture2DType(tt.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTexture2DType() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("newTexture2DType() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
