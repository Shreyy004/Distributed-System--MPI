#include <stdio.h>
#include <stdlib.h>
#include <mpi.h>

int main(int argc, char* argv[])
{
MPI_Init(&argc, &argv);
int my_rank;
int a[10];
for(int i=0;i<10;i++){
a[i]=i+1;
a[i]=a[i]*10;
}

MPI_Comm_rank(MPI_COMM_WORLD, &my_rank);
if(my_rank==0){
MPI_Bcast(a,10,MPI_INT,0,MPI_COMM_WORLD);
printf("Broad Casted from %d\n",my_rank);
}
else{
printf("Recieved by %d and the message contains ",my_rank);
for(int i=0;i<10;i++){
printf("%d ",a[i]);
}
printf("\n");

}
MPI_Finalize();
return EXIT_SUCCESS;
}
