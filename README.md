lightaws
==========

lightaws is a aws command line tool written in Golang.

### Motivation

[aws-cli](https://github.com/aws/aws-cli) covers almost all operations of AWS, but this requires python. Then it takes time to install by homebrew, and pip.
Today, we offen use aws-cli in some continuous integration service: (e.g.: TravisCI, CircleCI) 

It's a waste of time to install awscli every time. Further, this depends on platform. We believe that such tools should be provided in binary.

### install

```
$ go get github.com/openfresh/lightaws
```

### Usage

Command line interface of lightaws is same as aws-cli. For example, in Amazon ECR as follows

```
$ lightaws ecr get-login --region ap-northeast-1
```

### For contribution

lightaws used [cobra](https://github.com/spf13/cobra). You only have to add new command in `cmd` package.

### Roadmap

We don't think lightaws should cover all operations of AWS. lightaws will add necessary features each time. Your contribution is welcome!