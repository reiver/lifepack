package main



import "github.com/banthar/gl"
import "math"



func drawCircle(x, y, r  float64) {

//@TODO: Should speed this up by pre-computing all these cos() and sin() calculations and storing in a lookup table.
	gl.Begin(gl.TRIANGLE_FAN)
//		gl.Color3d(me.red, me.green, me.blue)

		gl.Vertex2d(x, y)

		pi2 := 2*math.Pi

		num := 36
		dTheta := pi2 / float64(num)
		theta := float64(0)
		for i:=0 ; i<num ; i++ {

			theta += dTheta

			dx := math.Cos(theta) * r
			dy := math.Sin(theta) * r

			gl.Vertex2d(x + dx, y + dy)
		}

		theta += dTheta

		dx := math.Cos(theta) * r
		dy := math.Sin(theta) * r

		gl.Vertex2d(x + dx, y + dy)

	gl.End()
}

func drawSquare(x, y, length float64) {

	x2 := x+length
	y2 := y+length

	gl.Begin(gl.QUADS)
		gl.Color3d(0.75, 0.35, 0.35)
		gl.Vertex2d(x  , y)
		gl.Vertex2d(x2 , y)
		gl.Vertex2d(x2 , y2)
		gl.Vertex2d(x  , y2)
	gl.End()
}


func draw6x8(x,y, length float64, mapp *[48]rune) {

	for k,v := range mapp {

		if ' ' != v {
			r := k % 6

			xx := x + float64(r)           * length
			yy := y + float64((k - r)/ 8) * length

			drawSquare(xx, yy, length)
		}

	}
}
