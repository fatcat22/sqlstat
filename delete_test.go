package sqlstat

import "testing"

func TestDelete(t *testing.T) {
	expectStat := "DELETE FROM aaa WHERE col1 = '2'"
	stat := DeleteFrom("aaa").
		Where(Column("col1").Equal(2)).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 没有 where
	expectStat = "DELETE FROM aaa"
	stat = DeleteFrom("aaa").
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}
}
