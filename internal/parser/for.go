package parser

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/charmbracelet/log"
)

func parseFor(id string, s string) (string, string, []string) {
	o := []string{}
	rx := regexp.MustCompile(`^for +(var +)?(.+) *:= *(.+) +(to|downto) +(.+?)( +step +(.+))? +do +(.+)`)
	name := rx.FindStringSubmatch(s)[2]
	val := rx.FindStringSubmatch(s)[3]
	end_val := rx.FindStringSubmatch(s)[5]
	to_or_downto := rx.FindStringSubmatch(s)[4]
	is_neg := to_or_downto == "downto"
	log.Debug("FOR", "id", id, "name", name)
	log.Debug("FOR", "id", id, "val", val)
	log.Debug("FOR", "id", id, "end_val", end_val)
	log.Debug("FOR", "id", id, "is_neg", is_neg)
	o = append(o, fmt.Sprintf("%s[\"Присвоить %s значение %s\"]", id+"assign", name, val))
	var r string
	step := 1
	if is_neg {
		r = fmt.Sprintf("%s{\"%s >= %s ?\"}", id, name, end_val)
	} else {
		r = fmt.Sprintf("%s{\"%s <= %s ?\"}", id, name, end_val)
	}
	if rx.FindStringSubmatch(s)[7] != "" {
		step, _ = strconv.Atoi(rx.FindStringSubmatch(s)[7])
	}
	log.Debug("FOR", "id", id, "step", step)
	o = append(o, r)
	o = append(o, fmt.Sprintf("%s-->%s", id+"assign", id))
	if is_neg {
		r = fmt.Sprintf("%s[\"Присвоить %s значение %s - %d\"]", id+"change", name, name, step)
	} else {
		r = fmt.Sprintf("%s[\"Присвоить %s значение %s + %d\"]", id+"change", name, name, step)
	}
	o = append(o, r)
	op_bid, op_eid, op_r := parseOperator(id+"body", rx.FindStringSubmatch(s)[8])
	o = append(o, op_r...)
	o = append(o, fmt.Sprintf("%s-->|тогда|%s", id, op_bid))
	o = append(o, fmt.Sprintf("%s-->%s", id+"change", id))
	o = append(o, fmt.Sprintf("%s-->%s", op_eid, id+"change"))
	return id + "assign", id, o
}
