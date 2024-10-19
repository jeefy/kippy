# Kippy

Your Kubernetes Insurance Policy (KIP-Y). But probably not for enterprises. 

## Quickstart

```
kubectl apply -f https://raw.githubusercontent.com/jeefy/kippy/refs/heads/main/deployment.yaml
```

Then edit the secret and cycle the pod. :) 

## What does it do?

KIPPY is meant to be simple. It watches all events emitted by the Kubernetes API Server. 

Any not-normal event emitted is then collected and sent to configured notification sinks every 60s.

### Notification Sinks

- Discord
- Slack
- Sendgrid
- Generic Webhook

## Prior art

https://github.com/redhat-cop/k8s-notify

https://github.com/opsgenie/kubernetes-event-exporter

## Why?

1. A lot of the previous examples are archived.
2. I like reinventing the wheel.
3. This matches my exact use-case.