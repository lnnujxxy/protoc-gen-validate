package goshared

const constTpl = `{{ $r := .Rules }}
	{{ if $r.Const }}
		if {{ accessor . }} != {{ lit $r.GetConst }} {
			return {{ err . "必须等于 " $r.GetConst }}
		}
	{{ end }}
`
