# toronto_time
Setting up the Environment:
1. Make sure MySQL and Go is installed.
2. Install Go Dependencies such as sql-driver
   
Creating a Database and table for storing the data:
Queries Used:
1.	To create DB: CREATE DATABASE go_api;
2.	Switch to the newly created DB: USE go_api;
3.	Create new table: CREATE TABLE time_log (id INT AUTO_INCREMENT PRIMARY KEY, timestamp DATETIME NOT NULL);
   
 ![image](https://github.com/user-attachments/assets/c1b3c4aa-d3dd-4e41-be37-9fe49739e055)

Running the Go Application:

 ![image](https://github.com/user-attachments/assets/4e87776b-d5e9-4b3e-bb20-d1287a7f3a9b)


Visiting the API endpoint: localhost:8080/current-time

 ![image](https://github.com/user-attachments/assets/c88fe00a-9bd6-404d-b6de-1f75a96e28a8)


Verifying if the data has been stored in the DB:
Query: SELECT * FROM time_log;
 
![image](https://github.com/user-attachments/assets/abcb2910-c9be-463b-9ac9-69baaf643f49)
