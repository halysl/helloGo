package basis

import "fmt"

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color Color
}

type BoxList []Box //a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func use_struct() {
	type person struct {
		name string
		age int
	}
	type student struct {
		person
		grade string
	}
	type teacher struct {
		person
		salary float64
	}
	type cool_man struct {
		person
		skill []string
	}

	p1 := person{"Ash", 17}
	p2 := person{age:25, name:"Bob"}
	s1 := student{person{"Bob", 12}, "2"}
	s2 := student{person: person{"Light", 18}, grade: "1"}
	t := teacher{person{"Lily", 21}, 3000.0}
	cm := cool_man{person{"qwer", 33}, []string{"aaa", "sss"}}
	cm.skill = append(cm.skill, "bbb")
	fmt.Printf("\np1.name=%s, p1.age=%d\n", p1.name, p1.age)
	fmt.Printf("p2.name=%s, p2.age=%d\n", p2.name, p2.age)
	fmt.Printf("s1.name=%s, s1.age=%d, s1.grade=%s\n", s1.name, s1.age, s1.grade)
	fmt.Printf("s2.name=%s, s2.age=%d, s2.grade=%s\n", s2.person.name, s2.person.age, s2.grade)
	fmt.Printf("t.name=%s, t.age=%d, t.salary=%f\n", t.name, t.age, t.salary)
	fmt.Printf("cm.name=%s, cm.age=%d, cm.skills=%s\n", cm.name, cm.age, cm.skill)
}

func Run_struct() {
	use_struct()
	use_method_inherit()
	use_method_rewrite()
}

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段
	school string
}

type Employee struct {
	Human //匿名字段
	company string
}

//在human上面定义了一个method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s, and my company is %s\n", e.name, e.phone, e.company)
}

func use_method_inherit() {
	h := Human{"light", 12, "1234567890"}
	s := Student{h, "high school"}

	h.SayHi()
	s.SayHi()

}

func use_method_rewrite() {
	h := Human{"light", 12, "1234567890"}
	e := Employee{h, "google"}
	h.SayHi()
	e.SayHi()
}