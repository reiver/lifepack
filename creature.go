package main



import "math"
import "math/rand"



type Creature struct {

	x,y   float64           // Where the creature is.
	dx,dy float64           // The velocity of the creature.
	θ     float64           // The angle the creature is facing. (In radians.)

	e     float64           // The energy.

	r     float64           // The radius of the creature. (The creature is basically a circle.)
	red,green,blue float64  // The color of the creature.

	name  string
}



func (me *Creature) X() float64 {
	return me.x
}
func (me *Creature) Y() float64 {
	return me.y
}



func (me *Creature) Dx() float64 {
	return me.dx
}
func (me *Creature) Dy() float64 {
	return me.dy
}

// Speed2 return the speed squared of the creature.
func (me *Creature) Speed2() float64 {
	speed2 := me.dx * me.dx  +  me.dy * me.dy

	return speed2
}



func (me *Creature) E() float64 {
	return me.e
}



func (me *Creature) R() float64 {
	return me.r
}



func (me *Creature) Red() float64 {
	return me.red
}
func (me *Creature) Green() float64 {
	return me.green
}
func (me *Creature) Blue() float64 {
	return me.blue
}



func (me *Creature) Mass() float64 {
	mass := me.r * me.r * math.Pi

	// For the density.
	mass *= 3

	return mass
}





func NewCreature(name string) (*Creature) {

	me := Creature{
		x: 0,
		y: 0,

		dx: 0,
		dy: 0,

		e: 300,

		r: 10,

		red:   0.8,
		green: 0.2,
		blue:  0.3,

		name: name,
	}

	return &me
}





func (me *Creature) RandomlyPlace(width, height float64) (*Creature) {

	me.x = rand.Float64()*width
	me.y = rand.Float64()*height

	me.dx = rand.Float64()*2 - 1
	me.dy = rand.Float64()*2 - 1

	me.θ  = rand.Float64()*2*math.Pi
	if 2*math.Pi == me.θ {
		me.θ = 0
	}

	me.red   = 0.5 + rand.Float64()*0.5
	me.green = 0.5 + rand.Float64()*0.5
	me.blue  = 0.5 + rand.Float64()*0.5

	me.r = 10 + rand.Float64()*30


	return me
}





func (me *Creature) Next(width, height float64) {


	// Update position.
		me.x += me.dx
		me.y += me.dy

//@TODO: If the creature is moving fast enough, this won't work!
		if (0 > me.x) {
			me.x += width
		} else if (width < me.x) {
			me.x -= width
		}

		if (0 > me.y) {
			me.y += height
		} else if (height < me.y) {
			me.y -= height
		}



	// Decrease energy.
		me.e -= 1
}
