import typing as t

from kombu import Exchange, Connection

from .base import BaseProvider
from ..config import config


class AMQPProvider(BaseProvider):

    def __init__(self, connection_url: str, exchange_name: str, exchange_type: str, **kwargs):
        super().__init__(connection_url)
        self.exchange = Exchange(name=exchange_name, type=exchange_type, **kwargs)

    def publish(self, payload: t.Dict[str, t.Any]) -> None:
        with Connection(self.connection_url) as conn:
            producer = conn.Producer(serializer='json')
            producer.publish(
                body=payload,
                routing_key=config.AMQP_ROUTE_KEY,
                exchange=self.exchange,
                declare=[self.exchange],
                retry=True,
            )

        return None
