from sqlalchemy import Column, Integer, String, DateTime, Date
# from sqlalchemy.sql import func
from datetime import datetime, timezone
from database import Base


# logs
class Logs(Base):
    __tablename__ = "logs"

    id = Column(Integer, primary_key=True, index=True)
    time = Column(DateTime)
    cantidad = Column(Integer)
    evento = Column(String)
    paradero = Column(String)

