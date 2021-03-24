# Assignment 3

#### Task 2
1) The program will not be executed concurrently which will make it slower.
2) The program will run forever because the waitgroup that is being passed now is not a pointer to the actual waitgroup but a copy of it and when wg.Done() is called it's referring to the copied waitgroup which means the waitgroup will not be finished.
3) When data is sent through unbuffered channels it will be blocked until the data has been received but when it come to unfuffered channels it will block only until the data gets copied in the buffer.
4) The program works correctly but only because the default case was never needed. 
  
#### Task 3
The parallell version is 2.7x faster than the normal one. 