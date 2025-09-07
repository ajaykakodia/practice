package main

import (
	"fmt"
	"sort"
)

type student struct {
	name       string
	class      int
	section    string
	rollNumber int
}

func (s student) String() string {
	return fmt.Sprintf("Class:%d, Section:%s, RollNumber:%d, Name: %s\n", s.class, s.section, s.rollNumber, s.name)
}

type customSort []student

func (cs customSort) Len() int      { return len(cs) }
func (cs customSort) Swap(i, j int) { cs[i], cs[j] = cs[j], cs[i] }
func (cs customSort) Less(i, j int) bool {
	if cs[i].class == cs[j].class {
		if cs[i].section == cs[j].section {
			return cs[i].rollNumber < cs[j].rollNumber
		}
		return cs[i].section < cs[j].section
	}
	return cs[i].class < cs[j].class
}

func main() {
	stds := []student{
		{"Ajay1", 2, "B", 3},
		{"Ajay2", 2, "B", 1},
		{"Ajay3", 2, "A", 2},
		{"Ajay4", 2, "A", 1},
		{"Ajay5", 1, "B", 2},
		{"Ajay6", 1, "B", 1},
		{"Ajay7", 1, "A", 2},
		{"Ajay8", 1, "C", 1},
	}
	fmt.Println(stds)

	sort.Sort(customSort(stds))

	fmt.Println(stds)

	students := []student{
		{"Ajay1", 2, "B", 3},
		{"Ajay2", 2, "B", 1},
		{"Ajay3", 2, "A", 2},
		{"Ajay4", 2, "A", 1},
		{"Ajay5", 1, "B", 2},
		{"Ajay6", 1, "B", 1},
		{"Ajay7", 1, "A", 2},
		{"Ajay8", 1, "C", 1},
	}

	sort.Slice(students, func(i, j int) bool {
		if students[i].class == students[j].class {
			if students[i].section == students[j].section {
				return students[i].rollNumber < students[j].rollNumber
			}
			return students[i].section < students[j].section
		}
		return students[i].class < students[j].class
	})

	fmt.Println(students)
}
