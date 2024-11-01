package parser

import (
	"fmt"
	"regexp"

	"github.com/charmbracelet/log"
)

func parseIf(id string, s string) (string, string, []string) {
	rx := regexp.MustCompile(`if +(.+) +then +(.+)`)
	cond := rx.FindStringSubmatch(s)[1]
	then := rx.FindStringSubmatch(s)[2]
	log.Debug("IF", "id", id, "cond", cond)
	log.Debug("IF", "id", id, "then", then)
	o := []string{}
	o = append(o, fmt.Sprintf("%s{\"%s ?\"}", id, cond))
	then_bid, then_eid, then_r := parseOperator(id+"then", then)
	o = append(o, then_r...)
	o = append(o, fmt.Sprintf("%s-->|тогда|%s", id, then_bid))
	o = append(o, fmt.Sprintf("%s-->|иначе|%s", id, id+"end"))
	o = append(o, fmt.Sprintf("%s-->%s", then_eid, id+"end"))
	o = append(o, fmt.Sprintf("%s[Конец условия]", id+"end"))
	return id, id + "end", o
}

func parseIfElse(id string, s string) (string, string, []string) {
	rx := regexp.MustCompile(`if +(.+) +then +(.+) +else +(.+)`)
	cond := rx.FindStringSubmatch(s)[1]
	then := rx.FindStringSubmatch(s)[2]
	els := rx.FindStringSubmatch(s)[3]
	log.Debug("IF_ELSE", "id", id, "cond", cond)
	log.Debug("IF_ELSE", "id", id, "then", then)
	log.Debug("IF_ELSE", "id", id, "else", els)
	o := []string{}
	o = append(o, fmt.Sprintf("%s{\"%s ?\"}", id, cond))
	then_bid, then_eid, then_r := parseOperator(id+"then", then)
	else_bid, else_eid, else_r := parseOperator(id+"else", els)
	o = append(o, then_r...)
	o = append(o, else_r...)
	o = append(o, fmt.Sprintf("%s-->|тогда|%s", id, then_bid))
	o = append(o, fmt.Sprintf("%s-->|иначе|%s", id, else_bid))
	o = append(o, fmt.Sprintf("%s-->%s", else_eid, id+"end"))
	o = append(o, fmt.Sprintf("%s-->%s", then_eid, id+"end"))
	o = append(o, fmt.Sprintf("%s[Конец условия]", id+"end"))
	return id, id + "end", o
}
