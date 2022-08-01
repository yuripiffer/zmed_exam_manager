# microservice zmed_exam_manager
This is a golang microservice study case. 
It makes part of a software architecture with microservices and AWS infrastructure prepared to handle exams for the fictional "ZMED" company.

## ZMED exams's architecture diagram
![](zmed_diagram.png)
#### microservices
 - **zmed_exam_manager** (Golang): manages exams registry in a dynamoDB table.
 - **zmed_patient_manager** (Golang): manages patients registry in a RDS(Postgres).

#### storage
- **zmed_exam_table** (DynamoDB): Stores exam's registry.
- **zmed_patient_table** (RDS): Stores patient's registry.
- **zmed_exams_results** (S3): Stores exam's results.


## Endpoints
### Register a new exam request
- path: "/exam/new"
- method: Post
- headers: none
- params: none
- body: 
  - document (string): patient's document
  - exam_type (int): the exam code

### Get exams from a patient
- path: "/exams/info"
- method: Get
- headers: none
- params: 
  - document (string): patient's document

### Register that an exam has started
- path: "/exam/start"
- method: Post
- headers: none
- params: nome
- body: 
  - document (string): patient's document
  - exam_id (string): uuid id from the exam
  - exam_type (int): the exam code


### Send a text message or an email to a patient that his/her exam is ready.
- path: "/exams/communicate"
- method: Post
- headers: none

## Parallel processing

## Installation


### Environment Variables


## License
This is an open-source and free software release.