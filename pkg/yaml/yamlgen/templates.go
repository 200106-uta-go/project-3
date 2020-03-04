package yamlgen

const deployment = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Name -}}-deployment
  labels:
    app: {{ .Name }}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{ .Name }}
  template:
    metadata:
      labels:
        app: {{ .Name }}
    spec:
      containers:
      - name: {{ .Name }}
      image: {{ .Image }}
      resources:
        limits:
          memory: "128Mi"
          cpu: "500m"
      ports:
		  {{- range $port := .ExposedPorts }}
        - containerPort: {{$port -}}
      {{- end }}`