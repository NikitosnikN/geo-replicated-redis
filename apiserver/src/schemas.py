import typing as t

from pydantic import BaseModel, Field, validator

from .config import SUPPORTED_COMMANDS


class SuccessResponse(BaseModel):
    success: bool = True


class CommandV1(BaseModel):
    command: str = Field(...)
    args: t.List[str] = Field(..., min_items=1)

    @validator('command')
    def validate_command(cls, v: str):
        if not v.isupper():
            raise ValueError('should be in upper case')

        if v not in SUPPORTED_COMMANDS:
            raise ValueError('command is not supported')

        return v
