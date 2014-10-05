package ini

import (
	"log"
	"testing"
)

func TestVar(t *testing.T) {
	conf := NewConf("test.ini")
	v1 := conf.String("section_2", "field1", "default")
	v2 := conf.String("section_1", "field1", "default")
	v3 := conf.String("section_3", "field1", "default")
	v4 := conf.String("section_4", "field1", "default")
	v5 := conf.Int(GLOBAL_SECTION, "global_2", 10)

	conf.Parse()

	log.Println(*v1)
	log.Println(*v2)
	log.Println(*v3)
	log.Println(*v4)
	log.Println(*v5)

}
