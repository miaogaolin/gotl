package yaml2go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYaml2Go_Convert(t *testing.T) {
	yamlContent := `
kind: test
metadata:
  name: cluster
  nullfield:
  nestedstruct:
  - nested:
      underscore_field: value
      field1:
      - 44.5
      - 43.6
      field2:
      - true
      - false
    nested2:
      - nested3:
          field1:
          - 44
          - 43
          fieldt:
          - true
          - false
          field3: value
abc:
  - def:
    - black
    - white
array1:
  - "string1"
  - "string2"
array2:
  - 2
  - 6
array3:
  - 3.14
  - 5.12
is_underscore: true

`
	yaml := New()
	_, err := yaml.Convert([]byte(yamlContent))
	assert.Nil(t, err)
}
