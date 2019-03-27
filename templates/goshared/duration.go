package goshared

const durationcmpTpl = `{{ $f := .Field }}{{ $r := .Rules }}
			{{  if $r.Const }}
				if dur != {{ durLit $r.Const }} {
					return {{ err . "必须等于 " (durStr $r.Const) }}
				}
			{{ end }}


			{{  if $r.Lt }}  lt  := {{ durLit $r.Lt }};  {{ end }}
			{{- if $r.Lte }} lte := {{ durLit $r.Lte }}; {{ end }}
			{{- if $r.Gt }}  gt  := {{ durLit $r.Gt }};  {{ end }}
			{{- if $r.Gte }} gte := {{ durLit $r.Gte }}; {{ end }}

			{{ if $r.Lt }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLt $r.GetGt }}
						if dur <= gt || dur >= lt {
							return {{ err . "必须在范围内 (" (durStr $r.GetGt) ", " (durStr $r.GetLt) ")" }}
						}
					{{ else }}
						if dur >= lt && dur <= gt {
							return {{ err . "必须在范围外 [" (durStr $r.GetLt) ", " (durStr $r.GetGt) "]" }}
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{  if durGt $r.GetLt $r.GetGte }}
						if dur < gte || dur >= lt {
							return {{ err . "必须在范围内 [" (durStr $r.GetGte) ", " (durStr $r.GetLt) ")" }}
						}
					{{ else }}
						if dur >= lt && dur < gte {
							return {{ err . "必须在范围外 [" (durStr $r.GetLt) ", " (durStr $r.GetGte) ")" }}
						}
					{{ end }}
				{{ else }}
					if dur >= lt {
						return {{ err . "必须小于 " (durStr $r.GetLt) }}
					}
				{{ end }}
			{{ else if $r.Lte }}
				{{ if $r.Gt }}
					{{  if durGt $r.GetLte $r.GetGt }}
						if dur <= gt || dur > lte {
							return {{ err . "必须在范围内 (" (durStr $r.GetGt) ", " (durStr $r.GetLte) "]" }}
						}
					{{ else }}
						if dur > lte && dur <= gt {
							return {{ err . "必须在范围外 (" (durStr $r.GetLte) ", " (durStr $r.GetGt) "]" }}
						}
					{{ end }}
				{{ else if $r.Gte }}
					{{ if durGt $r.GetLte $r.GetGte }}
						if dur < gte || dur > lte {
							return {{ err . "必须在范围内 [" (durStr $r.GetGte) ", " (durStr $r.GetLte) "]" }}
						}
					{{ else }}
						if dur > lte && dur < gte {
							return {{ err . "必须在范围外 (" (durStr $r.GetLte) ", " (durStr $r.GetGte) ")" }}
						}
					{{ end }}
				{{ else }}
					if dur > lte {
						return {{ err . "必须小于或等于 " (durStr $r.GetLte) }}
					}
				{{ end }}
			{{ else if $r.Gt }}
				if dur <= gt {
					return {{ err . "必须大于 " (durStr $r.GetGt) }}
				}
			{{ else if $r.Gte }}
				if dur < gte {
					return {{ err . "必须大于或等于 " (durStr $r.GetGte) }}
				}
			{{ end }}


			{{ if $r.In }}
				if _, ok := {{ lookup $f "InLookup" }}[dur]; !ok {
					return {{ err . "必须在列表中 " $r.In }}
				}
			{{ else if $r.NotIn }}
				if _, ok := {{ lookup $f "NotInLookup" }}[dur]; ok {
					return {{ err . "必须不能在列表中 " $r.NotIn }}
				}
			{{ end }}
`
