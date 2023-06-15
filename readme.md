# Zin Mail

Zin Mail is a full-stack project built using Vue3 on the frontend, Go on the backend, and the ZincSearch database. It allows users to search through a large dataset of 517k emails, providing an efficient and responsive experience.

## Overview

The project took over a week to complete and was a valuable learning experience. It covers various aspects of web development, from frontend and backend to working with databases and containerization.

~~The entire application is hosted on AWS, including the frontend, backend, database container, and the email dataset.~~

Unfortunately, AWS isn't free, so I had to take the project down. You can follow the installation instructions or watch a video demonstration of the app below.

## How to Install

### Frontend - Vue3

The frontend of Zin Mail is built using Vue3 with the Composition API, offering a more maintainable and scalable codebase. The application is styled with Tailwind CSS and focuses on desktop screen sizes (best viewed at a width of around 1000px). The frontend handles pagination and allows users to view the full email content with a single click on the view icon.

### Backend - Go/Go-chi

The backend is built using Go and the Go-chi library. It serves as a service layer that communicates with the ZincSearch database, retrieves data, filters relevant information, and sends it back to the frontend. A profiler (pprof) is also available to analyze the backend's performance and identify areas for improvement.

### ZincSearch

ZincSearch is an alternative to ElasticSearch, built using Go. The official website can be found here: [ZincSearch](https://zinc.dev/). To ingest the email dataset into ZincSearch, a Go script was created (located in the data folder). After ingesting the data, the backend communicates with ZincSearch through its endpoints to retrieve the necessary data.

## Optimizations

Checking the validation of JSON can be quite expensive in terms of performance. One potential optimization would be to avoid unmarshaling and copying the data directly to an interface. However, this method might not be safe.

Here are some performance metrics for searches with a maximum of 200 results, highlighting enabled, and 14.5k "hey" sent to the API:

| Word                      | Value (Return) | Runtime |
| ------------------------- | -------------: | ------: |
| 'example'                 |         11,288 |   20 ms |
| 'mutation'                |              9 |    3 ms |
| 'mutation, where are you' |              3 |    3 ms |
| 'how are you'             |          1,818 |  250 ms |
| 'how are'                 |          3,388 |  150 ms |
