# Avionics Propulsion Pipeline

> Automation of the post-processing of propulsion test data and the online archive for test results.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Svelte](https://img.shields.io/badge/svelte-%23f1413d.svg?style=for-the-badge&logo=svelte&logoColor=white)
![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

## To-Do List

### Test Analysis

- [ ] Detect test start and end based on:
  - Spike in chamber pressure.
  - Decrease in oxidizer pressure.
  - Decrease in oxidizer mass.
- [ ] Allow manual input for oxidizer shutoff time.
- [ ] Define test end when chamber pressure returns to normal.

### Data Handling

- [ ] Provide a way to download the filtered data.
- [ ] Display both unfiltered and filtered lines on the graph.
- [ ] Add a legend to the Plotly graph to toggle lines on/off.

### Graph Features

- [ ] Implement two data smoothing filters that do not distort the data.
- [ ] Add a button to enable full-screen mode for the graph.
- [ ] Display mass flow rate on the graph.
- [ ] (Optional) Plot fill mass and pressure over time.

### Website Features

- [ ] Integrate a home page into the application.
- [ ] Implement a file organization system:
  - [ ] Use a custom file extension or metadata to categorize files.
  - [ ] Enable filtering of files for different website sections.

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
