# ETIAssignment2
## API Endpoints:
| Microservice  | Port | Endpoint URL |
| ------------- | ---- | ------------ |
| Frontend  | 9030 | http://10.31.11.12:9030 |
| Tutor API  | 9091 | http://10.31.11.12:9031 |
| Testing API  | 9092 | http://10.31.11.12:9042 |

## Front End Webpage Routes
```sh
# 1. Comments and Ratings Dashboard
# (Comments and Ratings be posted and seen by first selecting target through this page)
http://10.31.11.12:9030/index.html

# 2. Specific student feedback
http://10.31.11.12:9040/profile.html
# 3. Personal posted and received ratings
http://10.31.11.12:9040/tutorsComments.html

```

## Documentation

### Tutor Data Structure
| Attribute | Data Type |
| --------- | ---- |
| TutorID | uint |
| Name | varchar(255) |
| Email | varchar(9) |
| Description | varchar(255) |

### Student Data Structure
| Attribute | Data Type |
| --------- | ---- |
| StudentID | varchar(255) |
| Name | varchar(255) |
| DateOfBirth | varchar(255) |
| Address | varchar(255) |
| Number | varchar(9) |

### Class Data Structure
| Attribute | Data Type |
| --------- | ---- |
| ClassID | uint |
| ModuleID | varchar(255) |
| ClassDate | varchar(255) |
| CLassStart | varchar(255) |
| ClassEnd | varchar(255) |
| Capacity | varchar(255) |
| TutorfName | varchar(255) |
| TutorID | uint |

### Module Data Structure
| Attribute | Data Type |
| --------- | ---- |
| Code | varchar(255) |
| Name | varchar(255) |
| LearningObjective | varchar(255) |
| Classes | List<Class> |
| AssignedTutor | varchar(255) |
| EnrolledStudent | List<Student> |

# Tutor Microservice API Documentation
### [GET] /api/v1/tutor
Test API if working
```
Endpoint
http://10.31.11.12:9031/api/v1/tutor
Response 
Status: Tutor API is working
```

### [GET] /api/v1/tutor/profile/{TutorID}
Get tutor by TutorID
```
Endpoint
http://10.31.11.12:9031/api/v1/tutor/profile/{TutorID}
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
Tutor 
```

## Sample of Comment jsonstring returned for GET request:
```
{
    "TutorID":"1",
    "Name":"John Lee",
    "Email": "John@np.com",
    "Description": "Math professor"
}
```
### [PUT] /api/v1/tutor/profile/{TutorID}
Update Tutor Account
```
Endpoint:
http://10.31.11.12:9031/api/v1/tutor/profile/{TutorID}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/v1/tutor/mod/{method}/{TutorID}
GET Mod assigned
```
Endpoint:
http://10.31.11.12:9031/api/comment/received/{"getMod"}/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```
## Sample of Comment jsonstring returned for GET request:
```
{
    "Code":"PRG1",
    "Name":"Programming 1",
    "LearningObjective": "Learn how to code in python",
    "Classes": [
        "ClassID": 1, 
        "ModuleID": "PRG", 
        "ClassDate": "25 Feb 2021", 
        "ClassStart": "8pm", 
        "ClassEnd": "9pm", 
        "Capacity": 50,
        "TutorfName": "John C maxwell",
        "TutorID": 1
    ],
    "AssginedTutor": "John C Maxwell",
    "EnrolledStudent": [
        "StudentID": "1",
        "Name": "Hormer",
        "DateOfBirth": "28 july 2000",
        "Address": "West vale",
        "Number": "91112321"
    ]
}
```
### [GET] /api/v1/tutor/mod/{method}/{TutorID}
GET Class Assigned by tutorID
```
Endpoint:
http://10.31.11.12:9031/api/v1/tutor/mod/"getClassAssigned"/{TutorID}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

## Sample of Comment jsonstring returned for GET request:
```
{
    "ClassID": 1, 
    "ModuleID": "PRG", 
    "ClassDate": "25 Feb 2021", 
    "ClassStart": "8pm", 
    "ClassEnd": "9pm", 
    "Capacity": 50,
    "TutorfName": "John C maxwell",
    "TutorID": 1
}


### [GET] /api/v1/tutor/mod/{method}/{TutorID}
Get user's posted comments
```
Endpoint:
http://10.31.11.12:9031/api/v1/tutor/mod/"enrolledStudent"/{TutorID}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```





### [GET] /api/comment/student/{id}
### [GET] /api/comment/tutor/{id}
### [GET] /api/comment/class/{id}
### [GET] /api/comment/module/{id}
Get comments for all students, tutor, classes or modules.
```
Endpoint:
http://10.31.11.12:9041/api/comment/student/{id}<br />
http://10.31.11.12:9041/api/comment/tutor/{id}<br />
http://10.31.11.12:9041/api/comment/class/{id}<br />
http://10.31.11.12:9041/api/comment/module/{id}<br />
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

## Sample of Comment jsonstring returned for GET request:
```
{
  "CommentID":1,
  "CreatorID":"1",
  "CreatorType":"Student",
  "TargetID":"2",
  "TargetType":"Student",
  "CommentData":"Great at working in teams.",
  "Anonymous":0,
  "DateTimePublished":"2022-02-04 04:14:06",
  "CreatorName":"Ethan",
  "TargetName":"Test"
 }
```

# Rating Microservice API Documentation

### [GET] /api/rating/{id}
Get Rating by RatingID
```
Endpoint:
http://10.31.11.12:9042/api/rating/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```


### [PUT] /api/rating/{id}
Update Rating 
```
Endpoint
http://10.31.11.12:9042/api/rating/{id}
Response
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [POST] /api/rating/create
Create Rating
```
Endpoint:
http://10.31.11.12:9042/api/rating/create
Response:
Status code 201 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```


### [GET] /api/rating/received/{type}/{id}
Get received ratings by user type and id
```
Endpoint:
http://10.31.11.12:9042/api/rating/received/{type}/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/rating/received/anonymous/{type}/{id}
Get received anonymous ratings by user type and id
```
Endpoint:
http://10.31.11.12:9042/api/rating/received/anonymous/{type}/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/rating/posted/{type}/{id}
Get poasted ratings by user type and id
```
Endpoint:
http://10.31.11.12:9042/api/rating/posted/{type}/{id}
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

### [GET] /api/rating/student/{id}
### [GET] /api/rating/tutor/{id}
### [GET] /api/rating/class/{id}
### [GET] /api/rating/module/{id}
Get ratings for all students, tutor, classes or modules.
```
Endpoint:
http://10.31.11.12:9042/api/rating/student/{id}<br />
http://10.31.11.12:9042/api/rating/tutor/{id}<br />
http://10.31.11.12:9042/api/rating/class/{id}<br />
http://10.31.11.12:9042/api/rating/module/{id}<br />
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```

## Sample of Rating jsonstring returned for GET request:
```
{"RatingID":1,
"CreatorID":"2",
"CreatorType":"Student",
"TargetID":"1",
"TargetType":"Student",
"RatingScore":3,
"Anonymous":1,
"DateTimePublished":"2022-02-04 04:14:06",
"CreatorName":"Kester",
"TargetName":"Ethan"}
```