// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/detectionjob"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/detectionjobinstance"
)

// DetectionJobInstance is the model entity for the DetectionJobInstance schema.
type DetectionJobInstance struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Value holds the value of the "value" field.
	Value float64 `json:"value,omitempty"`
	// Processed holds the value of the "processed" field.
	Processed bool `json:"processed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DetectionJobInstanceQuery when eager-loading is set.
	Edges                  DetectionJobInstanceEdges `json:"edges"`
	detection_job_instance *int
}

// DetectionJobInstanceEdges holds the relations/edges for other nodes in the graph.
type DetectionJobInstanceEdges struct {
	// Anomalies holds the value of the anomalies edge.
	Anomalies []*Anomaly `json:"anomalies,omitempty"`
	// DetectionJob holds the value of the detection_job edge.
	DetectionJob *DetectionJob `json:"detection_job,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AnomaliesOrErr returns the Anomalies value or an error if the edge
// was not loaded in eager-loading.
func (e DetectionJobInstanceEdges) AnomaliesOrErr() ([]*Anomaly, error) {
	if e.loadedTypes[0] {
		return e.Anomalies, nil
	}
	return nil, &NotLoadedError{edge: "anomalies"}
}

// DetectionJobOrErr returns the DetectionJob value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DetectionJobInstanceEdges) DetectionJobOrErr() (*DetectionJob, error) {
	if e.loadedTypes[1] {
		if e.DetectionJob == nil {
			// The edge detection_job was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: detectionjob.Label}
		}
		return e.DetectionJob, nil
	}
	return nil, &NotLoadedError{edge: "detection_job"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DetectionJobInstance) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case detectionjobinstance.FieldProcessed:
			values[i] = &sql.NullBool{}
		case detectionjobinstance.FieldValue:
			values[i] = &sql.NullFloat64{}
		case detectionjobinstance.FieldID:
			values[i] = &sql.NullInt64{}
		case detectionjobinstance.FieldType:
			values[i] = &sql.NullString{}
		case detectionjobinstance.FieldCreateTime, detectionjobinstance.FieldUpdateTime:
			values[i] = &sql.NullTime{}
		case detectionjobinstance.ForeignKeys[0]: // detection_job_instance
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type DetectionJobInstance", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DetectionJobInstance fields.
func (dji *DetectionJobInstance) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case detectionjobinstance.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dji.ID = int(value.Int64)
		case detectionjobinstance.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				dji.CreateTime = value.Time
			}
		case detectionjobinstance.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				dji.UpdateTime = value.Time
			}
		case detectionjobinstance.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				dji.Type = value.String
			}
		case detectionjobinstance.FieldValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				dji.Value = value.Float64
			}
		case detectionjobinstance.FieldProcessed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field processed", values[i])
			} else if value.Valid {
				dji.Processed = value.Bool
			}
		case detectionjobinstance.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field detection_job_instance", value)
			} else if value.Valid {
				dji.detection_job_instance = new(int)
				*dji.detection_job_instance = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAnomalies queries the "anomalies" edge of the DetectionJobInstance entity.
func (dji *DetectionJobInstance) QueryAnomalies() *AnomalyQuery {
	return (&DetectionJobInstanceClient{config: dji.config}).QueryAnomalies(dji)
}

// QueryDetectionJob queries the "detection_job" edge of the DetectionJobInstance entity.
func (dji *DetectionJobInstance) QueryDetectionJob() *DetectionJobQuery {
	return (&DetectionJobInstanceClient{config: dji.config}).QueryDetectionJob(dji)
}

// Update returns a builder for updating this DetectionJobInstance.
// Note that you need to call DetectionJobInstance.Unwrap() before calling this method if this DetectionJobInstance
// was returned from a transaction, and the transaction was committed or rolled back.
func (dji *DetectionJobInstance) Update() *DetectionJobInstanceUpdateOne {
	return (&DetectionJobInstanceClient{config: dji.config}).UpdateOne(dji)
}

// Unwrap unwraps the DetectionJobInstance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dji *DetectionJobInstance) Unwrap() *DetectionJobInstance {
	tx, ok := dji.config.driver.(*txDriver)
	if !ok {
		panic("ent: DetectionJobInstance is not a transactional entity")
	}
	dji.config.driver = tx.drv
	return dji
}

// String implements the fmt.Stringer.
func (dji *DetectionJobInstance) String() string {
	var builder strings.Builder
	builder.WriteString("DetectionJobInstance(")
	builder.WriteString(fmt.Sprintf("id=%v", dji.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(dji.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(dji.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", type=")
	builder.WriteString(dji.Type)
	builder.WriteString(", value=")
	builder.WriteString(fmt.Sprintf("%v", dji.Value))
	builder.WriteString(", processed=")
	builder.WriteString(fmt.Sprintf("%v", dji.Processed))
	builder.WriteByte(')')
	return builder.String()
}

// DetectionJobInstances is a parsable slice of DetectionJobInstance.
type DetectionJobInstances []*DetectionJobInstance

func (dji DetectionJobInstances) config(cfg config) {
	for _i := range dji {
		dji[_i].config = cfg
	}
}
