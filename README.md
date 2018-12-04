# rpn

RPN is a small reverse polish (programmable) calculator. It has 4 Stack levels and was written by myself in 2015. 
I use it every day to crunch some numbers on the commandline. You can invoke RPN from the terminal and let it parse its 
arguments or start the program and use it interactively.

This cute little piece of written go code is free software by any means. Do what you what you want with. Enjoy it. Change it and distribute it 
as long as you mention my name somewhere. Feel free to contact me. 

https://sebastiankind.de/post/terminalrechner

and 

https://github.com/zeppel13/rpn/blob/master/howto.org

The Link to my blog offers a German description.

- Install Go any version should do fine
- compile rpn with

```
cd ./rpn
go build 
```
- test it 

```
rpn 3 3 a 5 d
```

should give you 1.20

If you wish to compile RPN for older computers e.g. with a Pentium II or anything else which has 387 compatible 
floatingpoint processor, you need to set this in order to compile and run it correctly. Tested for a Pentium II on Windows XP and 
Debian Linux. Somewhere on my old computers HDD should exist a small C Port for a 286 Personal Computer. But it shouldn't be too
difficult to reprogram this software in any language by yourself.

```
GO386=387 go build
```


# Table of Contents

1.  [Kurzübersicht aller Befehle](#org9568ebd)
    1.  [Normale Befehle, die man auch braucht](#org98eba56)
    2.  [Sonderfunktionern](#org0ece62f)
    3.  [Ergebnisse Speichern](#orgb36b024):cool:
    4.  [Designkram](#org7357fac)
2.  [Achtung](#orgbcd37e2)



<a id="org9568ebd"></a>

# Kurzübersicht aller Befehle

Es lohnt sich in den Quelltext zu schauen.


<a id="org98eba56"></a>

## Normale Befehle, die man auch braucht

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-left" />

<col  class="org-left" />
</colgroup>
<thead>
<tr>
<th scope="col" class="org-left">Befehl</th>
<th scope="col" class="org-left">Funktion</th>
</tr>
</thead>

<tbody>
<tr>
<td class="org-left">[Enter]</td>
<td class="org-left">Befehl bearbeiten</td>
</tr>


<tr>
<td class="org-left">sw</td>
<td class="org-left">Registertauschen(1<->2)</td>
</tr>


<tr>
<td class="org-left">r</td>
<td class="org-left">Regiser rollen</td>
</tr>


<tr>
<td class="org-left">drop</td>
<td class="org-left">Letztes Register löschen</td>
</tr>


<tr>
<td class="org-left">reset</td>
<td class="org-left">Alle Register löschen</td>
</tr>
</tbody>
</table>

Operatoren und Operanden werden in eine Zeile eingegeben und mit
[Enter] abgearbeitet. Es lassen sich die Eingaben zum Erhalt der
Übersichtlichkeit oder aus praktischen Gründen auch nacheinander
jeweils mit Enter getrennt eingeben.

Die Eingaben erfolgen nach dem Prinzip der Umgekehrten Polnischen
Notation (UPN bzw in Englisch RPN) Wikipedia und YouTube helfen hier.

3 [Enter]
= 3

3 [Enter] 5 [Enter] sw [Enter]
= 3

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-right" />

<col  class="org-right" />

<col  class="org-left" />
</colgroup>
<tbody>
<tr>
<td class="org-right">Register</td>
<td class="org-right">Wert</td>
<td class="org-left">Fun Facts</td>
</tr>


<tr>
<td class="org-right">3</td>
<td class="org-right">0</td>
<td class="org-left">&#xa0;</td>
</tr>


<tr>
<td class="org-right">2</td>
<td class="org-right">0</td>
<td class="org-left">&#xa0;</td>
</tr>


<tr>
<td class="org-right">1</td>
<td class="org-right">5</td>
<td class="org-left">mit sw kann man [1] und [0] tauschen</td>
</tr>


<tr>
<td class="org-right">0</td>
<td class="org-right">3</td>
<td class="org-left">&#xa0;</td>
</tr>
</tbody>
</table>

Rechenoperationen, die sich auf das erste oder die ersten Register
beziehen.

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-left" />

<col  class="org-left" />
</colgroup>
<thead>
<tr>
<th scope="col" class="org-left">Befehl</th>
<th scope="col" class="org-left">Funktion    (Erklärung)</th>
</tr>
</thead>

<tbody>
<tr>
<td class="org-left">a</td>
<td class="org-left">Addieren</td>
</tr>


<tr>
<td class="org-left">s -</td>
<td class="org-left">Subtrahieren</td>
</tr>


<tr>
<td class="org-left">d /</td>
<td class="org-left">Dividieren</td>
</tr>


<tr>
<td class="org-left">m \*</td>
<td class="org-left">Muliplizieren</td>
</tr>


<tr>
<td class="org-left">mod %</td>
<td class="org-left">Rest bei Division</td>
</tr>


<tr>
<td class="org-left">ux</td>
<td class="org-left">Kehrwert bilden</td>
</tr>
</tbody>
</table>

Beispiel:

3 5 a
= 8

7 6 m
= 42

5 ux
= 0.2

7 4 mod
= 3

(3\*(4-2))

4 2 s 3 m
= 6

(34/5)\*(3-4\*2)

34 5 d 3 4 2 m s m
= -34

3+4+3+6+9

3 4 a 3 a 6 a 9 a 
= 25

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-left" />
</colgroup>
<thead>
<tr>
<th scope="col" class="org-left">Konstante</th>
</tr>
</thead>

<tbody>
<tr>
<td class="org-left">e</td>
</tr>


<tr>
<td class="org-left">pi</td>
</tr>
</tbody>
</table>

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-left" />
</colgroup>
<thead>
<tr>
<th scope="col" class="org-left">Trigonometrie</th>
</tr>
</thead>

<tbody>
<tr>
<td class="org-left">sin</td>
</tr>


<tr>
<td class="org-left">cos</td>
</tr>


<tr>
<td class="org-left">tan</td>
</tr>


<tr>
<td class="org-left">asin</td>
</tr>


<tr>
<td class="org-left">acos</td>
</tr>


<tr>
<td class="org-left">atan</td>
</tr>
</tbody>
</table>

Beispiel:

Komm Standardkram.

30 sin
= 0.5

0 sin
= 0

0 cos 
= r

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-left" />

<col  class="org-left" />
</colgroup>
<thead>
<tr>
<th scope="col" class="org-left">Logarithmen</th>
<th scope="col" class="org-left">Erklärung</th>
</tr>
</thead>

<tbody>
<tr>
<td class="org-left">lg</td>
<td class="org-left">Log zur Basis 10</td>
</tr>


<tr>
<td class="org-left">ln</td>
<td class="org-left">Log zur Bais e</td>
</tr>


<tr>
<td class="org-left">log</td>
<td class="org-left">log von [1] zur Basis in [0]</td>
</tr>
</tbody>
</table>

Beispiel:

1024 2 log
= 10

100 lg
= 2

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-left" />

<col  class="org-left" />
</colgroup>
<thead>
<tr>
<th scope="col" class="org-left">Potenzen</th>
<th scope="col" class="org-left">Erklärung</th>
</tr>
</thead>

<tbody>
<tr>
<td class="org-left">p</td>
<td class="org-left">Potenzieren    (basis[1] expont[0] p)</td>
</tr>


<tr>
<td class="org-left">x2</td>
<td class="org-left">Quadrieren</td>
</tr>


<tr>
<td class="org-left">sqrt</td>
<td class="org-left">Quadratwurzel</td>
</tr>


<tr>
<td class="org-left">cbrt</td>
<td class="org-left">Kubikwurzel</td>
</tr>


<tr>
<td class="org-left">xrt</td>
<td class="org-left">Nte-Wurzel   ([1] [0] xtr)</td>
</tr>
</tbody>
</table>

Beispiel:

2 10 p
= 1024

10 2 p
= 100

16 sqrt
= 4

27 cbrt
= 3

256 4 xrt 
= 4 


<a id="org0ece62f"></a>

## Sonderfunktionern

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-left" />

<col  class="org-left" />
</colgroup>
<tbody>
<tr>
<td class="org-left">Befehl</td>
<td class="org-left">Funktion</td>
</tr>


<tr>
<td class="org-left">pq</td>
<td class="org-left">Nullstellen mit pq-Formel errechnen</td>
</tr>


<tr>
<td class="org-left">&#xa0;</td>
<td class="org-left">[1] [0] pq</td>
</tr>


<tr>
<td class="org-left">distance</td>
<td class="org-left">Orthodromenstrecke berechnen</td>
</tr>
</tbody>
</table>

3 4 pq
= -1.5 in [1] und -1.5 [0]

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-right" />

<col  class="org-left" />
</colgroup>
<thead>
<tr>
<th scope="col" class="org-right">Register</th>
<th scope="col" class="org-left">Was?</th>
</tr>
</thead>

<tbody>
<tr>
<td class="org-right">3</td>
<td class="org-left">Breite<sub>1</sub></td>
</tr>


<tr>
<td class="org-right">2</td>
<td class="org-left">Länge<sub>1</sub></td>
</tr>


<tr>
<td class="org-right">1</td>
<td class="org-left">Breite<sub>2</sub></td>
</tr>


<tr>
<td class="org-right">0</td>
<td class="org-left">Länge<sub>2</sub></td>
</tr>
</tbody>
</table>

+48 +10 +40 -89 distance
= 7403 km (bis zu Joseph)


<a id="orgb36b024"></a>

## Ergebnisse Speichern     :cool:

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-left" />

<col  class="org-left" />
</colgroup>
<tbody>
<tr>
<td class="org-left">Befehl</td>
<td class="org-left">Funtktion, Paramenter(Freunde)</td>
</tr>


<tr>
<td class="org-left">-></td>
<td class="org-left">[Zahl oder Stack] -> [Variabelname]</td>
</tr>


<tr>
<td class="org-left">get</td>
<td class="org-left">get [Variabelname]</td>
</tr>


<tr>
<td class="org-left">show</td>
<td class="org-left">Alle Variabeln zeigen</td>
</tr>
</tbody>
</table>

3 -> my<sub>var</sub>
13 -> primzahl

3 in die Variabel my<sub>var</sub> speichern und 13 in primzahl speichern

get primzahl
= 13

show

my<sub>var</sub> : 3
primzahl : 13

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-left" />

<col  class="org-left" />

<col  class="org-left" />
</colgroup>
<tbody>
<tr>
<td class="org-left">Befehl</td>
<td class="org-left">Funtktion, Paramenter(Freunde)</td>
<td class="org-left">Fun Facts</td>
</tr>


<tr>
<td class="org-left">store</td>
<td class="org-left">store [Dateiname]</td>
<td class="org-left">Stack(Register) und Variablen in Textdatei speichern</td>
</tr>


<tr>
<td class="org-left">run</td>
<td class="org-left">run [Dateiname]</td>
<td class="org-left">Stack und Variablen aus Textdatei laden</td>
</tr>


<tr>
<td class="org-left">&#xa0;</td>
<td class="org-left">&#xa0;</td>
<td class="org-left">und Rechnungen durchführen</td>
</tr>


<tr>
<td class="org-left">stop</td>
<td class="org-left">&#xa0;</td>
<td class="org-left">beendet Programm</td>
</tr>
</tbody>
</table>


<a id="org7357fac"></a>

## Designkram

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">


<colgroup>
<col  class="org-left" />

<col  class="org-left" />
</colgroup>
<thead>
<tr>
<th scope="col" class="org-left">Befehl</th>
<th scope="col" class="org-left">Funtktion, Paramenter(Freunde)</th>
</tr>
</thead>

<tbody>
<tr>
<td class="org-left">view</td>
<td class="org-left">Toggle Stackanzeige</td>
</tr>


<tr>
<td class="org-left">print</td>
<td class="org-left">Stack einmalig zeigen</td>
</tr>


<tr>
<td class="org-left">bin</td>
<td class="org-left">Binäres (gerundetes) Ergbnis</td>
</tr>
</tbody>
</table>


<a id="orgbcd37e2"></a>

# Achtung

Es kann ab und an auch etwas schiefgehen. Die Benutzung ist frei und
vollkommen eigenverantwortlich.

