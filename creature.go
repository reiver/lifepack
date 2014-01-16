package main



import "math"
import "math/rand"



const (
	offset_x     = 0
	offset_y     = 1

	offset_dx    = 2
	offset_dy    = 3

	offset_θ     = 4

	offset_e     = 5

	offset_r     = 6

	offset_red   = 7
	offset_green = 8
	offset_blue  = 9
)


type Creature struct {


	// The fields for this struct are stored in an array instead
	// of individual fields so that it is easier to do matrix
	// multiplication with them.
	//
	// Conceptually, this is what is in there:
	//
	//	x,y   float64           // Where the creature is.
	//	dx,dy float64           // The velocity of the creature.
	//	θ     float64           // The angle the creature is facing. (In radians.)
	//
	//	e     float64           // The energy.
	//
	//	r     float64           // The radius of the creature. (The creature is basically a circle.)
	//	red,green,blue float64  // The color of the creature.
	//
	data  [10]float64

	name  string
}



func (me *Creature) X() float64 {
	return me.data[offset_x]
}
func (me *Creature) SetX(x float64) {
	me.data[offset_x] = x
}

func (me *Creature) Y() float64 {
	return me.data[offset_y]
}
func (me *Creature) SetY(y float64) {
	me.data[offset_y] = y
}



func (me *Creature) Dx() float64 {
	return me.data[offset_dx]
}
func (me *Creature) SetDx(dx float64) {
	me.data[offset_dx] = dx
}

func (me *Creature) Dy() float64 {
	return me.data[offset_dy]
}
func (me *Creature) SetDy(dy float64) {
	me.data[offset_dy] = dy
}

// Speed2 returns the speed squared of the creature.
func (me *Creature) Speed2() float64 {
	speed2 := me.Dx() * me.Dx()  +  me.Dy() * me.Dy()

	return speed2
}


func (me *Creature) θ() float64 {
	return me.data[offset_θ]
}
func (me *Creature) Setθ(theta float64) {
	me.data[offset_θ] = theta
}



func (me *Creature) E() float64 {
	return me.data[offset_e]
}



func (me *Creature) R() float64 {
	return me.data[offset_r]
}

func (me *Creature) Red() float64 {
	return me.data[offset_red]
}
func (me *Creature) Green() float64 {
	return me.data[offset_green]
}
func (me *Creature) Blue() float64 {
	return me.data[offset_blue]
}



func (me *Creature) Mass() float64 {
	mass := me.R() * me.R() * math.Pi

	// For the density.
	mass *= 3

	return mass
}





func NewCreature(name string) (*Creature) {

	me := Creature{
		name: name,
	}

	return &me
}





func (me *Creature) Randomize() (*Creature) {

	me.data[offset_red]   = 0.5 + rand.Float64()*0.5
	me.data[offset_green] = 0.5 + rand.Float64()*0.5
	me.data[offset_blue]  = 0.5 + rand.Float64()*0.5

	me.data[offset_r] = 10 + rand.Float64()*30


	return me
}

func (me *Creature) RandomlyPlace(width, height float64) (*Creature) {

	me.data[offset_x] = rand.Float64()*width
	me.data[offset_y] = rand.Float64()*height

	me.data[offset_dx] = rand.Float64()*2 - 1
	me.data[offset_dy] = rand.Float64()*2 - 1

	me.data[offset_θ]  = rand.Float64()*2*math.Pi
	if 2*math.Pi == me.data[offset_θ] {
		me.data[offset_θ] = 0
	}


	return me
}





func (me *Creature) Next(width, height float64) {


	// Update position.
		me.data[offset_x] += me.Dx()
		me.data[offset_y] += me.Dy()

//@TODO: If the creature is moving fast enough, this won't work!
		if (0 > me.data[offset_x]) {
			me.data[offset_x] += width
		} else if (width < me.data[offset_x]) {
			me.data[offset_x] -= width
		}

		if (0 > me.data[offset_y]) {
			me.data[offset_y] += height
		} else if (height < me.data[offset_y]) {
			me.data[offset_y] -= height
		}



	// Decrease energy.
		me.data[offset_e] -= 1
}
