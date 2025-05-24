#include <stdio.h>
#include <mpi.h>
#include <string.h>

int main(int argc, char** argv) {
  
   MPI_Init(&argc,&argv);
   int rank,size;
   MPI_Comm_rank(MPI_COMM_WORLD,&rank);
   MPI_Comm_size(MPI_COMM_WORLD,&size);
   char data[100];
   if (rank == 0) {
      
      
      scanf("%s",data);
   
      MPI_Send(data,strlen(data)+1,MPI_CHAR,1,0,MPI_COMM_WORLD);
      printf("Process %d receives %s\n",rank,data);
      
      MPI_Recv(data,100,MPI_CHAR,(size-1),0,MPI_COMM_WORLD,MPI_STATUS_IGNORE);
      printf("Process %d receives %s\n",rank,data);
   }
   
   else {
      MPI_Recv(data,100,MPI_CHAR,(rank-1),0,MPI_COMM_WORLD,MPI_STATUS_IGNORE);
      printf("Process %d receives %s\n",rank,data);
      
      int x = (rank+1) % size;
      MPI_Send(data,strlen(data)+1,MPI_CHAR,x,0,MPI_COMM_WORLD);
   }
   MPI_Finalize();
   
}
