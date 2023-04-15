FROM python:3.11-slim

ENV PYTHONUNBUFFERED=1 \
    \
    # pip
    PIP_NO_CACHE_DIR=off \
    PIP_DISABLE_PIP_VERSION_CHECK=on \
    PIP_DEFAULT_TIMEOUT=50 \
    \
    # poetry
    POETRY_VERSION=1.3.2 \
    POETRY_HOME="/opt/poetry" \
    POETRY_CACHE_DIR='/var/cache/pypoetry' \
    POETRY_VIRTUALENVS_CREATE=false \
    POETRY_VIRTUALENVS_IN_PROJECT=false \
    POETRY_NO_INTERACTION=1 \
    \
    # paths
    PYSETUP_PATH="/opt/pysetup"

RUN pip install "poetry==$POETRY_VERSION"

RUN mkdir /app/
WORKDIR /app/

COPY ./pyproject.toml ./poetry.lock /app/

# Install requirements
RUN poetry install --no-ansi --no-dev

COPY ./ /app/
