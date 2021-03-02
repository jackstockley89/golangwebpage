kubectl create -f cycling-namespace.yml
kubectl create -f cycling-db.yml
kubectl create -f cycling-app.yml
kubectl create -f cycling-app-service.yml
bash run_kube_jobs.bash
