# Sample project in GOLang - Rest API

<img style="float:left;" src="https://bit.ly/2HS3BiJ"/> <img style="float:left;" width="300" height="300" src="https://bit.ly/2S9c1qO"/>

## Description
In the context of HackDays for **TieFighters Team**, this project serves as a POC and sample project of a restful API written in GOLang. This serves also the purpose of introduction to a new language (GO).

## Tools / Languages Used
* Docker (https://www.docker.com/get-started)
* Docker Compose Specs (https://docs.docker.com/compose/)
* GoLang (https://golang.org/)
* MongoDB (https://www.mongodb.com/)
* Jetbrains GoLand (https://www.jetbrains.com/go/)

## The Objectives
* Introduction to GoLang
* Understand how to serve http requests in Go
* Define a basic project structure with packages
* Understand structs, types and vars in Go
* Introduction to a possible vendoring system, govendor (https://github.com/kardianos/govendor)
* Understand how to build and deploy a docker container with a Go Project

## Restful API Examples
### Get Tags
* GET http://localhost:8092/tag

###### Response
```json
[
  {
    "Id": 12,
    "Tag": "TestTAG2",
    "_id": "5c59998ea441d58cf318b50d"
  },
  {
    "Id": 13,
    "Tag": "TestTAG3",
    "_id": "5c59a2f0923b584830d26a72"
  },
  {
    "Id": 14,
    "Tag": "TestTAG4",
    "_id": "5c59a32d800d1e01dcbfbe5a"
  }
]
```

### Save a tag
* POST http://localhost:8092/tag

###### Body
```json
{
  "id": 10,
  "tag": "TestTAG"
}
```
###### Response
```json
Saved Successfully!
```