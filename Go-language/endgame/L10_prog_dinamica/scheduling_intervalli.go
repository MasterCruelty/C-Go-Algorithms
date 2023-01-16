package main


/*
 * Intervallo pesato definito come tripla di interi (i,f,v)
 * i = tempo di inizio
 * f = tempo di fine
 * v = valore/peso dell'intervallo.
 * Dato I insieme di intervalli, una soluzione è data da un sottoinsieme S contenuto in I di intervalli non sovrapposti tra loro.
 * Il valore di una soluzione S è dato dalla somma dei valori degli intervalli S.
 * 
 *
 *
 * Step 1: ordinare gli elementi di i in base al tempo di fine così da avere I= {(i1,f1,v1),(i2,f2,v2),...,(in,fn,vn)} con f1<=f2<=..<=fn.
 * Step 2: per ogni indice j tra 1 e n, definire p(j) come il più grande indice i < j  tale che l'intervallo di indice i non si sovrappone all'intervallo di indice j.
 * Step 3: Definire Opt(j) come il valore di una qualsiasi soluzione ottimale Sj costruita usando gli intervalli di indici{1,2,...,j}.
 * Quindi vale questa proprietà ricorsiva: Opt(j) = max{vj + Opt(p(j)),Opt(j-1)} (NB: se j appartiene a Sj vale il primo argomento, altrimenti il secondo.)
 *
*/
