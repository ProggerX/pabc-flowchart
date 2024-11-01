package parser

import (
	"fmt"
	"regexp"

	"github.com/charmbracelet/log"
)

func parseWhile(id string, s string) (string, string, []string) {
	o := []string{}
	rx := regexp.MustCompile(`^while +(.+) +do +(.+)`)
	cond := rx.FindStringSubmatch(s)[1]
	ops := rx.FindStringSubmatch(s)[2]
	log.Debug("WHILE", "id", id, "cond", cond)
	log.Debug("WHILE", "id", id, "ops", ops)
	op_bid, op_eid, op_r := parseOperator(id+"body", ops)
	o = append(o, fmt.Sprintf("%s{\"%s ? \"}", id, cond))
	o = append(o, op_r...)
	o = append(o, fmt.Sprintf("%s-->|тогда|%s", id, op_bid))
	o = append(o, fmt.Sprintf("%s-->%s", op_eid, id))
	return id, id, o
}
