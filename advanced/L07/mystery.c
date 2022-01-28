//Qual è lo scopo della funzione mystery? qual è l'obiettivo complessivo della funzione?
//
//risposta:
//Data come argomento una lista L e un intero N, la funzione scorre la lista fino a che il valore contenuto nel nodo corrente 
//è diverso dall'intero N passato come argomento.
//A quel punto il ciclo finisce e viene restituito il nodo precedente a quello con il valore uguale a N
struct node *mystery(List l, int value){
    struct node *current = l -> head;
    struct node *temp = NULL;
    while(current != NULL && current -> item != value){
	temp = current;
	current = current -> next;
    }
    return temp;
}
