package main

import "fmt"

// type SpoonOfJam interface {
// 	String() string
// }

// type Jam interface {
// 	GetOneSpoon() SpoonOfJam
// }

// type Bread struct {
// 	Var string
// }

// type StrawberryJam struct {
// }

// type SpoonOfStrawberryJam struct {
// }

// type OrangeJam struct {
// }
// type SpoonOfOrangeJam struct {
// }

// func (b *Bread) PutJam(jam Jam) {
// 	spoon := jam.GetOneSpoon()
// 	b.Var += spoon.String()
// }

// func (b *Bread) String() string {
// 	return "bread " + b.Var
// }

// func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
// 	return &SpoonOfStrawberryJam{}
// }

// func (s *SpoonOfStrawberryJam) String() string {
// 	return "+ strawberry"
// }

// func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
// 	return &SpoonOfOrangeJam{}
// }

// func (s *SpoonOfOrangeJam) String() string {
// 	return "+ Orange"
// }

/*
	1. 스트럭쳐형을 가져다 사용하는 func 는 상위에서 생성시 스트럭쳐.펑션을 가져다 사용할 수 있다.
	2. 행위로 구현된 인터페이스는 상속받은 타입에 여부와 상관없이 같은 func 명을 가진 메서드가 있다면 실행가능하다
	 - 입력받는 인자가 덕타입으로 실제, Move 든 Attac 이든 Action 메서드를 가지고있다면 실행되는듯하다.
	  >> 여기서 인터페이스와 덕타입에 대한 이해를 이런식으로 하면 될 것 같다????

	3. 구현시
		- 행위에 대한 데이터 적재 타입 1 >> 실제 비지니스에서 사용될 타입으로 주로 데이터 타입이 될거라 판단됨.
		- 공통으로 요청할 메서드타입의 인터페이스 1
		- 실제 비지니스 로직을 행할 메서드 1

		- 비지니스 로직에 추가되는, 개별 이벤트별 비지니스로직처리에 타입1, 액션1

	>> 앞으로 더 추가해 나갈 예정

*/

func main() {
	// bread := &Bread{}

	// //jam := &StrawberryJam{}
	// jam := &OrangeJam{}
	// bread.PutJam(jam)
	// fmt.Println(bread)

	game := &Game{}
	mClick := &Move{}
	game.Mclick(mClick)
	fmt.Println(game)

	aClick := &Attac{}
	game.Mclick(aClick)
	fmt.Println(game)

}

func (g *Game) Mclick(click Click) {

	ra := click.Action()

	g.ReAction += ra

}

type Game struct {
	ReAction string
}

type Click interface {
	Action() string
}

type Move struct{}

func (m *Move) Action() string {
	return "Move Action"
}

type Attac struct{}

func (a *Attac) Action() string {

	return "Attac Action"
}
