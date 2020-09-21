package sqlstat

import "testing"

func TestUpdate(t *testing.T) {
	expectStat := `UPDATE aaa SET col1 = 'abc',col2 = '2',col3 = 'mmm' WHERE col2 = 'dd'`
	stat := Update("aaa").
		With(Column("col1"), "abc").
		With(Column("col2"), 2).
		With(Column("col3"), "mmm").
		Where(Column("col2").Equal("dd")).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 没有 where
	expectStat = `UPDATE aaa SET col1 = 'abc',col2 = '2',col3 = 'mmm'`
	stat = Update("aaa").
		With(Column("col1"), "abc").
		With(Column("col2"), 2).
		With(Column("col3"), "mmm").
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 只有一列
	expectStat = `UPDATE aaa SET col1 = 'abc' WHERE col2 = 'dd'`
	stat = Update("aaa").
		With(Column("col1"), "abc").
		Where(Column("col2").Equal("dd")).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}
}
