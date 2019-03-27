package goshared

const ltgtTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLt $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val >= {{ $r.Lt }} {
					return {{ err . "必须在范围内 (" $r.GetGt ", " $r.GetLt ")" }}
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val <= {{ $r.Gt }} {
					return {{ err . "必须在范围之外 [" $r.GetLt ", " $r.GetGt "]" }}
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{  if gt $r.GetLt $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val >= {{ $r.Lt }} {
					return {{ err . "必须在范围内 [" $r.GetGte ", " $r.GetLt ")" }}
				}
			{{ else }}
				if val := {{ accessor . }}; val >= {{ $r.Lt }} && val < {{ $r.Gte }} {
					return {{ err . "必须在范围之外 [" $r.GetLt ", " $r.GetGte ")" }}
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} >= {{ $r.Lt }} {
				return {{ err . "必须小于 " $r.GetLt }}
			}
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLte $r.GetGt }}
				if val := {{ accessor . }};  val <= {{ $r.Gt }} || val > {{ $r.Lte }} {
					return {{ err . "必须在范围内 (" $r.GetGt ", " $r.GetLte "]" }}
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val <= {{ $r.Gt }} {
					return {{ err . "必须在范围之外 (" $r.GetLte ", " $r.GetGt "]" }}
				}
			{{ end }}
		{{ else if $r.Gte }}
			{{ if gt $r.GetLte $r.GetGte }}
				if val := {{ accessor . }};  val < {{ $r.Gte }} || val > {{ $r.Lte }} {
					return {{ err . "必须在范围内 [" $r.GetGte ", " $r.GetLte "]" }}
				}
			{{ else }}
				if val := {{ accessor . }}; val > {{ $r.Lte }} && val < {{ $r.Gte }} {
					return {{ err . "必须在范围之外 (" $r.GetLte ", " $r.GetGte ")" }}
				}
			{{ end }}
		{{ else }}
			if {{ accessor . }} > {{ $r.Lte }} {
				return {{ err . "必须小于或等于 " $r.GetLte }}
			}
		{{ end }}
	{{ else if $r.Gt }}
		if {{ accessor . }} <= {{ $r.Gt }} {
			return {{ err . "必须大于 " $r.GetGt }}
		}
	{{ else if $r.Gte }}
		if {{ accessor . }} < {{ $r.Gte }} {
			return {{ err . "必须大于或等于 " $r.GetGte }}
		}
	{{ end }}
`
