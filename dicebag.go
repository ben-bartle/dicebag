package dicebag

import (
	"regexp"
	"strconv"
)

type dicebag struct {
	dice []die;
}

//die matching regex
var re = regexp.MustCompile(`([1-9]\d*)d([1-9]\d*)`);

func DiceBag(descriptor string) *dicebag {
	diceSets := re.FindAllStringSubmatch(descriptor, -1);

	var dice = make([]die,0);
	if len(diceSets) > 0 {
		for _, diceSet := range diceSets {
			count, _ := strconv.Atoi(diceSet[1]);
			die, _ := strconv.Atoi(diceSet[2]);
			setDice := makeDiceNdX(count,die);
			dice = append(dice, setDice...);
		}

	}
	return &dicebag{dice: dice};
}

func DiceBagNdX(n, x int) *dicebag {
	dice := makeDiceNdX(n,x);
	return &dicebag{ dice: dice }
}


func (db *dicebag) Count() int {
	return len(db.dice);
}

func (db *dicebag) Roll() *dicebag {
	for i := range db.dice {
		db.dice[i].Roll();
	}
	return db;
}

func (db *dicebag) Values() []int {
	result := make([]int, db.Count());
	for i := range db.dice {
		result[i] = db.dice[i].Value();
	}
	return result;
}

func (db *dicebag) Sum() int {
	result := 0;
	for i := range db.dice {
		result += db.dice[i].Value();
	}
	return result;

}

func (db *dicebag) fixedRoll(rolls []int) {
	for i,roll := range rolls {
		db.dice[i].fixedRoll(roll);
	}
}