package types

import (
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type KippyHeartbeat struct {
	Events      []*v1.Event
	PodMetrics  []*v1beta1.PodMetrics
	NodeMetrics []*v1beta1.NodeMetrics
	Hash        string    `json:"hash,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

type KippyEvent struct {
	*v1.Event
}

type KippyPodMetrics struct {
	Name       string                  `json:"name,omitempty"`
	Namespace  string                  `json:"namespace,omitempty"`
	Containers []KippyContainerMetrics `json:"containers,omitempty"`
	Timestamp  time.Time               `json:"timestamp,omitempty"`
}

type KippyNodeMetrics struct {
	Name             string    `json:"name,omitempty"`
	CPU              string    `json:"cpu,omitempty"`
	Memory           string    `json:"memory,omitempty"`
	Storage          string    `json:"storage,omitempty"`
	StorageEphemeral string    `json:"storageEphemeral,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

type KippyContainerMetrics struct {
	Name   string `json:"name,omitempty"`
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
	Token  string `json:"token,omitempty"`
}

type KippySink struct {
	Type   string
	Config string
}

type KippySinkInterface interface {
	Send(string) error
	Payload() string
}

type KippyMessage struct {
	Kind      string    `json:"kind,omitempty"`
	Namespace string    `json:"namespace,omitempty"`
	Name      string    `json:"name,omitempty"`
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
