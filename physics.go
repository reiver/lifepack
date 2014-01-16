package main



func collision(a,b *Creature) bool {

	// Calculate the square of the change in x.
		Δx  := b.x - a.x
		Δx2 := Δx*Δx

	// Calculate the square of the change in y.
		Δy  := b.y - a.y
		Δy2 := Δy*Δy

	// Calculate the distance squared.
		d2 := Δx2 + Δy2

	// Calculate the threshold.
	// The threshold is the largest distance squared can be to be a collision.
		threshold2 := b.r + a.r
		threshold2 *= threshold2


	// See if there was a collision.
		result := threshold2 >= d2


	// Return.
		return result
}
