# ETIAssignment2
## API Endpoints:
| Microservice  | Port | Endpoint URL |
| ------------- | ---- | ------------ |
| Frontend  | 9030 | http://10.31.11.12:9030 |
| Tutor API  | 9091 | http://10.31.11.12:9031 |
| Testing API  | 9092 | http://10.31.11.12:9042 |

## 1. Design Consideration of your microservice
The microservices are loosely cooupled and working together. Therefore, after docomposing with stategic DDD, this packages rely heavily on other microservices to get the relevant data. This package is in charge of displaying all the relevant details to the tutor and updating the tutor's details. 
To ensure that the front-end would be able to display all the relevant information that can will be connected to other microservices. The back-end microservice would connect to the Moudle management, tutor managment, Class manaegement, Rating and comments and timetable microservice. This layer also acts as a layer to filter and relate the different class structure together before pushing it to the front-end to display. Lastly, this layer would act as an Anti-corruption layer as it would check the quality of the data before pushing it as JSON. This Microservice uses JSON, REST APIs to pass data from the backend to the front end.

## 2. Design Consideration of your microservice
![alt text](https://github.com/DicksonKuan/Assignment2_ETI/blob/main/diagram.png)

## 3. Link to container images
```
https://hub.docker.com/repository/docker/jiameelorh/eti-tutorfrontend   -    Front-end
https://hub.docker.com/repository/docker/jiameelorh/eti-tutorcontainer  -    Tutor's Microservice
```

## 4. Instructions for setting up and running your microservices
### 4.1. Update GIT 
1. cd S10192803 in server
2. git pull

### 4.2. Build image
1. docker-compose build
2. docker-compose up -d 

### 4.3. Access Microservice
1. http://10.31.11.12:9030/ in firefox

## Front End Webpage Routes
```sh
# 1. view tutor profile, Search tutor, view tutor mods, classes, enrolled students, Tutor's rating and comments, timetable
http://10.31.11.12:9030/index.html
# 2. View and update tutor's particular
http://10.31.11.12:9040/profile.html
# 3. View all tutor's rating and comment
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
```

### [GET] /api/v1/tutor/details
GET all Tutor
```
Endpoint:
http://10.31.11.12:9031/api/v1/tutor/details
Response:
Status code 200 if successful, else an error code with a corresponding status message will be returned if unsuccessful. 
```
## Sample of Comment jsonstring returned for GET request:
```
[
    {
        "TutorID":"1",
        "Name":"John Lee",
        "Email": "John@np.com",
        "Description": "Math professor"
    },
    {
        "TutorID":"2",
        "Name":"Susan Lee",
        "Email": "Susan@np.com",
        "Description": "Computing professor"
    }
]
```