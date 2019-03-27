package goshared

const strTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "const" . }}
	{{ template "in" . }}

	{{ if or $r.Len (and $r.MinLen $r.MaxLen (eq $r.GetMinLen $r.GetMaxLen)) }}
		{{ if $r.Len }}
			if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetLen }} {
			return {{ err . "长度必须是 " $r.GetLen " 字符" }}
		{{ else }}
		if utf8.RuneCountInString({{ accessor . }}) != {{ $r.GetMinLen }} {
			return {{ err . "长度必须是 " $r.GetMinLen " 字符" }}
		{{ end }}
	}
	{{ else if $r.MinLen }}
		{{ if $r.MaxLen }}
			if l := utf8.RuneCountInString({{ accessor . }}); l < {{ $r.GetMinLen }} || l > {{ $r.GetMaxLen }} {
				return {{ err . "长度必须在 " $r.GetMinLen " 和 " $r.GetMaxLen " 之间" }}
			}
		{{ else }}
			if utf8.RuneCountInString({{ accessor . }}) < {{ $r.GetMinLen }} {
				return {{ err . "长度必须最少为 " $r.GetMinLen " 字符" }}
			}
		{{ end }}
	{{ else if $r.MaxLen }}
		if utf8.RuneCountInString({{ accessor . }}) > {{ $r.GetMaxLen }} {
			return {{ err . "长度必须最多为 " $r.GetMaxLen " 字符" }}
		}
	{{ end }}

	{{ if or $r.LenBytes (and $r.MinBytes $r.MaxBytes (eq $r.GetMinBytes $r.GetMaxBytes)) }}
		{{ if $r.LenBytes }}
			if len({{ accessor . }}) != {{ $r.GetLenBytes }} {
				return {{ err . "长度必须等于 " $r.GetLenBytes " 字符" }}
			}
		{{ else }}
			if len({{ accessor . }}) != {{ $r.GetMinBytes }} {
				return {{ err . "长度必须等于 " $r.GetMinBytes " 字符" }}
			}
		{{ end }}
	{{ else if $r.MinBytes }}
		{{ if $r.MaxBytes }}
			if l := len({{ accessor . }}); l < {{ $r.GetMinBytes }} || l > {{ $r.GetMaxBytes }} {
					return {{ err . "值长度必须介于 " $r.GetMinBytes " - " $r.GetMaxBytes " 字符之间" }}
			}
		{{ else }}
			if len({{ accessor . }}) < {{ $r.GetMinBytes }} {
				return {{ err . "长度必须至少为 " $r.GetMinBytes " 字符" }}
			}
		{{ end }}
	{{ else if $r.MaxBytes }}
		if len({{ accessor . }}) > {{ $r.GetMaxBytes }} {
			return {{ err . "长度必须至多为 " $r.GetMaxBytes " 字符" }}
		}
	{{ end }}

	{{ if $r.Prefix }}
		if !strings.HasPrefix({{ accessor . }}, {{ lit $r.GetPrefix }}) {
			return {{ err . "没有前缀 " (lit $r.GetPrefix) }}
		}
	{{ end }}

	{{ if $r.Suffix }}
		if !strings.HasSuffix({{ accessor . }}, {{ lit $r.GetSuffix }}) {
			return {{ err . "没有后缀 " (lit $r.GetSuffix) }}
		}
	{{ end }}

	{{ if $r.Contains }}
		if !strings.Contains({{ accessor . }}, {{ lit $r.GetContains }}) {
			return {{ err . "不包含子字符串 " (lit $r.GetContains) }}
		}
	{{ end }}

	{{ if $r.GetIp }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil {
			return {{ err . "必须是合法ip" }}
		}
	{{ else if $r.GetIpv4 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() == nil {
			return {{ err . "必须是合法ipv4" }}
		}
	{{ else if $r.GetIpv6 }}
		if ip := net.ParseIP({{ accessor . }}); ip == nil || ip.To4() != nil {
			return {{ err . "必须是合法ipv6" }}
		}
	{{ else if $r.GetEmail }}
		if err := m._validateEmail({{ accessor . }}); err != nil {
			return {{ errCause . "err" "必须是合法email" }}
		}
	{{ else if $r.GetHostname }}
		if err := m._validateHostname({{ accessor . }}); err != nil {
			return {{ errCause . "err" "必须是合法主机名" }}
		}
	{{ else if $r.GetUri }}
		if uri, err := url.Parse({{ accessor . }}); err != nil {
			return {{ errCause . "err" "必须是合法URL" }}
		} else if !uri.IsAbs() {
			return {{ err . "必须是绝对值" }}
		}
	{{ else if $r.GetUriRef }}
		if _, err := url.Parse({{ accessor . }}); err != nil {
			return {{ errCause . "err" "必须是合法URL" }}
		}
	{{ end }}

	{{ if $r.Pattern }}
	if !{{ lookup $f "Pattern" }}.MatchString({{ accessor . }}) {
		return {{ err . "不能匹配正则表达式 " (lit $r.GetPattern) }}
	}
{{ end }}
`
