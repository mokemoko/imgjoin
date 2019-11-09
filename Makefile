deploy:
	gcloud builds submit --config cloudbuild.yml
	gcloud beta run deploy imgjoin --image gcr.io/rls-pdu-sandbox/imgjoin --platform managed
