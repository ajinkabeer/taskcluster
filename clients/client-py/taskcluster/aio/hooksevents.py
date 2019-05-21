# coding=utf-8
#####################################################
# THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT #
#####################################################
# noqa: E128,E201
from .asyncclient import AsyncBaseClient
from .asyncclient import createApiClient
from .asyncclient import config
from .asyncclient import createTemporaryCredentials
from .asyncclient import createSession
_defaultConfig = config


class HooksEvents(AsyncBaseClient):
    """
    The hooks service is responsible for creating tasks at specific times orin .  response to webhooks and API calls.Using this exchange allows us tomake hooks which repsond to particular pulse messagesThese exchanges provide notifications when a hook is created, updatedor deleted. This is so that the listener running in a different hooks process at the other end can direct another listener specified by`hookGroupId` and `hookId` to synchronize its bindings. But you are ofcourse welcome to use these for other purposes, monitoring changes for example.
    """

    classOptions = {
        "exchangePrefix": "exchange/taskcluster-hooks/v1/",
    }
    serviceName = 'hooks'
    apiVersion = 'v1'

    def hookCreated(self, *args, **kwargs):
        """
        Hook Created Messages

        Whenever the api receives a request to create apulse based hook, a message is posted to this exchange andthe receiver creates a listener with the bindings, to create a task

        This exchange takes the following keys:

         * reserved: Space reserved for future routing-key entries, you should always match this entry with `#`. As automatically done by our tooling, if not specified.
        """

        ref = {
            'exchange': 'hook-created',
            'name': 'hookCreated',
            'routingKey': [
                {
                    'multipleWords': True,
                    'name': 'reserved',
                },
            ],
            'schema': 'v1/pulse-hook-changed-message.json#',
        }
        return self._makeTopicExchange(ref, *args, **kwargs)

    def hookUpdated(self, *args, **kwargs):
        """
        Hook Updated Messages

        Whenever the api receives a request to update apulse based hook, a message is posted to this exchange andthe receiver updates the listener associated with that hook.

        This exchange takes the following keys:

         * reserved: Space reserved for future routing-key entries, you should always match this entry with `#`. As automatically done by our tooling, if not specified.
        """

        ref = {
            'exchange': 'hook-updated',
            'name': 'hookUpdated',
            'routingKey': [
                {
                    'multipleWords': True,
                    'name': 'reserved',
                },
            ],
            'schema': 'v1/pulse-hook-changed-message.json#',
        }
        return self._makeTopicExchange(ref, *args, **kwargs)

    def hookDeleted(self, *args, **kwargs):
        """
        Hook Deleted Messages

        Whenever the api receives a request to delete apulse based hook, a message is posted to this exchange andthe receiver deletes the listener associated with that hook.

        This exchange takes the following keys:

         * reserved: Space reserved for future routing-key entries, you should always match this entry with `#`. As automatically done by our tooling, if not specified.
        """

        ref = {
            'exchange': 'hook-deleted',
            'name': 'hookDeleted',
            'routingKey': [
                {
                    'multipleWords': True,
                    'name': 'reserved',
                },
            ],
            'schema': 'v1/pulse-hook-changed-message.json#',
        }
        return self._makeTopicExchange(ref, *args, **kwargs)

    funcinfo = {
    }


__all__ = ['createTemporaryCredentials', 'config', '_defaultConfig', 'createApiClient', 'createSession', 'HooksEvents']