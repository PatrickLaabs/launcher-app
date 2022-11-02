package footer

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Footer struct {
	app.Compo
}

func (f *Footer) Render() app.UI {
	return app.Footer().
		Class("footer").
		Style("background-color", "deepskyblue").
		Style("height", "50px").
		Style("position", "absolute").
		Style("bottom", "0").
		Text("Footer! bottom alignment")
}
