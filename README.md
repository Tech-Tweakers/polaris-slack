<h1 align="center">Tech Tweakers - Polaris Slack Chat API v1 </h1>
<p align="center"><i>Polaris Chatbot meets Slack with our seamless API integration.</i></p>

<div align="center">
  <a href="https://github.com/Tech-Tweakers/polaris-slack/stargazers"><img src="https://img.shields.io/github/stars/Tech-Tweakers/polaris-slack" alt="Stars Badge"/></a>
<a href="https://github.com/Tech-Tweakers/polaris-slack/network/members"><img src="https://img.shields.io/github/forks/Tech-Tweakers/polaris-slack" alt="Forks Badge"/></a>
<a href="https://github.com/Tech-Tweakers/polaris-slack/pulls"><img src="https://img.shields.io/github/issues-pr/Tech-Tweakers/polaris-slack" alt="Pull Requests Badge"/></a>
<a href="https://github.com/Tech-Tweakers/polaris-slack/issues"><img src="https://img.shields.io/github/issues/Tech-Tweakers/polaris-slack" alt="Issues Badge"/></a>
<a href="https://github.com/Tech-Tweakers/polaris-slack/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/Tech-Tweakers/polaris-slack?color=2b9348"></a>
<a href="https://github.com/Tech-Tweakers/polaris-slack/blob/master/LICENSE"><img src="https://img.shields.io/github/license/Tech-Tweakers/polaris-slack?color=2b9348" alt="License Badge"/></a>
</div>

<br>
<p align="center"><i>Have problems or a free time to help? Please open an <a href="https://github.com/Tech-Tweakers/polaris-slack/issues/new">Issue</a> to tell us!</i></p>

## About

This project is a simple REST API to make a bridge between Polaris and Slack.

## Setup

Create a **.env.local** file in the root of the project and add the following variables:

```
#
# Global settings
#

ENVIRONMENT=development
LOG_LEVEL=debug

#
# Slack integration settings
#

SLACK_CHANNEL_ID= # Slack Channel ID to send messages
SLACK_AUTH_TOKEN= # Slack OAuth Token
SLACK_APP_TOKEN=  # Slack App Token
SLACK_API_URL=https://api.slack.com/
```
## Run

```
go run cmd/main.go
```

## API Endpoints

```
GET: http://localhost:9990/entries

#
# POST not used since the AI is replying to the messages. For tests only.
#

POST: http://localhost:9990/slack-reply
    {
      "payloadtext": "reply to existing message",
      "payloadts": "1704049551.547199"
    }
```
## Polaris Projects

- Polaris Frontend: https://github.com/Tech-Tweakers/polaris-frontend :star:
- Polaris Backend: https://github.com/Tech-Tweakers/polaris-chatbot :star:

## Contributors / Special Thanks :heart:

In memoriam of **Anderson Roberto** - https://github.com/EuAndersonRoberto 
