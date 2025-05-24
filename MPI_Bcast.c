#include<stdio.h>
#include<mpi.h>

int main(int argc, char** argv) {
   
   MPI_Init(&argc, &argv);
   
   int rank, value = 0;
   
   MPI_Comm_rank(MPI_COMM_WORLD, &rank);
   
   while(value >= 0) {
    
    if(rank ==0) {
      scanf("%d", &value);
    }
    MPI_Bcast(&value, 1,MPI_INT,0,MPI_COMM_WORLD);
    printf("Processes %d got %d\n",rank,value);
   
   }
   MPI_Finalize();
   return 0;



}
