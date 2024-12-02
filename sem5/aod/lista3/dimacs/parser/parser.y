%{
    
package parser

import (
    "fmt"
    "dijkstra/graph"
    "strconv"
)

var result []interface{}
var problemGraph graph.Graph{}

func (y *YySymType) ChangeStr(str string) {
    y.str = str
}
%}

%union {
    str    string
    number int
}

%token <str> PROBLEM
%token <str> ARC
%token <str> PROBLEM_TYPE
%token <str> NUMBER
%token EOL

%type <str> ProblemLine ArcLine Line

%start Input

%%

Input:
    /* empty */
    | Input Line EOL
;

Line:
    ProblemLine
    | ArcLine
;


ProblemLine:
    PROBLEM PROBLEM_TYPE NUMBER NUMBER  {
        size, _ := strconv.Atoi($3)
        problemGraph = graph.NewGraph(size)
        fmt.Printf("Problem: Type=%s, Wierzchołki=%d, Łuki=%d\n", $2, $3, $4)
    }
;

ArcLine:
    ARC NUMBER NUMBER NUMBER  {
        from, _ := strconv.Atoi($2)
        to, _ := strconv.Atoi($3)
        cost, _ := strconv.Atoi($4)
        problemGraph.AddEdge(from, to, cost)
        fmt.Printf("Arc: From=%d, To=%d, Cost=%d\n", $2, $3, $4)
    }
;

%%

func yyError(str string) {
    fmt.Println("Error:", str)
}
