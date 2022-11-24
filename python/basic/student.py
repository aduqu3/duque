from person import Person

class Student(Person):
  def __init__(self, fname, lname, year):
    super().__init__(fname, lname)
    self.graduationyear = year
  
  

x = Student("Andres", "Duque", 2023)
print(x.firstname)
print(x.lastname)
print(x.graduationyear)