package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func parseOperator(id string, s string) (string, string, []string) {
	rx_block := regexp.MustCompile(`^begin +(.+) +end`)
	rx_if_else := regexp.MustCompile(`^if +(.+?) +then +((begin +)?.+(end *)?) +else +(.+)`)
	rx_if := regexp.MustCompile(`^if +(.+) +then +(.+)`)
	rx_assign := regexp.MustCompile(`^(\w+) +(\+|-|\*|\/|:)= +(.+)`)
	rx_for := regexp.MustCompile(`^for +(var +)?(.+) +:= +(.+) +(to|downto) +(.+?)( +step +(.+))? +do +(.+)`)
	rx_read_write := regexp.MustCompile(`^(write|read)(ln)?(\((.*)\))?`)
	rx_while := regexp.MustCompile(`^while +(.+) +(.+)`)
	is_block := rx_block.MatchString(s)
	is_if_else := rx_if_else.MatchString(s)
	is_if := rx_if.MatchString(s)
	is_assign := rx_assign.MatchString(s)
	is_for := rx_for.MatchString(s)
	is_read_write := rx_read_write.MatchString(s)
	is_while := rx_while.MatchString(s)
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
	if is_read_write {
		return id, id, []string{fmt.Sprintf("%s[/\"%s\"/]", id, s)}
	}
	if is_while {
		return parseWhile(id+"wh", s)
	}
	return id, id, []string{fmt.Sprintf("%s[\"%s\"]", id, s)}
}

func parseMainBlock(s string) []string {
	s = strings.TrimSpace(s)
	log.Debug("MAIN_BLOCK", "s", s)
	bid, eid, o := parseBlock("mb", s)
	o[0] = fmt.Sprintf("%s([НАЧАЛО])", bid)
	o[len(o)-1] = fmt.Sprintf("%s([КОНЕЦ])", eid)
	return o
}

func parseVarBlock(s string) []string {
	o := []string{}
	s = strings.TrimSpace(s)
	rx := regexp.MustCompile(`var +(((.+) *: *(.+) *; *)+) +begin`)
	s = rx.FindStringSubmatch(s)[1]
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
		log.Debug("VAR", "name", name, "type", typ)
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
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\r")
	log.Debug("CODE", "s", s)
	var_rx := regexp.MustCompile(`var +(((.+) *: *(.+) *; *)+) +begin`)
	if var_rx.MatchString(s) {
		o = append(o, parseVarBlock(s)...)
	}
	mb_rx := regexp.MustCompile(`begin(.*)end\.`)
	log.Debug("MAINBLOCK", "len", len(mb_rx.FindStringSubmatch(s)))
	o = append(o, parseMainBlock(mb_rx.FindStringSubmatch(s)[1])...)

	return o
}
