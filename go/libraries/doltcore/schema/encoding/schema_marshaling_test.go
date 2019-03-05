package encoding

import (
	"github.com/attic-labs/noms/go/spec"
	"github.com/attic-labs/noms/go/types"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/schema"
	"reflect"
	"testing"
)

func createTestSchema() schema.Schema {
	columns := []schema.Column{
		schema.NewColumn("id", 4, types.UUIDKind, true, schema.NotNullConstraint{}),
		schema.NewColumn("first", 1, types.StringKind, false),
		schema.NewColumn("last", 2, types.StringKind, false, schema.NotNullConstraint{}),
		schema.NewColumn("age", 3, types.UintKind, false),
	}

	colColl, _ := schema.NewColCollection(columns...)
	sch := schema.SchemaFromCols(colColl)

	return sch
}

func TestNomsMarshalling(t *testing.T) {
	tSchema := createTestSchema()
	dbSpec, err := spec.ForDatabase("mem")

	if err != nil {
		t.Fatal("Could not create in mem noms db.")
	}

	db := dbSpec.GetDatabase()
	val, err := MarshalAsNomsValue(db, tSchema)

	if err != nil {
		t.Fatal("Failed to marshal Schema as a types.Value.")
	}

	unMarshalled, err := UnmarshalNomsValue(val)

	if err != nil {
		t.Fatal("Failed to unmarshal types.Value as Schema")
	}

	if !reflect.DeepEqual(tSchema, unMarshalled) {
		t.Error("Value different after marshalling and unmarshalling.")
	}

	jsonStr, err := MarshalAsJson(tSchema)

	if err != nil {
		t.Fatal("Failed to marshal Schema as a types.Value.")
	}

	jsonUnmarshalled, err := UnmarshalJson(jsonStr)

	if err != nil {
		t.Fatal("Failed to unmarshal types.Value as Schema")
	}

	if !reflect.DeepEqual(tSchema, jsonUnmarshalled) {
		t.Error("Value different after marshalling and unmarshalling.")
	}
}