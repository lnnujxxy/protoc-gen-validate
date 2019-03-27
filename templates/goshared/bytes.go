package goshared

const bytesTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
			if len({{ accessor . }}) != {{ $r.GetLen }} {
				return {{ err . "长度必须是 " $r.GetLen " 字节" }}
			}
		{{ else }}
			if len({{ accessor . }}) != {{ $r.GetMinLen }} {
				return {{ err . "长度必须是 " $r.GetMinLen " 字节" }}
			}
		{{ end }}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				return {{ err . "长度必须介于 " $r.GetMinLen " - " $r.GetMaxLen " 字节之间" }}
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinLen }} {
				return {{ err . "长度至少 " $r.GetMinLen " 字节" }}
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if len({{ accessor . }}) > {{ $r.GetMaxLen }} {
			return {{ err . "长度至多 " $r.GetMaxLen " 字节" }}
		}
	{{ end }}

	{{ if $r.Prefix }}
		if !bytes.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			return {{ err . "不能有前缀 " (byteStr $r.GetPrefix) }}
		}
	{{ end }}

	{{ if $r.Suffix }}
		if !bytes.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			return {{ err . "不能有后缀 " (byteStr $r.GetSuffix) }}
		}
	{{ end }}

	{{ if $r.Contains }}
		if !bytes.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			return {{ err . "不能包含 " (byteStr $r.GetContains) }}
		}
	{{ end }}

	{{ if $r.In }}
		if _, ok := {{ lookup $f "InLookup" }}[string({{ accessor . }})]; !ok {
			return {{ err . "必须在列表中 " $r.In }}
		}
	{{ else if $r.NotIn }}
		if _, ok := {{ lookup $f "NotInLookup" }}[string({{ accessor . }})]; ok {
			return {{ err . "必须不能在列表中 " $r.NotIn }}
		}
	{{ end }}

	{{ if $r.Const }}
		if !bytes.Equal({{ accessor . }}, {{ lit $r.Const }}) {
			return {{ err . "必须等于 " $r.Const }}
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil {
			return {{ err . "必须是合法ip" }}
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.IP({{ accessor . }}); ip.To4() == nil {
			return {{ err . "必须是合法ipv4" }}
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.IP({{ accessor . }}); ip.To16() == nil || ip.To4() != nil {
			return {{ err . "必须是合法ipv6" }}
		}
	{{ end }}

	{{ if $r.Pattern }}
	if !{{ lookup $f "Pattern" }}.Match({{ accessor . }}) {
		return {{ err . "不能匹配正则 " (lit $r.GetPattern) }}
	}
	{{ end }}
`
