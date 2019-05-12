go-build:
	go build -o app .

go-run:
	./app

docker-build:
	docker build -f build/Dockerfile -t soichisumi0/k8s-monitoring-sample:v1.0 .

use-local-image:
	eval \$\(minikube docker-env\)

minikube-start:
	minikube start

minikube-delete:
	minikube delete

minikube-ip:
	minikube ip

prometheus-install:
	helm install --name prometheus -f deployments/helm/prometheus-customize-values.yaml stable/prometheus

prometheus-update:
	helm upgrade -f deployments/helm/prometheus-customize-values.yaml prometheus stable/prometheus

grafana-install:
	helm install --name grafana -f deployments/helm/grafana-customize-values.yaml stable/grafana

grafana-update:
	helm upgrade -f deployments/helm/grafana-customize-values.yaml grafana stable/grafana