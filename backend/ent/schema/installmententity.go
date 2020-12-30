package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
)

// InstallmentEntity holds the schema definition for the InstallmentEntity entity.
type InstallmentEntity struct {
	ent.Schema
}

// Fields of the InstallmentEntity.
func (InstallmentEntity) Fields() []ent.Field {
	return []ent.Field{
		field.Float("down payment"),
		field.Float("installment"),
		field.Int("numberOfPeriods"),
	}
}

// Edges of the InstallmentEntity.
func (InstallmentEntity) Edges() []ent.Edge {
	return nil
}
