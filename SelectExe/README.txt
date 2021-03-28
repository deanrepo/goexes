In Go language, the select statement is just like switch statement, but in the select statement,
case statement refers to communication, i.e. sent or receive operation on the channel. 

Important points:

    1) Select statement waits until the communication(send or receive operation) is prepared for some cases to begin. 
    2) If a select statement does not contain any case statement, then that select statement waits forever.
        Syntax:
            select{}
    3) The default statement in the select statement is used to protect select statement from blocking. This statement executes when there is no case statement is ready to proceed.    
    4) The blocking of select statement means when there is no case statement is ready and the select statement does not contain any default statement, then the select statement block until at least one case statement or communication can proceed.
    5) In select statement, if multiple cases are ready to proceed, then one of them can be selected randomly. (eg: multirandomly.go)