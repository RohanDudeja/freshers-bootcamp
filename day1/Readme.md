Exercises:

1. Create a struct called 'Matrix'. The Matrix struct has the following information:
number of rows of matrix\
number of columns of matrix\
elements of matrix (You can use 2D vector)\
The Matrix struct has methods for each of the following:
* get the number of rows
* get the number of columns
* set the elements of the matrix at a given position (i,j)
* adding two matrices. (2nd matrix can be taken as input to the method)
* print matrix structure as json
(Assume that the dimensions are correct)\
Expectation: Use structs, slices, methods

2. Create an expression tree using golang structs. For example string a + b - c on tree would look like the following
Once the tree is built, traverse the tree in preorder and postorder format and print the values.  
&nbsp;&nbsp;&nbsp;&nbsp;+  
&nbsp;&nbsp;/&nbsp;&nbsp;&nbsp;&nbsp;\  
a &nbsp;&nbsp;&nbsp;&nbsp;-    
&nbsp;&nbsp;&nbsp;&nbsp;/&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;\  
&nbsp;&nbsp;&nbsp;b&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;c    


Expectation: Use Recursion, Strings, Pointers, Struct

4. There are 3 types of employees in an organization:
Fulltime, Contractor, Freelancer\
Full-time employees are paid by the day and the Contractors are paid by the day.
Freelancers are paid by the hour.
Here is the salary structure.,

| Employee |  Basic |  Duration |
|----------|----------|----------|
| Full time | 500 | daily |
| Contractor | 100 | daily |
| Freelancer | 10 | hourly |

Assume there are 28 working days in a month. Total hours worked by the freelancer has to be configured.

Example:

Full-time employee gets paid 15000 per month
The contractor gets paid 3000 per month
Freelancer gets paid 2000 (if freelancer works 20 hours)

Define an interface that calculates the salary of an employee.

     

