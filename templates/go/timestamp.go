package golang

const timestampTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if or $r.Lt $r.Lte $r.Gt $r.Gte $r.LtNow $r.GtNow $r.Within $r.Const }}
		if t := {{ accessor . }}; t != nil {
			ts, err := ptypes.Timestamp(t)
			if err != nil { return {{ errCause . "err" "不是有效的时间戳" }} }

			{{ template "timestampcmp" . }}
		}
	{{ end }}
`
