package metrics

import "github.com/prometheus/client_golang/prometheus"

var processedByCRUDHandler *prometheus.CounterVec = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "processed_by_crud_handler",
		Help: "Num of processed By CRUD Handler",
	},
	[]string{"handler"}, // labels
)

func RegisterMetrics() {
	prometheus.MustRegister(processedByCRUDHandler)
}

func incrementByHandler(handler string, count int) {
	processedByCRUDHandler.With(prometheus.Labels{"handler": handler}).Add(float64(count))
}

func IncrementSuccessfulCreate(count int) {
	incrementByHandler("create", count)
}

func IncrementSuccessfulRead(count int) {
	incrementByHandler("read", count)
}

func IncrementSuccessfulUpdate(count int) {
	incrementByHandler("update", count)
}

func IncrementSuccessfulDelete(count int) {
	incrementByHandler("delete", count)
}
