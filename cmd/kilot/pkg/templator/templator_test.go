package templator

import "testing"

func TestTemplator_Execute(t *testing.T) {
	tpl := `func {{.funcName}}(){}`
	val := map[string]interface{}{
		"funcName": "FakeTest",
	}
	templator := New("func", tpl, true)
	execute, err := templator.Execute(val)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(execute.String())
}
