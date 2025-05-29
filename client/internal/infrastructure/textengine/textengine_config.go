package textengine

import (
	"fmt"
	"os"
)

type TextEngine struct {
	path string
	file map[string]*os.File
}

func NewTextEngine(path string) *TextEngine {
	return &TextEngine{
		path: path,
		file: make(map[string]*os.File),
	}
}

func (textEngine *TextEngine) AddTextFile(path, fileName, extension string) error {

	file, err := os.Create(fmt.Sprintf("%s/%s.%s", path, fileName, extension))

	if err != nil {
		return err
	}

	textEngine.file[path] = file

	return nil
}

func (textEngine *TextEngine) RemovePathFile(path string) {
	delete(textEngine.file, path)
}

func (textEngine *TextEngine) GetPathFile(path string) *os.File {
	return textEngine.file[path]
}
