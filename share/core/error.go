package core

func (app *ShareApp) GetErr() error {
	return app.Err
}
func (app *ShareApp) SetErr(err error) {
	app.Err = err
}
