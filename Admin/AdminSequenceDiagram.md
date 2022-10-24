<h1 align="center">Admin Sequence Diagram</h1> 

This is the diagram that shows how admin use and control the website

### 1. Log in function:
**Basic flow:**
This describes the steps:
1. The administrator enters username and password
2. The system checks for existing username/password in the database
3. The system accepts username and password
4. The administrator successfully logged into the website
5. The system adds administration session

<img src=/export/1.Login.png>
<p align="center"> Admin Login diagram </p>

### 2. Log out function:
**Basic flow:**
This describes the steps:
1. The administrator chooses to log out
2. The system asks for sign-out confirmation
3. The system deletes the user session

<img src=/export/2.Logout.png>
<p align="center"> Logout diagram </p>

### 3. Management function:
**Basic flow:**
This describes the steps:
1. Administrator locates the user's account or the test information in the database
2. The administrator selects the functions to perform on the discovered account or discovered test
3. The system receives the query and transmits it to the database
4. The query is executed by the database
5. The database informs the system of the outcome
6. The system shows the results to the administrator

<img src=/export/3.Management.png>
<p align="center"> Management function diagram </p>

### 4. View the performance function: 
**Basic flow:**
This describes the steps:
1. The administrator views the performance data
2. The administrator chooses to view users' performance or tests' performance
3. The system gets the information from the database to show the user's performance
4. The administrator chooses to view the other performance
5. The system gets the information from the database to show tests' performance

<img src=/export/4.Viewtheperformance.png>
<p align="center"> View the performance function diagram </p>

### 5. Judgement function: 
**Basic flow:**
This describes the steps:
1. The administrator chooses to judge user
2. The administrator chooses to grade the test or comment on the test 
3. The system gets the information from the database to show the test 
4. The administrator inputs the score or comments on the test
5. The system saves the score or comments on the test to the database and notifies the information has been saved

<img src=/export/5.Judgement.png>
<p align="center"> Judgement function diagram </p>

### 6.Assign a test function: 
**Basic flow:**
This describes the steps:
1. The administrator assigns the test
2. The administrator creates the group user
3. The system gets the user data from the database and show the users' list
4. The administrator selects the user and adds it to the group
5. The system saves the group and notifies message to the selected user.
6. The administrator selects the test and selects the group 
7. The system assigns to the group
8. The system saves the group and notifies message to the selected user 

<img src=/export/6.AssignATest.png>
<p align="center"> Assign A Test function diagram </p>