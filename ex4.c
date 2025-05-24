//This MPI program requires exactly 3 processes and distributes a 4×4 matrix in a way that odd numbers are sent first, followed by even numbers.



#include<stdio.h>
#include<mpi.h>
int r = 4;
int c = 4;

int main(int argc,char** argv) {

   MPI_Init(&argc,&argv);
   int rank , size;
   MPI_Comm_rank(MPI_COMM_WORLD,&rank);
   MPI_Comm_size(MPI_COMM_WORLD,&size);
   
   if (size != 3) {
     if(rank == 0) {
      printf("This program requires 3 MPI processes to run\n");
     }
     MPI_Finalize();
     return 1;
   }
   int a[4][4];
   int s_d[(r*c/3)+1];
   int p_d[r*c];
   if (rank == 0) {
   
      int index = 0;
      int data[4][4] = {{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16}};
      
      for (int i=0;i<r;i++) {
       for(int j=0;j<c;j++) {
         if(data[i][j] % 2 == 1) {
           p_d[index++] = data[i][j];
         }
       }
      }
      
      for(int i=0;i<r;i++) {
        for(int j=0;j<c;j++) {
            if (data[i][j] % 2 ==0 ) {
              p_d[index++] = data[i][j];
            }
        }
      }
   
     
   }
   
   MPI_Scatter(p_d, r*c/3, MPI_INT, s_d, r*c/3, MPI_INT, 0, MPI_COMM_WORLD);
   
   printf("Process %d received:\n",rank);
   for(int i=0;i<r*c/3;i++) {
     printf("%d ",s_d[i]);
   }
   printf("\n");
   
   MPI_Finalize();
   
   return 0;
   
   
}
