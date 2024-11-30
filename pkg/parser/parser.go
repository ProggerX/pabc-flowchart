package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
)

func parseOperator(id string, s string) (string, string, []string) {
	s = strings.TrimSpace(s)
	rx_block := regexp.MustCompile(`^begin;* +(.+) +end`)
	rx_if := regexp.MustCompile(`^if +(.+) +then +(.+)`)
	rx_assign := regexp.MustCompile(`^(\w+) +(\+|-|\*|\/|:)= +(.+)`)
	rx_for := regexp.MustCompile(`^for +(var +)?(.+) *:= *(.+) +(to|downto) +(.+?)( +step +(.+))? +do +(.+)`)
	rx_read_write := regexp.MustCompile(`^(write|read)(ln)?(\((.*)\))?`)
	rx_while := regexp.MustCompile(`^while +(.+) +(.+)`)
	is_block := rx_block.MatchString(s)
	is_if_else := detectIfElse(s)
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

func parsePreMain(string) []string {
	log.Debug("PreMain is not supported in this version")
	return []string{""}
}

func ParseFile(lines []string) []string {
	for i := 0; i < len(lines); i++ {
		rx := regexp.MustCompile(`\/\/.*`)
		lines[i] = string(rx.ReplaceAllString(lines[i], ""))
	}
	return parseCode(strings.Join(lines, " "))
}

func findMainBlock(s string) (int, int) {
	end := strings.Index(s, "end.") + 3
	scope := 0
	for i := end; i >= 0; i-- {
		l := s[i:]
		if strings.HasPrefix(l, "end ") || strings.HasPrefix(l, "end;") {
			scope++
		} else if strings.HasPrefix(l, "begin ") || strings.HasPrefix(l, "begin;") {
			if scope == 0 {
				return i, end
			}
			scope--
		}
	}
	return 0, end
}

func parseCode(s string) []string {
	o := []string{"flowchart TB"}
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\r")
	log.Debug("CODE", "s", s)
	premain_rx := regexp.MustCompile(`(.*?) +begin`)
	if premain_rx.MatchString(s) {
		o = append(o, parsePreMain(premain_rx.FindStringSubmatch(s)[1])...)
	}
	beg, end := findMainBlock(s)
	mb_s := s[beg:(end + 1)]
	mb_rx := regexp.MustCompile(`begin(.*)end\.`)
	o = append(o, parseMainBlock(mb_rx.FindStringSubmatch(mb_s)[1])...)

	return o
}
