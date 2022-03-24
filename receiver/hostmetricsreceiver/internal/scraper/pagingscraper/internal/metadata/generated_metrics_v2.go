// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/model/pdata"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for hostmetricsreceiver/paging metrics.
type MetricsSettings struct {
	SystemPagingFaults      MetricSettings `mapstructure:"system.paging.faults"`
	SystemPagingOperations  MetricSettings `mapstructure:"system.paging.operations"`
	SystemPagingUsage       MetricSettings `mapstructure:"system.paging.usage"`
	SystemPagingUtilization MetricSettings `mapstructure:"system.paging.utilization"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemPagingFaults: MetricSettings{
			Enabled: true,
		},
		SystemPagingOperations: MetricSettings{
			Enabled: true,
		},
		SystemPagingUsage: MetricSettings{
			Enabled: true,
		},
		SystemPagingUtilization: MetricSettings{
			Enabled: false,
		},
	}
}

type metricSystemPagingFaults struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.paging.faults metric with initial data.
func (m *metricSystemPagingFaults) init() {
	m.data.SetName("system.paging.faults")
	m.data.SetDescription("The number of page faults.")
	m.data.SetUnit("{faults}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemPagingFaults) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, typeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Type, pdata.NewValueString(typeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemPagingFaults) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemPagingFaults) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemPagingFaults(settings MetricSettings) metricSystemPagingFaults {
	m := metricSystemPagingFaults{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemPagingOperations struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.paging.operations metric with initial data.
func (m *metricSystemPagingOperations) init() {
	m.data.SetName("system.paging.operations")
	m.data.SetDescription("The number of paging operations.")
	m.data.SetUnit("{operations}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemPagingOperations) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, directionAttributeValue string, typeAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Direction, pdata.NewValueString(directionAttributeValue))
	dp.Attributes().Insert(A.Type, pdata.NewValueString(typeAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemPagingOperations) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemPagingOperations) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemPagingOperations(settings MetricSettings) metricSystemPagingOperations {
	m := metricSystemPagingOperations{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemPagingUsage struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.paging.usage metric with initial data.
func (m *metricSystemPagingUsage) init() {
	m.data.SetName("system.paging.usage")
	m.data.SetDescription("Swap (unix) or pagefile (windows) usage.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemPagingUsage) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, deviceAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.State, pdata.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemPagingUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemPagingUsage) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemPagingUsage(settings MetricSettings) metricSystemPagingUsage {
	m := metricSystemPagingUsage{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemPagingUtilization struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.paging.utilization metric with initial data.
func (m *metricSystemPagingUtilization) init() {
	m.data.SetName("system.paging.utilization")
	m.data.SetDescription("Swap (unix) or pagefile (windows) utilization.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemPagingUtilization) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, deviceAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.State, pdata.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemPagingUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemPagingUtilization) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemPagingUtilization(settings MetricSettings) metricSystemPagingUtilization {
	m := metricSystemPagingUtilization{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                     pdata.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity               int             // maximum observed number of metrics per resource.
	resourceCapacity              int             // maximum observed number of resource attributes.
	metricsBuffer                 pdata.Metrics   // accumulates metrics data before emitting.
	metricSystemPagingFaults      metricSystemPagingFaults
	metricSystemPagingOperations  metricSystemPagingOperations
	metricSystemPagingUsage       metricSystemPagingUsage
	metricSystemPagingUtilization metricSystemPagingUtilization
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pdata.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                     pdata.NewTimestampFromTime(time.Now()),
		metricsBuffer:                 pdata.NewMetrics(),
		metricSystemPagingFaults:      newMetricSystemPagingFaults(settings.SystemPagingFaults),
		metricSystemPagingOperations:  newMetricSystemPagingOperations(settings.SystemPagingOperations),
		metricSystemPagingUsage:       newMetricSystemPagingUsage(settings.SystemPagingUsage),
		metricSystemPagingUtilization: newMetricSystemPagingUtilization(settings.SystemPagingUtilization),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pdata.ResourceMetrics) {
	if mb.metricsCapacity < rm.InstrumentationLibraryMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.InstrumentationLibraryMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceOption applies changes to provided resource.
type ResourceOption func(pdata.Resource)

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead. Resource attributes should be provided as ResourceOption arguments.
func (mb *MetricsBuilder) EmitForResource(ro ...ResourceOption) {
	rm := pdata.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	for _, op := range ro {
		op(rm.Resource())
	}
	ils := rm.InstrumentationLibraryMetrics().AppendEmpty()
	ils.InstrumentationLibrary().SetName("otelcol/hostmetricsreceiver/paging")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricSystemPagingFaults.emit(ils.Metrics())
	mb.metricSystemPagingOperations.emit(ils.Metrics())
	mb.metricSystemPagingUsage.emit(ils.Metrics())
	mb.metricSystemPagingUtilization.emit(ils.Metrics())
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(ro ...ResourceOption) pdata.Metrics {
	mb.EmitForResource(ro...)
	metrics := pdata.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordSystemPagingFaultsDataPoint adds a data point to system.paging.faults metric.
func (mb *MetricsBuilder) RecordSystemPagingFaultsDataPoint(ts pdata.Timestamp, val int64, typeAttributeValue string) {
	mb.metricSystemPagingFaults.recordDataPoint(mb.startTime, ts, val, typeAttributeValue)
}

// RecordSystemPagingOperationsDataPoint adds a data point to system.paging.operations metric.
func (mb *MetricsBuilder) RecordSystemPagingOperationsDataPoint(ts pdata.Timestamp, val int64, directionAttributeValue string, typeAttributeValue string) {
	mb.metricSystemPagingOperations.recordDataPoint(mb.startTime, ts, val, directionAttributeValue, typeAttributeValue)
}

// RecordSystemPagingUsageDataPoint adds a data point to system.paging.usage metric.
func (mb *MetricsBuilder) RecordSystemPagingUsageDataPoint(ts pdata.Timestamp, val int64, deviceAttributeValue string, stateAttributeValue string) {
	mb.metricSystemPagingUsage.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, stateAttributeValue)
}

// RecordSystemPagingUtilizationDataPoint adds a data point to system.paging.utilization metric.
func (mb *MetricsBuilder) RecordSystemPagingUtilizationDataPoint(ts pdata.Timestamp, val float64, deviceAttributeValue string, stateAttributeValue string) {
	mb.metricSystemPagingUtilization.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, stateAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Device (Name of the page file.)
	Device string
	// Direction (Page In or Page Out.)
	Direction string
	// State (Breakdown of paging usage by type.)
	State string
	// Type (Type of fault.)
	Type string
}{
	"device",
	"direction",
	"state",
	"type",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeDirection are the possible values that the attribute "direction" can have.
var AttributeDirection = struct {
	PageIn  string
	PageOut string
}{
	"page_in",
	"page_out",
}

// AttributeState are the possible values that the attribute "state" can have.
var AttributeState = struct {
	Cached string
	Free   string
	Used   string
}{
	"cached",
	"free",
	"used",
}

// AttributeType are the possible values that the attribute "type" can have.
var AttributeType = struct {
	Major string
	Minor string
}{
	"major",
	"minor",
}
