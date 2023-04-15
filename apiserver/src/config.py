from pydantic import BaseSettings, Field

SUPPORTED_COMMANDS = ("SET", "DEL")


class ConfigClass(BaseSettings):
    HOST: str = Field(default="0.0.0.0")
    PORT: int = Field(default=8000)
    AMQP_CONNECTION_STRING: str
    AMQP_EXCHANGE_NAME: str
    AMQP_EXCHANGE_TYPE: str

    class Config:
        case_sensitive = True
        env_file = '.env', './config/.env'
        env_file_encoding = 'utf-8'


config = ConfigClass()
