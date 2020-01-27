package beamlattice

import "github.com/qmuntal/go3mf"

const ExtensionName = "http://schemas.microsoft.com/3dmanufacturing/beamlattice/2017/02"

// ClipMode defines the clipping modes for the beam lattices.
type ClipMode uint8

const (
	// ClipNone defines a beam lattice without clipping.
	ClipNone ClipMode = iota
	// ClipInside defines a beam lattice with clipping inside.
	ClipInside
	// ClipOutside defines a beam lattice with clipping outside.
	ClipOutside
)

func newClipMode(s string) (c ClipMode, ok bool) {
	c, ok = map[string]ClipMode{
		"none":    ClipNone,
		"inside":  ClipInside,
		"outside": ClipOutside,
	}[s]
	return
}

func (c ClipMode) String() string {
	return map[ClipMode]string{
		ClipNone:    "none",
		ClipInside:  "inside",
		ClipOutside: "outside",
	}[c]
}

// A CapMode is an enumerable for the different capping modes.
type CapMode uint8

const (
	// CapModeSphere when the capping is an sphere.
	CapModeSphere CapMode = iota
	// CapModeHemisphere when the capping is an hemisphere.
	CapModeHemisphere
	// CapModeButt when the capping is an butt.
	CapModeButt
)

func newCapMode(s string) (t CapMode, ok bool) {
	t, ok = map[string]CapMode{
		"sphere":     CapModeSphere,
		"hemisphere": CapModeHemisphere,
		"butt":       CapModeButt,
	}[s]
	return
}

func (b CapMode) String() string {
	return map[CapMode]string{
		CapModeSphere:     "sphere",
		CapModeHemisphere: "hemisphere",
		CapModeButt:       "butt",
	}[b]
}

// BeamLattice defines the Model Mesh BeamLattice Attributes class and is part of the BeamLattice extension to 3MF.
type BeamLattice struct {
	ClipMode                 ClipMode
	ClippingMeshID           uint32
	RepresentationMeshID     uint32
	Beams                    []Beam
	BeamSets                 []BeamSet
	MinLength, DefaultRadius float64
	CapMode                  CapMode
}

func (m *BeamLattice) checkSanity(nodeCount uint32) bool {
	for _, beam := range m.Beams {
		i0, i1 := beam.NodeIndices[0], beam.NodeIndices[1]
		if i0 == i1 {
			return false
		}
		if i0 >= nodeCount || i1 >= nodeCount {
			return false
		}
	}
	return true
}

// BeamSet defines a set of beams.
type BeamSet struct {
	Refs       []uint32
	Name       string
	Identifier string
}

// Beam defines a single beam.
type Beam struct {
	NodeIndices [2]uint32  // Indices of the two nodes that defines the beam.
	Radius      [2]float64 // Radius of both ends of the beam.
	CapMode     [2]CapMode // Capping mode.
}

// ExtensionBeamLattice extracts the BeamLattice attributes from an MeshResource.
// If it does not exist a new one is added.
func ExtensionBeamLattice(o *go3mf.MeshResource) *BeamLattice {
	if attr, ok := o.Extensions[ExtensionName]; ok {
		return attr.(*BeamLattice)
	}
	if o.Extensions == nil {
		o.Extensions = make(map[string]interface{})
	}
	attr := &BeamLattice{}
	o.Extensions[ExtensionName] = attr
	return attr
}

const (
	attrBeamLattice        = "beamlattice"
	attrRadius             = "radius"
	attrMinLength          = "minlength"
	attrPrecision          = "precision"
	attrClippingMode       = "clippingmode"
	attrClipping           = "clipping"
	attrClippingMesh       = "clippingmesh"
	attrRepresentationMesh = "representationmesh"
	attrCap                = "cap"
	attrBeams              = "beams"
	attrBeam               = "beam"
	attrBeamSets           = "beamsets"
	attrBeamSet            = "beamset"
	attrR1                 = "r1"
	attrR2                 = "r2"
	attrCap1               = "cap1"
	attrCap2               = "cap2"
	attrV1                 = "v1"
	attrV2                 = "v2"
	attrV3                 = "v3"
	attrName               = "name"
	attrIdentifier         = "identifier"
	attrRef                = "ref"
	attrIndex              = "index"
)
