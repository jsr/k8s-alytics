# k8s-alytics
Scans github for statistics on usage of Kubernetes resource objects 

# Getting started 

1. Create a [Personal API token on github](https://github.com/blog/1509-personal-api-tokens) (this is needed because you must be logged in to use the Search API)
2. Store your token in ~/.github_token
2. Install this tool with `go get github.com/jsr/k8s-alytics`
3. Run `$ k8s-alytics` 

It will output lines of CSV tuples (resource-name, count, url of details)

# Usage 


```shell
jsr in ~/go/src/github.com/jsr/k8s-alytics on master*
⚡ ls ~/.github_token
/Users/jsr/.github_token
jsr in ~/go/src/github.com/jsr/k8s-alytics on master*
⚡ ./k8s-alytics
DaemonSet, 3840, https://github.com/search?q=%22Kind%3A+DaemonSet%22+language%3Ayaml+language%3Ajson&type=Code
ExternalAdmissionHookConfiguration, 58, https://github.com/search?q=%22Kind%3A+ExternalAdmissionHookConfiguration%22+language%3Ayaml+language%3Ajson&type=Code
Volume, 1, https://github.com/search?q=%22apiVersion%3A+v1%22+%22Kind%3A+Volume%22+language%3Ayaml+language%3Ajson&type=Code
^C
```
