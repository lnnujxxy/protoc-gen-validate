package goshared

const enumTpl = `
		{{ $f := .Field }}{{ $r := .Rules }}
		{{ template "const" . }}
		{{ template "in" . }}

		{{ if $r.GetDefinedOnly }}
			if _, ok := {{ (typ $f).Element }}_name[int32({{ accessor . }})]; !ok {
				return {{ err . "必须是定义的枚举值之一" }}
			}
		{{ end }}
`
