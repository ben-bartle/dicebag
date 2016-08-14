package dicebag

import "testing"

func TestDie(t *testing.T) {
	d := newDie(6);

	// make sure the die is correct
	if (d.Pips() != 6) {
		t.Errorf("Wrong number of pips");
	}

	for i := 0; i < 100; i++ {
		d.Roll();
		if (d.Value() < 1 || d.Value() > 6) {
			t.Errorf("Bad dice value: %d", d.Value());
		}
	}

	d = newDie(6);

	d.fixedRoll(5);
	if (d.Value() != 5) {
		t.Errorf("Bad fixed roll (%s should be 5)", d.Value());
	}

}




func Test_makeDiceNdX(t *testing.T) {
    dice := makeDiceNdX(10,4);
    if len(dice) != 10 {
        t.Errorf("Not enough dice");
    }
    for _, die := range dice {
          if die.Pips() != 4 {
              t.Errorf("Die is not a 1d4");
              break;
          }
    }
}