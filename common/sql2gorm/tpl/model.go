package tpl

const ModelTpl = `
package {{.package}}
import (
	{{ range .imports}}
	"{{ . }}"
	{{ end }}
)

{{ range .struct }}
{{ . }}
{{ end }}
`
