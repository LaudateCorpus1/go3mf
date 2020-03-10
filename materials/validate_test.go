package materials

import (
	"image/color"
	"testing"

	"github.com/go-test/deep"
	"github.com/qmuntal/go3mf"
	specerr "github.com/qmuntal/go3mf/errors"
)

func TestValidate(t *testing.T) {
	rootPath := go3mf.DefaultModelPath
	type args struct {
		model *go3mf.Model
	}
	tests := []struct {
		name string
		args args
		want []error
	}{
		{"empty", args{new(go3mf.Model)}, []error{}},
		{"child", args{&go3mf.Model{Childs: map[string]*go3mf.ChildModel{
			"/other.model": &go3mf.ChildModel{Resources: go3mf.Resources{Assets: []go3mf.Asset{
				&ColorGroup{ID: 1},
			}}},
			"/that.model": &go3mf.ChildModel{Resources: go3mf.Resources{Assets: []go3mf.Asset{
				&MultiProperties{ID: 2},
			}}},
		}}}, []error{
			&specerr.AssetError{Path: "/other.model", Index: 0, Name: "ColorGroup", Err: specerr.ErrEmptyResourceProps},
			&specerr.AssetError{Path: "/that.model", Index: 0, Name: "MultiProperties", Err: &specerr.MissingFieldError{Name: attrPIDs}},
			&specerr.AssetError{Path: "/that.model", Index: 0, Name: "MultiProperties", Err: specerr.ErrMultiBlend},
			&specerr.AssetError{Path: "/that.model", Index: 0, Name: "MultiProperties", Err: specerr.ErrEmptyResourceProps},
		}},
		{"multi", args{&go3mf.Model{
			Resources: go3mf.Resources{Assets: []go3mf.Asset{
				&MultiProperties{ID: 4},
				&MultiProperties{ID: 5, Multis: []Multi{{PIndex: []uint32{}}}, PIDs: []uint32{4, 100}},
				&go3mf.BaseMaterials{ID: 1, Materials: []go3mf.Base{
					{Name: "a", Color: color.RGBA{R: 1}},
					{Name: "b", Color: color.RGBA{G: 1}},
				}},
				&ColorGroup{ID: 6, Colors: []color.RGBA{{R: 1}, {R: 2, G: 3, B: 4, A: 5}}},
				&CompositeMaterials{ID: 3, MaterialID: 1, Indices: []uint32{0, 1}, Composites: []Composite{{Values: []float32{1, 2}}}},
				&MultiProperties{ID: 2, Multis: []Multi{{PIndex: []uint32{1, 0}}}, PIDs: []uint32{1, 6}},
				&MultiProperties{ID: 7, Multis: []Multi{{PIndex: []uint32{1, 3}}}, PIDs: []uint32{1, 6}},
				&MultiProperties{ID: 8, Multis: []Multi{{PIndex: []uint32{}}}, PIDs: []uint32{6, 1, 6}},
				&MultiProperties{ID: 9, Multis: []Multi{{PIndex: []uint32{}}}, PIDs: []uint32{1, 3}},
			}},
		}}, []error{
			&specerr.AssetError{Path: rootPath, Index: 0, Name: "MultiProperties", Err: &specerr.MissingFieldError{Name: attrPIDs}},
			&specerr.AssetError{Path: rootPath, Index: 0, Name: "MultiProperties", Err: specerr.ErrMultiBlend},
			&specerr.AssetError{Path: rootPath, Index: 0, Name: "MultiProperties", Err: specerr.ErrEmptyResourceProps},
			&specerr.AssetError{Path: rootPath, Index: 1, Name: "MultiProperties", Err: specerr.ErrMultiRefMulti},
			&specerr.AssetError{Path: rootPath, Index: 1, Name: "MultiProperties", Err: specerr.ErrMissingResource},
			&specerr.AssetError{Path: rootPath, Index: 6, Name: "MultiProperties", Err: &specerr.IndexedError{
				Name:  attrMulti,
				Index: 0,
				Err:   specerr.ErrIndexOutOfBounds,
			}},
			&specerr.AssetError{Path: rootPath, Index: 7, Name: "MultiProperties", Err: specerr.ErrMaterialMulti},
			&specerr.AssetError{Path: rootPath, Index: 7, Name: "MultiProperties", Err: specerr.ErrMultiColors},
			&specerr.AssetError{Path: rootPath, Index: 8, Name: "MultiProperties", Err: specerr.ErrMaterialMulti},
		}},
		{"missingTextPart", args{&go3mf.Model{
			Resources: go3mf.Resources{Assets: []go3mf.Asset{
				&Texture2DResource{ID: 1},
				&Texture2DResource{ID: 2, ContentType: TextureTypePNG, Path: "/a.png"},
			}}},
		}, []error{
			&specerr.AssetError{Path: rootPath, Index: 0, Name: "Texture2DResource", Err: &specerr.MissingFieldError{Name: attrPath}},
			&specerr.AssetError{Path: rootPath, Index: 0, Name: "Texture2DResource", Err: &specerr.MissingFieldError{Name: attrContentType}},
			&specerr.AssetError{Path: rootPath, Index: 1, Name: "Texture2DResource", Err: specerr.ErrMissingTexturePart},
		}},
		{"textureGroup", args{&go3mf.Model{
			Attachments: []go3mf.Attachment{{Path: "/a.png"}},
			Resources: go3mf.Resources{Assets: []go3mf.Asset{
				&Texture2DResource{ID: 1, ContentType: TextureTypePNG, Path: "/a.png"},
				&Texture2DGroup{ID: 2},
				&Texture2DGroup{ID: 3, TextureID: 1, Coords: []TextureCoord{{}}},
				&Texture2DGroup{ID: 4, TextureID: 2, Coords: []TextureCoord{{}}},
				&Texture2DGroup{ID: 5, TextureID: 100, Coords: []TextureCoord{{}}},
			}}},
		}, []error{
			&specerr.AssetError{Path: rootPath, Index: 1, Name: "Texture2DGroup", Err: &specerr.MissingFieldError{Name: attrTexID}},
			&specerr.AssetError{Path: rootPath, Index: 1, Name: "Texture2DGroup", Err: specerr.ErrEmptyResourceProps},
			&specerr.AssetError{Path: rootPath, Index: 3, Name: "Texture2DGroup", Err: specerr.ErrTextureReference},
			&specerr.AssetError{Path: rootPath, Index: 4, Name: "Texture2DGroup", Err: specerr.ErrTextureReference},
		}},
		{"colorGroup", args{&go3mf.Model{
			Resources: go3mf.Resources{Assets: []go3mf.Asset{
				&ColorGroup{ID: 1},
				&ColorGroup{ID: 2, Colors: []color.RGBA{{R: 1}, {R: 2, G: 3, B: 4, A: 5}}},
				&ColorGroup{ID: 3, Colors: []color.RGBA{{R: 1}, {}}},
			}}}}, []error{
			&specerr.AssetError{Path: rootPath, Index: 0, Name: "ColorGroup", Err: specerr.ErrEmptyResourceProps},
			&specerr.AssetError{Path: rootPath, Index: 2, Name: "ColorGroup", Err: &specerr.IndexedError{
				Name:  attrColor,
				Index: 1,
				Err:   &specerr.MissingFieldError{Name: attrColor},
			}},
		}},
		{"composite", args{&go3mf.Model{
			Resources: go3mf.Resources{Assets: []go3mf.Asset{
				&go3mf.BaseMaterials{ID: 1, Materials: []go3mf.Base{
					{Name: "a", Color: color.RGBA{R: 1}},
					{Name: "b", Color: color.RGBA{G: 1}},
				}},
				&CompositeMaterials{ID: 2},
				&CompositeMaterials{ID: 3, MaterialID: 1, Indices: []uint32{0, 1}, Composites: []Composite{{Values: []float32{1, 2}}}},
				&CompositeMaterials{ID: 4, MaterialID: 1, Indices: []uint32{100, 100}, Composites: []Composite{{Values: []float32{1, 2}}}},
				&CompositeMaterials{ID: 5, MaterialID: 2, Indices: []uint32{0, 1}, Composites: []Composite{{Values: []float32{1, 2}}}},
				&CompositeMaterials{ID: 6, MaterialID: 100, Indices: []uint32{0, 1}, Composites: []Composite{{Values: []float32{1, 2}}}},
			}}}}, []error{
			&specerr.AssetError{Path: rootPath, Index: 1, Name: "CompositeMaterials", Err: &specerr.MissingFieldError{Name: attrMatID}},
			&specerr.AssetError{Path: rootPath, Index: 1, Name: "CompositeMaterials", Err: &specerr.MissingFieldError{Name: attrMatIndices}},
			&specerr.AssetError{Path: rootPath, Index: 1, Name: "CompositeMaterials", Err: specerr.ErrEmptyResourceProps},
			&specerr.AssetError{Path: rootPath, Index: 3, Name: "CompositeMaterials", Err: specerr.ErrIndexOutOfBounds},
			&specerr.AssetError{Path: rootPath, Index: 4, Name: "CompositeMaterials", Err: specerr.ErrCompositeBase},
			&specerr.AssetError{Path: rootPath, Index: 5, Name: "CompositeMaterials", Err: specerr.ErrMissingResource},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.model.Validate()
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("Validate() = %v", diff)
			}
		})
	}
}
