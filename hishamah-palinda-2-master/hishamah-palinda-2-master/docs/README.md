## Task 2

### What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?
The wait group has been added to wait for all the producers to finish sending the data before closing the channel but the problem is when the channel is closed first then no data can be send anymore through the channel which caused the error "Panic: send on closed channel".

### What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?
The channel will be closed right after one of the producers has finished sending data which will vary and that will prevent the other producers from sending the remaining data through the channel. 

### What happens if you remove the statement close(ch) completely?
Nothing will happens because it wasn't really necessary from the beginning. The statement was closing the channel after all data has been sent because of the waitGroup the blocks until the data has been sent.

### What happens if you increase the number of consumers from 2 to 4?
The same amount of data will be ditributed to 2 more consumers which means that data will now be recivied 4 consumers instead of 2 which will means less amount of data will be recieved by the first two consumers in the beginning.

### Can you be sure that all strings are printed before the program stops?
No because there is a possibility that the last strings has been sent but then the channel has closed or the program finished before the consumer function had time to print the last strings. The waitGroup is only blocking and waiting for the producers and will not wait for the consumers to finish that's why it would be a better to add a new waitgroup for the consumers.  
