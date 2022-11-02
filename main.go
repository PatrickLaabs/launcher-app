package main

import (
	"fmt"
	"log"
	"net/http"

	B "github.com/PatrickLaabs/launcher-app/bar"
	F "github.com/PatrickLaabs/launcher-app/footer"
	L "github.com/PatrickLaabs/launcher-app/leftside"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

type foo struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.H1().Text("Hello World!")
}

func (f *foo) Render() app.UI {
	return app.Div().
		Style("background-color", "deepskyblue").
		Body(
			app.H1().
				Class("title").
				Text("Build a GUI with Go"),
			app.P().
				Class("text").
				Text("Bla bla bla"),
			&B.Bar{},
			&myCompo{},
			&svg{},
			&homeButton{},
			&F.Footer{},
		)
}

type testPage struct {
	app.Compo
}

func (t *testPage) Render() app.UI {
	return app.Main().
		Body(
			app.H1().
				Class("title").
				Text("Test Page"),
			app.Section().
				Text("bla").
				Style("background-color", "yellow").
				Style("position", "absolute").
				Style("height", "70px").
				Style("width", "25px").
				Style("left", "0"),
			&L.Leftside{},
			&F.Footer{},
		)
}

type myCompo struct {
	app.Compo
}

func (c *myCompo) Render() app.UI {
	return app.Img().
		Alt("Gopher Image").
		Src("/web/_assets/gopher.png").
		Height(256).
		Width(256)
}

type svg struct {
	app.Compo
}

func (s *svg) Render() app.UI {
	return app.Raw(`
		<svg width="100" height="100">
		<circle cx="50" cy="50" r="40" stroke="green" stroke-width="4" fill="yellow" />
	</svg>
	`)
}

type homeButton struct {
	app.Compo
}

func (b *homeButton) Render() app.UI {
	return app.Button().
		Text("Test-Button").
		Style("width", "50px").
		Style("height", "50px").
		OnClick(b.OnClick)
}

func (b *homeButton) OnClick(ctx app.Context, e app.Event) {
	fmt.Println("Button pressed")

}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &hello{})
	app.Route("/foo", &foo{})
	app.Route("/gopher", &myCompo{})
	app.Route("/test", &testPage{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	err := app.GenerateStaticWebsite("./docs", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Resources:   app.GitHubPages("launcher-app"),
	})
	if err != nil {
		log.Fatal(err)
	}

	if err = http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
