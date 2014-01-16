package main



import "fmt"
import "github.com/banthar/gl"
import "github.com/jteeuwen/glfw"
import "math/rand"
import "os"
import "time"



func main() {

	// Set width and height.
		WIDTH  := 960
		HEIGHT := 720



	// Open up a Window (that would potentially be used for OpenGL) with glfw
		if err := glfw.Init(); err != nil {
			// Error.
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		defer glfw.Terminate()

		if err := glfw.OpenWindow(WIDTH, HEIGHT, 0, 0, 0, 0, 0, 0, glfw.Windowed); err != nil {
			// Error.
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		glfw.SetWindowTitle("LifePack")



	// Initialize.
		// Initialize (pseudo) random number generator (used in other files).
		rand.Seed( time.Now().Unix() )

		// We need a way for the user to exit the program.
		// This will be accomplished by making it so, when the user presses
		// the [ESC] key, that the program exists.
		//
		// The keyHandler function makes this happen, and will exit this
		// program if [ESC] is pressed.
		glfw.SetKeyCallback(keyHandler)

		// Set up OpenGL in a way that we can do 2D graphics (as opposed to 3D graphics).
		initialize(WIDTH, HEIGHT);



	// Run.
		run(WIDTH, HEIGHT);
}




func keyHandler(key, state int) {

	// Exit this program if [ESC] is pressed.
	if glfw.KeyEsc == key {
		os.Exit(0)
	}

}



func initialize(WIDTH, HEIGHT int) {

	gl.MatrixMode(gl.PROJECTION);
 
	gl.LoadIdentity();
 
	gl.Ortho(0, float64(WIDTH), float64(HEIGHT), 0, 0, 1);
 
	gl.MatrixMode(gl.MODELVIEW);





	 gl.Disable(gl.DEPTH_TEST)
}



func run (WIDTH, HEIGHT int) {

	// Initialize world.
		f64width  := float64(WIDTH)
		f64height := float64(HEIGHT)


		creatures := []*Creature{
			NewCreature( fmt.Sprintf("c%d", 1) ).RandomlyPlace(f64width, f64height),
			NewCreature( fmt.Sprintf("c%d", 2) ).RandomlyPlace(f64width, f64height),
			NewCreature( fmt.Sprintf("c%d", 3) ).RandomlyPlace(f64width, f64height),
			NewCreature( fmt.Sprintf("c%d", 4) ).RandomlyPlace(f64width, f64height),
			NewCreature( fmt.Sprintf("c%d", 5) ).RandomlyPlace(f64width, f64height),
		}




	// Make the world come alive.

		for i:=0 ; true ; i++ {

			// Logic.
				// Update state of each creature that results from the passage of time.
				for _,creature := range creatures {
					creature.Next(f64width, f64height)
				}

				// Deal with any potential collisions between creatures.
				for i:=0; i<len(creatures); i++ {
					for ii:=1+i; ii<len(creatures); ii++ {

						if collision(creatures[i], creatures[ii]) {

//@TODO: Make each of the creatures involved in the collision sense the collision.

							// Calculate the new velocities for each creature (that resulted from the collsion).
								mi  := creatures[i].Mass()  // The mass of the 1st creature
								mii := creatures[ii].Mass() // The mass of the 2nd creature.

								Σm := mi + mii // The total mass of the 2 creatures.
								Δm := mi - mii // The difference between the mass of the 2 creatures.

								newDx1 := (creatures[i].dx  * Δm  + (2 * mii * creatures[ii].dx)) / Σm;
								newDy1 := (creatures[i].dy  * Δm  + (2 * mii * creatures[ii].dy)) / Σm;
								newDx2 := (creatures[ii].dx * -Δm + (2 * mi  * creatures[i].dx))  / Σm;
								newDy2 := (creatures[ii].dy * -Δm + (2 * mi  * creatures[i].dy))  / Σm;


							// Update the new velocity's for each creature (that resulted from the collision).
								creatures[i].dx = newDx1;
								creatures[i].dy = newDy1;
								creatures[ii].dx = newDx2;
								creatures[ii].dy = newDy2;
						}

					}
				}




			// Draw.
				 gl.Clear(gl.COLOR_BUFFER_BIT)




				for _,creature := range creatures {
					creature.GlDraw()
				}




				glfw.SwapBuffers()





				// Wait for 0.05 seconds.
				time.Sleep(50 * time.Millisecond)
		} // for

}
