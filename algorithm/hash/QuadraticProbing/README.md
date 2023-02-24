# Quadratic Probing

## Konzept

Quadratic Probing ist ein offenes Adressierungsschema in der Computerprogrammierung zur Lösung von Hash-Kollisionen in Hash-Tabellen. Beim quadratischen Sondieren wird der ursprüngliche Hash-Index genommen und aufeinanderfolgende Werte eines beliebigen quadratischen Polynoms hinzugefügt, bis ein offener Slot gefunden wird.

Die quadratische Sondierung kann in einer offenen Adressierungstabelle ein effizienterer Algorithmus sein, da sie das Clustering-Problem, das bei der linearen Sondierung auftreten kann, besser vermeidet, obwohl sie nicht immun ist. Er bietet auch eine gute Zwischenspeicherung im Speicher, da er eine gewisse Lokalität des Verweises beibehält; die lineare Sondierung hat jedoch eine größere Lokalität und damit eine bessere Cache-Leistung.
