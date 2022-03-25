apiVersion: apps/v1
kind: Deployment
metadata:
  name: cycling-blog-app
  namespace: jacksapp-dev
  labels:
    app: cycling-blog-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cycling-blog-app
  template:
    metadata:
      labels:
        app: cycling-blog-app
    spec:
      containers:
      - name: cycling-blog-app
        image: ${ECR_URL}:${IMAGE_TAG}
        imagePullPolicy: Always
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "128Mi"
            cpu: "500m"
        ports:
        - containerPort: 8080
        env:
        - name: PGPASSFILE
          valueFrom: 
            secretKeyRef:
              name: pgpass
              key: connect
        - name: POSTGRES_USER
          valueFrom:
            secretKeyRef:
              name: pgpass
              key: username
        - name: PGPASSWORD
          valueFrom:
            secretKeyRef:
              name: pgpass
              key: password
        - name: POSTGRES_DB
          valueFrom:
            secretKeyRef:
              name: pgpass
              key: db_name
          