kubectl apply -f ./services/database/k8s/userdb/userdb-config-map.yaml
kubectl apply -f ./services/database/k8s/userdb/userdb-v.yaml
kubectl apply -f ./services/database/k8s/userdb/userdb-d.yaml
kubectl apply -f ./services/database/k8s/userdb/userdb-s.yaml

kubectl apply -f ../../services/database/k8s/userdb/userdb-config-map.yaml
kubectl apply -f ../../services/database/k8s/userdb/userdb-v.yaml
kubectl apply -f ../../services/database/k8s/userdb/userdb-d.yaml
kubectl apply -f ../../services/database/k8s/userdb/userdb-s.yaml
