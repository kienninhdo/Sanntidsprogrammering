
FORKLARING:

3. Ved GOMAXPROCS(1) får vi det samme resultatet hver gang gang som er 2000000. 
Ved GOPROCS(2) får vi ulikt resultat hver gang.

Fordi vi bruker en felles variabel i, får vi ikke 0 fordi threadsene ikke snakker sammen.
Dermed får vi race conditions.

4. Vi har valgt å bruke mutex, for vi forstår det slik at mutex brukes dersom vi opererer med få ressurser, mens semaphores brukes i mer komplekse systemer. Siden vi i dette tilfellet kun har en delt variabel, velger vi mutex.
Vi ser også at oppgaven spør oss om å "Acquire the lock", noe som kun er aktuelt ved bruk av mutex(semaphore bruker wait og signal).
Ved å bruke mutex til å lock og unlock unngår vi at increment og decrement bruker variabelen samtidig.