# Letter Expression Solver

![Main Menu](/images/ProblemStatement.png)

The problem statement above asks for a solution in which one specific set of expressions can be solved. I went down the route of creating a scalable application that can solve more than just one specific expression set. With this, it bumps the run time up a little bit but it is worth it.

## Getting Started

The building of this project is handled by a Makefile. It has the general commands of:

* all
* clean

and it has specific build commands depending on the OS you are using:

* darwin
* linux
* windows

### Prerequisites

* Download and Install Go - [Go Downloads](https://golang.org/dl/)
* Clone the project into your go src directory:

```bash
cd $HOME/go/src
git clone https://github.com/Cbuckles17/expressionsolver.git expressionsolver
```

### Build and Run the Executable

```bash
cd $HOME/go/src/expressionsolver
#make all (will not auto run)
make
#make darwin specific
make run-darwin
#make linux specific
make run-linux
#make windows specific
make run-windows
```

## Built With

* [Go](https://golang.org/doc/) - Coding Language
