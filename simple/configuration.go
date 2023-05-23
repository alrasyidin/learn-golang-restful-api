package simple

type Configuration struct {
	Name string
}

type Application struct {
	*Configuration
}

// Constructor for Application
func NewApplication() *Application {
	return &Application{
		Configuration: &Configuration{
			Name: "Hafidh Pradipta",
		},
	}
}
