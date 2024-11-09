package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
)

func detectIfElse(s string) bool {
	rx_common_sense := regexp.MustCompile(`^if +(.+?) +then +(.+) +else +(.+)`)
	if !rx_common_sense.MatchString(s) {
		return false
	}
	rx := regexp.MustCompile(`^if +(.+?) +then`)
	old_s := s
	s = rx.ReplaceAllString(s, "")
	s = strings.TrimSpace(s)

	scope := 0
	scope_if := 0
	for i := 0; i < len(s); i++ {
		l := s[i:]
		if (strings.HasPrefix(l, "begin ") || strings.HasPrefix(l, "begin;")) && (i == 0 || s[i-1] == ' ') {
			scope++
		} else if strings.HasPrefix(l, " end") {
			scope--
		} else if strings.HasPrefix(l, "if ") && scope == 0 && (i == 0 || s[i-1] == ' ') {
			scope_if++
		} else if strings.HasPrefix(l, " else ") && scope == 0 {
			if scope_if <= 0 {
				log.Debug("DETECTED_IF_ELSE", "s", old_s)
				return true
			}
			scope_if--
		}
	}
	return false
}

func parseIf(id string, s string) (string, string, []string) {
	rx := regexp.MustCompile(`if +(.+?) +then +(.+)`)
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
	rx := regexp.MustCompile(`^if (.+?) then`)
	cond := rx.FindStringSubmatch(s)[1]
	s = rx.ReplaceAllString(s, "")

	scope := 0
	scope_if := 0
	var then_op, else_op string
	for i := 0; i < len(s); i++ {
		l := s[i:]
		if (strings.HasPrefix(l, "begin ") || strings.HasPrefix(l, "begin;")) && (i == 0 || s[i-1] == ' ') {
			scope++
		} else if strings.HasPrefix(l, " end") {
			scope--
		} else if strings.HasPrefix(l, "if ") && scope == 0 && (i == 0 || s[i-1] == ' ') {
			scope_if++
		} else if strings.HasPrefix(l, " else ") && scope == 0 {
			if scope_if > 0 {
				scope_if--
			} else {
				log.Debug("delimeter", "i", i)
				then_op = s[1:i]
				else_op = s[i+6:]
				break
			}
		}
	}

	log.Debug("IF_ELSE", "id", id, "cond", cond)
	log.Debug("IF_ELSE", "id", id, "then_op", then_op)
	log.Debug("IF_ELSE", "id", id, "else_op", else_op)

	then_bid, then_eid, then_ops := parseOperator(id+"then", then_op)
	else_bid, else_eid, else_ops := parseOperator(id+"else", else_op)

	o := []string{}
	o = append(o, fmt.Sprintf("%s{\"%s ?\"}", id, cond))
	o = append(o, then_ops...)
	o = append(o, else_ops...)
	o = append(o, fmt.Sprintf("%s-->|тогда|%s", id, then_bid))
	o = append(o, fmt.Sprintf("%s-->|иначе|%s", id, else_bid))
	o = append(o, fmt.Sprintf("%s-->%s", else_eid, id+"end"))
	o = append(o, fmt.Sprintf("%s-->%s", then_eid, id+"end"))
	o = append(o, fmt.Sprintf("%s[Конец условия]", id+"end"))
	return id, id + "end", o
}
