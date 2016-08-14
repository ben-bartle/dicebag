package dicebag

import (
	"math/rand"
)

type die struct {
	pips int;
	value int;
}

func newDie (x int) *die {
	return &die{ pips: x };
}

func (d *die) Pips() int {
	return d.pips;
}

func (d *die) Value() int {
	return d.value;
}

func (d *die) Roll() {
	d.value = 1 + rand.Intn(d.pips);
}

func (d *die) fixedRoll(v int) {
	d.value = v;
}

func makeDiceNdX(n,x int) []die {
	var dice = make([]die, n);
	for d := 0; d < n; d++ {
		dice[d] = die{ pips: x};
	}
	return dice;
}