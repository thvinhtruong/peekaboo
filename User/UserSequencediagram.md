<h1 align="center">User Sequence Diagram</h1> 

This is the diagram that shows how user use website

### 1. Registration function:

**Basic flow:**
This describes the steps:

1. The user attempts to login
2. The user fills in the registration form
3. The user submits the registration form
4. The system checks the information
5. The system sends authentication code to the user
6. The user fills in the code
7. The account has been created

<img src=/export/1.Registration.png>
<p align="center"> Registration diagram </p>

### 2. Log in function:
**Basic flow:**
This describes the steps:
1. The user enters username and password
2. The system check for existing username/password in the database
3. The system accepts username and password
4. The user successfully logged into the website
5. The system adds user session 
6. The system checks previous session, then delete it

<img src=/export/2.Login.png>
<p align="center"> Login diagram </p>

### 3. Log out function:
**Basic flow:**
This describes the steps:
1. The user chooses to log out
2. The system asks for sign-out confirmation
3. The system deletes the user session

<img src=/export/3.Logout.png>
<p align="center"> Logout diagram </p>

### 4. Edit profile function:
**Basic flow:**
This describes the steps:
1. The user selects Edit profile function in the settings
2. The system changes to the Edit profile function
3. The user adjusts the information
4. The system sends authentication code to the user
5. The user enters the code
6. The system verifies the code
7. The information has successfully been updated

<img src=/export/4.EditProfile.png>
<p align="center"> Logout diagram </p>

### 5. Choose the test function:
**Basic flow:**
This describes the steps:
1. The user chooses the test
2. The system notifies user can choose test's types
3. The system notifies the choosen test

<img src=/export/5.Choosethetest.png>
<p align="center"> Choose the test function diagram</p>

### 6. Choose the Basic test function:
**Basic flow:**
This describes the steps:
1. The user starts to do the test
2. The user answers the question
3. The user submits the test
4. The system confirms to submit the test
5. The system checks the answer and shows the results

<img src=/export/6.BasicTest.png>
<p align="center"> Choose the Basic test function diagram </p>

### 7. Choose the IELTS test function:
**Basic flow:**
This describes the steps:
1. The user starts to do the test
2. The user answers the question
3. The user submits the test
4. The system confirms to submits the test
5. The system checks the answer and shows the results

<img src=/export/7.IELTStest.png>
<p align="center"> Choose the IELTS test function diagram </p>

### 8. Review function:
**Basic flow:**
This describes the steps:
1. The user wants to review the test
2. The user chooses to view wrong details in the previous test
3. The system gets the information from the database to show the results
4. The user chooses to view the comment by the administrator
5. The system gets the information from the database to show the comments

<img src=/export/8.Review.png>
<p align="center"> Review function diagram </p>

### 9. View the progress function:
**Basic flow:**
This describes the steps:
1. The user views the progress
2. The user chooses to view personal performance
3. The system gets the information from the database to show the user's performance
4. The user chooses to view the other performance
5. The system gets the information from the database to show others' performance

<img src=/export/9.Viewtheprogress.png>
<p align="center"> View the performance function diagram </p>
