#include <stdio.h>
#include <mpi.h>

int main(int argc, char** argv) {
  
   int size, rank;
   MPI_Init(&argc,&argv);
   MPI_Comm_rank(MPI_COMM_WORLD,&rank);
   MPI_Comm_size(MPI_COMM_WORLD,&size);
   
   int a1[4][4] = {{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16}};
   int r1[4];
   printf("Started\n");
   
   MPI_Scatter(a1,4,MPI_INT,r1,4,MPI_INT,0,MPI_COMM_WORLD);
   
   int gather[4][4];
   MPI_Gather(r1,4,MPI_INT,gather,4,MPI_INT,0,MPI_COMM_WORLD);
   
   if (rank == 0) {
     
     
     for(int i=0;i<4;i++) {
       for(int j=0;j<4;j++) {
          printf("%d ",gather[i][j]);
       }
       printf("\n");
     }
   
   }
   MPI_Finalize();
}
