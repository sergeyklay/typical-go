package typast

import (
	"regexp"
	"strings"
)

// ParseAnnots to parse godoc comment to list of annotation
func ParseAnnots(doc string) (annotations []*Annot) {
	if doc == "" {
		return
	}
	r, _ := regexp.Compile("\\[(.*?)\\]")
	for _, s := range r.FindAllString(doc, -1) {
		if a := ParseAnnotation(s); a != nil {
			annotations = append(annotations, a)
		}
	}
	return
}

// ParseAnnotation parse raw string to annotation
func ParseAnnotation(raw string) (a *Annot) {
	if raw[0] != '[' && raw[len(raw)-1] != ']' {
		return
	}
	raw = raw[1 : len(raw)-1]
	a = NewAnnot(name(raw))
	putAttr(a, rawAttribute(raw))
	return
}

func name(raw string) string {
	i := strings.IndexRune(raw, '(')
	if i < 0 {
		return raw
	}
	return raw[:i]
}

func rawAttribute(raw string) string {
	i0 := strings.IndexRune(raw, '(')
	i1 := strings.IndexRune(raw, ')')
	if i0 < 0 || i1 < 0 {
		return ""
	}
	return raw[i0+1 : i1]
}

func putAttr(a *Annot, rawAttr string) {
	rawAttr = strings.TrimSpace(rawAttr)
	if rawAttr == "" {
		return
	}
	var key, value string
	eq := strings.IndexRune(rawAttr, '=')
	space := strings.IndexRune(rawAttr, ' ')
	if space > 0 && (space < eq || eq < 1) {
		a.Put(rawAttr[:space], "")
		putAttr(a, rawAttr[space+1:])
		return
	}
	if eq > 0 {
		key = rawAttr[:eq]
		if eq == len(rawAttr)-1 {
			a.Put(key, "")
			return
		}
		value = rawAttr[eq+1:]
		if value[0] == '"' {
			value = value[1:]
			i := strings.IndexRune(value, '"')
			a.Put(key, value[:i])
			putAttr(a, value[i+1:])
			return
		}
		if value[0] == ' ' {
			a.Put(key, "")
			putAttr(a, value)
			return
		}
		if i := strings.IndexRune(value, ' '); i > 0 {
			a.Put(key, value[:i])
			putAttr(a, value[i+1:])
			return
		}
		a.Put(key, value)
		return
	}
	a.Put(rawAttr, "")
}
