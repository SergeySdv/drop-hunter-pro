package repository

import (
	"strings"
)

func Cols(tableCols []string, prefix string) string {

	tmp := make([]string, len(tableCols))

	for i := range tableCols {
		tmp[i] = prefix + tableCols[i]
	}

	pre := strings.Join(tmp, ",")

	return pre
}

func Join(strs ...string) string {
	return strings.Join(strs, "")
}

type Helper struct {
	cols []string

	preparedColsJoinComma string
	preparedColsJoinColon string
	preparedColsJoinPref  string
}

func NewHelper(cols []string) *Helper {
	return &Helper{
		cols:                  nil,
		preparedColsJoinComma: Cols(cols, ""),
		preparedColsJoinColon: Cols(cols, ":"),
		preparedColsJoinPref:  Cols(cols, "p."),
	}
}
func (h *Helper) Cols() string {
	a := h.preparedColsJoinComma
	return a
}

func (h *Helper) ColsColon() string {
	a := h.preparedColsJoinColon
	return a
}

func (h *Helper) ColsPref() string {
	a := h.preparedColsJoinPref
	return a
}
