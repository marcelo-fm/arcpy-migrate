package gen

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	START  = "CREATE TABLE"
	END    = ";"
	SDE    = "SDE"
	INDENT = "    "
)

func NewAddFieldParams(inTable, fieldName string, fieldType FieldType) *AddFieldParams {
	fieldAlias := strings.ReplaceAll(fieldName, "_", " ")
	return &AddFieldParams{
		inTable:    inTable,
		fieldName:  fieldName,
		fieldType:  string(fieldType),
		fieldAlias: &fieldAlias,
	}
}

type AddFieldParams struct {
	fieldPrecision  *string
	fieldScale      *string
	fieldLength     *string
	fieldAlias      *string
	fieldIsNullable *string
	fieldIsRequired *string
	fieldDomain     *string
	inTable         string
	fieldName       string
	fieldType       string
}

func (c *AddFieldParams) SetFieldIsNullable(val bool) {
	null := "NULLABLE"
	if !val {
		null = "NON_" + null
	}
	c.fieldIsNullable = &null
}

func (c *AddFieldParams) SetFieldIsRequired(val bool) {
	required := "REQUIRED"
	if !val {
		required = "NON_" + required
	}
	c.fieldIsRequired = &required
}

func (p *AddFieldParams) Command() string {
	c := "arcpy.management.AddField("
	c = fmt.Sprintf("%sin_table=\"%s\",", c, p.inTable)
	c = fmt.Sprintf("%sfield_name=\"%s\",", c, p.fieldName)
	c = fmt.Sprintf("%sfield_type=\"%s\",", c, p.fieldType)
	if p.fieldPrecision != nil {
		c = fmt.Sprintf("%sfield_precision=\"%s\",", c, *p.fieldPrecision)
	}
	if p.fieldScale != nil {
		c = fmt.Sprintf("%sfield_scale=\"%s\",", c, *p.fieldScale)
	}
	if p.fieldLength != nil {
		c = fmt.Sprintf("%sfield_length=\"%s\",", c, *p.fieldLength)
	}
	if p.fieldAlias != nil {
		c = fmt.Sprintf("%sfield_alias=\"%s\".title(),", c, *p.fieldAlias)
	}
	if p.fieldIsNullable != nil {
		c = fmt.Sprintf("%sfield_is_nullable=\"%s\",", c, *p.fieldIsNullable)
	}
	if p.fieldDomain != nil {
		c = fmt.Sprintf("%sfield_domain=\"%s\",", c, *p.fieldDomain)
	}
	c = fmt.Sprintf("%s)", c)
	return c
}

func fieldLine(lin []byte, tbName string) []byte {
	hasArcGISCols := (bytes.HasPrefix(lin, []byte("objectid")) ||
		bytes.HasPrefix(lin, []byte("globalid")) ||
		bytes.HasPrefix(lin, []byte("created_date")) ||
		bytes.HasPrefix(lin, []byte("last_edited_date")))
	if hasArcGISCols {
		return []byte{}
	}
	lin = bytes.Trim(lin, " ")
	lin = bytes.TrimSuffix(lin, []byte(","))
	args := bytes.Split(lin, []byte(" "))

	fieldName := string(args[0])
	var fieldType FieldType
	ft := strings.Trim(string(args[1]), " ")
	switch ft {
	case "float":
		fieldType = FLOAT
	case "integer":
		fieldType = LONG
	case "varchar":
		fieldType = TEXT
	case "date":
		fieldType = DATEONLY
	case "timestamp":
		fieldType = DATE
	}

	params := NewAddFieldParams(tbName, fieldName, fieldType)
	params.SetFieldIsNullable(bytes.Contains(lin, []byte("NOT NULL")))
	cmd := params.Command()
	cmd = fmt.Sprintf("%s%s", INDENT, cmd)
	return []byte(cmd)
}

func parseTbName(lin []byte) string {
	lin = bytes.Trim(lin, " ")
	words := bytes.Split(lin, []byte(" "))
	return string(words[2])
}

func createTbLine(tbName string) []byte {
	cmd := fmt.Sprintf("%sarcpy.management.CreateTable(", INDENT)
	cmd = fmt.Sprintf("%sout_path=%s,", cmd, SDE)
	cmd = fmt.Sprintf("%sout_name=\"%s\",", cmd, tbName)
	fieldAlias := strings.ReplaceAll(tbName, "_", " ")
	cmd = fmt.Sprintf("%sout_alias=\"%s\".title())", cmd, fieldAlias)
	return []byte(cmd)
}

func deleteTbLine(tbName string) []byte {
	cmd := fmt.Sprintf("%sarcpy.management.Delete(", INDENT)
	cmd = fmt.Sprintf("%sin_data=\"%s\")", cmd, tbName)
	return []byte(cmd)
}

func mainFn() []byte {
	mainFn := `def main():
    usage_msg = """
Error: Must provide arg up or down
Usage:
	arcpy_migrate.py up (Create tables)
	arcpy_migrate down (Delete tables)
"""
    if len(sys.argv) != 2:
        print(usage_msg)
        return sys.exit()
    arg = sys.argv[1]
    if arg not in {"up", "down"}:
        print(usage_msg)
        sys.exit(0)
    arcpy.env.workspace = SDE
    if arg == "up":
        up()
    if arg == "down":
        down()

if __name__ == "__main__":
    main()
`
	return []byte(mainFn)
}

func Generate(file io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(file)
	var tbName string
	var cmdLine []byte
	script := [][]byte{
		[]byte("import arcpy\nimport sys\nimport os"),
		[]byte(
			fmt.Sprintf(
				"%s = os.getenv(\"ARCPYMIGRATE_%s\", \"\")",
				SDE,
				SDE,
			),
		),
	}
	upFn := [][]byte{[]byte("def up():")}
	downFn := [][]byte{[]byte("def down():")}
	parsing := false

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Bytes()
		if bytes.HasPrefix(line, []byte(START)) {
			parsing = true
			tbName = parseTbName(line)
			upFn = append(upFn, createTbLine(tbName))
			downFn = append(downFn, deleteTbLine(tbName))
			continue
		}
		if !parsing {
			continue
		}
		if !bytes.HasSuffix(line, []byte(",")) {
			continue
		}
		cmdLine = fieldLine(line, tbName)
		upFn = append(upFn, cmdLine)
		if bytes.HasSuffix(line, []byte(";")) {
			parsing = false
			tbName = ""
		}
	}
	script = append(script, upFn...)
	script = append(script, downFn...)
	script = append(script, mainFn())

	return bytes.Join(script, []byte("\n")), nil
}

type FieldType string

const (
	SHORT             FieldType = "SHORT"
	LONG              FieldType = "LONG"
	BIGINTEGER        FieldType = "BIGINTEGER"
	FLOAT             FieldType = "FLOAT"
	DOUBLE            FieldType = "DOUBLE"
	TEXT              FieldType = "TEXT"
	DATE              FieldType = "DATE"
	DATEHIGHPRECISION FieldType = "DATEHIGHPRECISION"
	DATEONLY          FieldType = "DATEONLY"
	TIMEONLY          FieldType = "TIMEONLY"
	TIMESTAMPOFFSET   FieldType = "TIMESTAMPOFFSET"
	BLOB              FieldType = "BLOB"
	GUID              FieldType = "GUID"
	RASTER            FieldType = "RASTER"
)
