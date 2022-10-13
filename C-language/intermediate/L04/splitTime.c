#include <stdio.h>
#define N 5


//dato un intero time, viene calcolato quanto tempo in ore, minuti e secondi è passato dalla mezzanotte
void split_time(long int tot_sec, int *h, int *m, int *s);

int main(void){
    int time = 1800, h = 0, m = 0, s = 0;

    split_time(time,&h,&m,&s);
    printf("h = %d, m = %d, s = %d\n",h,m,s);
    return 0;
}

void split_time(long int tot_sec, int *h,int *m, int *s){
    *h = tot_sec /3600;
    tot_sec %= 3600;
    *m = tot_sec / 60;
    *s = tot_sec % 60;
}
