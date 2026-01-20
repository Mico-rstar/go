package animal

import "fmt"

// Animal 接口：所有动物必须能攻击其他动物
type Animal interface {
	Attack(target Animal)
	String() string // 用于打印动物名称
}

// Lion 类型
type Lion struct {
	Name string
}

func (l Lion) String() string {
	return "Lion(" + l.Name + ")"
}

func (l Lion) Attack(target Animal) {
	fmt.Printf("%s roars and pounces on %s!\n", l, target)
}

// Eagle 类型
type Eagle struct {
	Name string
}

func (e Eagle) String() string {
	return "Eagle(" + e.Name + ")"
}

func (e Eagle) Attack(target Animal) {
	fmt.Printf("%s dives from the sky and claws at %s!\n", e, target)
}

// Snake 类型
type Snake struct {
	Name string
}

func (s Snake) String() string {
	return "Snake(" + s.Name + ")"
}

func (s Snake) Attack(target Animal) {
	fmt.Printf("%s slithers silently and bites %s!\n", s, target)
}

// 模拟战斗场景的函数（体现多态）
func Battle(attacker, defender Animal) {
	fmt.Printf("%v 与 %v 发生战斗\n", attacker.String(), defender.String())
	attacker.Attack(defender)
}



func AnimalTest() {
	al := []Animal{
		Eagle{"老鹰"},
		Snake{"蛇"},
	}
		
	king := Lion{"狮子"}

	for _, animal := range al {
		Battle(king, animal)
	}
}