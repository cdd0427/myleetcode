select max(Salary) SecondHighestSalary
from employee
where Salary<(select max(Salary) SecondHighestSalary from employee)