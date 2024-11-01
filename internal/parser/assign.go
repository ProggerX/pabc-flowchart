package parser

import (
	"fmt"
	"regexp"

	"github.com/charmbracelet/log"
)

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
