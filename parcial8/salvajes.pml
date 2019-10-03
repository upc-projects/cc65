

#define wait(s) atomic{ s > 0 -> s--}
#define signal(s) s++

#define MAX 10

byte isEmpty = 1
byte cooking = 0
byte mutex = 0
int portions = MAX

active proctype cocinero(){
    do
    ::
        wait(cooking)
        printf("Cocinero cocinando, hay : %d\n", portions)
        portions++
        if
        :: portions == MAX -> signal(isEmpty)
        :: else -> signal(cooking)
        fi  
    od
}

active proctype salvaje(){
    do
    ::
        wait(isEmpty)
        printf("Salvaje comiendo, quedan:  %d porciones\n", portions)
        portions--
        if
        :: portions == 0 -> signal(cooking) 
        :: else -> signal(isEmpty)
        fi
    od
}