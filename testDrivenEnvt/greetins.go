package testdrivenenvt

func greet(name string) string {
	if (name == ""){
		return "Welcome, Stranger"
	}
	return "Welcome, "+ name
}