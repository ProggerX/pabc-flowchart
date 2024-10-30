package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func parseOperator(i int, s string, t int) []string {
	return []string{}
}

func parseBlock(i int, s string) []string {
	id := strconv.Itoa(i) + "beg(Начало блока)"
	eid := strconv.Itoa(i) + "end(Конец блока)"
	o := []string{id}
	scope := 0
	ops := []string{}
	left := 0
	for i := 0; i < len(s); i++ {
		l := s[i:]
		if strings.HasSuffix(l, "begin ") {
			scope++
		} else if strings.HasSuffix(l, "end;") || strings.HasSuffix(l, "end.") {
			scope--
		}
		if l[0] == ';' && scope == 0 {
			ops = append(ops, s[left:i])
			left = i + 1
		}
	}
	log.Debug("operators:", "ops", ops, "len", len(ops))
	o = append(o, eid)
	return o
}

func parseMainBlock(s string) []string {
	o := []string{"mainbeg(Начало)"}
	rx := regexp.MustCompile(`begin(.*)end\.`)
	s = rx.FindStringSubmatch(s)[1]
	s = strings.TrimSpace(s)
	log.Debug("Parsing main block,", "s", s)
	o = append(o, parseBlock(0, s)...)
	o = append(o, "mainend(Конец)")
	return o
}

func parseVarBlock(s string) []string {
	o := []string{}
	rx := regexp.MustCompile(`begin.*end\.`)
	s = rx.ReplaceAllString(s, "")
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "var ")
	log.Debug("Parsing var block,", "s", s)
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
	log.Debug("Parsing code,", "code", s)
	if strings.HasPrefix(s, "var") {
		o = append(o, parseVarBlock(s)...)
	}
	o = append(o, parseMainBlock(s)...)

	return o
}
