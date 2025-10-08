## Pod - K8s

É o menor, uma caixinha onde o container fica rodando.

- Criar o primeiro pod.

 - `kind delete cluster --name fullcycle`   
 - `kind create cluster --name fullcycle`
 - `kubectl get pods`
 - `kubectl apply -f k8s/pod.yaml`
 - `kubectl delete pod goserver`