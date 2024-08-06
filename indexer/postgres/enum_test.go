package postgres

import (
	"os"

	"cosmossdk.io/indexer/postgres/internal/testdata"
)

func ExampleCreateEnumTypeSql() {
	err := CreateEnumTypeSql(os.Stdout, "test", testdata.MyEnum)
	if err != nil {
		panic(err)
	}
	// Output:
	// CREATE TYPE "test_my_enum" AS ENUM ('a', 'b', 'c');
}
