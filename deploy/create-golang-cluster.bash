kubectl create -f cycling-namespace.yml
kubectl create -f cycling-db.yml
kubectl create -f cycling-app.yml
kubectl create -f cycling-service.yml
sleep 10
bash run_kube_jobs.bash
