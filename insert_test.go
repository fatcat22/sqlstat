package sqlstat

import "testing"

func TestInsert(t *testing.T) {
	expectStat := "INSERT INTO aaa(col1,col2,col3) VALUES ('str','1','33')"
	stat := InsertInto("aaa").
		With(Column("col1"), "str").
		With(Column("col2"), 1).
		With(Column("col3"), int64(33)).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	expectStat = "INSERT INTO aaa(col1) VALUES ('str')"
	stat = InsertInto("aaa").
		With(Column("col1"), "str").
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}
	expectStat = ""
	stat = InsertInto("aaa").
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	expectStat = "INSERT INTO aaa(col1,col2) VALUES ('1','yourname') ON DUPLICATE KEY UPDATE col3 = '33',col4 = 'xxx'"
	stat = InsertInto("aaa").
		With(Column("col1"), 1).
		With(Column("col2"), "yourname").
		UpdateWith(Column("col3"), 33).
		UpdateWith(Column("col4"), "xxx").
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}
}
