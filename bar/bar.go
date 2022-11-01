package bar

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Bar struct {
	app.Compo
}

func (b *Bar) Render() app.UI {
	return app.Text("Bar!\n")
}
