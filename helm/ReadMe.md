## https://soichisumi.net/2019/05/setup-prom-and-grafana-on-k8s/?fbclid=IwAR3AgGi9-b6UdPEOAY_MqlRcLIEGTNRLCIdP_MeKMcaug_apCxgW7yjN7Ww
## Deploy Prometheus
Create the following yaml file in deployments / helm / prometheus-customize-values.yaml to get the cluster and application metrics correctly.
$ brew install kubernetes-helm
$ kubectl -n kube-system get pods
$ kubectl create serviceaccount -n kube-system tiller
$ kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
$ helm init --service-account tiller
$ helm init --upgrade --service-account tiller
$ kubectl --namespace kube-system get pods | grep tiller

## Installing helm
#download helm
$ curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get > install-helm.sh

#Make instalation script executable
$ chmod u+x install-helm.sh

#Install helm
$ ./install-helm.sh

$ curl -L https://git.io/get_helm.sh | bash

## Link Helm with tiller
#Create tiller service account
$ kubectl -n kube-system create serviceaccount tiller

#Create cluster role binding for tiller
$ kubectl create clusterrolebinding tiller --clusterrole cluster-admin --serviceaccount=kube-system:tiller

#Initialize tiller
$ helm init --service-account tiller

$ helm init --history-max 200 --service-account tiller
$ helm init
$ helm install --name prometheus -f helm/prometheus-customize-values.yaml stable/prometheus

