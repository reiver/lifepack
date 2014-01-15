package main



import "github.com/banthar/gl"
import "math"



// GlDraw draws the creature to the OpenGL screen.
func (me *Creature) GlDraw() {


	// Draw circle.
//@TODO: Should speed this up by pre-computing all these cos() and sin() calculations and storing in a lookup table.
		gl.Begin(gl.TRIANGLE_FAN)
			gl.Color3d(me.red, me.green, me.blue)

			gl.Vertex2d(me.x, me.y)

			pi2 := 2*math.Pi

			num := 36
			dTheta := pi2 / float64(num)
			theta := float64(0)
			for i:=0 ; i<num ; i++ {

				theta += dTheta

				dx := math.Cos(theta) * me.r
				dy := math.Sin(theta) * me.r

				gl.Vertex2d(me.x + dx, me.y + dy)
			}

			theta += dTheta

			dx := math.Cos(theta) * me.r
			dy := math.Sin(theta) * me.r

			gl.Vertex2d(me.x + dx, me.y + dy)

		gl.End()
}

