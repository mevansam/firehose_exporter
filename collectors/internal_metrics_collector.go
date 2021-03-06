package collectors

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/cloudfoundry-community/firehose_exporter/metrics"
)

type InternalMetricsCollector struct {
	namespace                                string
	metricsStore                             *metrics.Store
	totalEnvelopesReceivedDesc               *prometheus.Desc
	lastEnvelopeReceivedTimestampDesc        *prometheus.Desc
	totalMetricsReceivedDesc                 *prometheus.Desc
	lastMetricReceivedTimestampDesc          *prometheus.Desc
	totalContainerMetricsReceivedDesc        *prometheus.Desc
	totalContainerMetricsProcessedDesc       *prometheus.Desc
	lastContainerMetricReceivedTimestampDesc *prometheus.Desc
	totalCounterEventsReceivedDesc           *prometheus.Desc
	totalCounterEventsProcessedDesc          *prometheus.Desc
	lastCounterEventReceivedTimestampDesc    *prometheus.Desc
	totalValueMetricsReceivedDesc            *prometheus.Desc
	totalValueMetricsProcessedDesc           *prometheus.Desc
	lastValueMetricReceivedTimestampDesc     *prometheus.Desc
	slowConsumerAlertDesc                    *prometheus.Desc
	lastSlowConsumerAlertTimestampDesc       *prometheus.Desc
}

func NewInternalMetricsCollector(
	namespace string,
	metricsStore *metrics.Store,
) *InternalMetricsCollector {
	totalEnvelopesReceivedDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_envelopes_received"),
		"Total number of envelopes received from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	lastEnvelopeReceivedTimestampDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "last_envelope_received_timestamp"),
		"Number of seconds since 1970 since last envelope received from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	totalMetricsReceivedDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_metrics_received"),
		"Total number of metrics received from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	lastMetricReceivedTimestampDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "last_metric_received_timestamp"),
		"Number of seconds since 1970 since last metric received from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	totalContainerMetricsReceivedDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_container_metrics_received"),
		"Total number of container metrics received from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	totalContainerMetricsProcessedDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_container_metrics_processed"),
		"Total number of container metrics processed from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	lastContainerMetricReceivedTimestampDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "last_container_metric_received_timestamp"),
		"Number of seconds since 1970 since last container metric received from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	totalCounterEventsReceivedDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_counter_events_received"),
		"Total number of counter events received from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	totalCounterEventsProcessedDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_counter_events_processed"),
		"Total number of counter events processed from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	lastCounterEventReceivedTimestampDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "last_counter_event_received_timestamp"),
		"Number of seconds since 1970 since last counter event received from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	totalValueMetricsReceivedDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_value_metrics_received"),
		"Total number of value metrics received from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	totalValueMetricsProcessedDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_value_metrics_processed"),
		"Total number of value metrics processed from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	lastValueMetricReceivedTimestampDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "last_value_metric_received_timestamp"),
		"Number of seconds since 1970 since last value metric received from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	slowConsumerAlertDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "slow_consumer_alert"),
		"Nozzle could not keep up with Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	lastSlowConsumerAlertTimestampDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "last_slow_consumer_alert_timestamp"),
		"Number of seconds since 1970 since last slow consumer alert received from Cloud Foundry Firehose.",
		[]string{},
		nil,
	)

	collector := &InternalMetricsCollector{
		namespace:                                namespace,
		metricsStore:                             metricsStore,
		totalEnvelopesReceivedDesc:               totalEnvelopesReceivedDesc,
		lastEnvelopeReceivedTimestampDesc:        lastEnvelopeReceivedTimestampDesc,
		totalMetricsReceivedDesc:                 totalMetricsReceivedDesc,
		lastMetricReceivedTimestampDesc:          lastMetricReceivedTimestampDesc,
		totalContainerMetricsReceivedDesc:        totalContainerMetricsReceivedDesc,
		totalContainerMetricsProcessedDesc:       totalContainerMetricsProcessedDesc,
		lastContainerMetricReceivedTimestampDesc: lastContainerMetricReceivedTimestampDesc,
		totalCounterEventsReceivedDesc:           totalCounterEventsReceivedDesc,
		totalCounterEventsProcessedDesc:          totalCounterEventsProcessedDesc,
		lastCounterEventReceivedTimestampDesc:    lastCounterEventReceivedTimestampDesc,
		totalValueMetricsReceivedDesc:            totalValueMetricsReceivedDesc,
		totalValueMetricsProcessedDesc:           totalValueMetricsProcessedDesc,
		lastValueMetricReceivedTimestampDesc:     lastValueMetricReceivedTimestampDesc,
		slowConsumerAlertDesc:                    slowConsumerAlertDesc,
		lastSlowConsumerAlertTimestampDesc:       lastSlowConsumerAlertTimestampDesc,
	}
	return collector
}

