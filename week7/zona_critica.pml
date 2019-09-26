#define wait(s) atomic { s > 0 -> s-- }
#define signal(s) s++

byte counter = 0
byte roomEmpty = 1
byte mutex = 1

active proctype writer(){
    do
    :: 
        wait(roomEmpty)
        // para ver existencia
        printf("Writting stuff 1\n")
        printf("Writting stuff 2\n")
        printf("Writting stuff 3\n")
        signal(roomEmpty)
    od
}


active[2] proctype reader(){
    do
    :: 
        wait(mutex)
        counter++
        if
        :: counter == 1 -> wait(roomEmpty)
        :: else -> 
        fi
        signal(mutex)
        // para ver existencia
        printf("Reading stuff 1\n")
        printf("Reading stuff 2\n")
        printf("Reading stuff 3\n")

        wait(mutex)
        counter--
        if
        :: counter == 0 -> signal(roomEmpty)
        :: else -> 
        fi
        signal(mutex)
    od
}