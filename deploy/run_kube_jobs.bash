#!/bin/bash

kubectl apply -f db-create-job.yml
sleep 10
kubectl apply -f db-insert-job.yml
