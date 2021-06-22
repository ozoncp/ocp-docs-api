package metrics

import "github.com/prometheus/client_golang/prometheus"

var processedByCRUDHandler *prometheus.CounterVec = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "processed_by_crud_handler",
		Help: "Num of processed By CRUD Handler",
	},
	[]string{"handler", "status"}, // labels
)

func RegisterMetrics() {
	prometheus.MustRegister(processedByCRUDHandler)
}

func incrementByHandler(handler string, status string, count int) {
	processedByCRUDHandler.With(prometheus.Labels{"handler": handler, "status": status}).Add(float64(count))
}

func IncrementCreate(count int, status string) {
	incrementByHandler("create", status, count)
}

func IncrementRead(count int, status string) {
	incrementByHandler("read", status, count)
}

func IncrementUpdate(count int, status string) {
	incrementByHandler("update", status, count)
}

func IncrementDelete(count int, status string) {
	incrementByHandler("delete", status, count)
}
