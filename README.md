# ARGOS

## Table of Contents

- [Summary](#summary)
- [Prerequisites](#prerequisites)
  - [Required](#required)
- [Development](#development)
  - [How to run](#how-to-run)
- [Deployment](#deployment)
  - [Compile](#compile)
  - [Docker](#docker)
- [Compile](#compile)
- [Built With](#built-with)
- [Authors](#authors)
- [License](#license)

## Summary

### Prerequisites

#### Required

The following tools are required for anyone who wishes to compile or run Argos,
regardless of the mode. Please ensure you meet the requirements below:

- [Golang compiler](https://golang.org/)

## Development

This section provides a guide on getting the project up and running for development. You’ll
find a brief and straightforward explanation on managing configuration files and setting up
MongoDB on your local machine. However, if possible, we recommend connecting the project
to MongoDB Atlas for a faster setup process.

### How to run

Before running the project, ensure you have met all the [prerequisites](#prerequisites).
Once done, you can start the project by running `go run cmd/main.go <central system url>`.
Be sure to read the terminal output to confirm everything is initialized properly before
proceeding. You can refer to the example below.

``` bash
# Bash Shell
# You can run one of either and the
# result will be the same

user@domain:~$ go run main.go <central system url>
```

## Deployment

### Compile

### Docker

Before running Argos make sure that you have read [Compile](#compile)
as it will contain important information regarding deployment environments.
If you have read the section and know where you will be shipping the code
then there is not much to it. We have pasted a easy to follow example bellow,
and do make sure that you have installed Docker before building the image.

```bash
# Bash Shell
# This section explains how to build a docker image that you can use for
# testing or deployment.
# make sure to import the right env file when you are deploying
user@domain:~$ docker build  -f docker/dockerfile -t <image name> .
```

### Built With

- [Golang](https://golang.org/)
- [Docker](https://docker.com)
- [Editor config](https://editorconfig.org/)
- [GIT](https://git-scm.com/)
- [Visual Studio Code](https://code.visualstudio.com/)

### Authors

- **Mcs Unity** - _Initial work_ - [Mcs unity](https://github.com/mcs-unity)

See also the list of [contributors](https://github.com/mcs-unity/argos/graphs/contributors)
who participated in this project.

### License

This project is licensed under the GNU GENERAL PUBLIC V3 - see the [LICENSE](LICENSE) file for details