CREATE DATABASE IF NOT EXISTS company;

USE company;

CREATE TABLE Departments (
    Department_ID INT PRIMARY KEY,
    Department_Name VARCHAR(255),
    Department_Address VARCHAR(255)
);

CREATE TABLE Colaborators (
    Colaborator_ID INT PRIMARY KEY,
    File_Number INT,
    Last_Name VARCHAR(255),
    First_Name VARCHAR(255),
    Date_of_Birth DATE,
    Date_of_Incorporation DATE,
    Position VARCHAR(255),
    Salary DECIMAL(19, 2),
    fk_Department_ID INT,
    FOREIGN KEY (fk_Department_ID) REFERENCES Departments(Department_ID)
);

INSERT INTO Departments (Department_ID, Department_Name, Department_Address) VALUES
(1, 'Department 1', 'Street 1'),
(2, 'Department 2', 'Street 2'),
(3, 'Department 3', 'Street 3'),
(4, 'Department 4', 'Street 4'),
(5, 'Department 5', 'Street 5');

INSERT INTO Colaborators (
  Colaborator_ID, 
  File_Number, 
  Last_Name, 
  First_Name, 
  Date_of_Birth, 
  Date_of_Incorporation,
  Position,
  Salary,
  fk_Department_ID
  ) VALUES
(1, 10001, 'Last Name 1', 'First Name 1', '1990-01-01', '2020-01-01', 'Position 1', 1000.00, 1),
(2, 10002, 'Last Name 2', 'First Name 2', '1990-02-02', '2021-01-02', 'Position 1', 1000.00, 2),
(3, 10003, 'Last Name 3', 'First Name 3', '1990-02-03', '2022-01-03', 'Position 1', 1000.00, 3),
(4, 10004, 'Last Name 4', 'First Name 4', '1990-04-04', '2023-01-04', 'Position 1', 1000.00, 4),
(5, 10005, 'Last Name 5', 'First Name 5', '1990-05-05', '2024-01-05', 'Position 1', 1000.00, 5),
(6, 10006, 'Last Name 6', 'First Name 6', '1990-01-01', '2020-01-06', 'Position 2', 2000.00, 1),
(7, 10007, 'Last Name 7', 'First Name 7', '1990-02-02', '2021-01-07', 'Position 2', 2000.00, 2),
(8, 10008, 'Last Name 8', 'First Name 8', '1990-03-03', '2022-01-08', 'Position 2', 2000.00, 3),
(9, 10009, 'Last Name 9', 'First Name 9', '1990-04-04', '2023-01-09', 'Position 2', 2000.00, 4),
(10, 10010, 'Last Name 10', 'First Name 10', '1990-05-05', '2024-01-10', 'Position 2', 2000.00, 5),
(11, 10011, 'Last Name 11', 'First Name 11', '1990-01-01', '2020-01-11', 'Position 3', 3000.00, 1),
(12, 10012, 'Last Name 12', 'First Name 12', '1990-02-02', '2021-01-12', 'Position 3', 3000.00, 2),
(13, 10013, 'Last Name 13', 'First Name 13', '1990-03-03', '2022-01-13', 'Position 3', 3000.00, 3),
(14, 10014, 'Last Name 14', 'First Name 14', '1990-04-04', '2023-01-14', 'Position 3', 3000.00, 4),
(15, 10015, 'Last Name 15', 'First Name 15', '1990-05-05', '2024-01-15', 'Position 3', 3000.00, 5);

SELECT 
  Departments.*, 
  Colaborators.Colaborator_ID,
  Colaborators.File_Number,
  Colaborators.Last_Name,
  Colaborators.First_Name,
  Colaborators.Date_of_Birth,
  Colaborators.Date_of_Incorporation,
  Colaborators.Position,
  Colaborators.Salary
FROM Departments 
LEFT JOIN Colaborators ON Departments.Department_ID = Colaborators.fk_Department_ID 
WHERE Departments.Department_ID = 1 OR Departments.Department_ID = 2;

-- The following fails due to the foreign key constraint
INSERT INTO Colaborators (
  Colaborator_ID, 
  File_Number, 
  Last_Name, 
  First_Name, 
  Date_of_Birth, 
  Date_of_Incorporation,
  Position,
  Salary,
  fk_Department_ID
  ) VALUES
(1, 10001, 'Last Name 1', 'First Name 1', '2000-01-01', '2020-01-01', 'Position 1', 1000.00, 1000);