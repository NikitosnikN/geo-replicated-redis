from fastapi import responses
from fastapi.exceptions import HTTPException as FastAPIHTTPException
from fastapi.exceptions import RequestValidationError
from pydantic import ValidationError
from starlette import status
from starlette.requests import Request

__all__ = ["exception_handlers"]


async def pydantic_exception_handler_func(request: Request, exception: ValidationError):
    errors = []
    for error in exception.errors():
        errors.append(
            {
                "field": error["loc"][-1],
                "message": error.get("msg"),
                "type": error.get("type"),
            }
        )

    return responses.JSONResponse(
        content={"success": False, "message": "validation error", "details": errors},
        status_code=status.HTTP_422_UNPROCESSABLE_ENTITY,
    )


async def http_exc_handler(request: Request, exception: FastAPIHTTPException):
    headers = {"Access-Control-Allow-Origin": "*"}

    if exception.headers:
        headers.update(**exception.headers)

    return responses.JSONResponse(
        content={"success": False, "message": exception.detail, "details": {}},
        status_code=getattr(exception, "status_code", status.HTTP_500_INTERNAL_SERVER_ERROR),
        headers=headers,
    )


exception_handlers = {
    FastAPIHTTPException: http_exc_handler,
    RequestValidationError: pydantic_exception_handler_func,
    ValidationError: pydantic_exception_handler_func,
}
