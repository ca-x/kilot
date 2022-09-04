package formater

import "testing"

func TestString_ToCamel(t *testing.T) {
	camel := From("User").ToCamel()
	t.Log(camel)
}
