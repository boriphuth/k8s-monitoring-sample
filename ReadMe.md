## Install Minikube https://github.com/kubernetes/minikube/releases
$ curl -Lo minikube https://storage.googleapis.com/minikube/releases/v1.4.0/minikube-darwin-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/
$ minikube start --cpus=2 --memory=2000
$ minikube version
$ minikube ip

## Install kubectl
$ brew install kubectl
$ kubectl api-versions

## Docker for Mac https://kubernetes.github.io/ingress-nginx/deploy/#docker-for-mac
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud-generic.yaml
$ kubectl create ns ingress-nginx

## Helm
$ brew install kubernetes-helm
$ curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get > get_helm.sh
$ chmod 700 get_helm.sh
$ ./get_helm.sh
$ kubectl create -f clusterrole.yaml
$ kubectl create serviceaccount -n kube-system tiller
$ kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
$ helm init --service-account tiller
$ kubectl --namespace kube-system get pods

kubectl create ns monitoring
kubectl create ns ingress-nginx

## Deploy Prometheus
$ helm install --name prometheus -f helm/prometheus-customize-values.yaml stable/prometheus

## Deploy Grafana
$ helm install --name grafana -f helm/grafana-customize-values.yaml stable/grafana

$ kubectl get deploy,svc,ing

## Enable ingress
## Docker for Mac
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud-generic.yaml

## Minikube
$ minikube addons enable ingress

## Domain assignment
$ code . /etc/hosts
192.168.99.105  prometheus.local grafana.local app.local

$ kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
aftNwz94xEYml34dE05EtsKl3kEdhkyFCRF4XIlj

$ helm del --purge prometheus
$ helm del --purge grafana


$ kubectl delete --all pods --namespace=default


เปิด dashboard
$ kubectl proxy
จากนั้นเปิดลิงค์ http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/แสดง pods
$ kubectl get podsดูรายละเอียด pods
$ kubectl describe pods
$ kubectl describe pods demo-nginx // แสดงเฉพาะ demo-nginxแสดง services
$ kubectl get services
$ kubectl get services demo-nginx // แสดงเฉพาะ demo-nginxเข้าถึง containers ผ่าน shell
$ kubectl exec -it demo-nginx-548685f5cc-v7rmc shแสดง logs
$ kubectl logs -f demo-nginx-548685f5cc-v7rmc