#include<stdio.h>
#include<mpi.h>


int main(int argc,char** argv) {
   
   int size,rank;
   MPI_Init(&argc,&argv);
   MPI_Comm_rank(MPI_COMM_WORLD,&rank);
   MPI_Comm_size(MPI_COMM_WORLD,&size);
   
   int R = size;
   int C = size;
   int count = rank * C + 1;
   int a[C];
   for(int i=0;i<C;i++) {
     a[i] = ++count;
   }
   
   int gather_data[R][C];
   MPI_Gather(a,C,MPI_INT,gather_data,C,MPI_INT,0,MPI_COMM_WORLD);
   int transpose[R][C];
   if(rank==0) {
        
        printf("Received matrix at process 0:\n");
        for(int i=0;i<R;i++) {
         for(int j=0;j<C;j++) {
           printf("%d ",gather_data[i][j]);
         
         }
         printf("\n");
        }
        for(int i=0;i<R;i++) {
         for(int j=0;j<C;j++) {
           transpose[j][i] = gather_data[i][j];
         }
        }
       printf("Transpose matrix:\n");
        for(int i=0;i<R;i++) {
         for(int j=0;j<C;j++) {
           printf("%d ",transpose[i][j]);
         
         }
         printf("\n");
        } 
      
   
   } 
   MPI_Bcast(transpose, R, MPI_INT, 0,MPI_COMM_WORLD);
   
   MPI_Finalize();
}
