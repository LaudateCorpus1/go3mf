package materials

import (
	"image/color"
	"strings"

	"github.com/qmuntal/go3mf"
	specerr "github.com/qmuntal/go3mf/errors"
)

func (r *ColorGroup) Validate(m *go3mf.Model, _ string) []error {
	var errs []error
	if r.ID == 0 {
		errs = append(errs, specerr.ErrMissingID)
	}
	if len(r.Colors) == 0 {
		errs = append(errs, specerr.ErrEmptyResourceProps)
	}
	for j, c := range r.Colors {
		if c == (color.RGBA{}) {
			errs = append(errs, &specerr.IndexedError{
				Name:  attrColor,
				Index: j,
				Err:   &specerr.MissingFieldError{Name: attrColor},
			})
		}
	}
	return errs
}

func (r *Texture2DGroup) Validate(m *go3mf.Model, path string) []error {
	var errs []error
	if r.ID == 0 {
		errs = append(errs, specerr.ErrMissingID)
	}
	if r.TextureID == 0 {
		errs = append(errs, &specerr.MissingFieldError{Name: attrTexID})
	} else if text, ok := m.FindAsset(path, r.TextureID); ok {
		if _, ok := text.(*Texture2DResource); !ok {
			errs = append(errs, specerr.ErrTextureReference)
		}
	} else {
		errs = append(errs, specerr.ErrTextureReference)
	}
	if len(r.Coords) == 0 {
		errs = append(errs, specerr.ErrEmptyResourceProps)
	}
	return errs
}

func (r *Texture2DResource) Validate(m *go3mf.Model, path string) []error {
	var errs []error
	if r.ID == 0 {
		errs = append(errs, specerr.ErrMissingID)
	}
	if r.Path == "" {
		errs = append(errs, &specerr.MissingFieldError{Name: attrPath})
	} else {
		var hasTexture bool
		for _, a := range m.Attachments {
			if strings.EqualFold(a.Path, r.Path) {
				hasTexture = true
				break
			}
		}
		if !hasTexture {
			errs = append(errs, specerr.ErrMissingTexturePart)
		}
	}
	if r.ContentType == 0 {
		errs = append(errs, &specerr.MissingFieldError{Name: attrContentType})
	}
	return errs
}

func (r *MultiProperties) Validate(m *go3mf.Model, path string) []error {
	var errs []error
	if r.ID == 0 {
		errs = append(errs, specerr.ErrMissingID)
	}
	if len(r.PIDs) == 0 {
		errs = append(errs, &specerr.MissingFieldError{Name: attrPIDs})
	}
	if len(r.BlendMethods) > len(r.PIDs)-1 {
		errs = append(errs, specerr.ErrMultiBlend)
	}
	if len(r.Multis) == 0 {
		errs = append(errs, specerr.ErrEmptyResourceProps)
	}
	var (
		colorCount        int
		resourceUndefined bool
		lengths           = make([]int, len(r.PIDs))
	)
	for j, pid := range r.PIDs {
		if pr, ok := m.FindAsset(path, pid); ok {
			switch pr := pr.(type) {
			case *go3mf.BaseMaterials:
				if j != 0 {
					errs = append(errs, specerr.ErrMaterialMulti)
				}
				lengths[j] = len(pr.Materials)
			case *CompositeMaterials:
				if j != 0 {
					errs = append(errs, specerr.ErrMaterialMulti)
				}
				lengths[j] = len(pr.Composites)
			case *MultiProperties:
				errs = append(errs, specerr.ErrMultiRefMulti)
			case *ColorGroup:
				if colorCount == 1 {
					errs = append(errs, specerr.ErrMultiColors)
				}
				colorCount++
				lengths[j] = len(pr.Colors)
			}
		} else if !resourceUndefined {
			resourceUndefined = true
			errs = append(errs, specerr.ErrMissingResource)
		}
	}
	for j, m := range r.Multis {
		for k, index := range m.PIndex {
			if k < len(r.PIDs) && lengths[k] < int(index) {
				errs = append(errs, &specerr.IndexedError{
					Name:  attrMulti,
					Index: j,
					Err:   specerr.ErrIndexOutOfBounds,
				})
				break
			}
		}
	}
	return errs
}

func (r *CompositeMaterials) Validate(m *go3mf.Model, path string) []error {
	var errs []error
	if r.ID == 0 {
		errs = append(errs, specerr.ErrMissingID)
	}
	if r.MaterialID == 0 {
		errs = append(errs, &specerr.MissingFieldError{Name: attrMatID})
	} else if mat, ok := m.FindAsset(path, r.MaterialID); ok {
		if bm, ok := mat.(*go3mf.BaseMaterials); ok {
			for _, index := range r.Indices {
				if int(index) > len(bm.Materials) {
					errs = append(errs, specerr.ErrIndexOutOfBounds)
					break
				}
			}
		} else {
			errs = append(errs, specerr.ErrCompositeBase)
		}
	} else {
		errs = append(errs, specerr.ErrMissingResource)
	}
	if len(r.Indices) == 0 {
		errs = append(errs, &specerr.MissingFieldError{Name: attrMatIndices})
	}
	if len(r.Composites) == 0 {
		errs = append(errs, specerr.ErrEmptyResourceProps)
	}
	return errs
}
