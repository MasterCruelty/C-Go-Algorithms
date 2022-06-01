#include <stdio.h>
#include <stdlib.h>
#include "queue.h"


int main(void){
	Queue q=NULL;
	
	new_queue(&q,3);
	enqueue(q,3);
	enqueue(q,7);
	enqueue(q,1);
	print_queue(q);
	
	printf("\nCoda vuota? %s", ( (is_empty_queue(q) == 0) ? "No" : "Si") );
	printf("\nIl primo elemento della coda è: %d\n", frontValue(q));
	
	printf("Tolgo %d dalla coda\n",dequeue(q));
	printf("Tolgo %d dalla coda\n",dequeue(q));
	printf("Tolgo %d dalla coda",dequeue(q));
	print_queue(q);

	
	printf("\nCoda vuota? %s", ( (is_empty_queue(q) == 0) ? "No" : "Si") );
	printf("\n");

	destroy(q);

	return 0;
}
