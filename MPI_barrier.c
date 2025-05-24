/*
 
 is a synchronization fn in mpi that ensures all processes in a communicator reach the barrier before any can proceed.It acts like a checkpoint, making sure that all processes wait at the barrier until every process has reached it.
 
 
 MPI_Barrier(MPI_Comm comm);
 communicator that defines the group of processes that must synchronize.
 
 why?
 - ensure order
 - data consistency
 - perfromance

#include<stdio.h>
#include<mpi.h>

int main(int argc, char** argv) {
  
  MPI_Init(&argc, &argv);
  
  int rank;
  MPI_Comm_rank(MPI_COMM_WORLD, &rank);
  
  printf("Procss %d reached this point, rank);
  
  MPI_Finalize():
  return 0;


}

not synchronized


*/


#include<stdio.h>
#include<mpi.h>

int main(int argc, char** argv) {
  
  MPI_Init(&argc, &argv);
  
  int rank;
  MPI_Comm_rank(MPI_COMM_WORLD, &rank);
  
  MPI_Barrier(MPI_COMM_WORLD); // exec with an without and chek the difference
  printf("Procss %d reached this point\n", rank);
  
  MPI_Finalize();
  return 0;


}
