// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"shopular/ent/category"
	"shopular/ent/subcategory"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// SubCategory is the model entity for the SubCategory schema.
type SubCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubCategoryQuery when eager-loading is set.
	Edges         SubCategoryEdges `json:"edges"`
	category_subs *int
}

// SubCategoryEdges holds the relations/edges for other nodes in the graph.
type SubCategoryEdges struct {
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// Product holds the value of the product edge.
	Product []*Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubCategoryEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[0] {
		if e.Category == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e SubCategoryEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[1] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SubCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subcategory.FieldID:
			values[i] = new(sql.NullInt64)
		case subcategory.FieldTitle, subcategory.FieldDescription:
			values[i] = new(sql.NullString)
		case subcategory.FieldCreatedAt, subcategory.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case subcategory.ForeignKeys[0]: // category_subs
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SubCategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SubCategory fields.
func (sc *SubCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sc.ID = int(value.Int64)
		case subcategory.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				sc.Title = value.String
			}
		case subcategory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				sc.Description = value.String
			}
		case subcategory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sc.CreatedAt = value.Time
			}
		case subcategory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sc.UpdatedAt = value.Time
			}
		case subcategory.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field category_subs", value)
			} else if value.Valid {
				sc.category_subs = new(int)
				*sc.category_subs = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCategory queries the "category" edge of the SubCategory entity.
func (sc *SubCategory) QueryCategory() *CategoryQuery {
	return (&SubCategoryClient{config: sc.config}).QueryCategory(sc)
}

// QueryProduct queries the "product" edge of the SubCategory entity.
func (sc *SubCategory) QueryProduct() *ProductQuery {
	return (&SubCategoryClient{config: sc.config}).QueryProduct(sc)
}

// Update returns a builder for updating this SubCategory.
// Note that you need to call SubCategory.Unwrap() before calling this method if this SubCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *SubCategory) Update() *SubCategoryUpdateOne {
	return (&SubCategoryClient{config: sc.config}).UpdateOne(sc)
}

// Unwrap unwraps the SubCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *SubCategory) Unwrap() *SubCategory {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: SubCategory is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *SubCategory) String() string {
	var builder strings.Builder
	builder.WriteString("SubCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("title=")
	builder.WriteString(sc.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(sc.Description)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sc.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SubCategories is a parsable slice of SubCategory.
type SubCategories []*SubCategory

func (sc SubCategories) config(cfg config) {
	for _i := range sc {
		sc[_i].config = cfg
	}
}