package sqlstat

import "testing"

func TestBitwise(t *testing.T) {
	expectStat := `(col1&3) = '3'`
	stat := Column("col1").
		BitAnd(3).
		Equal(3).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}
}
