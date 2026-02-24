.PHONY: vector-install vector-uninstall

vector-uninstall:
	helm uninstall vector --namespace vector

vector-install:
	helm install vector vector/vector --namespace vector --create-namespace --values vector-values.yaml
