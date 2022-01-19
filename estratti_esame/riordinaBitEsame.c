#include <stdio.h>

//dato un n da input, inserire n zeri e uni da input in disordine. l'output dovrà essere solo di zeri a sinistra e uni a destra.
int main(void){
    int n;
    scanf("%d",&n);
    int a[n];
    //popolo l'array con zeri e uni
    for(int i = 0;i < n; i++){
        scanf("%d",&a[i]);
    }
    while(1){
        int sx = 0;
        int dx = n -1;
	//controllo se partendo da sinistra ho uno zero, in caso vado avanti altrimenti se trovo uno mi fermo
        while(a[sx] == 0)
            sx++;
	//controllo se partendo da destra ho un uno, in caso vado avanti altrimenti se trovo zero mi fermo
	while(a[dx] == 1)
	    dx--;
	//se sx è maggiore di dx ho già superato la metà dell'array quindi ho finito
	if(sx > dx)
	    break;
	//scambio lo zero con l'uno se scorrendo ho trovato lo zero e l'uno in posizione sbagliata
	a[sx] = 0;
	a[dx] = 1;
    }
    for(int i = 0;i < n; i++){
        printf("%d",a[i]);
    }
    printf("\n");
    return 0;
}
