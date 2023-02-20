# Bridge

Bridge ist ein strukturelles Entwurfsmuster, das die Geschäftslogik oder eine große Klasse in separate Klassenhierarchien unterteilt, die unabhängig voneinander entwickelt werden können.

Eine dieser Hierarchien (oft Abstraktion genannt), erhält einen Verweis auf ein Objekt der zweiten Hierarchie (Implementierung). Die Abstraktion ist in der Lage, einige (manchmal auch die meisten) ihrer Aufrufe an das Implementierungsobjekt zu delegieren. Da alle Implementierungen eine gemeinsame Schnittstelle haben, sind sie innerhalb der Abstraktion austauschbar.
