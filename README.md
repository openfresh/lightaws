lightaws
==========

[![Circle CI](https://circleci.com/gh/openfresh/lightaws.svg?style=shield&circle-token=6aafa4f2691fc1b6cebc6efc62d50d03d834ad36)](https://circleci.com/gh/openfresh/lightaws)
[![Language](http://img.shields.io/badge/language-go-brightgreen.svg?style=flat)](https://golang.org/)
[![issues](https://img.shields.io/github/issues/openfresh/lightaws.svg?style=flat)](https://github.com/openfresh/lightaws/issues?state=open)
[![License: MIT](http://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)

lightaws is a aws command line tool written in Golang.

## Motivation

[aws-cli](https://github.com/aws/aws-cli) covers almost all operations of AWS, but this requires python. Then it takes time to install by homebrew, and pip.
Today, we offen use aws-cli in some continuous integration service: (e.g.: TravisCI, CircleCI) 

It's a waste of time to install awscli every time. Further, this depends on platform. We believe that such tools should be provided in binary.

## install

```
$ go get github.com/openfresh/lightaws
```

## Usage

### EC2 Container Registry (ECR)

Command line interface of lightaws is same as aws-cli. For example, in Amazon ECR as follows

```
$ lightaws ecr get-login --region ap-northeast-1
```

### EC2 Metadata

```
$ lightaws metadata region 
ap-northeast-1
```

```
$ lightaws metadata instance-id 
i-xxxxxxxxxxxxxxxx
```

## For contribution

lightaws used [cobra](https://github.com/spf13/cobra). You only have to add new command in `cmd` package.

## Roadmap

We don't think lightaws should cover all operations of AWS. lightaws will add necessary features each time. Your contribution is welcome!
