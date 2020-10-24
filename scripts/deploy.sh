VERSION=$1
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: snarecc-pages
  labels:
    app: snarecc-pages
spec:
  replicas: 1
  selector:
    matchLabels:
      app: snarecc-pages
  template:
    metadata:
      labels:
        app: snarecc-pages
    spec:
      containers:
      - name: snarecc-pages
        image: snarecc/pages:$VERSION
        ports:
        - containerPort: 5000
---
apiVersion: v1
kind: Service
metadata:
  name: snarecc-pages
spec:
  type: NodePort
  ports:
  - port: 8080
    targetPort: 5000
  selector:
    app: snarecc-pages
EOF
