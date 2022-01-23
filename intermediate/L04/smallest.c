#include <stdio.h>
#define N 100

int *smallest(int a[],int n);

int main(void){
    int num[N];
    int n;
    int *small;
    printf("Inserire dimensione array: ");
    scanf("%d",&n);
    printf("Inserire i numeri: ");
    for(int i = 0; i < n;i++){
        scanf("%d",&num[i]);
    }
    small = smallest(num,n);
    printf("lo smallest è: %d\n",*small);
    return 0;
}

int *smallest(int a[],int n){
    int *save;
    save = &a[0];
    for(int i = 0;i < n;i++){
        if(a[i] < *save){
            save = &a[i];
        }
    }    
    return save;
}