func (c InternalMetricsCollector) Collect(ch chan<- prometheus.Metric) {
	internalMetrics := c.metricsStore.GetInternalMetrics()

	ch <- prometheus.MustNewConstMetric(
		c.totalEnvelopesReceivedDesc,
		prometheus.CounterValue,
		float64(internalMetrics.TotalEnvelopesReceived),
	)

	ch <- prometheus.MustNewConstMetric(
		c.lastEnvelopeReceivedTimestampDesc,
		prometheus.GaugeValue,
		float64(internalMetrics.LastEnvelopReceivedTimestamp),
	)

	ch <- prometheus.MustNewConstMetric(
		c.totalMetricsReceivedDesc,
		prometheus.CounterValue,
		float64(internalMetrics.TotalMetricsReceived),
	)

	ch <- prometheus.MustNewConstMetric(
		c.lastMetricReceivedTimestampDesc,
		prometheus.GaugeValue,
		float64(internalMetrics.LastMetricReceivedTimestamp),
	)

	ch <- prometheus.MustNewConstMetric(
		c.totalContainerMetricsReceivedDesc,
		prometheus.CounterValue,
		float64(internalMetrics.TotalContainerMetricsReceived),
	)

	ch <- prometheus.MustNewConstMetric(
		c.totalContainerMetricsProcessedDesc,
		prometheus.CounterValue,
		float64(internalMetrics.TotalContainerMetricsProcessed),
	)

	ch <- prometheus.MustNewConstMetric(
		c.lastContainerMetricReceivedTimestampDesc,
		prometheus.GaugeValue,
		float64(internalMetrics.LastContainerMetricReceivedTimestamp),
	)

	ch <- prometheus.MustNewConstMetric(
		c.totalCounterEventsReceivedDesc,
		prometheus.CounterValue,
		float64(internalMetrics.TotalCounterEventsReceived),
	)

	ch <- prometheus.MustNewConstMetric(
		c.totalCounterEventsProcessedDesc,
		prometheus.CounterValue,
		float64(internalMetrics.TotalCounterEventsProcessed),
	)

	ch <- prometheus.MustNewConstMetric(
		c.lastCounterEventReceivedTimestampDesc,
		prometheus.GaugeValue,
		float64(internalMetrics.LastCounterEventReceivedTimestamp),
	)

	ch <- prometheus.MustNewConstMetric(
		c.totalValueMetricsReceivedDesc,
		prometheus.CounterValue,
		float64(internalMetrics.TotalValueMetricsReceived),
	)

	ch <- prometheus.MustNewConstMetric(
		c.totalValueMetricsProcessedDesc,
		prometheus.CounterValue,
		float64(internalMetrics.TotalValueMetricsProcessed),
	)

	ch <- prometheus.MustNewConstMetric(
		c.lastValueMetricReceivedTimestampDesc,
		prometheus.GaugeValue,
		float64(internalMetrics.LastValueMetricReceivedTimestamp),
	)

	if internalMetrics.SlowConsumerAlert {
		ch <- prometheus.MustNewConstMetric(
			c.slowConsumerAlertDesc,
			prometheus.UntypedValue,
			1,
		)
	} else {
		ch <- prometheus.MustNewConstMetric(
			c.slowConsumerAlertDesc,
			prometheus.UntypedValue,
			0,
		)
	}

	ch <- prometheus.MustNewConstMetric(
		c.lastSlowConsumerAlertTimestampDesc,
		prometheus.GaugeValue,
		float64(internalMetrics.LastSlowConsumerAlertTimestamp),
	)
}

func (c InternalMetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.totalEnvelopesReceivedDesc
	ch <- c.lastEnvelopeReceivedTimestampDesc
	ch <- c.totalMetricsReceivedDesc
	ch <- c.lastMetricReceivedTimestampDesc
	ch <- c.totalContainerMetricsReceivedDesc
	ch <- c.totalContainerMetricsProcessedDesc
	ch <- c.lastContainerMetricReceivedTimestampDesc
	ch <- c.totalCounterEventsReceivedDesc
	ch <- c.totalCounterEventsProcessedDesc
	ch <- c.lastCounterEventReceivedTimestampDesc
	ch <- c.totalValueMetricsReceivedDesc
	ch <- c.totalValueMetricsProcessedDesc
	ch <- c.lastValueMetricReceivedTimestampDesc
	ch <- c.slowConsumerAlertDesc
	ch <- c.lastSlowConsumerAlertTimestampDesc
}
