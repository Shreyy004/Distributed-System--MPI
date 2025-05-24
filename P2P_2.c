#include<stdio.h>
#include<mpi.h>

int main(int argc, char** argv) {
 
   MPI_Init(&argc,&argv);
   int rank , size;
   MPI_Comm_rank(MPI_COMM_WORLD, &rank);
   MPI_Comm_size(MPI_COMM_WORLD, &size);
   
   if (size < 2) {
     printf("Run with atleast 2 processes !\n");
     MPI_Finalize();
     return 0;
   }
   
   int data;
   MPI_Request request;
   if (rank ==0) {
     data = 200;
     printf("Process %d sending data: %d to Process 1 (N-B)\n", rank , data);
     MPI_Isend(&data, 1, MPI_INT, 1, 0, MPI_COMM_WORLD, &request);
     MPI_Wait(&request, MPI_STATUS_IGNORE);
   }
   else if (rank == 1) {
     MPI_Irecv(&data, 1, MPI_INT, 0, 0, MPI_COMM_WORLD, &request);
     MPI_Wait(&request, MPI_STATUS_IGNORE);
     printf("Process %d received data: %d from process 0 (N_B)\n", rank, data);
     
   }
   MPI_Finalize();
   return 0;

}
