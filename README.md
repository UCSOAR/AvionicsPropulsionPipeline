# Avionics Propulsion Pipeline

> Automation of the post-processing of propulsion test data and the online archive for test results.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Vue.js](https://img.shields.io/badge/vuejs-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D)
![Google Cloud](https://img.shields.io/badge/GoogleCloud-%234285F4.svg?style=for-the-badge&logo=google-cloud&logoColor=white)

## Development environment

To sign into GCP use the CLI command:

```bash
gcloud auth login
```

Enable API:

```bash
gcloud services enable cloudfunctions.googleapis.com
```

Set the project:

```bash
gcloud config set project avionics-propulsion-pipeline
```

Before testing the functions, you need to set the environment variables for the functions:

```bash
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/storage-admin.json"
```

Deploying a function:

```bash
gcloud functions deploy <FunctionName> \
  --runtime go123 \
  --trigger-http \
  --allow-unauthenticated \
  --region us-west1
```

It is important that `<FunctionName>` is the name of the function you want to deploy and must be part of a package called `function` in a file called `main.go`.

## LabVIEW Measurement to Cache Tree Diagram

A **cache tree** is represented as follows for multiple X columns.

```plaintext
test.lvm
  |
  v
  test/
  ├─ PreviewMetadata
  ├─ x/
  │  ├─ (X) someYColumnName
  │  ├─ (X) anotherYColumnName
  │  └─ (X) yetAnotherYColumnName
  └─ y/
     ├─ someYColumnName
     ├─ anotherYColumnName
     └─ yetAnotherYColumnName
```

_Or_, if there is only one X column.

```plaintext
test.lvm
  |
  v
  test/
  ├─ PreviewMetadata
  ├─ x/
  │  └─ X_Value
  └─ y/
     ├─ someYColumnName
     ├─ anotherYColumnName
     └─ yetAnotherYColumnName
```

Where each of these files that is not a directory is a binary data containing the data for the corresponding column or metadata. With this structure, only the necessary data is loaded into memory, and the data is only loaded when it is requested.
