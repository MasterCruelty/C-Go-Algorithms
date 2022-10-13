Programmazione dinamica

Problemi di ottimizzazione. Ci possono essere più soluzioni possibili rappresentate da un valore e a noi ci interessa la soluzione ottimale di valore massimo o minimo.
La soluzione ottimale contiene al suo interno le soluzioni ottimali dei suoi sotto-problemi.
Ogni sotto-problema viene risolto una volta sola e ne viene memorizzata la soluzione in una tabella(memoization).


Problema: scheduling di intervalli pesati.

Un intervallo è definito come una tripla (i,f,v) con i tempo di inizio, f tempo di fine, v valore dell'intervallo.
Avendo un insieme I di intervalli, una soluzione al problema è un insieme S contenuto in I di intervalli che non si sovrappongono tra loro. 
Il valore di una soluzione S è dato dalla somma dei valori degli intervalli contenuti in S.

Risoluzione:

Ordino l'insieme degli intervalli in base al tempo di fine, quindi:
I = {(i_1,f_1,v_1),(i_2,f_2,v_2), . . . , (i_n,f_n,v_n)}

con f_1 <= f_2 <= ... <= f_n

Definisco p(j) con j compreso tra 1 e n. p(j) mi sputa fuori l'indice i < j più grande tale che l'intervallo di indice i non si sovrappone all'intervallo di indice j.

1 |--------| w_1 = 2

2   |--------------| w_2 = 4

3              |------------| w_3 = 4

4     |------------------------| w_4 = 7 

5                            |-----| w_5= 2

6                             |-----| w_6= 1

ad esempio avendo questi intervalli p(1) = 0 perché non vi è nessun intervallo con indice minore di 1 che non si sovrappone.
p(3) = 1 perché l'intervallo di indice maggiore che non si sovrappone prima di 3 è di indice 1.
p(5) = 3 perché l'intervallo di indice maggiore che non si sovrappone prima di 5 è di infice 3.


Definito con Opt(j) il valore di una soluzione ottimale S_j costruita usando gli intervalli di indice {1,2,..,j}, vale la seguente relazione ricorsiva:

Opt(j) = max{v_j + Opt(p(j)), Opt(j-1)}

se j appartiene a S_j, allora Opt(j) = v_j + Opt(p(j))
altrimenti Opt(j) = Opt(j-1).

Per evitare le troppe chiamate ricorsive e dover ricalcolare tante volte gli stessi sotto-problemi, attuiamo la memoization.
Memorizziamo le soluzioni dei sotto-problemi in un vettore.

Compute-Opt(j):

if j == 0 then
	return 0
else if M[j] contiene già qualcosa(ho già risolto questo sotto problema)
	return M[j]
else
	M[j] = max(v_j + Compute-Opt(p(j)), Compute-Opt(j-1))
	return M[j]
endif

e ora:

FindSolution(j):

if j == 0 then
	output vuoto
else
	if v_j + M[p(j)] >= M[j-1] then
		Stampa j con il risultato di FindSolution(p(j))
	else
		Stampa il risultato di FindSolution(j-1)
	endif
endif



Versione iterativa per la costruzione del vettore delle sottosoluzioni:

M[0] = 0
for j = 1,2,..,n
	M[j] = max(v_j + M[p(j)], M[j-1])
endfor
