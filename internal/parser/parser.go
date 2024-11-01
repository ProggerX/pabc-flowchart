package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

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

func parseAssign(id string, s string) (string, string, []string) {
	rx := regexp.MustCompile(`^(.+) +(\+|-|\*|\/|:)= +(.+)`)
	name := rx.FindStringSubmatch(s)[1]
	oper := rx.FindStringSubmatch(s)[2]
	val := rx.FindStringSubmatch(s)[3]
	log.Debug("ASSIGN", "id", id, "name", name)
	log.Debug("ASSIGN", "id", id, "oper", oper)
	log.Debug("ASSIGN", "id", id, "val", val)
	if oper == ":" {
		return id, id, []string{fmt.Sprintf("%s[\"Присвоить %s значение %s\"]", id, name, val)}
	}
	return id, id, []string{fmt.Sprintf("%s[\"Присвоить %s значение %s %s %s\"]", id, name, name, oper, val)}
}

func parseFor(id string, s string) (string, string, []string) {

}

func parseOperator(id string, s string) (string, string, []string) {
	rx_block := regexp.MustCompile(`^begin +(.+) +end`)
	rx_if_else := regexp.MustCompile(`^if +(.+) +then +(.+) else +(.+)`)
	rx_if := regexp.MustCompile(`^if +(.+) +then +(.+)`)
	rx_assign := regexp.MustCompile(`^(.+) +(\+|-|\*|\/|:)= +(.+)`)
	rx_for := regexp.MustCompile(`^for +(var)? +(.+) +:= +(.+) +(to|downto) +do`)
	is_block := rx_block.MatchString(s)
	is_if_else := rx_if_else.MatchString(s)
	is_if := rx_if.MatchString(s)
	is_assign := rx_assign.MatchString(s)
	is_for := rx_for.MatchString(s)
	if is_block {
		return parseBlock(id+"b", rx_block.FindStringSubmatch(s)[1])
	}
	if is_if_else {
		return parseIfElse(id+"ie", s)
	}
	if is_if {
		return parseIf(id+"if", s)
	}
	if is_assign {
		return parseAssign(id+"as", s)
	}
	if is_for {
		return parseFor(id+"for", s)
	}
	return id, id, []string{fmt.Sprintf("%s[\"%s\"]", id, s)}
}

func parseBlock(id string, s string) (string, string, []string) {
	bid := id + "beg"
	eid := id + "end"
	o := []string{bid + "([Начало блока])"}
	scope := 0
	ops := []string{}
	left := 0
	for i := 0; i < len(s); i++ {
		l := s[i:]
		if strings.HasPrefix(l, "begin ") {
			scope++
		} else if strings.HasPrefix(l, "end") {
			scope--
		}
		if l[0] == ';' && scope == 0 {
			ops = append(ops, s[left:i])
			left = i + 1
		}
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
	o = append(o, eid+"([Конец блока])")
	o = append(o, fmt.Sprintf("%s-->%s", prev_eid, eid))
	return bid, eid, o
}

func parseMainBlock(s string) []string {
	rx := regexp.MustCompile(`begin(.*)end\.`)
	s = rx.FindStringSubmatch(s)[1]
	s = strings.TrimSpace(s)
	log.Debug("MAIN_BLOCK", "s", s)
	_, _, o := parseBlock("mb", s)
	return o
}

func parseVarBlock(s string) []string {
	o := []string{}
	rx := regexp.MustCompile(`begin.*end\.`)
	s = rx.ReplaceAllString(s, "")
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "var ")
	log.Debug("VAR_BLOCK", "s", s)
	vars := strings.Split(s, ";")
	vars = vars[:len(vars)-1]
	for i := 0; i < len(vars); i++ {
		rx1 := regexp.MustCompile(`.* *:`)
		rx2 := regexp.MustCompile(`: *.*`)
		name := strings.Trim(rx1.FindString(vars[i]), " :")
		typ := strings.Trim(rx2.FindString(vars[i]), " :")
		id := "var" + strconv.Itoa(i)
		if i < len(vars)-1 {
			idnext := "var" + strconv.Itoa(i+1)
			o = append(o, fmt.Sprintf("%s-->%s", id, idnext))
		}
		o = append(o, fmt.Sprintf("%s[Объявить %s типа %s]", id, name, typ))
	}
	return o
}

func ParseFile(lines []string) []string {
	for i := 0; i < len(lines); i++ {
		rx := regexp.MustCompile(`\/\/.*`)
		lines[i] = string(rx.ReplaceAllString(lines[i], ""))
	}
	return parseCode(strings.Join(lines, " "))
}

func parseCode(s string) []string {
	o := []string{"flowchart TB"}
	s = strings.TrimSpace(s)
	log.Debug("CODE", "s", s)
	if strings.HasPrefix(s, "var") {
		o = append(o, parseVarBlock(s)...)
	}
	o = append(o, parseMainBlock(s)...)

	return o
}
