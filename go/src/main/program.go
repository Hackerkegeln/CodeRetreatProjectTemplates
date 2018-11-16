package main

func main() {
	g := new(greater)
	g.name = "Julian"
	print(g.hello())
}

type greater struct {
	name string
}

func (me *greater) hello() string {
	return "Hello " + me.name;
}

func hello(name string) string {
	return "Hello " + name
}