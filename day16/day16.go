package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	low  int
	high int
}

type Field struct {
	name   string
	ranges []*Range
}

type Ticket struct {
	fields []*Field
}

func (r *Range) IsIn(value int) bool {
	return value >= r.low && value <= r.high
}

func (c *Field) IsValid(value int) bool {
	for _, classRange := range c.ranges {
		if classRange.IsIn(value) {
			return true
		}
	}
	return false
}

func (t *Ticket) IsValid(value int) bool {
	for _, field := range t.fields {
		if field.IsValid(value) {
			return true
		}
	}
	return false
}

func ParseField(input string) *Field {
	split := strings.Split(input, ": ")
	rangeStr := strings.Split(split[1], " or ")
	var ranges []*Range
	for _, str := range rangeStr {
		low := 0
		high := 0
		fmt.Sscanf(str, "%d-%d", &low, &high)
		ranges = append(ranges, &Range{low: low, high: high})
	}
	return &Field{name: split[0], ranges: ranges}
}

func ParseInput(input []string) (*Ticket, string, []string) {
	var fields []*Field
	for _, line := range input {
		if len(line) == 0 {
			break
		}
		fields = append(fields, ParseField(line))
	}
	myticket := input[len(fields)+2 : len(fields)+3][0]
	scannedTickets := input[len(fields)+5:]
	return &Ticket{fields: fields}, myticket, scannedTickets
}

func CalcErrorRate(input []string) int {
	ticket, _, scannedTickets := ParseInput(input)
	errorRate := 0
	for _, scan := range scannedTickets {
		ids := strings.Split(scan, ",")
		for _, id := range ids {
			val, _ := strconv.Atoi(id)
			if !ticket.IsValid(val) {
				errorRate += val
			}
		}
	}
	return errorRate
}

func (t *Ticket) IsValidTicket(ticket string) bool {
	ids := strings.Split(ticket, ",")
	for _, id := range ids {
		val, _ := strconv.Atoi(id)
		if !t.IsValid(val) {
			return false
		}
	}
	return true
}

func (t *Ticket) GetValidTickets(input []string) []string {
	var validTickets []string
	for _, scan := range input {
		if t.IsValidTicket(scan) {
			validTickets = append(validTickets, scan)
		}
	}
	return validTickets
}

func CalcField(fieldMap map[string]*Field, values []string) []string {

	for _, valStr := range values {
		remain := make(map[string]*Field)
		val, _ := strconv.Atoi(valStr)
		for _, field := range fieldMap {
			if field.IsValid(val) {
				remain[field.name] = field
			}
		}
		fieldMap = remain
	}

	rows := []string{}
	for key := range fieldMap {
		rows = append(rows, key)
	}
	return rows
}

func CalcOrder(ticket *Ticket, scannedTickets []string) map[int][]string {
	rows := make(map[int][]string)
	for i := 0; i < len(ticket.fields); i++ {
		fieldMap := make(map[string]*Field)
		for _, field := range ticket.fields {
			fieldMap[field.name] = field
		}
		var col []string
		for _, tick := range scannedTickets {
			parts := strings.Split(tick, ",")
			col = append(col, parts[i])
		}
		rows[i] = CalcField(fieldMap, col)
	}
	for !done(rows) {
		for _, row := range rows {
			if len(row) == 1 {
				// fmt.Println(row[0])
				toRemove := row[0]
				for j, alterRows := range rows {
					if len(alterRows) > 1 {
						newRow := []string{}
						for _, id := range alterRows {
							if id != toRemove {
								newRow = append(newRow, id)
							}
						}
						rows[j] = newRow
					}
				}
			}
		}
	}
	return rows
}

func done(rows map[int][]string) bool {
	for _, row := range rows {
		if len(row) > 1 {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("Part 1:", CalcErrorRate(input))
	ticket, myTicket, scannedTickets := ParseInput(input)
	scannedTickets = ticket.GetValidTickets(scannedTickets)
	rows := CalcOrder(ticket, scannedTickets)
	// fmt.Println(rows)
	var myValues []int
	for _, valStr := range strings.Split(myTicket, ",") {
		val, _ := strconv.Atoi(valStr)
		myValues = append(myValues, val)
	}
	// fmt.Println(myValues)
	total := 1
	for i, fieldName := range rows {
		if strings.HasPrefix(fieldName[0], "departure") {
			// fmt.Println(fieldName, myValues[i])
			total *= myValues[i]
		}
	}
	fmt.Println("Part 2:", total)
}
