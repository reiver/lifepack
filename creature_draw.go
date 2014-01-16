package main



import "fmt"
import "github.com/banthar/gl"
import "math"



// GlDraw draws the creature to the OpenGL screen.
func (me *Creature) GlDraw() {


	// Draw the creature's body, which is just a circle.
		gl.Color3d(me.red, me.green, me.blue)
		drawCircle(me.x, me.y, me.r)

	// Draw circles for the creature's eyes.
	// (This is so we can tell what direction the creature is facing.)
	//
	// We'll make the eye size 7% of the total radius.
	// We'll also put the eyes
		rr := float64(0.07) * me.r

//@TODO: Should speed this up by pre-computing all these cos() and sin() calculations and storing in a lookup table.
		xx := me.x + ((float64(0.8) * me.r) * math.Cos(me.θ - 0.4))
		yy := me.y + ((float64(0.8) * me.r) * math.Sin(me.θ - 0.4))

		gl.Color3d(0, 0, 0)
		drawCircle(xx, yy, rr)

//@TODO: Should speed this up by pre-computing all these cos() and sin() calculations and storing in a lookup table.
		xx = me.x + ((float64(0.8) * me.r) * math.Cos(me.θ + 0.4))
		yy = me.y + ((float64(0.8) * me.r) * math.Sin(me.θ + 0.4))

		gl.Color3d(0, 0, 0)
		drawCircle(xx, yy, rr)

	// Draw the creature's name.
		xx = me.x - me.r - 10
		yy = me.y - me.r - 10

		drawString(xx,yy, me.name)


	// Draw the creature's energy.
		xx = me.x + me.r
		yy = me.y - me.r - 10

		s := fmt.Sprintf("energy: %v", me.e)

		drawString(xx,yy, s)


	// Draw the creature's x position.
		xx = me.x + me.r
		yy = me.y - me.r + 30

		s = fmt.Sprintf("x: %v", me.x)

		drawString(xx,yy, s)


	// Draw the creature's y position.
		xx = me.x + me.r
		yy = me.y - me.r + 45

		s = fmt.Sprintf("y: %v", me.y)

		drawString(xx,yy, s)


	// Draw the creature's x speed.
		xx = me.x + me.r
		yy = me.y - me.r + 60

		s = fmt.Sprintf("dx: %v", me.dx)

		drawString(xx,yy, s)


	// Draw the creature's y speed.
		xx = me.x + me.r
		yy = me.y - me.r + 75

		s = fmt.Sprintf("dy: %v", me.dy)

		drawString(xx,yy, s)


	// Draw the creature's orientation.
		xx = me.x + me.r
		yy = me.y - me.r + 90

		s = fmt.Sprintf("θ: %v degrees", me.θ/(2*math.Pi))

		drawString(xx,yy, s)
}

