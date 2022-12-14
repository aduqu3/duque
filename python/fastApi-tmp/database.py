from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.declarative import declarative_base
from config import settings

# SQLALCHEMY_DATABASE_URL = "sqlite:///./database.db"
# SQLALCHEMY_DATABASE_URL = "postgresql://postgres:postgres@localhost/postgres"

engine = create_engine(settings.SQLALCHEMY_DATABASE_URL)

# engine = create_engine(
#     SQLALCHEMY_DATABASE_URL, connect_args={"check_same_thread": False}
# )

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

Base = declarative_base()
