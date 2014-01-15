package main



import "github.com/banthar/gl"



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
