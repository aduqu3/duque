from fastapi import FastAPI, HTTPException, Depends, Query
from pydantic import BaseModel, Field
from . import models
from datetime import datetime, date
from .database import engine, SessionLocal
from sqlalchemy.orm import Session
from enum import Enum


app = FastAPI()

models.Base.metadata.create_all(bind=engine)


def get_db():
    try:
        db = SessionLocal()
        yield db
    finally:
        db.close()

class Logs(BaseModel):
    cantidad: int
    evento : str
    paradero : str

paraderos = Query("Centro", enum=[
        "Parque de los Estudiantes", 
        "Unicentro",
        "Colegio Normal",
        "Colegio Industrial",
        "Makro",
        "Alborada",
        "Universidad Cooperativa",
        "Gasolinera Ocoa",
        "Barcelona"
    ])

@app.get("/")
def read_api(db: Session = Depends(get_db)):
    return db.query(models.Logs).all()


@app.post("/add_fila")
def add_fila(fila: Logs, db: Session = Depends(get_db), paradero: str = paraderos):
    log = models.Logs()
    log.time = datetime.now()
    log.cantidad = 1
    log.evento = "Estudiante entra a la fila"
    log.paradero = paradero
    db.add(log)
    db.commit()
    return log

@app.post("/out_fila")
def out_fila(fila: Logs, db: Session = Depends(get_db), paradero: str = paraderos):
    log = models.Logs()
    log.time = datetime.now()
    log.cantidad = 1
    log.evento = "Estudiante sale de la fila"
    log.paradero = paradero
    db.add(log)
    db.commit()
    return log

@app.post("/start_trayecto")
def start_trayecto(fila: Logs, db: Session = Depends(get_db), paradero: str = paraderos):
    log = models.Logs()
    log.time = datetime.now()
    log.evento = "Inicia trayecto"
    log.paradero = paradero
    db.add(log)
    db.commit()
    return log

@app.post("/parada_trayecto_centro_u")
def parada_trayecto_centro_u(fila: Logs, db: Session = Depends(get_db), paradero: str = paraderos):
    log = models.Logs()
    log.time = datetime.now()
    log.evento = "Parada intermedia Centro a Universidad"
    log.paradero = paradero
    db.add(log)
    db.commit()
    return log

@app.post("/finish_trayecto")
def finish_trayecto(fila: Logs, db: Session = Depends(get_db), paradero: str = paraderos):
    log = models.Logs()
    log.time = datetime.now()
    log.evento = "Finaliza trayecto"
    log.paradero = paradero
    db.add(log)
    db.commit()
    return log

@app.post("/sube_bus")
def sube_bus(fila: Logs, db: Session = Depends(get_db), paradero: str = paraderos):
    log = models.Logs()
    log.time = datetime.now()
    log.cantidad = 1
    log.evento = "Sube al bus"
    log.paradero = paradero
    db.add(log)
    db.commit()
    return log

@app.post("/baja_bus")
def baja_bus(fila: Logs, db: Session = Depends(get_db), paradero: str = paraderos):
    log = models.Logs()
    log.time = datetime.now()
    log.cantidad = 1
    log.evento = "Baja del bus"
    log.paradero = paradero
    db.add(log)
    db.commit()
    return log


@app.post("/parada_trayecto_u_centro")
def parada_trayecto_u_centro(fila: Logs, db: Session = Depends(get_db), paradero: str = paraderos):
    log = models.Logs()
    log.time = datetime.now()
    log.evento = "Parada intermedia Centro a Universidad"
    log.paradero = paradero
    db.add(log)
    db.commit()
    return log

# @app.post("/1/")
# def create_fila_with_1(db: Session = Depends(get_db)):

#     fila_model = models.Fila()
#     fila_model.cantidad = 1

#     db.add(fila_model)
#     db.commit()

#     return fila_model


# @app.put("/{book_id}")
# def update_book(book_id: int, fila: Book, db: Session = Depends(get_db)):

#     fila_model = db.query(models.Books).filter(models.Books.id == book_id).first()

#     if fila_model is None:
#         raise HTTPException(
#             status_code=404,
#             detail=f"ID {book_id} : Does not exist"
#         )

#     fila_model.title = fila.title
#     fila_model.author = fila.author
#     fila_model.description = fila.description
#     fila_model.rating = fila.rating

#     db.add(fila_model)
#     db.commit()

#     return fila


# @app.delete("/{book_id}")
# def delete_book(book_id: int, db: Session = Depends(get_db)):

#     fila_model = db.query(models.Books).filter(models.Books.id == book_id).first()

#     if fila_model is None:
#         raise HTTPException(
#             status_code=404,
#             detail=f"ID {book_id} : Does not exist"
#         )

#     db.query(models.Books).filter(models.Books.id == book_id).delete()

#     db.commit()
