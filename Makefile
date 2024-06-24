export GOOGLE_APPLICATION_CREDENTIALS="./paotang-core-sit-d50437475f5b.json"
export PROJECT_ID="paotang-core-sit"

activate-gcloud-sit:
	gcloud auth activate-service-account --key-file=./paotang-core-sit-d50437475f5b.json

generate-token:
	export TOKEN=$(gcloud auth print-identity-token)
