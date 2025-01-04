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

Set the project:

```bash
gcloud config set project avionics-propulsion-pipeline
```

Before testing the functions, you need to set the environment variables for the functions:

```bash
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/storage-admin.json"
```
