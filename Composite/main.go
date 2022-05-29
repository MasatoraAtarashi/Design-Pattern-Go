package main

func main() {
	rootDir := NewDirectory("root")
	binDir := NewDirectory("bin")
	tmpDir := NewDirectory("tmp")

	err := rootDir.Add(binDir)
	if err != nil {
		panic(err)
	}

	err = rootDir.Add(tmpDir)
	if err != nil {
		panic(err)
	}

	err = binDir.Add(NewFile("vi", 10000))
	if err != nil {
		panic(err)
	}

	err = binDir.Add(NewFile("latex", 20000))
	if err != nil {
		panic(err)
	}

	rootDir.PrintList("")
}
