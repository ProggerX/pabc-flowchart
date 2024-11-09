package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func parseBlock(id string, s string) (string, string, []string) {
	bid := id + "beg"
	eid := id + "end"
	o := []string{bid + "([Начало блока])"}
	scope := 0
	ops := []string{}
	left := 0
	for i := 0; i < len(s); i++ {
		l := s[i:]
		if strings.HasPrefix(l, "begin ") || strings.HasPrefix(l, "begin;") {
			scope++
		} else if strings.HasPrefix(l, "end") {
			scope--
		}
		if l[0] == ';' && scope == 0 {
			ops = append(ops, s[left:i])
			left = i + 1
		}
	}
	if s[len(s)-1] != ';' {
		var i int
		for i = len(s) - 1; s[i] != ';'; i-- {
		}
		last := s[i+2:]
		log.Debug("BLOCK", "id", id, "last_op", last)
		ops = append(ops, last)
	}
	log.Debug("BLOCK", "id", id, "len(ops)", len(ops))
	prev_eid := bid
	for i := 0; i < len(ops); i++ {
		ops[i] = strings.TrimSpace(ops[i])
		log.Debug("BLOCK", "id", id, "i", i, "ops[i]", ops[i])
		o_bid, o_eid, o_r := parseOperator(id+strconv.Itoa(i), ops[i])
		o = append(o, o_r...)
		o = append(o, fmt.Sprintf("%s-->%s", prev_eid, o_bid))
		prev_eid = o_eid
	}
	o = append(o, fmt.Sprintf("%s-->%s", prev_eid, eid))
	o = append(o, eid+"([Конец блока])")
	return bid, eid, o
}
