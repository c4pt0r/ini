ini
===

Parse INI file like "flag" in Go


Example:

```
conf := NewConf("test.ini")
v1 := conf.String("section_2", "field1", "default")
v2 := conf.String("section_1", "field1", "default")
v3 := conf.String("section_3", "field1", "default")
v4 := conf.String("section_4", "field1", "default")
v5 := conf.Int(GLOBAL_SECTION, "global_1", 0)

conf.Parse()

```
