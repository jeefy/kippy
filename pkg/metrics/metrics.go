package metrics

import (
	"context"
	"sync"

	"github.com/jeefy/kippy/pkg/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func GetMetrics(config *rest.Config, HeartBeat *types.KippyHeartbeat, HeartBeatLock *sync.Mutex) {
	// creates the clientset
	clientset, err := metricsv.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.MetricsV1beta1().PodMetricses("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	PodMetrics(pods, HeartBeat, HeartBeatLock)

	nodes, err := clientset.MetricsV1beta1().NodeMetricses().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	NodeMetrics(nodes, HeartBeat, HeartBeatLock)

}

func PodMetrics(podMetrics *v1beta1.PodMetricsList, HeartBeat *types.KippyHeartbeat, HeartBeatLock *sync.Mutex) {
	for _, metric := range podMetrics.Items {
		HeartBeatLock.Lock()
		HeartBeat.PodMetrics = append(HeartBeat.PodMetrics, &metric)
		HeartBeatLock.Unlock()
	}
}

func NodeMetrics(nodeMetrics *v1beta1.NodeMetricsList, HeartBeat *types.KippyHeartbeat, HeartBeatLock *sync.Mutex) {
	for _, metric := range nodeMetrics.Items {
		HeartBeatLock.Lock()
		HeartBeat.NodeMetrics = append(HeartBeat.NodeMetrics, &metric)
		HeartBeatLock.Unlock()
	}
}
