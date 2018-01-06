#!/usr/bin/python3

import os
import json
import unittest
from unittest.mock import Mock, patch

import send


HERE = os.path.dirname(__file__)


class TestSlack(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        super(TestSlack, cls).setUpClass()

        with open(os.path.join(HERE, '../', 'messages.txt'), 'r') as f:
            cls.messages = f.readlines()

    @patch.object(send, 'urlopen')
    def test_slack_sending(self, mock_open):
        """
        Assert that given a fixed input, we get the desired output
        """
        mock_open.return_value.status = 200
        mock_open.return_value.read.return_value = 'OK'

        fake_arg = Mock(
            webhook='http://test.com',
            input=os.path.join(HERE, '../', 'messages.txt'),
            channel='channel',
            username='username',
            emoji='emoji'
        )

        send.send_messages(fake_arg)

        request, payload = mock_open.call_args[0]

        self.assertEqual(request.headers.get('Content-type'), 'application/json; charset=utf-8')
        self.assertEqual(mock_open.call_count, len(self.messages))

        for call, expected in zip(mock_open.call_args_list, self.messages):
            message = json.loads(call[0][1])
            self.assertEqual(message['text'], expected.strip())
            self.assertEqual(message['channel'], fake_arg.channel)
            self.assertEqual(message['username'], fake_arg.username)
            self.assertEqual(message['icon_emoji'], fake_arg.emoji)


if __name__ == '__main__':
    unittest.main()
