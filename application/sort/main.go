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
	students := []student{
		student{"Ajay1", 2, "B", 3},
		student{"Ajay2", 2, "B", 1},
		student{"Ajay3", 2, "A", 2},
		student{"Ajay4", 2, "A", 1},
		student{"Ajay5", 1, "B", 2},
		student{"Ajay6", 1, "B", 1},
		student{"Ajay7", 1, "A", 2},
		student{"Ajay8", 1, "C", 1},
	}
	fmt.Println(students)

	sort.Sort(customSort(students))

	fmt.Println(students)
}
