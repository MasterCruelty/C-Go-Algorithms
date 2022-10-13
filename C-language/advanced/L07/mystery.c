//Qual è lo scopo della funzione mystery? qual è l'obiettivo complessivo della funzione?
//
//risposta:
//la funzione non fa altro che cercare l'elemento valure all'interno della lista e lo restituisce quando l'ha trovato.
struct node *mystery(List l, int value){
    struct node *current = l -> head;
    struct node *temp = NULL;
    while(current != NULL && current -> item != value){
	temp = current;
	current = current -> next;
    }
    return temp;
}
