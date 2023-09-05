
# Cloudraft Key-Value Store

Cloudraft Key-Value Store is a simple and efficient key-value store service built using Go and designed for Kubernetes deployment. This repository provides the source code, Dockerfile, and Kubernetes manifests to deploy the service on your Kubernetes cluster.

## Table of Contents 
### Features

- Simple RESTful API for key-value operations.
- Support for setting, getting, and searching for keys.
- Prometheus metrics endpoint for monitoring.
- Scalable and containerized for Kubernetes deployment.

### Prerequisites

- Before you begin, ensure you have the following prerequisites:
- Go (>=1.14) installed on your local development machine.
- Docker installed to build the Docker image.
- Kubernetes cluster configured and accessible.
- kubectl command-line tool installed for interacting with your Kubernetes cluster.

### Run the services and tests locally

- clone the repository 
   https://github.com/RkReddy11/Cloudraft-KV-Store.git
- Run the service and tests locally
   - go run kvs.go 
   - go test
### Usage
The Cloudraft Key-Value Store provides the following endpoints:

- GET /get/key: Retrieve the value associated with a given key.
- POST /set: Set a key-value pair.
- GET /search?prefix=<prefix>&suffix=<suffix>: Search for keys with a specific prefix or suffix.
- GET /metrics: Prometheus metrics endpoint for monitoring

### Build the Docker image,test it and push it
- docker build -t customkvs:latest .
- docker run --name kvs -p 8000:8000 customkvs:latest
- docker push [docker-hub-username]/customkvs:latest [tag it according to your wish]

### Deploy customkvs on Kubernetes

- To create deployment, service and ingress.

  - kubectl apply -f k8s-manifest/
    
### Implement or describe how you will put Observability for the customkvs service?

To achieve observability for our cloud-native customkvs service, we can use a combination of essential tools and practices. Firstly, we implement monitoring using Prometheus for collecting and storing metrics and Grafana for visualizing them. Next, we enable tracing with Jaeger or OpenTelemetry to monitor request flows. We implement centralized logging using Elasticsearch, Fluentd, and Kibana (EFK/ELK Stack) to capture and analyze logs, including error reporting using Sentry or Rollbar. We integrate our service with Istio or Linkerd for service mesh capabilities, including telemetry and security. We ensure observability checks are part of our CI/CD pipeline. For security observability, we consider tools like Falco. We monitor and optimize cloud costs using cloud-native cost management tools. We document observability practices, provide training, and establish an incident response process using observability data. We store observability data for long-term analysis, and we continuously improve our observability strategy by staying updated with the latest tools and best practices. This holistic approach prioritizes the health and performance of our customkvs service in a cloud-native environment.

### What other tests you would perform on the service before releasing to production and how?

#### Set Operation Test:

##### Functionality: Verify that the service correctly sets key-value pairs.

- Set a new key-value pair.
- Update an existing key with a new value.
- Set a key with an empty value.
- Set a key with special characters in the key and value.
- Test with different data types for keys and values (e.g., string, integer).

#### Get Operation Test:

Functionality: Ensure that the service retrieves values correctly based on the provided key.

- Retrieving an existing key.
- Attempting to retrieve a non-existent key.
- Handling special characters in keys.
- Testing concurrent get operations.

#### Search Operation Test:

Functionality: Confirm that the search operation returns expected results based on prefix and suffix.

- Searching for keys with a given prefix.
- Searching for keys with a given suffix.
- Combining prefix and suffix searches.
- Handling cases where no keys match the search criteria.

#### Error Handling Test:

Functionality: Ensure that the service handles errors gracefully.

- Setting a key with an invalid format.
- Retrieving a key that exceeds the maximum length.
- Handling unexpected server errors and returning appropriate error responses (e.g., 4xx and 5xx status codes).

#### Concurrency Test:

Functionality: Validate that the service can handle concurrent requests without data corruption or race conditions.

- How: Create test cases that simulate multiple clients concurrently setting and getting keys. Ensure data consistency and correctness under high concurrency.

#### Usability Test:

Functionality: Evaluate the service's ease of use from an end-user perspective.

- How: Involve users or stakeholders in testing the service's functionality to ensure it meets their requirements and expectations.

#### User Acceptance Testing (UAT):

Functionality: Conduct user acceptance testing with real users to validate that the service aligns with their needs.

- How: Collaborate with users to define test cases and scenarios. Allow them to interact with the service and gather feedback

By performing these tests along with whole lot of other tests(Scalability, security, load etc), we can ensure that our customkvs service's core functionalities are robust, reliable, and ready for production release.














  

  


