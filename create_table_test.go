package sqlstat

import "testing"

func TestCreateTable(t *testing.T) {
	// 全部字段
	expectStat := `CREATE TABLE aaa(col1 VARCHAR(255) DEFAULT '',col2 VARCHAR(255) NOT NULL,col3 INTEGER NOT NULL DEFAULT 0,col4 VARCHAR(255) DEFAULT '',col5 BIGINT NOT NULL,col6 BIGINT NOT NULL AUTO_INCREMENT,INDEX (col1,col2),PRIMARY KEY (col4,col5,col6)) ENGINE=AAA  DEFAULT CHARSET=BBB AUTO_INCREMENT = 111`
	stat := CreateTable("aaa").
		AppendColumn(Column("col1").Attr().Type("VARCHAR(255)").DefaultEmptyString()).
		AppendColumn(Column("col2").Attr().Type("VARCHAR(255)").NotNULL()).
		AppendColumn(Column("col3").Attr().Type("INTEGER").NotNULL().DefaultInt(0)).
		AppendColumn(Column("col4").Attr().Type("VARCHAR(255)").DefaultEmptyString()).
		AppendColumn(Column("col5").Attr().Type("BIGINT").NotNULL()).
		AppendColumn(Column("col6").Attr().Type("BIGINT").NotNULL().AutoIncrement()).
		Index([]Column{Column("col1"), Column("col2")}).
		PrimaryKey([]Column{Column("col4"), Column("col5"), Column("col6")}).
		Engine("AAA").
		DefaultCharset("BBB").
		AutoIncrementBegin(111).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 没有设置 ENGINE, DEFAULT CHARSET, AUTO_INCREMENT
	expectStat = `CREATE TABLE aaa(col1 VARCHAR(255) DEFAULT '',col2 VARCHAR(255) NOT NULL,col3 INTEGER NOT NULL DEFAULT 0,col4 VARCHAR(255) DEFAULT '',col5 BIGINT NOT NULL,col6 BIGINT NOT NULL AUTO_INCREMENT,INDEX (col1,col2),PRIMARY KEY (col4,col5,col6)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT = 1`
	stat = CreateTable("aaa").
		AppendColumn(Column("col1").Attr().Type("VARCHAR(255)").DefaultEmptyString()).
		AppendColumn(Column("col2").Attr().Type("VARCHAR(255)").NotNULL()).
		AppendColumn(Column("col3").Attr().Type("INTEGER").NotNULL().DefaultInt(0)).
		AppendColumn(Column("col4").Attr().Type("VARCHAR(255)").DefaultEmptyString()).
		AppendColumn(Column("col5").Attr().Type("BIGINT").NotNULL()).
		AppendColumn(Column("col6").Attr().Type("BIGINT").NotNULL().AutoIncrement()).
		Index([]Column{Column("col1"), Column("col2")}).
		PrimaryKey([]Column{Column("col4"), Column("col5"), Column("col6")}).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 没有 INDEX
	expectStat = `CREATE TABLE aaa(col1 VARCHAR(255) DEFAULT '',col2 VARCHAR(255) NOT NULL,col3 INTEGER NOT NULL DEFAULT 0,col4 VARCHAR(255) DEFAULT '',col5 BIGINT NOT NULL,col6 BIGINT NOT NULL AUTO_INCREMENT,PRIMARY KEY (col4,col5,col6)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT = 1`
	stat = CreateTable("aaa").
		AppendColumn(Column("col1").Attr().Type("VARCHAR(255)").DefaultEmptyString()).
		AppendColumn(Column("col2").Attr().Type("VARCHAR(255)").NotNULL()).
		AppendColumn(Column("col3").Attr().Type("INTEGER").NotNULL().DefaultInt(0)).
		AppendColumn(Column("col4").Attr().Type("VARCHAR(255)").DefaultEmptyString()).
		AppendColumn(Column("col5").Attr().Type("BIGINT").NotNULL()).
		AppendColumn(Column("col6").Attr().Type("BIGINT").NotNULL().AutoIncrement()).
		PrimaryKey([]Column{Column("col4"), Column("col5"), Column("col6")}).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 没有 PRIMARY KEY
	expectStat = `CREATE TABLE aaa(col1 VARCHAR(255) DEFAULT '',col2 VARCHAR(255) NOT NULL,col3 INTEGER NOT NULL DEFAULT 0,col4 VARCHAR(255) DEFAULT '',col5 BIGINT NOT NULL,col6 BIGINT NOT NULL AUTO_INCREMENT,INDEX (col1,col2)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT = 1`
	stat = CreateTable("aaa").
		AppendColumn(Column("col1").Attr().Type("VARCHAR(255)").DefaultEmptyString()).
		AppendColumn(Column("col2").Attr().Type("VARCHAR(255)").NotNULL()).
		AppendColumn(Column("col3").Attr().Type("INTEGER").NotNULL().DefaultInt(0)).
		AppendColumn(Column("col4").Attr().Type("VARCHAR(255)").DefaultEmptyString()).
		AppendColumn(Column("col5").Attr().Type("BIGINT").NotNULL()).
		AppendColumn(Column("col6").Attr().Type("BIGINT").NotNULL().AutoIncrement()).
		Index([]Column{Column("col1"), Column("col2")}).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 只有列属性，没有 INDEX 和 PRIMARY KEY
	expectStat = `CREATE TABLE aaa(col1 VARCHAR(255) DEFAULT '',col2 INTEGER) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT = 1`
	stat = CreateTable("aaa").
		AppendColumn(Column("col1").Attr().Type("VARCHAR(255)").DefaultEmptyString()).
		AppendColumn(Column("col2").Attr().Type("INTEGER")).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}

	// 只有一列
	expectStat = `CREATE TABLE aaa(col1 VARCHAR(255) DEFAULT '') ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT = 1`
	stat = CreateTable("aaa").
		AppendColumn(Column("col1").Attr().Type("VARCHAR(255)").DefaultEmptyString()).
		String()
	if stat != expectStat {
		t.Fatalf("expect \n%q \nbut got \n%q\n", expectStat, stat)
	}
}
