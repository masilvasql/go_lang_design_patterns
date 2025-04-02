package component

type File struct {
	Name string
}

func (f *File) Show(indent string) {
	println(indent + "- " + f.Name)
}
