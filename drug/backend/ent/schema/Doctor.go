package schema
 
import (
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent"
)

// User holds the schema definition for the Doctor entity.
type Doctor struct {
    ent.Schema
}

// Fields of the Doctor.
func (Doctor) Fields() []ent.Field {
    return []ent.Field{
        field.String("Doctor_Email").NotEmpty(),
        field.String("Doctor_Password").NotEmpty(),
        field.String("Doctor_Name").NotEmpty(),
        field.String("Doctor_Tel").NotEmpty(),
    }
 }

 
 // Edges of the Doctor.
 func (Doctor) Edges() []ent.Edge {
    return []ent.Edge{
        edge. To("Doctor_DrugAllergy",DrugAllergy.Type).StorageKey(edge.Column("doctor_id")),
    }
 }

 
