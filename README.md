我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# Slacking playground
A command line script that (given a webhook), reads messages from a file, and posts them to slack, written in
Python3, Go, and SCALA


## Setup

First setup an incoming webhook in slack and grab the url, (see http://bit.ly/2EapumJ for detailed instructions)

## Usage

### Python

#### Testing
    ./python/test.py

#### Running
    ./python/send.py --help
    ./python/send.py --webhook https://hooks.slack.com/services/XXXXX/XXXXX

### Go
    cd go
    go run send.go --help
    go run send.go --webhook https://hooks.slack.com/services/XXXXX/XXXXX

### SCALA
    cd scala
    ./send.sh https://hooks.slack.com/services/XXXXX/XXXXX

### Futrue Enhancements

- Add tests for Go Implementation
- Allow Go script to run from anywhere (fix relative path lookup)
- Add tests for SCALA Implementation
- Tidy up SCALA implementation, extend App class
