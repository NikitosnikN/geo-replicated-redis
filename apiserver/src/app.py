from fastapi import FastAPI

from .api.handlers import api_router
from .utils import exception_handlers


def get_app() -> FastAPI:
    docs_urls = {"docs_url": "/docs/", "redoc_url": "/redocs/", "openapi_url": '/docs/openapi.json'}

    _app = FastAPI(exception_handlers=exception_handlers, **docs_urls)

    ###################
    # Routes
    ####################

    _app.include_router(api_router)

    return _app


app = get_app()
