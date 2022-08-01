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
