// Code generated by entc, DO NOT EDIT.

package datafile

import (
	"Backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HasHeader applies the HasEdge predicate on the "header" edge.
func HasHeader() predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HeaderTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, HeaderTable, HeaderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHeaderWith applies the HasEdge predicate on the "header" edge with a given conditions (other predicates).
func HasHeaderWith(preds ...predicate.Header) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HeaderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, HeaderTable, HeaderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGames applies the HasEdge predicate on the "games" edge.
func HasGames() predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GamesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GamesTable, GamesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGamesWith applies the HasEdge predicate on the "games" edge with a given conditions (other predicates).
func HasGamesWith(preds ...predicate.Game) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GamesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GamesTable, GamesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Datafile) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Datafile) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
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
func Not(p predicate.Datafile) predicate.Datafile {
	return predicate.Datafile(func(s *sql.Selector) {
		p(s.Not())
	})
}