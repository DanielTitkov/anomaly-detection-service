// Code generated by entc, DO NOT EDIT.

package detectionjobinstance

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
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
func IDGT(id int) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// FinishedAt applies equality check predicate on the "finished_at" field. It's identical to FinishedAtEQ.
func FinishedAt(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFinishedAt), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.DetectionJobInstance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.DetectionJobInstance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.DetectionJobInstance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.DetectionJobInstance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.DetectionJobInstance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartedAt), v...))
	})
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.DetectionJobInstance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartedAt), v...))
	})
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartedAt), v))
	})
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartedAt), v))
	})
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtIsNil applies the IsNil predicate on the "started_at" field.
func StartedAtIsNil() predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartedAt)))
	})
}

// StartedAtNotNil applies the NotNil predicate on the "started_at" field.
func StartedAtNotNil() predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartedAt)))
	})
}

// FinishedAtEQ applies the EQ predicate on the "finished_at" field.
func FinishedAtEQ(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFinishedAt), v))
	})
}

// FinishedAtNEQ applies the NEQ predicate on the "finished_at" field.
func FinishedAtNEQ(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFinishedAt), v))
	})
}

// FinishedAtIn applies the In predicate on the "finished_at" field.
func FinishedAtIn(vs ...time.Time) predicate.DetectionJobInstance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFinishedAt), v...))
	})
}

// FinishedAtNotIn applies the NotIn predicate on the "finished_at" field.
func FinishedAtNotIn(vs ...time.Time) predicate.DetectionJobInstance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFinishedAt), v...))
	})
}

// FinishedAtGT applies the GT predicate on the "finished_at" field.
func FinishedAtGT(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFinishedAt), v))
	})
}

// FinishedAtGTE applies the GTE predicate on the "finished_at" field.
func FinishedAtGTE(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFinishedAt), v))
	})
}

// FinishedAtLT applies the LT predicate on the "finished_at" field.
func FinishedAtLT(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFinishedAt), v))
	})
}

// FinishedAtLTE applies the LTE predicate on the "finished_at" field.
func FinishedAtLTE(v time.Time) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFinishedAt), v))
	})
}

// FinishedAtIsNil applies the IsNil predicate on the "finished_at" field.
func FinishedAtIsNil() predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFinishedAt)))
	})
}

// FinishedAtNotNil applies the NotNil predicate on the "finished_at" field.
func FinishedAtNotNil() predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFinishedAt)))
	})
}

// HasAnomalies applies the HasEdge predicate on the "anomalies" edge.
func HasAnomalies() predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnomaliesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AnomaliesTable, AnomaliesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnomaliesWith applies the HasEdge predicate on the "anomalies" edge with a given conditions (other predicates).
func HasAnomaliesWith(preds ...predicate.Anomaly) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnomaliesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AnomaliesTable, AnomaliesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDetectionJob applies the HasEdge predicate on the "detection_job" edge.
func HasDetectionJob() predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DetectionJobTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DetectionJobTable, DetectionJobColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDetectionJobWith applies the HasEdge predicate on the "detection_job" edge with a given conditions (other predicates).
func HasDetectionJobWith(preds ...predicate.DetectionJob) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DetectionJobInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DetectionJobTable, DetectionJobColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DetectionJobInstance) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DetectionJobInstance) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
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
func Not(p predicate.DetectionJobInstance) predicate.DetectionJobInstance {
	return predicate.DetectionJobInstance(func(s *sql.Selector) {
		p(s.Not())
	})
}
