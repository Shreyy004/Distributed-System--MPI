#include <stdio.h>
#include "mpi.h"


////////////
//MPI_Isend
//MPI_REQUEST_NULL////////////
//
// int MPI_Isend(const void *buf, int count, MPI_Datatype datatype, int dest, int tag,
//              MPI_Comm comm, MPI_Request *request)
//
// This example uses MPI_Isend to do a non-blocking send of information from the root process to a destination process.
// The destination process is set as a variable in the code and must be less than the number of processes started.
//
// example usage:
//              compile: mpicc -o mpi_isend mpi_isend.c
//              run: mpirun -n 4 mpi_isend
//
int main(argc, argv)
int argc;
char **argv;
{
    int rank, size;
    int tag, destination, count;
    int buffer; //value to send
int buf;
    tag = 1234;
    destination = 2; //destination process
    count = 1; //number of elements in buffer

    MPI_Status status;
    MPI_Request request = MPI_REQUEST_NULL,request1=MPI_REQUEST_NULL;

    MPI_Init(&argc, &argv);

    MPI_Comm_size(MPI_COMM_WORLD, &size); //number of processes
    MPI_Comm_rank(MPI_COMM_WORLD, &rank); //rank of current process
    printf("Enter a value to send to processor %d:\n", destination);
        scanf("%d", &buffer);

    if (rank == 0) {
       
        MPI_Isend(&buffer, count, MPI_INT, destination, tag, MPI_COMM_WORLD, &request1); //non blocking send to destination process
       MPI_Irecv(&buf, count, MPI_INT, 2, tag, MPI_COMM_WORLD, &request); //destination process receives
        printf("hello");
    }

    if (rank == destination) {
       MPI_Issend(&buffer, count, MPI_INT, 0,tag, MPI_COMM_WORLD, &request);
        MPI_Irecv(&buffer, count, MPI_INT, 0, tag, MPI_COMM_WORLD, &request1); //destination process receives
        printf("done");
}

 MPI_Wait(&request1,&status); 
  MPI_Wait(&request, &status);
 
    if (rank == 0) {
        printf("processor %d sent %d\n", rank, buf);
    }
    if (rank == destination) {
        printf("processor %d got %d\n", rank, buffer);
    }

    MPI_Finalize();

        return 0;
}

