# Go REST API Example
I'm working on learning Go and this creates a really basic API using the Gin library. The API has endpoints to create and retrieve details about residents in my building based on their unit. The backend is Google Cloud Firestore.


# Running
Before running, ensure that you're authenticated to Google Cloud and you've set the following env vars:
`GCP_PROJECT` -> Your Google Cloud project id
`COLLECTION_NAME` -> The name of the Firestore collection you want to use for storing data

