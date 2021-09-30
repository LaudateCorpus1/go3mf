// © Copyright 2021 HP Development Company, L.P.
// SPDX-License Identifier: BSD-2-Clause

package production

import (
	"encoding/xml"

	"github.com/hpinc/go3mf/spec"
)

// Marshal3MFAttr encodes the resource attributes.
func (u *BuildAttr) Marshal3MFAttr(_ spec.Encoder, start *xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: Namespace, Local: attrProdUUID}, Value: u.UUID})
	return nil
}

// Marshal3MFAttr encodes the resource attributes.
func (u *ObjectAttr) Marshal3MFAttr(_ spec.Encoder, start *xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: Namespace, Local: attrProdUUID}, Value: u.UUID})
	return nil
}

// Marshal3MFAttr encodes the resource attributes.
func (u *ItemAttr) Marshal3MFAttr(_ spec.Encoder, start *xml.StartElement) error {
	if u.Path != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: Namespace, Local: attrPath}, Value: u.Path})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: Namespace, Local: attrProdUUID}, Value: u.UUID})
	return nil
}

// Marshal3MFAttr encodes the resource attributes.
func (u *ComponentAttr) Marshal3MFAttr(_ spec.Encoder, start *xml.StartElement) error {
	if u.Path != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: Namespace, Local: attrPath}, Value: u.Path})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: Namespace, Local: attrProdUUID}, Value: u.UUID})
	return nil
}
