## Prometheus
$ helm inspect values stable/prometheus > prometheus-values.yaml
$ helm install --name prometheus --namespace monitoring -f prometheus-values.yaml stable/prometheus

## Grafana
$ helm install --name grafana --namespace monitoring -f grafana-values.yaml stable/grafana
$ kubectl get secret --namespace monitoring grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

$ helm del --purge prometheus
$ helm del --purge grafana