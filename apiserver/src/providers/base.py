import typing as t


class BaseProvider:

    def __init__(self, connection_url: str):
        self.connection_url = connection_url

    def publish(self, payload: t.Dict[str, t.Any]) -> None:
        ...
