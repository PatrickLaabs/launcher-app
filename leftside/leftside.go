package leftside

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Leftside struct {
	app.Compo
}

func (b *Leftside) Render() app.UI {
	return app.Div().
		Style("background-color", "black").
		Style("position", "absolute").
		Style("left", "0")
}
