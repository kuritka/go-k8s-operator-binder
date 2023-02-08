# k8s-map-binder 
My kubernetes operators usually read the annotations of operated resourcers.
This GO library provides simple way to parse annotations into custom structures. 
It supports slices in addition to primitive types. Besides the actual parsing, 
it also offers ways to bind fields such as default values, require, 
value protection, private fields and nested structures.

### QuickStart
