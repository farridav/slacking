#!/usr/bin/env python3

import os
import argparse
import json
from urllib.request import Request, urlopen


HERE = os.path.dirname(__file__)


def send_messages(cmd_args):
    """
    A Script for reading a text file of messages, and sending them to a Slack webhook

    # Setup

    Go to https://<subdomain>.slack.com/services/B8NRSBB0V and setup an integration (grab the webhook_url)

        ./send.py --webhook '<your url from above>'

    """
    messages = open(cmd_args.input, 'rb').readlines()

    for message in messages:
        payload = {
            'text': message.strip().decode('utf-8'),
            'username': cmd_args.username,
            'icon_emoji': cmd_args.emoji
        }

        if cmd_args.channel:
            payload['channel'] = cmd_args.channel

        payload = json.dumps(payload)
        request = Request(cmd_args.webhook)
        request.add_header('Content-Type', 'application/json; charset=utf-8')
        request.add_header('Content-Length', len(payload))

        response = urlopen(request, payload.encode('utf-8'))

        print('\n\n')
        print('Status: {}'.format(response.status))
        print('Content: {}'.format(response.read()))
        print('\n\n')


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("--webhook", help="Your slack webhook url, see http://bit.ly/2EapumJ", required=True)
    parser.add_argument("--channel", help="The channel to post to", required=False)

    parser.add_argument("--input", help="Path to your message file", default=os.path.join(HERE, '../', 'messages.txt'))
    parser.add_argument("--username", help="The username to use", default='Mr Shipit')
    parser.add_argument("--emoji", help="Your icon emoji", default=':shipit:')

    args = parser.parse_args()
    send_messages(args)
