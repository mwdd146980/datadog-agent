The kubernetes manifests found in this directory have been automatically generated
from the [helm chart `datadog/datadog`](https://github.com/DataDog/helm-charts/tree/master/charts/datadog)
version 3.3.1 with the following `values.yaml`:

```yaml
datadog:
  collectEvents: true
  apm:
    socketEnabled: false
  processAgent:
    enabled: false
  clusterChecks:
    enabled: true
clusterAgent:
  enabled: true
  metricsProvider:
    enabled: true
clusterChecksRunner:
  enabled: true
```
