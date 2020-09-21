package sqlstat

import "testing"

func TestSelect(t *testing.T) {
	// 全部
	expectStat := `SELECT col1,col2,col3 FROM aaa WHERE (col1 = 'x')AND(col2 = '2') ORDER BY col1 DESC,col2 DESC LIMIT 3, 10`
	stat := Select(Column("col1"), Column("col2"), Column("col3")).
		From("aaa").
		Where(Logic(Column("col1").Equal("x")).AND(Column("col2").Equal(2))).
		OrderBy(Column("col1"), DESC).
		OrderBy(Column("col2"), DESC).
		Limit(3, 10).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 没有 where
	expectStat = `SELECT col1,col2,col3 FROM aaa ORDER BY col1 DESC,col2 DESC LIMIT 3, 10`
	stat = Select(Column("col1"), Column("col2"), Column("col3")).
		From("aaa").
		OrderBy(Column("col1"), DESC).
		OrderBy(Column("col2"), DESC).
		Limit(3, 10).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 没有 order by
	expectStat = `SELECT col1,col2,col3 FROM aaa WHERE (col1 = 'x')AND(col2 = '2') LIMIT 3, 10`
	stat = Select(Column("col1"), Column("col2"), Column("col3")).
		From("aaa").
		Where(Logic(Column("col1").Equal("x")).AND(Column("col2").Equal(2))).
		Limit(3, 10).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 没有 limit offset
	expectStat = `SELECT col1,col2,col3 FROM aaa WHERE (col1 = 'x')AND(col2 = '2') ORDER BY col1 DESC,col2 DESC LIMIT 10`
	stat = Select(Column("col1"), Column("col2"), Column("col3")).
		From("aaa").
		Where(Logic(Column("col1").Equal("x")).AND(Column("col2").Equal(2))).
		OrderBy(Column("col1"), DESC).
		OrderBy(Column("col2"), DESC).
		LimitCount(10).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 没有 limit
	expectStat = `SELECT col1,col2,col3 FROM aaa WHERE (col1 = 'x')AND(col2 = '2') ORDER BY col1 DESC,col2 DESC`
	stat = Select(Column("col1"), Column("col2"), Column("col3")).
		From("aaa").
		Where(Logic(Column("col1").Equal("x")).AND(Column("col2").Equal(2))).
		OrderBy(Column("col1"), DESC).
		OrderBy(Column("col2"), DESC).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 所有条件都没有，只有最基本的
	expectStat = `SELECT col1,col2,col3 FROM aaa`
	stat = Select(Column("col1"), Column("col2"), Column("col3")).
		From("aaa").
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 只有一列
	expectStat = `SELECT col1 FROM aaa WHERE col1 = 'x' ORDER BY col2 DESC LIMIT 3, 10`
	stat = Select(Column("col1")).
		From("aaa").
		Where(Column("col1").Equal("x")).
		OrderBy(Column("col2"), DESC).
		Limit(3, 10).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// select count
	expectStat = `SELECT COUNT(col1,col2,col3) FROM aaa WHERE (col1 = 'x')AND(col2 = '2') ORDER BY col1 DESC,col2 DESC LIMIT 3, 10`
	stat = SelectCount(Column("col1"), Column("col2"), Column("col3")).
		From("aaa").
		Where(Logic(Column("col1").Equal("x")).AND(Column("col2").Equal(2))).
		OrderBy(Column("col1"), DESC).
		OrderBy(Column("col2"), DESC).
		Limit(3, 10).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// select count *
	expectStat = `SELECT COUNT(*) FROM aaa WHERE (col1 = 'x')AND(col2 = '2') ORDER BY col1 DESC,col2 DESC LIMIT 3, 10`
	stat = SelectCount(Column("*")).
		From("aaa").
		Where(Logic(Column("col1").Equal("x")).AND(Column("col2").Equal(2))).
		OrderBy(Column("col1"), DESC).
		OrderBy(Column("col2"), DESC).
		Limit(3, 10).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}
}
