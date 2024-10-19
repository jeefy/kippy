package events

import (
	"context"
	"log"
	"sync"

	"github.com/jeefy/kippy/pkg/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func WatchEvents(config *rest.Config, HeartBeat *types.KippyHeartbeat, HeartBeatLock *sync.Mutex) {
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	events, err := clientset.CoreV1().Events("").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for {
		e := <-events.ResultChan()
		if e.Object == nil {
			log.Println("Event object from channel is nil. Restarting watch.")
			events, err = clientset.CoreV1().Events("").Watch(context.TODO(), metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			continue
		}
		event := e.Object.(*v1.Event)
		HeartBeatLock.Lock()
		if event.Type != "Normal" {
			HeartBeat.Events = append(HeartBeat.Events, event)
		}
		HeartBeatLock.Unlock()
	}
}
