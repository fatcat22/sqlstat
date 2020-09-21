package sqlstat

import "testing"

func TestColumnAttr(t *testing.T) {
	expectStat := "aaa INT NOT NULL DEFAULT '' AUTO_INCREMENT"
	stat := NewColumn("aaa").Attr().Type("INT").NotNULL().DefaultEmptyString().AutoIncrement().String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	expectStat = "aaa INT DEFAULT '' AUTO_INCREMENT"
	stat = NewColumn("aaa").Attr().Type("INT").DefaultEmptyString().AutoIncrement().String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	expectStat = "aaa INT NOT NULL DEFAULT 33 AUTO_INCREMENT"
	stat = NewColumn("aaa").Attr().Type("INT").NotNULL().DefaultInt(33).AutoIncrement().String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	expectStat = "aaa INT NOT NULL AUTO_INCREMENT"
	stat = NewColumn("aaa").Attr().Type("INT").NotNULL().AutoIncrement().String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	expectStat = "aaa INT NOT NULL DEFAULT ''"
	stat = NewColumn("aaa").Attr().Type("INT").NotNULL().DefaultEmptyString().String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	expectStat = "aaa INT"
	stat = NewColumn("aaa").Attr().Type("INT").String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}
}
