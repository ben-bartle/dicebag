package dicebag


import (
	"testing"
)

func TestDiceBag(t *testing.T) {
	dicebag := DiceBag("4d6");

	if (dicebag.Count() != 4) {
		t.Errorf("Dicebag doesn't have the right number of dice (%d should be %d)",dicebag.Count(),4)
	}

	dicebag = DiceBag("10d6 4d5 3d2");

	if (dicebag.Count() != 17) {
		t.Errorf("Dicebag of 10d6, 4d5, and 3d2 should have 17 dice")
	}
}

func TestDiceBagNdX(t *testing.T) {
	n := 7;
	dicebag := DiceBagNdX(n,6);

	if (dicebag.Count() != n) {
		t.Errorf("Dicebag doesn't have the right number of dice (%d should be %d)",dicebag.Count(),n)
	}
}

func TestRollValues(t *testing.T) {
	dicebag := DiceBag("4d6");
	dicebag.Roll();

	values := dicebag.Values();
	if len(values) != 4 {
		t.Errorf("4d6 should have 4 values");
	}

	for i,value := range values {
		if (value < 1 || value > 6){
			t.Errorf("Bad dice roll #%d: %d", i, value);
		}
	}
}

func TestFixedRoll(t *testing.T) {
	dicebag := DiceBag("4d6");
	dicebag.fixedRoll([]int{5,3,1,6});

	values := dicebag.Values();
	if len(values) != 4 {
		t.Errorf("4d6 should have 4 values");
	}

	if values[0] != 5 {
		t.Errorf("value #1 should be 5");
	}

	if values[1] != 3 {
		t.Errorf("value #2 should be 3");
	}

	if values[2] != 1 {
		t.Errorf("value #3 should be 1");
	}

	if values[3] != 6 {
		t.Errorf("value #4 should be 6");
	}
}

func TestSum(t *testing.T) {
	dicebag := DiceBag("4d6");
	dicebag.fixedRoll([]int{5,3,4,1});

	if (dicebag.Sum() != 5+3+4+1) {
		t.Errorf("Bad sum");
	}
}