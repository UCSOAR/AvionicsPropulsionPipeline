# Avionics Propulsion Pipeline

> Automation of the post-processing of propulsion test data and the online archive for test results.

## What we need

A cloud platform that provides file/object storage, basic compute, and event-driven functions.
At the moment, a database is not projected to be required for this application.

## Possible candidates

![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon&logoColor=white)
![Google Cloud](https://img.shields.io/badge/GoogleCloud-%234285F4.svg?style=for-the-badge&logo=google-cloud&logoColor=white)
![Supabase](https://img.shields.io/badge/Supabase-3ECF8E?style=for-the-badge&logo=supabase&logoColor=white)

### AWS

![NodeJS](https://img.shields.io/badge/node.js-6DA55F?style=for-the-badge&logo=node.js&logoColor=white)
![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Java](https://img.shields.io/badge/java-%23ED8B00.svg?style=for-the-badge&logo=openjdk&logoColor=white)
![C#](https://img.shields.io/badge/c%23-%23239120.svg?style=for-the-badge&logo=csharp&logoColor=white)
![Ruby](https://img.shields.io/badge/ruby-%23CC342D.svg?style=for-the-badge&logo=ruby&logoColor=white)
![PHP](https://img.shields.io/badge/php-%23777BB4.svg?style=for-the-badge&logo=php&logoColor=white)
![PowerShell](https://img.shields.io/badge/PowerShell-%235391FE.svg?style=for-the-badge&logo=powershell&logoColor=white)

- Support for file storage with **S3**. ☑️
- Support for triggering functions on file upload event with **Lamda** functions. ☑️
- **Lambda free tier** (refreshes monthly)
  - 1 million free requests.
  - 400 000 GB-seconds or 3.2 million seconds of compute time.
- **S3 free tier** (12 months free, refreshes monthly)
  - 5 GB in S3 Standard storage class.
  - 20 000 `GET` requests.
  - 2000 `PUT`, `COPY`, `POST`, or `LIST` requests.

### GCP

![NodeJS](https://img.shields.io/badge/node.js-6DA55F?style=for-the-badge&logo=node.js&logoColor=white)
![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Java](https://img.shields.io/badge/java-%23ED8B00.svg?style=for-the-badge&logo=openjdk&logoColor=white)
![C#](https://img.shields.io/badge/c%23-%23239120.svg?style=for-the-badge&logo=csharp&logoColor=white)
![Ruby](https://img.shields.io/badge/ruby-%23CC342D.svg?style=for-the-badge&logo=ruby&logoColor=white)
![PHP](https://img.shields.io/badge/php-%23777BB4.svg?style=for-the-badge&logo=php&logoColor=white)

- Support for file storage with **Cloud Storage**. ☑️
- Support for triggering functions on file upload event with **Cloud Functions**. ☑️
- **Cloud functions free tier** (per month)
  - 2 million total invocations.
  - 400 000 GB-seconds, 200 000 GHz-seconds of compute time.
  - 5 GB of outbound data transfer.
- **Cloud storage**
  - See <a href="https://cloud.google.com/storage?_gl=1*xbsxwu*_up*MQ..&gclid=Cj0KCQiAgdC6BhCgARIsAPWNWH1j71W_CebJk7Pk1C9a4ZV-6Eou-4d87-X6XkKMoFi9fVzs6Lf6NxwaAqgKEALw_wcB&gclsrc=aw.ds&hl=en">here</a>.

### Supabase

![Deno JS](https://img.shields.io/badge/deno%20js-000000?style=for-the-badge&logo=deno&logoColor=white)

- Support for file storage. ☑️
- Support for triggering functions on file upload event with **Edge Functions**. ☑️
- Edge functions only support the **Deno** JS runtime 😕
- Highly tied to PostgreSQL database. Not needed for our purposes.
- **Free tier** (per month)
  - 1 GB file storage
  - 500 000 edge function invocations
  - 20 MB script size
  - 25 total edge functions allowed

## Comparison table

| Feature                          | AWS                                                          | GCP                                                        | Supabase                                                 |
| -------------------------------- | ------------------------------------------------------------ | ---------------------------------------------------------- | -------------------------------------------------------- |
| File Storage                     | S3                                                           | Cloud Storage                                              | File storage                                             |
| Function Triggers on File Upload | Lambda functions                                             | Cloud Functions                                            | Edge Functions                                           |
| Free Tier                        | 12 months free, refreshes monthly                            | Monthly free tier                                          | Monthly free tier                                        |
| Free Tier Details                | - 5 GB in S3 Standard storage class.                         | - 2 million total invocations.                             | - 1 GB file storage                                      |
|                                  | - 20,000 `GET` requests.                                     | - 400,000 GB-seconds, 200,000 GHz-seconds of compute time. | - 500,000 edge function invocations                      |
|                                  | - 2,000 `PUT`, `COPY`, `POST`, or `LIST` requests.           | - 5 GB of outbound data transfer.                          | - 20 MB script size                                      |
|                                  | - 1 million free Lambda requests.                            |                                                            | - 25 total edge functions allowed                        |
|                                  | - 400,000 GB-seconds or 3.2 million seconds of compute time. |                                                            |                                                          |
| Long-term Storage Needs          | Suitable for long-term storage                               | Suitable for long-term storage                             | Limited by free tier, additional storage may incur costs |
