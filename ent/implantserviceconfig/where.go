// Code generated by entc, DO NOT EDIT.

package implantserviceconfig

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kcarretto/realm/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// ExecutablePath applies equality check predicate on the "executablePath" field. It's identical to ExecutablePathEQ.
func ExecutablePath(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExecutablePath), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ImplantServiceConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ImplantServiceConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ImplantServiceConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ImplantServiceConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// ExecutablePathEQ applies the EQ predicate on the "executablePath" field.
func ExecutablePathEQ(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExecutablePath), v))
	})
}

// ExecutablePathNEQ applies the NEQ predicate on the "executablePath" field.
func ExecutablePathNEQ(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExecutablePath), v))
	})
}

// ExecutablePathIn applies the In predicate on the "executablePath" field.
func ExecutablePathIn(vs ...string) predicate.ImplantServiceConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExecutablePath), v...))
	})
}

// ExecutablePathNotIn applies the NotIn predicate on the "executablePath" field.
func ExecutablePathNotIn(vs ...string) predicate.ImplantServiceConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExecutablePath), v...))
	})
}

// ExecutablePathGT applies the GT predicate on the "executablePath" field.
func ExecutablePathGT(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExecutablePath), v))
	})
}

// ExecutablePathGTE applies the GTE predicate on the "executablePath" field.
func ExecutablePathGTE(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExecutablePath), v))
	})
}

// ExecutablePathLT applies the LT predicate on the "executablePath" field.
func ExecutablePathLT(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExecutablePath), v))
	})
}

// ExecutablePathLTE applies the LTE predicate on the "executablePath" field.
func ExecutablePathLTE(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExecutablePath), v))
	})
}

// ExecutablePathContains applies the Contains predicate on the "executablePath" field.
func ExecutablePathContains(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExecutablePath), v))
	})
}

// ExecutablePathHasPrefix applies the HasPrefix predicate on the "executablePath" field.
func ExecutablePathHasPrefix(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExecutablePath), v))
	})
}

// ExecutablePathHasSuffix applies the HasSuffix predicate on the "executablePath" field.
func ExecutablePathHasSuffix(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExecutablePath), v))
	})
}

// ExecutablePathEqualFold applies the EqualFold predicate on the "executablePath" field.
func ExecutablePathEqualFold(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExecutablePath), v))
	})
}

// ExecutablePathContainsFold applies the ContainsFold predicate on the "executablePath" field.
func ExecutablePathContainsFold(v string) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExecutablePath), v))
	})
}

// HasImplantConfigs applies the HasEdge predicate on the "implantConfigs" edge.
func HasImplantConfigs() predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ImplantConfigsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ImplantConfigsTable, ImplantConfigsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImplantConfigsWith applies the HasEdge predicate on the "implantConfigs" edge with a given conditions (other predicates).
func HasImplantConfigsWith(preds ...predicate.ImplantConfig) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ImplantConfigsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ImplantConfigsTable, ImplantConfigsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ImplantServiceConfig) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ImplantServiceConfig) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ImplantServiceConfig) predicate.ImplantServiceConfig {
	return predicate.ImplantServiceConfig(func(s *sql.Selector) {
		p(s.Not())
	})
}