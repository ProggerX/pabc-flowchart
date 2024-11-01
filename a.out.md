flowchart TB
var0[Объявить a типа Integer]
mbbeg([НАЧАЛО])
mb0["writeln('Give me a number')"]
mbbeg-->mb0
mb1["readln(a)"]
mb0-->mb1
mb2["writeln('Your number is ', a, '!')"]
mb1-->mb2
mb3if{"a mod 2 = 0 ?"}
mb3ifthen["writeln('wow! your num is even!')"]
mb3if-->|тогда|mb3ifthen
mb3if-->|иначе|mb3ifend
mb3ifthen-->mb3ifend
mb3ifend[Конец условия]
mb2-->mb3if
mb4as["Присвоить a значение a + 2"]
mb3ifend-->mb4as
mb5["writeln('Your num + 2 = ', a)"]
mb4as-->mb5
mb6forassign["Присвоить a значение 0"]
mb6for{"a <= 10 ?"}
mb6forassign-->mb6for
mb6forchange["Присвоить a значение a + 2"]
mb6for-->|тогда|mb6forchange
mb6forbody["writeln(a)"]
mb6forchange-->mb6forbody
mb6forbody-->mb6for
mb5-->mb6forassign
mb7forassign["Присвоить a значение 1"]
mb7for{"a <= 9 ?"}
mb7forassign-->mb7for
mb7forchange["Присвоить a значение a + 2"]
mb7for-->|тогда|mb7forchange
mb7forbody["writeln(a)"]
mb7forchange-->mb7forbody
mb7forbody-->mb7for
mb6for-->mb7forassign
mb7for-->mbend
mbend([КОНЕЦ])