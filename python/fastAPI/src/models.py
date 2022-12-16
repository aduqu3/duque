from sqlalchemy import Column, Integer, String, DateTime, Date
# from sqlalchemy.sql import func
from datetime import datetime, timezone
from database import Base


# fila entrada
class Fila(Base):
    __tablename__ = "fila"

    id = Column(Integer, primary_key=True, index=True)
    time = Column(DateTime, default=datetime.now())
    cantidad = Column(Integer)


# fila salida
