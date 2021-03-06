// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package material

import (
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/math32"
)

// Point material is normally used for single point sprites
type Point struct {
	Standard // Embedded standard material
}

// NewPoint creates and returns a pointer to a new point material
func NewPoint(color *math32.Color) *Point {

	pm := new(Point)
	pm.Standard.Init("shaderPoint", color)

	// Sets uniform's initial values
	pm.uni.SetColor(vEmissive, color)
	pm.uni.SetPos(pSize, 1.0)
	pm.uni.SetPos(pRotationZ, 0)
	return pm
}

// SetEmissiveColor sets the material emissive color
// The default is {0,0,0}
func (pm *Point) SetEmissiveColor(color *math32.Color) {

	pm.uni.SetColor(vEmissive, color)
}

// SetSize sets the point size
func (pm *Point) SetSize(size float32) {

	pm.uni.SetPos(pSize, size)
}

// SetRotationZ sets the point rotation around the Z axis.
func (pm *Point) SetRotationZ(rot float32) {

	pm.uni.SetPos(pRotationZ, rot)
}

// RenderSetup is called by the engine before drawing the object
// which uses this material
func (pm *Point) RenderSetup(gs *gls.GLS) {

	pm.Material.RenderSetup(gs)
	pm.uni.Transfer(gs)
}
