from fastapi import FastAPI, HTTPException, Depends
from pydantic import BaseModel, Field
import models
from datetime import datetime, date
from database import engine, SessionLocal
from sqlalchemy.orm import Session

app = FastAPI()

models.Base.metadata.create_all(bind=engine)


def get_db():
    try:
        db = SessionLocal()
        yield db
    finally:
        db.close()


class Fila(BaseModel):
    cantidad: int = Field(gt=-1, lt=101)  

# BOOKS = []


@app.get("/")
def read_api(db: Session = Depends(get_db)):
    return db.query(models.Fila).all()


@app.post("/")
def create_fila(fila: Fila, db: Session = Depends(get_db)):

    fila_model = models.Fila()
    fila_model.cantidad = fila.cantidad

    db.add(fila_model)
    db.commit()

    return fila

@app.post("/1/")
def create_fila_with_1(db: Session = Depends(get_db)):

    fila_model = models.Fila()
    fila_model.cantidad = 1

    db.add(fila_model)
    db.commit()

    return fila_model


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
