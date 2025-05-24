#include <stdio.h>
#include <mpi.h>
int a = 10;
int main(int argc, char** argv) {

   MPI_Init(&argc,&argv);
   int rank,size;
   MPI_Request request;
   MPI_Status status;
   MPI_Comm_rank(MPI_COMM_WORLD, &rank);
   MPI_Comm_size(MPI_COMM_WORLD, &size);
   
   if (rank == 0) {
      
     int x = 12;
     int  y = 10;
    //int MPI_Isend(void *buf, int count, MPI_Datatype datatype, int dest, int tag, MPI_Comm comm, MPI_Request *request);

    MPI_Isend(&a, 1, MPI_INT, 1, 5, MPI_COMM_WORLD,&request);
     int b = x*x;
     int c = y*y*y;
     int d = b+c;
     MPI_Wait(&request,&status);
     int e = a+d;
     a = d;
     printf("For Rank 0 %d\n",a);
   }
   else if (rank == 1) {
    
    MPI_Irecv(&a, 1,MPI_INT, 0, 5, MPI_COMM_WORLD,&request);
    MPI_Wait(&request,&status);
    printf("%d\n",a);
   
   }
   
   MPI_Finalize();	
   
  }















