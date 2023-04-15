import typing as t

from fastapi import APIRouter, Body

from ..config import config
from ..providers import AMQPProvider
from ..schemas import CommandV1, SuccessResponse

api_router = APIRouter(prefix="/api/v1")


@api_router.post("/command", response_model=SuccessResponse)
async def command_v1(payload: t.Union[CommandV1, t.List[CommandV1]] = Body(...), ):
    provider = AMQPProvider(
        connection_url=config.AMQP_CONNECTION_STRING,
        exchange_name=config.AMQP_EXCHANGE_NAME,
        exchange_type=config.AMQP_EXCHANGE_TYPE
    )
    if isinstance(payload, list):
        for i in payload:
            provider.publish(payload=i.dict())
    else:
        provider.publish(payload=payload.dict())
    return SuccessResponse()
