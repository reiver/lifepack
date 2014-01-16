package main



import "fmt"
import "github.com/banthar/gl"
import "math"



// GlDraw draws the creature to the OpenGL screen.
func (me *Creature) GlDraw() {


	// Draw the creature's body, which is just a circle.
		gl.Color3d(me.Red(), me.Green(), me.Blue())
		drawCircle(me.X(), me.Y(), me.R())

	// Draw circles for the creature's eyes.
	// (This is so we can tell what direction the creature is facing.)
	//
	// We'll make the eye size 7% of the total radius.
	// We'll also put the eyes
		rr := float64(0.07) * me.R()

//@TODO: Should speed this up by pre-computing all these cos() and sin() calculations and storing in a lookup table.
		xx := me.X() + ((float64(0.8) * me.R()) * math.Cos(me.θ() - 0.4))
		yy := me.Y() + ((float64(0.8) * me.R()) * math.Sin(me.θ() - 0.4))

		gl.Color3d(0, 0, 0)
		drawCircle(xx, yy, rr)

//@TODO: Should speed this up by pre-computing all these cos() and sin() calculations and storing in a lookup table.
		xx = me.X() + ((float64(0.8) * me.R()) * math.Cos(me.θ() + 0.4))
		yy = me.Y() + ((float64(0.8) * me.R()) * math.Sin(me.θ() + 0.4))

		gl.Color3d(0, 0, 0)
		drawCircle(xx, yy, rr)

	// Draw creature's thrusters.
		rr = float64(0.25) * me.R()

//@TODO: Should speed this up by pre-computing all these cos() and sin() calculations and storing in a lookup table.
		xx = me.X() + ((float64(0.6) * me.R()) * math.Cos(me.θ() - 2.4))
		yy = me.Y() + ((float64(0.6) * me.R()) * math.Sin(me.θ() - 2.4))

		gl.Color3d(0, 0, 0)
		drawCircle(xx, yy, rr)

//@TODO: Should speed this up by pre-computing all these cos() and sin() calculations and storing in a lookup table.
		xx = me.X() + ((float64(0.6) * me.R()) * math.Cos(me.θ() + 2.4))
		yy = me.Y() + ((float64(0.6) * me.R()) * math.Sin(me.θ() + 2.4))

		gl.Color3d(0, 0, 0)
		drawCircle(xx, yy, rr)


	// Draw the creature's name.
		xx = me.X() - me.R() - 10
		yy = me.Y() - me.R() - 10

		drawString(xx,yy, me.name)


	// Draw the creature's energy.
		xx = me.X() + me.R()
		yy = me.Y() - me.R() - 10

		s := fmt.Sprintf("energy: %v", me.E())

		drawString(xx,yy, s)


	// Draw the creature's x position.
		xx = me.X() + me.R()
		yy = me.Y() - me.R() + 30

		s = fmt.Sprintf("x: %v", me.X())

		drawString(xx,yy, s)


	// Draw the creature's y position.
		xx = me.X() + me.R()
		yy = me.Y() - me.R() + 45

		s = fmt.Sprintf("y: %v", me.Y())

		drawString(xx,yy, s)


	// Draw the creature's x speed.
		xx = me.X() + me.R()
		yy = me.Y() - me.R() + 60

		s = fmt.Sprintf("dx: %v", me.Dx())

		drawString(xx,yy, s)


	// Draw the creature's y speed.
		xx = me.X() + me.R()
		yy = me.Y() - me.R() + 75

		s = fmt.Sprintf("dy: %v", me.Dy())

		drawString(xx,yy, s)


	// Draw the creature's orientation.
		xx = me.X() + me.R()
		yy = me.Y() - me.R() + 90

		s = fmt.Sprintf("θ: %v degrees", 360*me.θ()/(2*math.Pi))

		drawString(xx,yy, s)
}

