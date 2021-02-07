// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/anomaly"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/detectionjob"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/detectionjobinstance"
	"github.com/DanielTitkov/anomaly-detection-service/internal/repository/entgo/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	anomalyMixin := schema.Anomaly{}.Mixin()
	anomalyMixinFields0 := anomalyMixin[0].Fields()
	_ = anomalyMixinFields0
	anomalyFields := schema.Anomaly{}.Fields()
	_ = anomalyFields
	// anomalyDescCreateTime is the schema descriptor for create_time field.
	anomalyDescCreateTime := anomalyMixinFields0[0].Descriptor()
	// anomaly.DefaultCreateTime holds the default value on creation for the create_time field.
	anomaly.DefaultCreateTime = anomalyDescCreateTime.Default.(func() time.Time)
	// anomalyDescUpdateTime is the schema descriptor for update_time field.
	anomalyDescUpdateTime := anomalyMixinFields0[1].Descriptor()
	// anomaly.DefaultUpdateTime holds the default value on creation for the update_time field.
	anomaly.DefaultUpdateTime = anomalyDescUpdateTime.Default.(func() time.Time)
	// anomaly.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	anomaly.UpdateDefaultUpdateTime = anomalyDescUpdateTime.UpdateDefault.(func() time.Time)
	// anomalyDescType is the schema descriptor for type field.
	anomalyDescType := anomalyFields[0].Descriptor()
	// anomaly.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	anomaly.TypeValidator = anomalyDescType.Validators[0].(func(string) error)
	// anomalyDescProcessed is the schema descriptor for processed field.
	anomalyDescProcessed := anomalyFields[2].Descriptor()
	// anomaly.DefaultProcessed holds the default value on creation for the processed field.
	anomaly.DefaultProcessed = anomalyDescProcessed.Default.(bool)
	detectionjobMixin := schema.DetectionJob{}.Mixin()
	detectionjobMixinFields0 := detectionjobMixin[0].Fields()
	_ = detectionjobMixinFields0
	detectionjobFields := schema.DetectionJob{}.Fields()
	_ = detectionjobFields
	// detectionjobDescCreateTime is the schema descriptor for create_time field.
	detectionjobDescCreateTime := detectionjobMixinFields0[0].Descriptor()
	// detectionjob.DefaultCreateTime holds the default value on creation for the create_time field.
	detectionjob.DefaultCreateTime = detectionjobDescCreateTime.Default.(func() time.Time)
	// detectionjobDescUpdateTime is the schema descriptor for update_time field.
	detectionjobDescUpdateTime := detectionjobMixinFields0[1].Descriptor()
	// detectionjob.DefaultUpdateTime holds the default value on creation for the update_time field.
	detectionjob.DefaultUpdateTime = detectionjobDescUpdateTime.Default.(func() time.Time)
	// detectionjob.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	detectionjob.UpdateDefaultUpdateTime = detectionjobDescUpdateTime.UpdateDefault.(func() time.Time)
	// detectionjobDescMethod is the schema descriptor for method field.
	detectionjobDescMethod := detectionjobFields[1].Descriptor()
	// detectionjob.MethodValidator is a validator for the "method" field. It is called by the builders before save.
	detectionjob.MethodValidator = detectionjobDescMethod.Validators[0].(func(string) error)
	// detectionjobDescSiteID is the schema descriptor for site_id field.
	detectionjobDescSiteID := detectionjobFields[2].Descriptor()
	// detectionjob.SiteIDValidator is a validator for the "site_id" field. It is called by the builders before save.
	detectionjob.SiteIDValidator = detectionjobDescSiteID.Validators[0].(func(string) error)
	// detectionjobDescMetric is the schema descriptor for metric field.
	detectionjobDescMetric := detectionjobFields[3].Descriptor()
	// detectionjob.MetricValidator is a validator for the "metric" field. It is called by the builders before save.
	detectionjob.MetricValidator = detectionjobDescMetric.Validators[0].(func(string) error)
	// detectionjobDescAttribute is the schema descriptor for attribute field.
	detectionjobDescAttribute := detectionjobFields[4].Descriptor()
	// detectionjob.AttributeValidator is a validator for the "attribute" field. It is called by the builders before save.
	detectionjob.AttributeValidator = detectionjobDescAttribute.Validators[0].(func(string) error)
	// detectionjobDescTimeAgo is the schema descriptor for time_ago field.
	detectionjobDescTimeAgo := detectionjobFields[5].Descriptor()
	// detectionjob.TimeAgoValidator is a validator for the "time_ago" field. It is called by the builders before save.
	detectionjob.TimeAgoValidator = detectionjobDescTimeAgo.Validators[0].(func(string) error)
	// detectionjobDescTimeStep is the schema descriptor for time_step field.
	detectionjobDescTimeStep := detectionjobFields[6].Descriptor()
	// detectionjob.TimeStepValidator is a validator for the "time_step" field. It is called by the builders before save.
	detectionjob.TimeStepValidator = detectionjobDescTimeStep.Validators[0].(func(string) error)
	detectionjobinstanceMixin := schema.DetectionJobInstance{}.Mixin()
	detectionjobinstanceMixinFields0 := detectionjobinstanceMixin[0].Fields()
	_ = detectionjobinstanceMixinFields0
	detectionjobinstanceFields := schema.DetectionJobInstance{}.Fields()
	_ = detectionjobinstanceFields
	// detectionjobinstanceDescCreateTime is the schema descriptor for create_time field.
	detectionjobinstanceDescCreateTime := detectionjobinstanceMixinFields0[0].Descriptor()
	// detectionjobinstance.DefaultCreateTime holds the default value on creation for the create_time field.
	detectionjobinstance.DefaultCreateTime = detectionjobinstanceDescCreateTime.Default.(func() time.Time)
	// detectionjobinstanceDescUpdateTime is the schema descriptor for update_time field.
	detectionjobinstanceDescUpdateTime := detectionjobinstanceMixinFields0[1].Descriptor()
	// detectionjobinstance.DefaultUpdateTime holds the default value on creation for the update_time field.
	detectionjobinstance.DefaultUpdateTime = detectionjobinstanceDescUpdateTime.Default.(func() time.Time)
	// detectionjobinstance.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	detectionjobinstance.UpdateDefaultUpdateTime = detectionjobinstanceDescUpdateTime.UpdateDefault.(func() time.Time)
}
