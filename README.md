# allocate-memory

Allocate memory is web application supposed to dynamically allocate memory according to web request.

## web API
There is API provided:

|URI Path|Description|
| :-------------: | :--------------: |
|/| return "OK"|
|/ping|return "pong"|
|/_ping|return "pong"|
|/memory| allocate 256KB memery and release it after 10s|
|/memory/:size/action/allocate|allocate size MB meory|
|/cpu|consume as much CPU as it can|


## Listen Port

This application will listen on port 8080.

## Participating

You can contribute to this project in several different ways:

* To report a problem or request a feature, please feel free to file an issue.

* Of course, we welcome pull requests and patches. Setting up a local this project development environment and submitting PRs is described here.


## Copyright and license
Copyright © 2016. All rights reserved