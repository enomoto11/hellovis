// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "last_name", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString},
		{Name: "grade", Type: field.TypeInt},
		{Name: "is_high_school", Type: field.TypeBool, Default: true},
		{Name: "manavis_code", Type: field.TypeString, Unique: true},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
	}
	// StudentCheckinsColumns holds the columns for the "student_checkins" table.
	StudentCheckinsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "checkin_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "student_id", Type: field.TypeUUID},
	}
	// StudentCheckinsTable holds the schema information for the "student_checkins" table.
	StudentCheckinsTable = &schema.Table{
		Name:       "student_checkins",
		Columns:    StudentCheckinsColumns,
		PrimaryKey: []*schema.Column{StudentCheckinsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "student_checkins_students_checkins",
				Columns:    []*schema.Column{StudentCheckinsColumns[4]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// StudentCheckoutsColumns holds the columns for the "student_checkouts" table.
	StudentCheckoutsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "checkout_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime(6)"}},
		{Name: "student_id", Type: field.TypeUUID},
	}
	// StudentCheckoutsTable holds the schema information for the "student_checkouts" table.
	StudentCheckoutsTable = &schema.Table{
		Name:       "student_checkouts",
		Columns:    StudentCheckoutsColumns,
		PrimaryKey: []*schema.Column{StudentCheckoutsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "student_checkouts_students_checkouts",
				Columns:    []*schema.Column{StudentCheckoutsColumns[4]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StudentsTable,
		StudentCheckinsTable,
		StudentCheckoutsTable,
	}
)

func init() {
	StudentCheckinsTable.ForeignKeys[0].RefTable = StudentsTable
	StudentCheckoutsTable.ForeignKeys[0].RefTable = StudentsTable
}
