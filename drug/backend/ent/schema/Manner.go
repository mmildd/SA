package schema
 
import (
   "github.com/facebookincubator/ent/schema/edge"
   "github.com/facebookincubator/ent/schema/field"
   "github.com/facebookincubator/ent"
)
// Manner holds the schema definition for the Manner entity.
type Manner struct {
    ent.Schema
}
// Fields of the Manner.
 func (Manner) Fields() []ent.Field {
    return []ent.Field{
        field.String("Manner_Name").NotEmpty(),
    }
 }
 //Edges of the Manner.
 func (Manner) Edges() []ent.Edge {
   return []ent.Edge{
		edge. To("Manner_DrugAllergy",DrugAllergy.Type).StorageKey(edge.Column("manner_id")),
 }
 }