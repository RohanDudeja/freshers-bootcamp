Exercises: 

1. For a given set of strings, find out the frequency of each letter in all the strings parallelly. For example, if given the following input.,\
Expectations: Slices, Goroutines, Channels [“quick”, “brown”, “fox”, “lazy”, “dog”]. 
The output should be.,
>{\
   > “a”: 1,\
   > “b”: 1,\
   > “c”: 1,\
   > …\
   > “z”: 1\
}
2. Suppose you are a class monitor and you have to take the performance rating of the class teacher from your classmates. Write a program to take the ratings from all of your classmates and then print the average rating.\
Expectation: use waitgroups\
Assumptions:
+ There are 200 students in the class.\ 
+ Every student will take a random amount of time to respond.\
+ You can simulate the random response time of your classmates by using the math/rand package.\ 


3. Write a program that starts with a single bank account with a starting balance of Rs.500. It should be possible to deposit and withdraw money concurrently. However, the balance update should happen only once at any point in time. If the withdrawal action should fail if the balance is reaching negative.\
Expectations: Use Goroutines, Mutex 
