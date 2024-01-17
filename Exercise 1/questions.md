Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
Concurrency er når det virker som om to hendelser skjer samtidig, men i realiteten så 
skifter man bare utrolig hurtig mellom dem. Ved parallelism har man to hendelser som faktisk skjer samtidig. 

What is the difference between a *race condition* and a *data race*? 
Race condition definerer når resultatet avhenger av rekkefølgen på prosesser/hendelser. Data race er en type race condition der flere threads aksesserer samme minne samtidig.
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
En scheduler setter opp en oversikt over hvilken thread som er neste den skal kjøre. Det gjør den ved å swappe og starte threads.


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
Vi bruker flere threads fordi da får vi at funksjonene ikke delayer hverandre, men gir mulighet til å lage concurrency

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
Fibers er når vi deler en thread inn i flere mindre oppgaver, og alle disse oppgavene gjør eget arbeid. Grunnen til at denne er bedre enn thread er fordi den selv vet når det er fornuftig å stoppe i motsetning til en thread som kan bli forstyrret

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
Det blir mer koding, men programmet kjører som du ser for deg når du bruker concurrency, så selv om du må kode mer, så er det for det beste

What do you think is best - *shared variables* or *message passing*?
Message passing, fordi du ønsker så mye uavhengighet som mulig


