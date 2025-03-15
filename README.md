# Avionics Propulsion Pipeline

> Automation of the post-processing of propulsion test data and the online archive for test results.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Vue.js](https://img.shields.io/badge/vuejs-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D)
![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

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