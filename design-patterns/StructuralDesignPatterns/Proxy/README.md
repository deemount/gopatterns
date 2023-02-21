# Proxy

Proxy ist ein strukturelles Entwurfsmuster, das ein Objekt bereitstellt, das als Ersatz für ein echtes, von einem Client verwendeten Dienstobjekt fungiert. Ein Proxy empfängt Client-Anfragen, erledigt einige Aufgaben (Zugriffskontrolle, Zwischenspeicherung usw.) und leitet die Anfrage dann an ein Dienstobjekt weiter.

Das Proxy-Objekt hat dieselbe Schnittstelle wie ein Dienst, so dass es mit einem echten Objekt austauschbar ist, wenn es an einen Client übergeben wird.
