package schema
 
import (
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent"
)
// Patient holds the schema definition for the Patient entity.
type Patient struct {
    ent.Schema
}
// Fields of the Patient.
func (Patient) Fields() []ent.Field {
    return []ent.Field{
        field.String("Patient_Name").NotEmpty(),
    }
 }
 //Edges of the Patient.
 func (Patient) Edges() []ent.Edge {
    return []ent.Edge{
		edge. To("Patient_DrugAllergy",DrugAllergy.Type).StorageKey(edge.Column("patient_id")),
 }
 }