#!/bin/bash

kubectl apply -f db-create-job.yml
sleep 10
kubectl apply -f db-insert-job.yml
sleep 10
kubectl -n golang-app delete job cycling-blog-app-create
kubectl -n golang-app delete job cycling-blog-app-insert
