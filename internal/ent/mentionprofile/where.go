// Code generated by ent, DO NOT EDIT.

package mentionprofile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Southclaws/storyden/internal/ent/predicate"
	"github.com/rs/xid"
)

// ID filters vertices based on their ID field.
func ID(id xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldEQ(FieldCreatedAt, v))
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldEQ(FieldAccountID, v))
}

// PostID applies equality check predicate on the "post_id" field. It's identical to PostIDEQ.
func PostID(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldEQ(FieldPostID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldLTE(FieldCreatedAt, v))
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldEQ(FieldAccountID, v))
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldNEQ(FieldAccountID, v))
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldIn(FieldAccountID, vs...))
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldNotIn(FieldAccountID, vs...))
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldGT(FieldAccountID, v))
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldGTE(FieldAccountID, v))
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldLT(FieldAccountID, v))
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldLTE(FieldAccountID, v))
}

// AccountIDContains applies the Contains predicate on the "account_id" field.
func AccountIDContains(v xid.ID) predicate.MentionProfile {
	vc := v.String()
	return predicate.MentionProfile(sql.FieldContains(FieldAccountID, vc))
}

// AccountIDHasPrefix applies the HasPrefix predicate on the "account_id" field.
func AccountIDHasPrefix(v xid.ID) predicate.MentionProfile {
	vc := v.String()
	return predicate.MentionProfile(sql.FieldHasPrefix(FieldAccountID, vc))
}

// AccountIDHasSuffix applies the HasSuffix predicate on the "account_id" field.
func AccountIDHasSuffix(v xid.ID) predicate.MentionProfile {
	vc := v.String()
	return predicate.MentionProfile(sql.FieldHasSuffix(FieldAccountID, vc))
}

// AccountIDEqualFold applies the EqualFold predicate on the "account_id" field.
func AccountIDEqualFold(v xid.ID) predicate.MentionProfile {
	vc := v.String()
	return predicate.MentionProfile(sql.FieldEqualFold(FieldAccountID, vc))
}

// AccountIDContainsFold applies the ContainsFold predicate on the "account_id" field.
func AccountIDContainsFold(v xid.ID) predicate.MentionProfile {
	vc := v.String()
	return predicate.MentionProfile(sql.FieldContainsFold(FieldAccountID, vc))
}

// PostIDEQ applies the EQ predicate on the "post_id" field.
func PostIDEQ(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldEQ(FieldPostID, v))
}

// PostIDNEQ applies the NEQ predicate on the "post_id" field.
func PostIDNEQ(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldNEQ(FieldPostID, v))
}

// PostIDIn applies the In predicate on the "post_id" field.
func PostIDIn(vs ...xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldIn(FieldPostID, vs...))
}

// PostIDNotIn applies the NotIn predicate on the "post_id" field.
func PostIDNotIn(vs ...xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldNotIn(FieldPostID, vs...))
}

// PostIDGT applies the GT predicate on the "post_id" field.
func PostIDGT(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldGT(FieldPostID, v))
}

// PostIDGTE applies the GTE predicate on the "post_id" field.
func PostIDGTE(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldGTE(FieldPostID, v))
}

// PostIDLT applies the LT predicate on the "post_id" field.
func PostIDLT(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldLT(FieldPostID, v))
}

// PostIDLTE applies the LTE predicate on the "post_id" field.
func PostIDLTE(v xid.ID) predicate.MentionProfile {
	return predicate.MentionProfile(sql.FieldLTE(FieldPostID, v))
}

// PostIDContains applies the Contains predicate on the "post_id" field.
func PostIDContains(v xid.ID) predicate.MentionProfile {
	vc := v.String()
	return predicate.MentionProfile(sql.FieldContains(FieldPostID, vc))
}

// PostIDHasPrefix applies the HasPrefix predicate on the "post_id" field.
func PostIDHasPrefix(v xid.ID) predicate.MentionProfile {
	vc := v.String()
	return predicate.MentionProfile(sql.FieldHasPrefix(FieldPostID, vc))
}

// PostIDHasSuffix applies the HasSuffix predicate on the "post_id" field.
func PostIDHasSuffix(v xid.ID) predicate.MentionProfile {
	vc := v.String()
	return predicate.MentionProfile(sql.FieldHasSuffix(FieldPostID, vc))
}

// PostIDEqualFold applies the EqualFold predicate on the "post_id" field.
func PostIDEqualFold(v xid.ID) predicate.MentionProfile {
	vc := v.String()
	return predicate.MentionProfile(sql.FieldEqualFold(FieldPostID, vc))
}

// PostIDContainsFold applies the ContainsFold predicate on the "post_id" field.
func PostIDContainsFold(v xid.ID) predicate.MentionProfile {
	vc := v.String()
	return predicate.MentionProfile(sql.FieldContainsFold(FieldPostID, vc))
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.MentionProfile {
	return predicate.MentionProfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.MentionProfile {
	return predicate.MentionProfile(func(s *sql.Selector) {
		step := newAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPost applies the HasEdge predicate on the "Post" edge.
func HasPost() predicate.MentionProfile {
	return predicate.MentionProfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PostTable, PostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostWith applies the HasEdge predicate on the "Post" edge with a given conditions (other predicates).
func HasPostWith(preds ...predicate.Post) predicate.MentionProfile {
	return predicate.MentionProfile(func(s *sql.Selector) {
		step := newPostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MentionProfile) predicate.MentionProfile {
	return predicate.MentionProfile(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MentionProfile) predicate.MentionProfile {
	return predicate.MentionProfile(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MentionProfile) predicate.MentionProfile {
	return predicate.MentionProfile(sql.NotPredicates(p))
}
