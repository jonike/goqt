package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/visualfc/goqt/ui"
)

type CalclatorForm struct {
	*ui.QWidget
	spinBox1    *ui.QSpinBox
	spinBox2    *ui.QSpinBox
	outputLable *ui.QLabel
}

func IsValidDriver(v ui.Driver) bool {
	return !reflect.ValueOf(v).IsNil()
}

func NewCalclatorForm() (*CalclatorForm, error) {
	w := &CalclatorForm{}
	w.QWidget = ui.NewQWidget()

	file := ui.NewQFileWithName(":/forms/calculatorform.ui")
	if !file.Open(ui.QIODevice_ReadOnly) {
		return nil, errors.New("error load ui")
	}

	loader := ui.NewQUiLoader()
	formWidget := loader.Load(file)
	if formWidget == nil {
		return nil, errors.New("error load form widget")
	}

	w.spinBox1 = ui.NewQSpinBoxFromDriver(formWidget.FindChild("inputSpinBox1"))
	w.spinBox2 = ui.NewQSpinBoxFromDriver(formWidget.FindChild("inputSpinBox2"))
	w.outputLable = ui.NewQLabelFromDriver(formWidget.FindChild("outputWidget"))

	if ui.IsValidDriver(w.spinBox1) && ui.IsValidDriver(w.spinBox2) && ui.IsValidDriver(w.outputLable) {
		fnChanged := func() {
			w.outputLable.SetText(fmt.Sprintf("%d", w.spinBox1.Value()+w.spinBox2.Value()))
		}

		w.spinBox1.OnValueChanged(func(string) {
			fnChanged()
		})
		w.spinBox2.OnValueChanged(func(string) {
			fnChanged()
		})
	}

	layout := ui.NewQVBoxLayout()
	layout.AddWidget(formWidget)
	w.SetLayout(layout)

	w.SetWindowTitle("Calculator Builder")
	return w, nil
}
