package services

import (
	"fmt"
	"github.com/DiegoJCordeiro/golang-study/activity/client/internal/infrastructure/textengine"
)

type TextEngineService interface {
	WriteText(path, text string) error
}

type TextEngineServiceImpl struct {
	textEngine *textengine.TextEngine
}

func NewTextEngineServiceImpl(textEngine *textengine.TextEngine) TextEngineService {
	return &TextEngineServiceImpl{textEngine: textEngine}
}

func (service *TextEngineServiceImpl) WriteText(path, text string) error {

	file := service.textEngine.GetPathFile(path)

	defer file.Close()
	defer service.textEngine.RemovePathFile(path)

	_, err := file.WriteString(fmt.Sprintf("Dólar: %s\n", text))

	if err != nil {
		return err
	}

	defer service.textEngine.AddTextFile(path, file.Name(), ".txt")

	return nil
}
