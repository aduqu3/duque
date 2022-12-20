# app/config.py

import os

from pydantic import BaseSettings, Field


class Settings(BaseSettings):
    SQLALCHEMY_DATABASE_URL: str = Field(..., env='DATABASE_URL')
settings = Settings()