package sqlstat

import "testing"

func TestCompare(t *testing.T) {
	// string equal
	expectStat := "col1 = 'hello'"
	stat := Column("col1").Equal("hello").String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// int equal
	expectStat = "col1 = '10'"
	stat = Column("col1").Equal(10).String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// string not equal
	expectStat = "col1 != 'aaa'"
	stat = Column("col1").NotEqual("aaa").String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// int not equal
	expectStat = "col1 != '10'"
	stat = Column("col1").NotEqual(10).String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// int <=
	expectStat = "col1 <= '10'"
	stat = Column("col1").LessOrEqual(10).String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// int >
	expectStat = "col1 > '10'"
	stat = Column("col1").Great(10).String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}
}

func TestLogic(t *testing.T) {
	expectStat := `((col1 = 'hello')AND(col2 = '2'))OR(col3 = 'abc')`
	stat := Logic(Logic(Column("col1").Equal("hello")).AND(Column("col2").Equal(2))).
		OR(Column("col3").Equal("abc")).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}
}
