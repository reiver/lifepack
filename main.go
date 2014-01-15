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

	// Make the world come alive.

		for i:=0 ; true ; i++ {

			// Draw.
				 gl.Clear(gl.COLOR_BUFFER_BIT)




				glfw.SwapBuffers()





				// Wait for 0.05 seconds.
				time.Sleep(50 * time.Millisecond)
		} // for

}
