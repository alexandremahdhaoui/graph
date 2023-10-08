# graph

Simple graph datastructures for Go

## TODO

- [ ] Recreate the GRAPH datastructure
- [ ] Recreate the topological sorting

## Graph for DI

    // How to design the stuff in di:
    // - We have a Tree or a Graph that lays relationship between functions and other stuffs
    //
    // We can view our project as a mere attempt to copy the AST. But what we want to achieve is producing a Graph or Tree
    // of the executed code, not syntax.
    //
    // Therefor, we will have multiple kinds:
    // - Function
    // - Func Body ? StatementList?
    // - Variable (can be created from an assignment or a var decl obviously or could also be a func declaration)
    // - ...
