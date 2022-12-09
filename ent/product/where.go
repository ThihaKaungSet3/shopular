// Code generated by ent, DO NOT EDIT.

package product

import (
	"shopular/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// ImageUrl applies equality check predicate on the "imageUrl" field. It's identical to ImageUrlEQ.
func ImageUrl(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageUrl), v))
	})
}

// Summary applies equality check predicate on the "summary" field. It's identical to SummaryEQ.
func Summary(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSummary), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// ImageUrlEQ applies the EQ predicate on the "imageUrl" field.
func ImageUrlEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageUrl), v))
	})
}

// ImageUrlNEQ applies the NEQ predicate on the "imageUrl" field.
func ImageUrlNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImageUrl), v))
	})
}

// ImageUrlIn applies the In predicate on the "imageUrl" field.
func ImageUrlIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldImageUrl), v...))
	})
}

// ImageUrlNotIn applies the NotIn predicate on the "imageUrl" field.
func ImageUrlNotIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldImageUrl), v...))
	})
}

// ImageUrlGT applies the GT predicate on the "imageUrl" field.
func ImageUrlGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImageUrl), v))
	})
}

// ImageUrlGTE applies the GTE predicate on the "imageUrl" field.
func ImageUrlGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImageUrl), v))
	})
}

// ImageUrlLT applies the LT predicate on the "imageUrl" field.
func ImageUrlLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImageUrl), v))
	})
}

// ImageUrlLTE applies the LTE predicate on the "imageUrl" field.
func ImageUrlLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImageUrl), v))
	})
}

// ImageUrlContains applies the Contains predicate on the "imageUrl" field.
func ImageUrlContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImageUrl), v))
	})
}

// ImageUrlHasPrefix applies the HasPrefix predicate on the "imageUrl" field.
func ImageUrlHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImageUrl), v))
	})
}

// ImageUrlHasSuffix applies the HasSuffix predicate on the "imageUrl" field.
func ImageUrlHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImageUrl), v))
	})
}

// ImageUrlEqualFold applies the EqualFold predicate on the "imageUrl" field.
func ImageUrlEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImageUrl), v))
	})
}

// ImageUrlContainsFold applies the ContainsFold predicate on the "imageUrl" field.
func ImageUrlContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImageUrl), v))
	})
}

// SummaryEQ applies the EQ predicate on the "summary" field.
func SummaryEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSummary), v))
	})
}

// SummaryNEQ applies the NEQ predicate on the "summary" field.
func SummaryNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSummary), v))
	})
}

// SummaryIn applies the In predicate on the "summary" field.
func SummaryIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSummary), v...))
	})
}

// SummaryNotIn applies the NotIn predicate on the "summary" field.
func SummaryNotIn(vs ...string) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSummary), v...))
	})
}

// SummaryGT applies the GT predicate on the "summary" field.
func SummaryGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSummary), v))
	})
}

// SummaryGTE applies the GTE predicate on the "summary" field.
func SummaryGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSummary), v))
	})
}

// SummaryLT applies the LT predicate on the "summary" field.
func SummaryLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSummary), v))
	})
}

// SummaryLTE applies the LTE predicate on the "summary" field.
func SummaryLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSummary), v))
	})
}

// SummaryContains applies the Contains predicate on the "summary" field.
func SummaryContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSummary), v))
	})
}

// SummaryHasPrefix applies the HasPrefix predicate on the "summary" field.
func SummaryHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSummary), v))
	})
}

// SummaryHasSuffix applies the HasSuffix predicate on the "summary" field.
func SummaryHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSummary), v))
	})
}

// SummaryEqualFold applies the EqualFold predicate on the "summary" field.
func SummaryEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSummary), v))
	})
}

// SummaryContainsFold applies the ContainsFold predicate on the "summary" field.
func SummaryContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSummary), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuantity), v))
	})
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldQuantity), v...))
	})
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.Product {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldQuantity), v...))
	})
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQuantity), v))
	})
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQuantity), v))
	})
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQuantity), v))
	})
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQuantity), v))
	})
}

// HasSub applies the HasEdge predicate on the "sub" edge.
func HasSub() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SubTable, SubPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubWith applies the HasEdge predicate on the "sub" edge with a given conditions (other predicates).
func HasSubWith(preds ...predicate.SubCategory) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SubTable, SubPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
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
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		p(s.Not())
	})
}
