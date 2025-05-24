#include "mpi.h"


#include<stdio.h>


int main(int argc,char *argv[])


{


int rank,size,i;


int b[10];


MPI_Status status;


MPI_Init(&argc,&argv);


MPI_Comm_size(MPI_COMM_WORLD,&size);


MPI_Comm_rank(MPI_COMM_WORLD,&rank);


if(size<2){


printf("run with 2 processes");


fflush(stdout);


MPI_Finalize();


}


if(rank==0){


for(i=0;i<10;i++)


b[i]=i;


MPI_Ssend(b,10,MPI_INT,1,123,MPI_COMM_WORLD);


}

 

if(rank==1){


for(i=0;i<10;i++)


b[i]=-1; 


MPI_Recv(b,10,MPI_INT,0,123,MPI_COMM_WORLD,&status);


for(i=0;i<10;i++){


printf("%d \t",b[i]);


if(b[i]!=i){


printf("Error:buffer[%d] = %d but is exepcted to be %d \n",i,b[i],i);


}


}


fflush(stdout);


}


MPI_Finalize();


}


