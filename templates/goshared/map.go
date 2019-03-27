package goshared

const mapTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.GetMinPairs }}
		{{ if eq $r.GetMinPairs $r.GetMaxPairs }}
			if len({{ accessor . }}) != {{ $r.GetMinPairs }} {
				return {{ err . "必须包含 " $r.GetMinPairs " 键值对" }}
			}
		{{ else if $r.MaxPairs }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinPairs }} || l > {{ $r.GetMaxPairs }} {
			 	return {{ err . "必须包含在 " $r.GetMinPairs " 和 " $r.GetMaxPairs " 键值对" }}
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinPairs }} {
				return {{ err . "必须至少包含 " $r.GetMinPairs " 键值对" }}
			}
		{{ end }}
	{{ else if $r.MaxPairs }}
		if len({{ accessor . }}) > {{ $r.GetMaxPairs }} {
			return {{ err . "包含不超过 " $r.GetMaxPairs " 键值对" }}
		}
	{{ end }}

	{{ if or $r.GetNoSparse (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
		for key, val := range {{ accessor . }} {
			_ = val

			{{ if $r.GetNoSparse }}
				if val == nil {
					return {{ errIdx . "key" "不能是空值" }}
				}
			{{ end }}

			{{ render (.Key "key" "key") }}

			{{ render (.Elem "val" "key") }}
		}
	{{ end }}
`
