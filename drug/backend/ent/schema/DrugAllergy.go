package schema
 
import (
   "github.com/facebookincubator/ent/schema/edge"
   "github.com/facebookincubator/ent"
)
// DrugAllergy holds the schema definition for the DrugAllergy entity.
type DrugAllergy struct {
    ent.Schema
}
// Fields of the DrugAllergy.
 func (DrugAllergy) Fields() []ent.Field {
    return []ent.Field{
  
    }
 }

 //Edges of the DrugAllergy.
 func (DrugAllergy) Edges() []ent.Edge {
   return []ent.Edge{
      edge. From("doctor",Doctor.Type).Ref("Doctor_DrugAllergy").Unique(),
      edge. From("patient",Patient.Type).Ref("Patient_DrugAllergy").Unique(),
      edge. From("medicine",Medicine.Type).Ref("Medicine_DrugAllergy").Unique(),
      edge. From("manner",Manner.Type).Ref("Manner_DrugAllergy").Unique(),
   }
 }