# Avionics Propulsion Pipeline

> Automation of the post-processing of propulsion test data and the online archive for test results.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Svelte](https://img.shields.io/badge/svelte-%23f1413d.svg?style=for-the-badge&logo=svelte&logoColor=white)
![OAuth2](https://img.shields.io/badge/oauth2-4285F4?style=for-the-badge&logo=google&logoColor=white)
![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

## To-Do List

### Test Analysis

- [ ] Detect test start and end based on:
  - Spike in chamber pressure.
  - Decrease in oxidizer pressure.
  - Decrease in oxidizer mass.
- [x] Allow manual input for test start and end.
- [ ] Define test end when chamber pressure returns to normal.

### Data Handling

- [ ] Provide a way to download the filtered data. (Needs review)
- [ ] Display both unfiltered and filtered lines on the graph.
- [x] Add a legend to the Plotly graph to toggle lines on/off.

### Graph Features

- [x] Implement two data smoothing filters that do not distort the data.
- [ ] Integrate the data smoothing filters into the graph.
- [x] Add a button to enable full-screen mode for the graph.
- [ ] Display mass flow rate on the graph. (Needs review)
- [ ] (Optional) Plot fill mass and pressure over time.

### Website Features

- [x] Integrate a home page into the application.
- [ ] Implement a file organization system:
  - [x] Use a custom file extension or metadata to categorize files.
  - [ ] Enable filtering of files for different website sections.

## Example `.env.toml` file

```toml
# API keys.
google_client_id="..."
google_client_secret="..."

# This is used to sign authorization tokens. It should be a long random string.
signing_key="..."

# Set this to true in a deployment environment.
in_production=false

# Place GMail addresses that are allowed to use the platform here.
whitelist = ["example@gmail.com", "example2@gmail.com"]

# Development configuration.
[dev]
host = "http://localhost"
port = "8080"
allowedorigins = ["http://localhost:5173"]

# Production configuration.
[prod]
host = "https://api.soarpipeline.com"
port = "8080"
allowedorigins = ["https://soarpipeline.com", "https://api.soarpipeline.com"]
```

## Running the backend

```bash
go run cmd/soarpipeline/soarpipeline.go
```

## Running with Docker

Simply run the following command to start the application.

```bash
docker-compose -f docker-compose.<dev | prod>.yaml up
```

or

```bash
docker-compose -f docker-compose.<dev | prod>.yaml up --build
```

to rebuild the containers from scratch.

## Stopping the Application

To stop the application and remove the containers, run:

```bash
docker-compose -f docker-compose.<dev | prod>.yaml down
```

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
cd
